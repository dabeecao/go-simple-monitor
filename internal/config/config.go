package config

import (
	"os"
	"strings"
	"github.com/joho/godotenv"
)

var (
	AdminUser         = "admin"
	AdminPass         = "admin"
	JwtSecret         = []byte("super_secret_jwt_key_change_me_now_is_secure")
	JwtExpirationMins = 120
	Port              = "5000"
	TrustedProxies    []string
)

func Load() {
	godotenv.Load()
	if env := os.Getenv("ADMIN_USER"); env != "" { AdminUser = env }
	if env := os.Getenv("ADMIN_PASS"); env != "" { AdminPass = env }
	if env := os.Getenv("SECRET_TOKEN"); env != "" { JwtSecret = []byte(env) }
	if env := os.Getenv("PORT"); env != "" { Port = env }
	
	if env := os.Getenv("TRUSTED_PROXIES"); env != "" {
		proxies := strings.Split(env, ",")
		for _, p := range proxies {
			TrustedProxies = append(TrustedProxies, strings.TrimSpace(p))
		}
	}
}