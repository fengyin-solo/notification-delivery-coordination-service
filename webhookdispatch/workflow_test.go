package webhookdispatch_test

import (
    "errors"
    "testing"
    "notification/webhooktransport"
    "notification/webhookdispatch"
)

func TestCase011Flow011(t *testing.T) {
    gateway := webhooktransport.NewGateway(webhooktransport.ErrRejected, nil)
    coordinator := webhookdispatch.NewCoordinator(gateway)
    err := coordinator.Run("case-011")
    if !errors.Is(err, webhooktransport.ErrRejected) { t.Errorf("Webhook 投递应保留拒绝错误，实际 %v", err) }
    if gateway.Calls() != 1 { t.Errorf("Webhook 投递遇到拒绝后调用次数应为 1，实际 %d", gateway.Calls()) }
    if coordinator.Committed() != 0 { t.Errorf("Webhook 投递被拒后提交数应为 0，实际 %d", coordinator.Committed()) }
}
