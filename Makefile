# 변수 설정
BINARY_NAME=myapp

build-linux:
	$env:GOOS="linux"; $env:GOARCH="amd64"; go build -o bin/$(BINARY_NAME)-linux main.go

build-windows:
	$env:GOOS="windows"; $env:GOARCH="amd64"; go build -o bin/$(BINARY_NAME).exe main.go

build-mac:
	$env:GOOS="darwin"; $env:GOARCH="arm64"; go build -o bin/$(BINARY_NAME)-mac main.go

build-all: build-linux build-windows build-mac