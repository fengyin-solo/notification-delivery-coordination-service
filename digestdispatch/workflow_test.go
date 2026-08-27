package digestdispatch_test

import (
    "errors"
    "testing"
    "notification/digesttransport"
    "notification/digestdispatch"
)

func TestCase017Flow017(t *testing.T) {
    gateway := digesttransport.NewGateway(digesttransport.ErrRejected, nil)
    coordinator := digestdispatch.NewCoordinator(gateway)
    err := coordinator.Run("case-017")
    if !errors.Is(err, digesttransport.ErrRejected) { t.Errorf("摘要通知投递应保留拒绝错误，实际 %v", err) }
    if gateway.Calls() != 1 { t.Errorf("摘要通知投递遇到拒绝后调用次数应为 1，实际 %d", gateway.Calls()) }
    if coordinator.Committed() != 0 { t.Errorf("摘要通知投递被拒后提交数应为 0，实际 %d", coordinator.Committed()) }
}
