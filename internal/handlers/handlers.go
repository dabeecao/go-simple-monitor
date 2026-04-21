package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	stdnet "net"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"go-simple-monitor/internal/config"
	"go-simple-monitor/internal/db"
	"go-simple-monitor/internal/monitor"
	"go-simple-monitor/internal/pty"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/shirou/gopsutil/v3/process"
	"golang.org/x/time/rate"
)

var (
	limiterMap   = make(map[string]*rate.Limiter)
	limiterMutex sync.Mutex
)

func round(num float64) float64 {
	return math.Round(num*100) / 100
}

func getLocalIP() string {
	conn, err := stdnet.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "127.0.0.1"
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*stdnet.UDPAddr)
	return localAddr.IP.String()
}

// --- MIDDLEWARES ---
func getLimiter(ip string) *rate.Limiter {
	limiterMutex.Lock()
	defer limiterMutex.Unlock()
	limiter, exists := limiterMap[ip]
	if !exists {
		limiter = rate.NewLimiter(rate.Every(15*time.Minute/5), 5)
		limiterMap[ip] = limiter
	}
	return limiter
}

func rateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !getLimiter(c.ClientIP()).Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Quá nhiều yêu cầu"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"detail": "Missing Token"})
			c.Abort()
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) { return config.JwtSecret, nil })
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"detail": "Token không hợp lệ hoặc đã hết hạn"})
			c.Abort()
			return
		}
		c.Next()
	}
}

// --- ROUTES ĐĂNG KÝ VÀO GIN ---
func RegisterRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) { c.HTML(http.StatusOK, "index.html", nil) })
	
	r.POST("/api/login", rateLimit(), func(c *gin.Context) {
		var req struct { Username string `json:"username"`; Password string `json:"password"` }
		if c.ShouldBindJSON(&req) == nil && req.Username == config.AdminUser && req.Password == config.AdminPass {
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"sub": req.Username, "exp": time.Now().Add(time.Duration(config.JwtExpirationMins) * time.Minute).Unix(),
			})
			tokenString, _ := token.SignedString(config.JwtSecret)
			c.JSON(http.StatusOK, gin.H{"token": tokenString})
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{"detail": "Sai tài khoản hoặc mật khẩu"})
	})

	api := r.Group("/api")
	api.Use(authMiddleware())
	api.GET("/stats", getStats)
	api.GET("/system", getSystem)
	api.GET("/network", getNetwork)
	api.GET("/ports", getPorts)
	api.POST("/kill/:pid", killProcess)
	api.GET("/settings", getSettings)
	api.POST("/settings", setSettings)
	api.POST("/settings/test", testTgSettings)
	api.GET("/cron", getCrons)
	api.POST("/cron", addCron)
	api.DELETE("/cron/:job_id", deleteCron)

	r.GET("/ws/terminal", pty.TerminalSocket)
}

