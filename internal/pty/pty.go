package pty

import (
	"os/exec"
	"strconv"
	"strings"
	"net/http"

	"go-simple-monitor/internal/config"
	"github.com/creack/pty"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

func TerminalSocket(c *gin.Context) {
	tokenStr := c.Query("token")
	token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) { return config.JwtSecret, nil })

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil { return }
	defer ws.Close()

	if token == nil || !token.Valid {
		ws.WriteMessage(websocket.TextMessage, []byte("❌ Authentication Failed: Invalid or Expired Token.\r\n"))
		return
	}

	cmd := exec.Command("/bin/bash")
	ptmx, err := pty.Start(cmd)
	if err != nil { return }
	defer ptmx.Close()
	defer cmd.Process.Kill()

	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := ptmx.Read(buf)
			if err != nil { return }
			ws.WriteMessage(websocket.TextMessage, buf[:n])
		}
	}()

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil { break }
		msgStr := string(msg)
		if strings.HasPrefix(msgStr, "__RESIZE__:") {
			parts := strings.Split(strings.TrimPrefix(msgStr, "__RESIZE__:"), ",")
			if len(parts) == 2 {
				cols, _ := strconv.Atoi(parts[0])
				rows, _ := strconv.Atoi(parts[1])
				pty.Setsize(ptmx, &pty.Winsize{Rows: uint16(rows), Cols: uint16(cols)})
			}
		} else {
			ptmx.Write(msg)
		}
	}
}