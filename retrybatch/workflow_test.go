package retrybatch_test

import (
    "errors"
    "testing"
    "notification/retrypool"
    "notification/retrybatch"
)

func TestCase030Flow030(t *testing.T) {
    pool := retrypool.NewPool(2)
    batch := retrybatch.NewBatch(pool)
    succeeded, err := batch.Process([]error{nil, errors.New("case-030-rejected"), nil})
    if err != nil { t.Errorf("失败消息批量重试不应耗尽会话资源，实际 %v", err) }
    if succeeded != 2 { t.Errorf("失败消息批量重试成功数应为 2，实际 %d", succeeded) }
    if pool.Committed() != 2 { t.Errorf("失败消息批量重试提交数应排除失败项，实际 %d", pool.Committed()) }
    if pool.Open() != 0 { t.Errorf("失败消息批量重试结束后未释放会话，剩余 %d", pool.Open()) }
}
