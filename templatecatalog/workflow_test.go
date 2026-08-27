package templatecatalog_test

import (
    "testing"
    "notification/templatedecoder"
    "notification/templatecatalog"
)

func TestCase023Flow023(t *testing.T) {
    decoder := templatedecoder.NewDecoder()
    cache := templatecatalog.NewCache()
    first := decoder.Decode("case-023-first")
    _ = decoder.Decode("case-023-later")
    cache.Put("case-023", first)
    got := cache.Get("case-023")
    if string(got) != "case-023-first" { t.Errorf("模板变量缓存应保留首批内容 %q，实际 %q", "case-023-first", string(got)) }
    if len(got) > 0 { got[0] = 'X' }
    again := cache.Get("case-023")
    if string(again) != "case-023-first" { t.Errorf("模板变量缓存读取方修改后缓存被污染为 %q", string(again)) }
}
