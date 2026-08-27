package broadcastdispatch

import (
    "errors"
    "notification/broadcasttransport"
)

type Coordinator struct { gateway *broadcasttransport.Gateway; committed int }

func NewCoordinator(gateway *broadcasttransport.Gateway) *Coordinator { return &Coordinator{gateway: gateway} }

func (c *Coordinator) Run(key string) error {
    var last error
    for attempt := 0; attempt < 3; attempt++ {
        err := c.gateway.Send(key)
        if err == nil { c.committed++; return nil }
        // A rejection is terminal: never retry, never leave a committed record.
        if errors.Is(err, broadcasttransport.ErrRejected) { return err }
        last = err
    }
    return last
}

func (c *Coordinator) Committed() int { return c.committed }
