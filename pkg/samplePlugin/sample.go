package samplePlugin

import (
	"context"
	"fmt"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	fwk "k8s.io/kube-scheduler/framework"
	framework "k8s.io/kubernetes/pkg/scheduler/framework"
	// "sigs.k8s.io/scheduler-plugins/apis/config"
	// "sigs.k8s.io/scheduler-plugins/apis/config/validation"
)

// 1. 플러그인 구조체 정의
type SamplePlugin struct {
	logger          klog.Logger
	handle          framework.Handle
	scheduleTimeout *time.Duration
	pgBackoff       *time.Duration
}

var _ framework.FilterPlugin = &SamplePlugin{}

const Name = "SamplePlugin"

func (s *SamplePlugin) Name() string {
	return Name
}

func (s *SamplePlugin) Filter(ctx context.Context, state fwk.CycleState, pod *v1.Pod, nodeInfo fwk.NodeInfo) *fwk.Status {
	node := nodeInfo.Node()

	if node == nil {
		return fwk.NewStatus(fwk.Error, "node not found")
	}

	// 예시: 노드에 "allow-custom=true" 레이블이 있어야만 스케줄링 허용
	if val, ok := node.Labels["allow-custom"]; !ok || val != "true" {
		return fwk.NewStatus(fwk.Unschedulable, fmt.Sprintf("Node %s is missing required label", node.Name))
	}

	return fwk.NewStatus(fwk.Success, "")
}

func (s *SamplePlugin) Score() {

}

// 3. 플러그인 생성자
func New(ctx context.Context, obj runtime.Object, h framework.Handle) (framework.Plugin, error) {
	lh := klog.FromContext(ctx).WithValues("plugin", Name)
	lh.V(5).Info("creating new Sample plugin")

	return &SamplePlugin{handle: h, logger: lh}, nil
}