// --- IMPLEMENTATIONS ---
func getStats(c *gin.Context) {
	vm, _ := mem.VirtualMemory(); swap, _ := mem.SwapMemory()
	netIO, _ := net.IOCounters(false); diskIO, _ := disk.IOCounters()
	cpuPercents, _ := cpu.Percent(0, false); cpuPerCore, _ := cpu.Percent(0, true)
	cpuCountsCore, _ := cpu.Counts(false); cpuCountsThread, _ := cpu.Counts(true)
	cpuInfo, _ := cpu.Info(); hostStat, _ := host.Info()

	var diskRead, diskWrite float64
	for _, stat := range diskIO { diskRead += float64(stat.ReadBytes); diskWrite += float64(stat.WriteBytes) }

	cpuModel := "Unknown CPU"; freqMax := 0.0
	if len(cpuInfo) > 0 { cpuModel = cpuInfo[0].ModelName; freqMax = cpuInfo[0].Mhz }

	partitions, _ := disk.Partitions(false)
	var disksInfo []map[string]interface{}
	var totalDisk, usedDisk uint64

	for _, part := range partitions {
		if part.Fstype == "" || part.Fstype == "squashfs" || part.Fstype == "tmpfs" || part.Fstype == "overlay" || part.Fstype == "loop" || strings.HasPrefix(part.Mountpoint, "/boot") || strings.HasPrefix(part.Mountpoint, "/snap") { continue }
		if usage, err := disk.Usage(part.Mountpoint); err == nil && usage.Total > 0 {
			totalDisk += usage.Total; usedDisk += usage.Used
			disksInfo = append(disksInfo, map[string]interface{}{ 
                "mountpoint": part.Mountpoint, 
                "percent": round(usage.UsedPercent), 
                "total": round(float64(usage.Total) / (1024 * 1024 * 1024)), 
                "used": round(float64(usage.Used) / (1024 * 1024 * 1024)), 
                "free": round(float64(usage.Free) / (1024 * 1024 * 1024)),
            })
		}
	}
	aggDiskPercent := 0.0
	if totalDisk > 0 { aggDiskPercent = (float64(usedDisk) / float64(totalDisk)) * 100 }
	uptimeSec := int(time.Now().Unix()) - int(monitor.BootTime)

    var cpuTotal float64
    if len(cpuPercents) > 0 { cpuTotal = round(cpuPercents[0]) }

    var roundedCpuPerCore []float64
    for _, coreP := range cpuPerCore {
        roundedCpuPerCore = append(roundedCpuPerCore, round(coreP))
    }

    var netSent, netRecv float64
    if len(netIO) > 0 {
        netSent = round(float64(netIO[0].BytesSent) / (1024 * 1024))
        netRecv = round(float64(netIO[0].BytesRecv) / (1024 * 1024))
    }

	c.JSON(http.StatusOK, gin.H{
		"cpu": cpuTotal, 
        "cpu_per_core": roundedCpuPerCore, 
        "cpu_cores": cpuCountsCore, 
        "cpu_threads": cpuCountsThread,
        "cpu_model": cpuModel, 
        "cpu_freq": map[string]float64{"current": round(freqMax), "max": round(freqMax)},
		"ram": round(vm.UsedPercent), 
        "ram_total": round(float64(vm.Total) / (1024 * 1024 * 1024)), 
        "ram_used": round(float64(vm.Used) / (1024 * 1024 * 1024)),
        "ram_active": round(float64(vm.Active) / (1024 * 1024 * 1024)),
        "ram_cached": round(float64(vm.Cached) / (1024 * 1024 * 1024)),
        "ram_buffers": round(float64(vm.Buffers) / (1024 * 1024 * 1024)),
		"swap_percent": round(swap.UsedPercent), 
        "swap_total": round(float64(swap.Total) / (1024 * 1024 * 1024)),
        "swap_used": round(float64(swap.Used) / (1024 * 1024 * 1024)),
        "disk": round(aggDiskPercent), 
        "disk_total": round(float64(totalDisk) / (1024 * 1024 * 1024)),
        "disk_used": round(float64(usedDisk) / (1024 * 1024 * 1024)),
        "disk_io": map[string]float64{"read": round(diskRead / (1024 * 1024 * 1024)), "write": round(diskWrite / (1024 * 1024 * 1024))},
		"disks": disksInfo, 
        "uptime": fmt.Sprintf("%dh %dm", uptimeSec/3600, (uptimeSec%3600)/60),
		"sys_info": map[string]string{ "os": getPrettyOS(), "hostname": hostStat.Hostname, "ip": getLocalIP(),},
		"net_io": map[string]float64{"sent": netSent, "recv": netRecv},
	})
}

func getSystem(c *gin.Context) {
	procs, _ := process.Processes()
	var res []map[string]interface{}
	for _, p := range procs {
		name, _ := p.Name(); cmdline, _ := p.Cmdline(); username, _ := p.Username()
		cpuP, _ := p.CPUPercent(); memP, _ := p.MemoryPercent(); memI, _ := p.MemoryInfo()
		if cmdline != "" { name = cmdline }
		memMb := 0.0
		if memI != nil { memMb = float64(memI.RSS) / (1024 * 1024) }
		res = append(res, map[string]interface{}{ "pid": p.Pid, "name": name, "username": username, "cpu_percent": round(cpuP), "memory_percent": round(float64(memP)), "memory_mb": round(memMb) })
	}
	c.JSON(http.StatusOK, res)
}

