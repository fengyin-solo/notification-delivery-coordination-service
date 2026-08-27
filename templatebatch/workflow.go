package templatebatch

import "notification/templatepool"

type Batch struct { pool *templatepool.Pool }

func NewBatch(pool *templatepool.Pool) *Batch { return &Batch{pool: pool} }

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
