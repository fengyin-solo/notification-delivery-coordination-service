package templatepreview_test

import (
    "context"
    "errors"
    "testing"
    "time"
    "notification/renderprobe"
    "notification/templatepreview"
)

func TestCase005Flow005(t *testing.T) {
    probe := renderprobe.NewProbe(120 * time.Millisecond)
    runner := templatepreview.NewRunner(probe)
    ctx, cancel := context.WithCancel(context.Background())
    cancel()
    started := time.Now()
    err := runner.Execute(ctx)
    elapsed := time.Since(started)
    if !errors.Is(err, context.Canceled) { t.Errorf("模板预览取消后应返回 context.Canceled，实际 %v", err) }
    if elapsed > 60*time.Millisecond { t.Errorf("模板预览取消后仍等待下游，耗时 %s", elapsed) }
    if probe.Calls() != 1 { t.Errorf("模板预览取消后下游调用次数应为 1，实际 %d", probe.Calls()) }
}
