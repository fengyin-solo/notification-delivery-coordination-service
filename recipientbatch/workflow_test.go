package recipientbatch_test

import (
    "errors"
    "testing"
    "notification/recipientpool"
    "notification/recipientbatch"
)

func TestCase026Flow026(t *testing.T) {
    pool := recipientpool.NewPool(2)
    batch := recipientbatch.NewBatch(pool)
    succeeded, err := batch.Process([]error{nil, errors.New("case-026-rejected"), nil})
    if err != nil { t.Errorf("接收人批量投递不应耗尽会话资源，实际 %v", err) }
    if succeeded != 2 { t.Errorf("接收人批量投递成功数应为 2，实际 %d", succeeded) }
    if pool.Committed() != 2 { t.Errorf("接收人批量投递提交数应排除失败项，实际 %d", pool.Committed()) }
    if pool.Open() != 0 { t.Errorf("接收人批量投递结束后未释放会话，剩余 %d", pool.Open()) }
}
