package escalationdispatch

import "notification/escalationtransport"

type Coordinator struct { gateway *escalationtransport.Gateway; committed int }

func NewCoordinator(gateway *escalationtransport.Gateway) *Coordinator { return &Coordinator{gateway: gateway} }

func (c *Coordinator) Run(key string) error {
    var last error
    for attempt := 0; attempt < 3; attempt++ {
        err := c.gateway.Send(key)
        if err == nil { c.committed++; return nil }
        last = err
    }
    return last
}

func (c *Coordinator) Committed() int { return c.committed }
