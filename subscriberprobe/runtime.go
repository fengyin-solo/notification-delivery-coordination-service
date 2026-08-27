package subscriberprobe

import (
    "context"
    "sync/atomic"
    "time"
)

type Probe struct { delay time.Duration; calls atomic.Int64 }

func NewProbe(delay time.Duration) *Probe { return &Probe{delay: delay} }

func (p *Probe) Fetch(ctx context.Context) error {
    p.calls.Add(1)
    time.Sleep(p.delay)
    return nil
}

func (p *Probe) Calls() int64 { return p.calls.Load() }
