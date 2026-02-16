package main

import (
	"os"

	"github.com/ccppoo/k8s-custom-scheduler/pkg/samplePlugin"
	"k8s.io/component-base/cli"
	"k8s.io/component-base/logs"
	"k8s.io/kubernetes/cmd/kube-scheduler/app"
)

func main() {
	// 기본 스케줄러 커맨드에 커스텀 플러그인 주입
	command := app.NewSchedulerCommand(
		app.WithPlugin(samplePlugin.Name, samplePlugin.New),
	)

	logs.InitLogs()
	defer logs.FlushLogs()

	code := cli.Run(command)
	os.Exit(code)
}
