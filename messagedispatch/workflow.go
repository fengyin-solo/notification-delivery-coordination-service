package messagedispatch

import (
    "context"
    "notification/deliveryprobe"
)

type Runner struct { probe *deliveryprobe.Probe }

func NewRunner(probe *deliveryprobe.Probe) *Runner { return &Runner{probe: probe} }

func (r *Runner) Execute(ctx context.Context) error {
    var last error
    for attempt := 0; attempt < 3; attempt++ {
        err := r.probe.Fetch(context.Background())
        if err == nil { return nil }
        last = err
    }
    return last
}
