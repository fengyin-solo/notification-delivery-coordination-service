package channelcatalog_test

import (
    "testing"
    "notification/channeldecoder"
    "notification/channelcatalog"
)

func TestCase022Flow022(t *testing.T) {
    decoder := channeldecoder.NewDecoder()
    cache := channelcatalog.NewCache()
    first := decoder.Decode("case-022-first")
    _ = decoder.Decode("case-022-later")
    cache.Put("case-022", first)
    got := cache.Get("case-022")
    if string(got) != "case-022-first" { t.Errorf("渠道参数缓存应保留首批内容 %q，实际 %q", "case-022-first", string(got)) }
    if len(got) > 0 { got[0] = 'X' }
    again := cache.Get("case-022")
    if string(again) != "case-022-first" { t.Errorf("渠道参数缓存读取方修改后缓存被污染为 %q", string(again)) }
}
