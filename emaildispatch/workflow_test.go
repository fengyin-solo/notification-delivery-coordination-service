package emaildispatch_test

import (
    "errors"
    "testing"
    "notification/emailtransport"
    "notification/emaildispatch"
)

func TestCase012Flow012(t *testing.T) {
    gateway := emailtransport.NewGateway(emailtransport.ErrRejected, nil)
    coordinator := emaildispatch.NewCoordinator(gateway)
    err := coordinator.Run("case-012")
    if !errors.Is(err, emailtransport.ErrRejected) { t.Errorf("邮件渠道投递应保留拒绝错误，实际 %v", err) }
    if gateway.Calls() != 1 { t.Errorf("邮件渠道投递遇到拒绝后调用次数应为 1，实际 %d", gateway.Calls()) }
    if coordinator.Committed() != 0 { t.Errorf("邮件渠道投递被拒后提交数应为 0，实际 %d", coordinator.Committed()) }
}
