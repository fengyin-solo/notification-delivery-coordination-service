package recipientpool

import "errors"

var ErrCapacity = errors.New("session capacity exhausted")

type Pool struct{ max, open, committed int }
type Session struct {
    pool   *Pool
    closed bool
}

func NewPool(max int) *Pool { return &Pool{max: max} }

func (p *Pool) Acquire() (*Session, error) {
    if p.open >= p.max {
        return nil, ErrCapacity
    }
    p.open++
    return &Session{pool: p}, nil
}

// Close 释放会话占用的容量。仅当 success 为 true（投递成功）时
// 才计入 committed；投递失败只释放容量，不计入成功统计。
func (s *Session) Close(success bool) {
    if s.closed {
        return
    }
    s.closed = true
    if success {
        s.pool.committed++
    }
    s.pool.open--
}

func (p *Pool) Open() int      { return p.open }
func (p *Pool) Committed() int { return p.committed }
