package topicfilter

import "errors"

type Gate interface { Allow(string) error }
type ruleGate struct { blocked string }

func NewGate(enabled bool) Gate {
    if !enabled { var gate *ruleGate; return gate }
    return &ruleGate{blocked: "blocked"}
}

func (g *ruleGate) Allow(value string) error {
    if g.blocked == value { return errors.New("value blocked") }
    return nil
}
