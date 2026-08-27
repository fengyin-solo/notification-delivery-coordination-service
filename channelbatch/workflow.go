package channelbatch

import "notification/channelpool"

type Batch struct { pool *channelpool.Pool }

func NewBatch(pool *channelpool.Pool) *Batch { return &Batch{pool: pool} }

func (b *Batch) Process(outcomes []error) (int, error) {
    succeeded := 0
    for _, outcome := range outcomes {
        session, err := b.pool.Acquire()
        if err != nil { return succeeded, err }
        success := outcome == nil
        defer session.Close(success)
        if success { succeeded++ }
    }
    return succeeded, nil
}
