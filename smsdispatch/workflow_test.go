package smsdispatch_test

import (
    "errors"
    "testing"
    "notification/smstransport"
    "notification/smsdispatch"
)

func TestCase013Flow013(t *testing.T) {
    gateway := smstransport.NewGateway(smstransport.ErrRejected, nil)
    coordinator := smsdispatch.NewCoordinator(gateway)
    err := coordinator.Run("case-013")
    if !errors.Is(err, smstransport.ErrRejected) { t.Errorf("短信渠道投递应保留拒绝错误，实际 %v", err) }
    if gateway.Calls() != 1 { t.Errorf("短信渠道投递遇到拒绝后调用次数应为 1，实际 %d", gateway.Calls()) }
    if coordinator.Committed() != 0 { t.Errorf("短信渠道投递被拒后提交数应为 0，实际 %d", coordinator.Committed()) }
}
