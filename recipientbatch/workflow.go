package recipientbatch

import "notification/recipientpool"

type Batch struct { pool *recipientpool.Pool }

func NewBatch(pool *recipientpool.Pool) *Batch { return &Batch{pool: pool} }

// Process 对每个接收人投递一次。每个接收人结束后立即释放会话，
// 只统计投递成功的接收人；单个接收人失败不影响后续接收人继续投递。
func (b *Batch) Process(outcomes []error) (int, error) {
    succeeded := 0
    for _, outcome := range outcomes {
        session, err := b.pool.Acquire()
        if err != nil {
            return succeeded, err
        }
        success := outcome == nil
        session.Close(success) // 立即释放，避免占用容量影响下一个接收人
        if success {
            succeeded++
        }
    }
    return succeeded, nil
}