func getPrettyOS() string {
    if runtime.GOOS == "linux" {
        data, err := os.ReadFile("/etc/os-release")
        if err == nil {
            var prettyName, codename string

            for _, line := range strings.Split(string(data), "\n") {
                if strings.HasPrefix(line, "PRETTY_NAME=") {
                    val := strings.TrimPrefix(line, "PRETTY_NAME=")
                    prettyName = strings.Trim(val, `"`)
                }

                if strings.HasPrefix(line, "VERSION_CODENAME=") {
                    val := strings.TrimPrefix(line, "VERSION_CODENAME=")
                    codename = strings.Trim(val, `"`)
                }
            }

            arch := runtime.GOARCH
            if arch == "amd64" {
                arch = "x86_64"
            }

            if prettyName != "" {
                if codename != "" {
                    return fmt.Sprintf("%s (%s) %s", prettyName, codename, arch)
                }
                return fmt.Sprintf("%s %s", prettyName, arch)
            }
        }
    }

    hostStat, _ := host.Info()
    arch := runtime.GOARCH
    if arch == "amd64" {
        arch = "x86_64"
    }

    return fmt.Sprintf("%s %s %s",
        hostStat.Platform,
        hostStat.PlatformVersion,
        arch,
    )
}

func getNetwork(c *gin.Context) {
	conns, _ := net.Connections("inet")
	currentNethogs := monitor.GetCurrentNethogs()
	var res []map[string]interface{}
	for _, conn := range conns {
		name, laddr, tx, rx := "-", "-", 0.0, 0.0
		if conn.Pid != 0 {
			if p, err := process.NewProcess(conn.Pid); err == nil {
				if cmd, err := p.Cmdline(); err == nil && cmd != "" { name = cmd } else if n, err := p.Name(); err == nil { name = n }
			}
			if nh, ok := currentNethogs[conn.Pid]; ok { tx = nh.Tx; rx = nh.Rx }
		}
		if conn.Laddr.IP != "" { laddr = fmt.Sprintf("%s:%d", conn.Laddr.IP, conn.Laddr.Port) }
		res = append(res, map[string]interface{}{ "pid": conn.Pid, "name": name, "status": conn.Status, "laddr": laddr, "tx": round(tx), "rx": round(rx) })
	}
	c.JSON(http.StatusOK, res)
}

func getPorts(c *gin.Context) {
	conns, _ := net.Connections("inet")
	portsMap := make(map[uint32]map[string]interface{})
	for _, conn := range conns {
		if conn.Status == "LISTEN" && conn.Laddr.IP != "" {
			name := "System"
			if p, err := process.NewProcess(conn.Pid); err == nil && conn.Pid != 0 { name, _ = p.Name() }
			portsMap[conn.Laddr.Port] = map[string]interface{}{ "port": conn.Laddr.Port, "ip": conn.Laddr.IP, "pid": conn.Pid, "name": name }
		}
	}
	var res []map[string]interface{}
	for _, v := range portsMap { res = append(res, v) }
	c.JSON(http.StatusOK, res)
}

func killProcess(c *gin.Context) {
	pid, _ := strconv.Atoi(c.Param("pid"))
	if p, err := process.NewProcess(int32(pid)); err == nil { p.Kill(); c.JSON(http.StatusOK, gin.H{"status": "success"}) } else { c.JSON(http.StatusBadRequest, gin.H{"detail": "Cannot kill"}) }
}

func getSettings(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"tg_token": db.GetSetting("tg_token", ""), "tg_chat_id": db.GetSetting("tg_chat_id", ""),
		"tg_cpu_enabled": db.GetSetting("tg_cpu_enabled", "0") == "1", "tg_cpu_threshold": db.ParseFloatSetting("tg_cpu_threshold", 90),
		"tg_ram_enabled": db.GetSetting("tg_ram_enabled", "0") == "1", "tg_ram_threshold": db.ParseFloatSetting("tg_ram_threshold", 90),
		"tg_disk_enabled": db.GetSetting("tg_disk_enabled", "0") == "1", "tg_disk_threshold": db.ParseFloatSetting("tg_disk_threshold", 90),
	})
}

