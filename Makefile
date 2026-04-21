run:
	go run cmd/server/main.go

build:
	go build -o bin/go-simple-monitor cmd/server/main.go

tidy:
	go mod tidy