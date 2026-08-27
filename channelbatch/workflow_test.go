package channelbatch_test

import (
    "errors"
    "testing"
    "notification/channelpool"
    "notification/channelbatch"
)

func TestCase027Flow027(t *testing.T) {
    pool := channelpool.NewPool(2)
    batch := channelbatch.NewBatch(pool)
    succeeded, err := batch.Process([]error{nil, errors.New("case-027-rejected"), nil})
    if err != nil { t.Errorf("渠道批量探测不应耗尽会话资源，实际 %v", err) }
    if succeeded != 2 { t.Errorf("渠道批量探测成功数应为 2，实际 %d", succeeded) }
    if pool.Committed() != 2 { t.Errorf("渠道批量探测提交数应排除失败项，实际 %d", pool.Committed()) }
    if pool.Open() != 0 { t.Errorf("渠道批量探测结束后未释放会话，剩余 %d", pool.Open()) }
}
