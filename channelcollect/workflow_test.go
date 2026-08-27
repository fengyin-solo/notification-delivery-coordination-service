package channelcollect_test

import (
    "context"
    "errors"
    "testing"
    "time"
    "notification/channelstream"
    "notification/channelcollect"
)

func TestCase007Flow007(t *testing.T) {
    ctx, cancel := context.WithTimeout(context.Background(), 180*time.Millisecond)
    defer cancel()
    type result struct { items []string; err error }
    done := make(chan result, 1)
    go func() {
        items, err := channelcollect.Collect(ctx, []string{"case-007-one", "case-007-two", "case-007-three"}, 1)
        done <- result{items: items, err: err}
    }()
    select {
    case got := <-done:
        if !errors.Is(got.err, channelstream.ErrSource) { t.Errorf("渠道流汇集应返回来源错误，实际 %v", got.err) }
        if len(got.items) != 1 { t.Errorf("渠道流汇集应保留一条已收结果，实际 %d", len(got.items)) }
    case <-time.After(240 * time.Millisecond):
        t.Fatalf("渠道流汇集来源失败后没有在超时边界内结束")
    }
}
