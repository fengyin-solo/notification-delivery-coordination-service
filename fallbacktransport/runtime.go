package fallbacktransport

import (
	"errors"
	"fmt"
)

var ErrRejected = errors.New("request rejected")
var ErrTemporary = errors.New("temporary unavailable")

// Gateway 模拟备用渠道网关，按 outcomes 序列返回预设结果。
// 一旦遭遇拒绝（ErrRejected），后续相同 key 的发送必须继续返回拒绝，
// 而非因 outcomes 耗尽退化为 nil——拒绝是终态，回执需保留。
type Gateway struct {
	outcomes []error
	calls    int
	rejected map[string]bool
}

func NewGateway(outcomes ...error) *Gateway {
	return &Gateway{outcomes: outcomes, rejected: make(map[string]bool)}
}

func (g *Gateway) Send(key string) error {
	g.calls++
	// 拒绝是终态：已被拒绝的 key 不再退化为成功，保留拒绝回执。
	if g.rejected[key] {
		return fmt.Errorf("send %s: %w", key, ErrRejected)
	}
	index := g.calls - 1
	var err error
	if index < len(g.outcomes) {
		err = g.outcomes[index]
	} else {
		err = nil
	}
	if err != nil {
		if errors.Is(err, ErrRejected) {
			g.rejected[key] = true
		}
		return fmt.Errorf("send %s: %w", key, err)
	}
	return nil
}

func (g *Gateway) Calls() int { return g.calls }
