package templatebatch_test

import (
    "errors"
    "testing"
    "notification/templatepool"
    "notification/templatebatch"
)

func TestCase028Flow028(t *testing.T) {
    pool := templatepool.NewPool(2)
    batch := templatebatch.NewBatch(pool)
    succeeded, err := batch.Process([]error{nil, errors.New("case-028-rejected"), nil})
    if err != nil { t.Errorf("模板批量渲染不应耗尽会话资源，实际 %v", err) }
    if succeeded != 2 { t.Errorf("模板批量渲染成功数应为 2，实际 %d", succeeded) }
    if pool.Committed() != 2 { t.Errorf("模板批量渲染提交数应排除失败项，实际 %d", pool.Committed()) }
    if pool.Open() != 0 { t.Errorf("模板批量渲染结束后未释放会话，剩余 %d", pool.Open()) }
}
