# 🚀 Go Simple Monitor

**Go Simple Monitor** là một công cụ giám sát hệ thống Linux siêu nhẹ, hiệu suất cao được viết bằng **Golang**.  
Nhờ khả năng build thành **single binary**, bạn có thể triển khai cực kỳ nhanh chóng chỉ với một file duy nhất.

---

## ✨ Tính năng nổi bật

- 📊 **Giám sát thời gian thực**  
  Theo dõi CPU, RAM, Disk, Swap và tốc độ mạng (Tx/Rx) chi tiết theo từng tiến trình.

- 💻 **Web Terminal**  
  Truy cập Terminal trực tiếp trên trình duyệt qua WebSocket, hỗ trợ resize tự động.

- 🔔 **Cảnh báo Telegram**  
  Tự động gửi thông báo khi tài nguyên vượt ngưỡng cấu hình.

- ⚙️ **Quản lý tiến trình & cổng**  
  Xem tiến trình, kiểm tra port đang listen và kill tiến trình ngốn tài nguyên.

- ⏳ **Quản lý Cronjob**  
  Thêm / sửa / xóa tác vụ cron trực quan.

- 📦 **Single Binary**  
  Deploy nhanh chóng chỉ với 1 file.

- 🛡️ **Bảo mật**  
  - JWT Authentication  
  - Rate Limit  
  - Trusted Proxies

---

## ⚙️ Yêu cầu hệ thống

- **Hệ điều hành:** Linux (Ubuntu, Debian, CentOS, Fedora...)
- **Phụ thuộc bắt buộc:**

### Ubuntu / Debian

```bash
sudo apt update && sudo apt install nethogs -y
```
### CentOS / RHEL
```bash
sudo yum install nethogs -y
````

---

## 🚀 Cài đặt & triển khai

### ⚡ Cách 1: Dùng bản build sẵn (khuyên dùng)

1. Truy cập **Releases** trên GitHub
2. Tải file phù hợp (ví dụ: `go-simple-monitor-linux-amd64`)
3. Cấp quyền và chạy:

```bash
chmod +x go-simple-monitor
./go-simple-monitor
```

---

### 🛠️ Cách 2: Tự build từ source

> ⚠️ **Yêu cầu:** Phải cài **Go 1.20+** trước khi build. Xem cách cài tại: [https://go.dev/dl/](https://go.dev/dl/)

#### 1. Clone project

```bash
git clone https://github.com/dabeecao/go-simple-monitor.git
cd go-simple-monitor
```

#### 2. Sử dụng Makefile


#### Cài dependencies
```bash
make tidy # (tương đương: go mod tidy)
```

#### Chạy thử (dev)
```bash
make run # (tương đương: go run cmd/server/main.go)
```

#### Build binary
```bash
make build # (tương đương: go build -o bin/go-simple-monitor cmd/server/main.go)
```

👉 File sau khi build nằm trong thư mục:

```
bin/go-simple-monitor
```

---

## 🛠 Cấu hình (.env)

Đặt file `env.example` cùng cấp với binary và đổi tên thành .env:

```env
# Tài khoản quản trị
ADMIN_USER=admin
ADMIN_PASS=mat_khau_cua_ban

# JWT Secret (nên đổi)
SECRET_TOKEN=thay_doi_ma_nay_de_bao_mat_he_thong

# Port server
PORT=5000

# TRUSTED PROXIES
# Để trống (hoặc không khai báo) = Không tin tưởng proxy nào, lấy IP trực tiếp (an toàn nhất)
# Nếu dùng Nginx làm reverse proxy, điền: 127.0.0.1
# Nếu có nhiều proxy, cách nhau bằng dấu phẩy: 127.0.0.1,192.168.1.100
TRUSTED_PROXIES=
```

### ⚠️ Mặc định (nếu không có .env)

* Username: `admin`
* Password: `admin`
* Port: `5000`

---

## 🛡️ Chạy nền với Systemd (khuyên dùng)

### 1. Tạo service

```bash
sudo nano /etc/systemd/system/gosmonitor.service
```

### 2. Nội dung service

```ini
[Unit]
Description=Go Simple Monitor Service
After=network.target

[Service]
Type=simple
User=root
WorkingDirectory=/opt/go-simple-monitor
ExecStart=/opt/go-simple-monitor/go-simple-monitor
Restart=on-failure
RestartSec=5

[Install]
WantedBy=multi-user.target
```

> 🔁 Nhớ chỉnh lại đường dẫn `/opt/go-simple-monitor` cho đúng thực tế

---

### 3. Kích hoạt

```bash
sudo systemctl daemon-reload
sudo systemctl enable gosmonitor
sudo systemctl start gosmonitor
```

---

## 📝 License

Dự án sử dụng **MIT License**