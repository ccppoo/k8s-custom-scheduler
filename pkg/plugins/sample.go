package plugins

import (
	"context"
	"fmt"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	fwk "k8s.io/kube-scheduler/framework"
	framework "k8s.io/kubernetes/pkg/scheduler/framework"
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

// PostFilter(ctx context.Context, state fwk.CycleState, pod *v1.Pod, filteredNodeStatusMap NodeToStatusReader) (*PostFilterResult, *fwk.Status)
// 2. Filter 로직 구현
// Filter(ctx context.Context, state fwk.CycleState, pod *v1.Pod, nodeInfo fwk.NodeInfo) *fwk.Status
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

// 3. 플러그인 생성자
func New(obj runtime.Object, h framework.Handle) (framework.Plugin, error) {
	return &SamplePlugin{handle: h}, nil
}
