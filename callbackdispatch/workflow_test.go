package callbackdispatch_test

import (
    "errors"
    "testing"
    "notification/callbacktransport"
    "notification/callbackdispatch"
)

func TestCase015Flow015(t *testing.T) {
    gateway := callbacktransport.NewGateway(callbacktransport.ErrRejected, nil)
    coordinator := callbackdispatch.NewCoordinator(gateway)
    err := coordinator.Run("case-015")
    if !errors.Is(err, callbacktransport.ErrRejected) { t.Errorf("回调通知投递应保留拒绝错误，实际 %v", err) }
    if gateway.Calls() != 1 { t.Errorf("回调通知投递遇到拒绝后调用次数应为 1，实际 %d", gateway.Calls()) }
    if coordinator.Committed() != 0 { t.Errorf("回调通知投递被拒后提交数应为 0，实际 %d", coordinator.Committed()) }
}
