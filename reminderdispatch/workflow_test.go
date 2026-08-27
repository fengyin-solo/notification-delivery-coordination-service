package reminderdispatch_test

import (
    "errors"
    "testing"
    "notification/remindertransport"
    "notification/reminderdispatch"
)

func TestCase019Flow019(t *testing.T) {
    gateway := remindertransport.NewGateway(remindertransport.ErrRejected, nil)
    coordinator := reminderdispatch.NewCoordinator(gateway)
    err := coordinator.Run("case-019")
    if !errors.Is(err, remindertransport.ErrRejected) { t.Errorf("提醒通知投递应保留拒绝错误，实际 %v", err) }
    if gateway.Calls() != 1 { t.Errorf("提醒通知投递遇到拒绝后调用次数应为 1，实际 %d", gateway.Calls()) }
    if coordinator.Committed() != 0 { t.Errorf("提醒通知投递被拒后提交数应为 0，实际 %d", coordinator.Committed()) }
}
