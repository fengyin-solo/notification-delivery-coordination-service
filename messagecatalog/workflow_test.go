package messagecatalog_test

import (
    "testing"
    "notification/messagedecoder"
    "notification/messagecatalog"
)

func TestCase025Flow025(t *testing.T) {
    decoder := messagedecoder.NewDecoder()
    cache := messagecatalog.NewCache()
    first := decoder.Decode("case-025-first")
    _ = decoder.Decode("case-025-later")
    cache.Put("case-025", first)
    got := cache.Get("case-025")
    if string(got) != "case-025-first" { t.Errorf("消息载荷缓存应保留首批内容 %q，实际 %q", "case-025-first", string(got)) }
    if len(got) > 0 { got[0] = 'X' }
    again := cache.Get("case-025")
    if string(again) != "case-025-first" { t.Errorf("消息载荷缓存读取方修改后缓存被污染为 %q", string(again)) }
}
