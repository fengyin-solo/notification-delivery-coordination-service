package topiccatalog_test

import (
    "testing"
    "notification/topicdecoder"
    "notification/topiccatalog"
)

func TestCase024Flow024(t *testing.T) {
    decoder := topicdecoder.NewDecoder()
    cache := topiccatalog.NewCache()
    first := decoder.Decode("case-024-first")
    _ = decoder.Decode("case-024-later")
    cache.Put("case-024", first)
    got := cache.Get("case-024")
    if string(got) != "case-024-first" { t.Errorf("主题标识缓存应保留首批内容 %q，实际 %q", "case-024-first", string(got)) }
    if len(got) > 0 { got[0] = 'X' }
    again := cache.Get("case-024")
    if string(again) != "case-024-first" { t.Errorf("主题标识缓存读取方修改后缓存被污染为 %q", string(again)) }
}
