package schedulepool

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

// Close 归还执行会话：无论成功与否都释放占用（open--），
// 但只有成功的定时任务才计入提交（committed++）；被拒绝的任务不提交。
func (s *Session) Close(success bool) {
    if s.closed { return }
    s.closed = true
    if success {
        s.pool.committed++
    }
    s.pool.open--
}

func (p *Pool) Open() int { return p.open }
func (p *Pool) Committed() int { return p.committed }
