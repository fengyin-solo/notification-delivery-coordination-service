package subscriptionrefresh_test

import (
    "context"
    "errors"
    "testing"
    "time"
    "notification/subscriberprobe"
    "notification/subscriptionrefresh"
)

func TestCase002Flow002(t *testing.T) {
    probe := subscriberprobe.NewProbe(120 * time.Millisecond)
    runner := subscriptionrefresh.NewRunner(probe)
    ctx, cancel := context.WithCancel(context.Background())
    cancel()
    started := time.Now()
    err := runner.Execute(ctx)
    elapsed := time.Since(started)
    if !errors.Is(err, context.Canceled) { t.Errorf("订阅刷新取消后应返回 context.Canceled，实际 %v", err) }
    if elapsed > 60*time.Millisecond { t.Errorf("订阅刷新取消后仍等待下游，耗时 %s", elapsed) }
    if probe.Calls() != 1 { t.Errorf("订阅刷新取消后下游调用次数应为 1，实际 %d", probe.Calls()) }
}
