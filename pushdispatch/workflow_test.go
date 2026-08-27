package pushdispatch_test

import (
    "errors"
    "testing"
    "notification/pushtransport"
    "notification/pushdispatch"
)

func TestCase014Flow014(t *testing.T) {
    gateway := pushtransport.NewGateway(pushtransport.ErrRejected, nil)
    coordinator := pushdispatch.NewCoordinator(gateway)
    err := coordinator.Run("case-014")
    if !errors.Is(err, pushtransport.ErrRejected) { t.Errorf("应用推送投递应保留拒绝错误，实际 %v", err) }
    if gateway.Calls() != 1 { t.Errorf("应用推送投递遇到拒绝后调用次数应为 1，实际 %d", gateway.Calls()) }
    if coordinator.Committed() != 0 { t.Errorf("应用推送投递被拒后提交数应为 0，实际 %d", coordinator.Committed()) }
}