func setSettings(c *gin.Context) {
	var d struct {
		TgToken string `json:"tg_token"`; TgChatID string `json:"tg_chat_id"`
		TgCpuEnabled bool `json:"tg_cpu_enabled"`; TgCpuThreshold float64 `json:"tg_cpu_threshold"`
		TgRamEnabled bool `json:"tg_ram_enabled"`; TgRamThreshold float64 `json:"tg_ram_threshold"`
		TgDiskEnabled bool `json:"tg_disk_enabled"`; TgDiskThreshold float64 `json:"tg_disk_threshold"`
	}
	if c.ShouldBindJSON(&d) != nil { c.JSON(http.StatusBadRequest, gin.H{"detail": "Invalid Data"}); return }
	db.SetSetting("tg_token", d.TgToken); db.SetSetting("tg_chat_id", d.TgChatID)
	db.SetSetting("tg_cpu_enabled", map[bool]string{true: "1", false: "0"}[d.TgCpuEnabled]); db.SetSetting("tg_cpu_threshold", fmt.Sprintf("%v", d.TgCpuThreshold))
	db.SetSetting("tg_ram_enabled", map[bool]string{true: "1", false: "0"}[d.TgRamEnabled]); db.SetSetting("tg_ram_threshold", fmt.Sprintf("%v", d.TgRamThreshold))
	db.SetSetting("tg_disk_enabled", map[bool]string{true: "1", false: "0"}[d.TgDiskEnabled]); db.SetSetting("tg_disk_threshold", fmt.Sprintf("%v", d.TgDiskThreshold))
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func testTgSettings(c *gin.Context) {
	var d struct { TgToken string `json:"tg_token"`; TgChatID string `json:"tg_chat_id"` }
	if c.ShouldBindJSON(&d) != nil || d.TgToken == "" || d.TgChatID == "" { c.JSON(http.StatusBadRequest, gin.H{"detail": "Thiếu dữ liệu"}); return }
	payload := map[string]interface{}{"chat_id": d.TgChatID, "text": "✅ <b>Test Message</b>\n\nHoạt động tốt!", "parse_mode": "HTML"}
	jsonData, _ := json.Marshal(payload)
	resp, err := http.Post(fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", d.TgToken), "application/json", bytes.NewBuffer(jsonData))
	if err != nil || resp.StatusCode != 200 { c.JSON(http.StatusBadRequest, gin.H{"detail": "Error from Telegram"}); return }
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func getCrontabLines() []string {
	if out, err := exec.Command("crontab", "-l").Output(); err == nil { return strings.Split(strings.TrimSpace(string(out)), "\n") }
	return []string{}
}
func writeCrontabLines(lines []string) {
	cmd := exec.Command("crontab", "-"); cmd.Stdin = strings.NewReader(strings.Join(lines, "\n") + "\n"); cmd.Run()
}

func getCrons(c *gin.Context) {
	var res []map[string]interface{}
	for i, line := range getCrontabLines() {
		if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "#") { continue }
		parts := strings.Fields(line)
		if len(parts) >= 6 { res = append(res, map[string]interface{}{"id": i, "schedule": strings.Join(parts[:5], " "), "command": strings.Join(parts[5:], " "), "comment": ""}) }
	}
	c.JSON(http.StatusOK, res)
}

func addCron(c *gin.Context) {
	var item struct { Schedule string `json:"schedule"`; Command string `json:"command"`; Comment string `json:"comment"` }
	if c.ShouldBindJSON(&item) != nil { c.JSON(http.StatusBadRequest, gin.H{"detail": "Invalid Data"}); return }
	lines := getCrontabLines()
	newLine := fmt.Sprintf("%s %s", item.Schedule, item.Command)
	if item.Comment != "" { newLine += " # " + item.Comment }
	writeCrontabLines(append(lines, newLine))
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func deleteCron(c *gin.Context) {
	jobID, _ := strconv.Atoi(c.Param("job_id"))
	lines := getCrontabLines()
	if jobID >= 0 && jobID < len(lines) { writeCrontabLines(append(lines[:jobID], lines[jobID+1:]...)) }
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}