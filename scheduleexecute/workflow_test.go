package scheduleexecute_test

import (
    "context"
    "errors"
    "testing"
    "time"
    "notification/scheduleprobe"
    "notification/scheduleexecute"
)

func TestCase003Flow003(t *testing.T) {
    probe := scheduleprobe.NewProbe(120 * time.Millisecond)
    runner := scheduleexecute.NewRunner(probe)
    ctx, cancel := context.WithCancel(context.Background())
    cancel()
    started := time.Now()
    err := runner.Execute(ctx)
    elapsed := time.Since(started)
    if !errors.Is(err, context.Canceled) { t.Errorf("定时发送执行取消后应返回 context.Canceled，实际 %v", err) }
    if elapsed > 60*time.Millisecond { t.Errorf("定时发送执行取消后仍等待下游，耗时 %s", elapsed) }
    if probe.Calls() != 1 { t.Errorf("定时发送执行取消后下游调用次数应为 1，实际 %d", probe.Calls()) }
}
