package fallbackdispatch_test

import (
    "errors"
    "testing"
    "notification/fallbacktransport"
    "notification/fallbackdispatch"
)

func TestCase016Flow016(t *testing.T) {
    gateway := fallbacktransport.NewGateway(fallbacktransport.ErrRejected, nil)
    coordinator := fallbackdispatch.NewCoordinator(gateway)
    err := coordinator.Run("case-016")
    if !errors.Is(err, fallbacktransport.ErrRejected) { t.Errorf("备用渠道投递应保留拒绝错误，实际 %v", err) }
    if gateway.Calls() != 1 { t.Errorf("备用渠道投递遇到拒绝后调用次数应为 1，实际 %d", gateway.Calls()) }
    if coordinator.Committed() != 0 { t.Errorf("备用渠道投递被拒后提交数应为 0，实际 %d", coordinator.Committed()) }
}
