package db

import (
	"database/sql"
	"log"
	"strconv"
	"time"
	_ "modernc.org/sqlite"
)

var DB *sql.DB

func Init() {
	var err error
	DB, err = sql.Open("sqlite", "simple_monitor.db")
	if err != nil { log.Fatalf("DB Error: %v", err) }

	DB.Exec(`PRAGMA journal_mode=WAL`)
	DB.Exec(`CREATE TABLE IF NOT EXISTS settings (key TEXT PRIMARY KEY, value TEXT)`)
	DB.Exec(`CREATE TABLE IF NOT EXISTS stats_history (timestamp INTEGER, cpu REAL, ram REAL, disk REAL)`)
	DB.Exec(`CREATE INDEX IF NOT EXISTS idx_stats_timestamp ON stats_history(timestamp)`)

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

func SaveStats(cpu, ram, disk float64) {
	DB.Exec("INSERT INTO stats_history (timestamp, cpu, ram, disk) VALUES (?, ?, ?, ?)", time.Now().Unix(), cpu, ram, disk)
	// Tự động xóa dữ liệu cũ hơn 7 ngày để tránh đầy ổ cứng
	DB.Exec("DELETE FROM stats_history WHERE timestamp < ?", time.Now().Unix()-7*24*3600)
}

func GetStatsHistory(limit int) []map[string]interface{} {
	rows, err := DB.Query("SELECT timestamp, cpu, ram, disk FROM stats_history ORDER BY timestamp DESC LIMIT ?", limit)
	if err != nil { return nil }
	defer rows.Close()
	var res []map[string]interface{}
	for rows.Next() {
		var ts int64
		var c, r, d float64
		if err := rows.Scan(&ts, &c, &r, &d); err == nil {
			res = append(res, map[string]interface{}{"timestamp": ts, "cpu": c, "ram": r, "disk": d})
		}
	}
	return res
}