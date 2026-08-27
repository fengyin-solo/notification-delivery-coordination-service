package retrypool

import "errors"

var ErrCapacity = errors.New("session capacity exhausted")

type Pool struct { max, open, committed int }
type Session struct { pool *Pool; closed bool }

func NewPool(max int) *Pool { return &Pool{max: max} }

func (p *Pool) Acquire() (*Session, error) {
    if p.open >= p.max { return nil, ErrCapacity }
    p.open++
    return &Session{pool: p}, nil
}

func (s *Session) Close(success bool) {
    if s.closed { return }
    s.closed = true
    // 释放资源对所有结果都生效；但只有成功项才计入完成数，
    // 失败的重试必须被排除，避免无效重试被算作完成。
    s.pool.open--
    if success {
        s.pool.committed++
    }
}

func (p *Pool) Open() int { return p.open }
func (p *Pool) Committed() int { return p.committed }
