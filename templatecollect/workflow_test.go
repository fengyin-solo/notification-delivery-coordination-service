package templatecollect_test

import (
    "context"
    "errors"
    "testing"
    "time"
    "notification/templatestream"
    "notification/templatecollect"
)

func TestCase010Flow010(t *testing.T) {
    ctx, cancel := context.WithTimeout(context.Background(), 180*time.Millisecond)
    defer cancel()
    type result struct { items []string; err error }
    done := make(chan result, 1)
    go func() {
        items, err := templatecollect.Collect(ctx, []string{"case-010-one", "case-010-two", "case-010-three"}, 1)
        done <- result{items: items, err: err}
    }()
    select {
    case got := <-done:
        if !errors.Is(got.err, templatestream.ErrSource) { t.Errorf("模板流汇集应返回来源错误，实际 %v", got.err) }
        if len(got.items) != 1 { t.Errorf("模板流汇集应保留一条已收结果，实际 %d", len(got.items)) }
    case <-time.After(240 * time.Millisecond):
        t.Fatalf("模板流汇集来源失败后没有在超时边界内结束")
    }
}
