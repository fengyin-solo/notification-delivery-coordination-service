package digestdispatch

import (
	"errors"

	"notification/digesttransport"
)

type Coordinator struct { gateway *digesttransport.Gateway; committed int }

func NewCoordinator(gateway *digesttransport.Gateway) *Coordinator { return &Coordinator{gateway: gateway} }

func (c *Coordinator) Run(key string) error {
	var last error
	for attempt := 0; attempt < 3; attempt++ {
		err := c.gateway.Send(key)
		if err == nil { c.committed++; return nil }
		// 退回是永久性失败：停止发送，不计入成功。
		if errors.Is(err, digesttransport.ErrRejected) { return err }
		last = err
	}
	return last
}

func (c *Coordinator) Committed() int { return c.committed }
