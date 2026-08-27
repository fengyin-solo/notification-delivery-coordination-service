package recipientcatalog_test

import (
    "testing"
    "notification/recipientdecoder"
    "notification/recipientcatalog"
)

func TestCase021Flow021(t *testing.T) {
    decoder := recipientdecoder.NewDecoder()
    cache := recipientcatalog.NewCache()
    first := decoder.Decode("case-021-first")
    _ = decoder.Decode("case-021-later")
    cache.Put("case-021", first)
    got := cache.Get("case-021")
    if string(got) != "case-021-first" { t.Errorf("接收地址缓存应保留首批内容 %q，实际 %q", "case-021-first", string(got)) }
    if len(got) > 0 { got[0] = 'X' }
    again := cache.Get("case-021")
    if string(again) != "case-021-first" { t.Errorf("接收地址缓存读取方修改后缓存被污染为 %q", string(again)) }
}
