package smstransport

import (
    "errors"
    "fmt"
)

var ErrRejected = errors.New("request rejected")
var ErrTemporary = errors.New("temporary unavailable")

type Gateway struct { outcomes []error; calls int }

func NewGateway(outcomes ...error) *Gateway { return &Gateway{outcomes: outcomes} }

func (g *Gateway) Send(key string) error {
    g.calls++
    index := g.calls - 1
    if index >= len(g.outcomes) { return nil }
    if err := g.outcomes[index]; err != nil { return fmt.Errorf("send %s: %v", key, err) }
    return nil
}

func (g *Gateway) Calls() int { return g.calls }
