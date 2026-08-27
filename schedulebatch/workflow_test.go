package schedulebatch_test

import (
    "errors"
    "testing"
    "notification/schedulepool"
    "notification/schedulebatch"
)

func TestCase029Flow029(t *testing.T) {
    pool := schedulepool.NewPool(2)
    batch := schedulebatch.NewBatch(pool)
    succeeded, err := batch.Process([]error{nil, errors.New("case-029-rejected"), nil})
    if err != nil { t.Errorf("定时任务批量执行不应耗尽会话资源，实际 %v", err) }
    if succeeded != 2 { t.Errorf("定时任务批量执行成功数应为 2，实际 %d", succeeded) }
    if pool.Committed() != 2 { t.Errorf("定时任务批量执行提交数应排除失败项，实际 %d", pool.Committed()) }
    if pool.Open() != 0 { t.Errorf("定时任务批量执行结束后未释放会话，剩余 %d", pool.Open()) }
}
