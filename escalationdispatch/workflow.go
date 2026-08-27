package escalationdispatch

import (
    "errors"

    "notification/escalationtransport"
)

type Coordinator struct { gateway *escalationtransport.Gateway; committed int }

func NewCoordinator(gateway *escalationtransport.Gateway) *Coordinator { return &Coordinator{gateway: gateway} }

func (c *Coordinator) Run(key string) error {
    var last error
    for attempt := 0; attempt < 3; attempt++ {
        err := c.gateway.Send(key)
        if err == nil { c.committed++; return nil }
        // 被接收端拒绝属于终态：立即结束升级通知，不再重试，也不计入已提交。
        if errors.Is(err, escalationtransport.ErrRejected) { return err }
        last = err
    }
    return last
}

func (c *Coordinator) Committed() int { return c.committed }
