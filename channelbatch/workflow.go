package channelbatch

import "notification/channelpool"

type Batch struct { pool *channelpool.Pool }

func NewBatch(pool *channelpool.Pool) *Batch { return &Batch{pool: pool} }

func (b *Batch) Process(outcomes []error) (int, error) {
    succeeded := 0
    for _, outcome := range outcomes {
        session, err := b.pool.Acquire()
        if err != nil {
            // 单个渠道探测失败：跳过该项继续后续渠道，不占用资源、不提交。
            continue
        }
        success := outcome == nil
        // 单个渠道结束立即归还资源，避免后续渠道误报资源耗尽；
        // 失败项在 Close 中不会计入提交数。
        session.Close(success)
        if success {
            succeeded++
        }
    }
    return succeeded, nil
}
