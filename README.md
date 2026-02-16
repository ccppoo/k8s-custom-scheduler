# k8s-custom-scheduler

[kubernetes-sigs/scheduler-plugins](https://github.com/kubernetes-sigs/scheduler-plugins) 이거 기반으로 스케줄러 공부중

# build

```bash
go build -o /bin ./cmd/scheduler/main.go

$env:GOOS="windows"; $env:GOARCH="amd64"; go build -o bin/$(BINARY_NAME).exe main.go
$env:GOOS="linux"; $env:GOARCH="amd64"; go build -o bin/$(BINARY_NAME)-linux main.go
$env:GOOS="linux"; $env:GOARCH="amd64"; go build -o bin/run ./cmd/scheduler/main.go
```