package retrybatch

import "notification/retrypool"

type Batch struct { pool *retrypool.Pool }

func NewBatch(pool *retrypool.Pool) *Batch { return &Batch{pool: pool} }

func (b *Batch) Process(outcomes []error) (int, error) {
    succeeded := 0
    for _, outcome := range outcomes {
        session, err := b.pool.Acquire()
        if err != nil { return succeeded, err }
        success := outcome == nil
        // 每条重试处理完立即释放资源，否则资源会一直占用到函数返回，
        // 导致后续消息因池容量耗尽拿不到资源。失败项不算完成。
        session.Close(success)
        if success { succeeded++ }
    }
    return succeeded, nil
}
