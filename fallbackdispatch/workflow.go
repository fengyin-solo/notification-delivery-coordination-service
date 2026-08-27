package fallbackdispatch

import (
	"errors"

	"notification/fallbacktransport"
)

// maxAttempts 临时性失败的最大重试次数。
const maxAttempts = 3

// Coordinator 协调备用渠道的发送与重试。
type Coordinator struct {
	gateway   *fallbacktransport.Gateway
	committed int
}

func NewCoordinator(gateway *fallbacktransport.Gateway) *Coordinator {
	return &Coordinator{gateway: gateway}
}

// Run 按 key 发送并按错误类型决定是否重试。
// 拒绝（ErrRejected）是终态：立即停止、保留拒绝回执、不递增 committed。
// 临时失败（ErrTemporary）可重试，耗尽重试次数后返回最后一次回执。
func (c *Coordinator) Run(key string) error {
	var last error
	for attempt := 0; attempt < maxAttempts; attempt++ {
		err := c.gateway.Send(key)
		if err == nil {
			c.committed++
			return nil
		}
		last = err
		// 明确拒绝不可重试：回执需保留，提交数应为零。
		if errors.Is(err, fallbacktransport.ErrRejected) {
			return last
		}
	}
	return last
}

func (c *Coordinator) Committed() int { return c.committed }
