package main

import (
	"fmt"
	"html/template"
	"io/fs"
	"net/http"

	"go-simple-monitor/internal/config"
	"go-simple-monitor/internal/db"
	"go-simple-monitor/internal/handlers"
	"go-simple-monitor/internal/monitor"
	"go-simple-monitor/web"

	"github.com/gin-gonic/gin"
)

// 👇 thêm dòng này
var version = "dev"

func main() {
	fmt.Println("Starting go-simple-monitor version:", version)

	// 1. Tải cấu hình (.env)
	config.Load()

	// 2. Mở kết nối Database
	db.Init()

	// 3. Chạy Background Monitors
	monitor.Start()

	// 4. Cấu hình Web Server (Gin)
	r := gin.Default()
	r.SetTrustedProxies(config.TrustedProxies)

	templ := template.Must(template.ParseFS(web.FS, "templates/*"))
	r.SetHTMLTemplate(templ)

	staticFS, err := fs.Sub(web.FS, "static")
	if err != nil {
		panic(err)
	}

	r.StaticFS("/static", http.FS(staticFS))

	// 5. Đăng ký các API Routes
	handlers.RegisterRoutes(r)

	// 6. Chạy Server
	r.Run("0.0.0.0:" + config.Port)
}