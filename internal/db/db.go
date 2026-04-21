package db

import (
	"database/sql"
	"log"
	"strconv"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Init() {
	var err error
	DB, err = sql.Open("sqlite3", "simple_monitor.db")
	if err != nil { log.Fatalf("DB Error: %v", err) }

	DB.Exec(`CREATE TABLE IF NOT EXISTS settings (key TEXT PRIMARY KEY, value TEXT)`)

	defaults := map[string]string{
		"tg_token": "", "tg_chat_id": "",
		"tg_cpu_enabled": "1", "tg_cpu_threshold": "90",
		"tg_ram_enabled": "1", "tg_ram_threshold": "90",
		"tg_disk_enabled": "1", "tg_disk_threshold": "90",
	}
	for k, v := range defaults {
		DB.Exec("INSERT OR IGNORE INTO settings (key, value) VALUES (?, ?)", k, v)
	}
}

func GetSetting(key string, def string) string {
	var value string
	if err := DB.QueryRow("SELECT value FROM settings WHERE key=?", key).Scan(&value); err != nil { return def }
	return value
}

func SetSetting(key string, value string) {
	DB.Exec("INSERT OR REPLACE INTO settings (key, value) VALUES (?, ?)", key, value)
}

func ParseFloatSetting(key string, def float64) float64 {
	valStr := GetSetting(key, "")
	val, err := strconv.ParseFloat(valStr, 64)
	if err != nil { return def }
	return val
}