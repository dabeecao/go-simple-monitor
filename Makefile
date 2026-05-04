run:
	go run cmd/server/main.go

build:
	go build -o bin/go-simple-monitor cmd/server/main.go

build-linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o bin/go-simple-monitor-linux-amd64 cmd/server/main.go

build-arm:
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -ldflags="-s -w" -o bin/go-simple-monitor-linux-arm64 cmd/server/main.go

tidy:
	go mod tidy