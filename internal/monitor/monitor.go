package monitor

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	"go-simple-monitor/internal/db"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

type NetHogData struct { Tx, Rx float64 }

var (
	NethogsData = make(map[int32]NetHogData)
	nethogsLock sync.RWMutex
	BootTime    uint64
)

func Start() {
	stat, _ := host.BootTime()
	BootTime = stat
	go parseNethogs()
	go monitorSystem()
}

func GetCurrentNethogs() map[int32]NetHogData {
	nethogsLock.RLock()
	defer nethogsLock.RUnlock()
	copyMap := make(map[int32]NetHogData)
	for k, v := range NethogsData { copyMap[k] = v }
	return copyMap
}

func parseNethogs() {
	for {
		cmd := exec.Command("nethogs", "-t", "-d", "2")
		stdout, err := cmd.StdoutPipe()
		if err != nil { time.Sleep(5 * time.Second); continue }
		cmd.Start()
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if line == "" || strings.HasPrefix(line, "Refreshing:") { continue }
			parts := strings.Split(line, "\t")
			if len(parts) >= 3 {
				tx, _ := strconv.ParseFloat(parts[1], 64)
				rx, _ := strconv.ParseFloat(parts[2], 64)
				infoParts := strings.Split(parts[0], "/")
				if len(infoParts) >= 2 {
					if pid, err := strconv.Atoi(infoParts[len(infoParts)-2]); err == nil {
						nethogsLock.Lock()
						NethogsData[int32(pid)] = NetHogData{Tx: tx, Rx: rx}
						nethogsLock.Unlock()
					}
				}
			}
		}
		cmd.Wait()
		time.Sleep(2 * time.Second)
	}
}

func monitorSystem() {
	var lastAlertTime int64 = 0
	for {
		cpuPercents, _ := cpu.Percent(time.Second, false)
		cpuVal := 0.0
		if len(cpuPercents) > 0 { cpuVal = cpuPercents[0] }
		vmStat, _ := mem.VirtualMemory()
		now := time.Now().Unix()

		if now-lastAlertTime > 300 {
			var alerts []string
			if db.GetSetting("tg_cpu_enabled", "0") == "1" && cpuVal >= db.ParseFloatSetting("tg_cpu_threshold", 90) {
				alerts = append(alerts, fmt.Sprintf("🔴 <b>CPU Overload:</b> %.2f%%", cpuVal))
			}
			if db.GetSetting("tg_ram_enabled", "0") == "1" && vmStat.UsedPercent >= db.ParseFloatSetting("tg_ram_threshold", 90) {
				alerts = append(alerts, fmt.Sprintf("🔴 <b>RAM Overload:</b> %.2f%%", vmStat.UsedPercent))
			}
			if db.GetSetting("tg_disk_enabled", "0") == "1" {
				diskThreshold := db.ParseFloatSetting("tg_disk_threshold", 90)
				partitions, _ := disk.Partitions(false)
				var fullDisks []string
				for _, part := range partitions {
					if part.Fstype == "" || part.Fstype == "squashfs" || part.Fstype == "tmpfs" || part.Fstype == "overlay" || part.Fstype == "loop" || strings.HasPrefix(part.Mountpoint, "/boot") { continue }
					usage, err := disk.Usage(part.Mountpoint)
					if err == nil && usage.Total > 0 && usage.UsedPercent >= diskThreshold {
						fullDisks = append(fullDisks, fmt.Sprintf("%s (%.2f%%)", part.Mountpoint, usage.UsedPercent))
					}
				}
				if len(fullDisks) > 0 { alerts = append(alerts, fmt.Sprintf("🔴 <b>Disk Almost Full:</b> %s", strings.Join(fullDisks, ", "))) }
			}

			if len(alerts) > 0 {
				token, chatID := db.GetSetting("tg_token", ""), db.GetSetting("tg_chat_id", "")
				if token != "" && chatID != "" {
					payload := map[string]interface{}{"chat_id": chatID, "text": "⚠️ <b>VPS ALERT</b> ⚠️\n\n" + strings.Join(alerts, "\n"), "parse_mode": "HTML"}
					jsonData, _ := json.Marshal(payload)
					http.Post(fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token), "application/json", bytes.NewBuffer(jsonData))
				}
				lastAlertTime = now
			}
		}
		time.Sleep(10 * time.Second)
	}
}