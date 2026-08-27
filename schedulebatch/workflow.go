package schedulebatch

import "notification/schedulepool"

type Batch struct { pool *schedulepool.Pool }

func NewBatch(pool *schedulepool.Pool) *Batch { return &Batch{pool: pool} }

// Process 批量执行定时任务。outcomes 中 nil 表示成功可提交，非 nil 表示被拒绝的任务。
// 每个任务获取的执行会话在当轮迭代结束时立即归还，避免会话长期占用导致后续任务容量耗尽；
// 被拒绝的任务不会计入提交（committed），但仍会释放占用的会话。
func (b *Batch) Process(outcomes []error) (int, error) {
    succeeded := 0
    for _, outcome := range outcomes {
        session, err := b.pool.Acquire()
        if err != nil { return succeeded, err }
        success := outcome == nil
        if success { succeeded++ }
        // 立即归还执行会话，而非 defer 到函数返回，防止会话泄漏与提前触发容量耗尽。
        session.Close(success)
    }
    return succeeded, nil
}
