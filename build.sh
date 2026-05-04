#!/bin/bash

mkdir -p bin

APP_NAME="go-simple-monitor"
MAIN_PATH="cmd/server/main.go"
LDFLAGS="-s -w"

echo "Bắt đầu biên dịch..."

# Linux AMD64
echo "Building Linux AMD64..."
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="$LDFLAGS" -o bin/$APP_NAME-linux-amd64 $MAIN_PATH

# Linux ARM64
echo "Building Linux ARM64..."
GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -ldflags="$LDFLAGS" -o bin/$APP_NAME-linux-arm64 $MAIN_PATH

# Linux ARMv7
echo "Building Linux ARMv7..."
GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=0 go build -ldflags="$LDFLAGS" -o bin/$APP_NAME-linux-armv7 $MAIN_PATH

echo "Đã xong! Các file nằm trong thư mục bin/"
chmod +x bin/*
