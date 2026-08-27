package broadcastdispatch_test

import (
    "errors"
    "testing"
    "notification/broadcasttransport"
    "notification/broadcastdispatch"
)

func TestCase018Flow018(t *testing.T) {
    gateway := broadcasttransport.NewGateway(broadcasttransport.ErrRejected, nil)
    coordinator := broadcastdispatch.NewCoordinator(gateway)
    err := coordinator.Run("case-018")
    if !errors.Is(err, broadcasttransport.ErrRejected) { t.Errorf("广播通知投递应保留拒绝错误，实际 %v", err) }
    if gateway.Calls() != 1 { t.Errorf("广播通知投递遇到拒绝后调用次数应为 1，实际 %d", gateway.Calls()) }
    if coordinator.Committed() != 0 { t.Errorf("广播通知投递被拒后提交数应为 0，实际 %d", coordinator.Committed()) }
}
