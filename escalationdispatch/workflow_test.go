package escalationdispatch_test

import (
    "errors"
    "testing"
    "notification/escalationtransport"
    "notification/escalationdispatch"
)

func TestCase020Flow020(t *testing.T) {
    gateway := escalationtransport.NewGateway(escalationtransport.ErrRejected, nil)
    coordinator := escalationdispatch.NewCoordinator(gateway)
    err := coordinator.Run("case-020")
    if !errors.Is(err, escalationtransport.ErrRejected) { t.Errorf("升级通知投递应保留拒绝错误，实际 %v", err) }
    if gateway.Calls() != 1 { t.Errorf("升级通知投递遇到拒绝后调用次数应为 1，实际 %d", gateway.Calls()) }
    if coordinator.Committed() != 0 { t.Errorf("升级通知投递被拒后提交数应为 0，实际 %d", coordinator.Committed()) }
}
