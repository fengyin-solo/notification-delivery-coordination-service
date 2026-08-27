package recipientadmission

import "notification/recipientfilter"

type Coordinator struct { gate recipientfilter.Gate; labels map[string]string }

func NewCoordinator(enabled bool) *Coordinator {
    return &Coordinator{gate: recipientfilter.NewGate(enabled)}
}

func (c *Coordinator) Apply(key string) error {
    if c.gate != nil {
        if err := c.gate.Allow(key); err != nil { return err }
    }
    c.labels[key] = "active"
    return nil
}

func (c *Coordinator) Label(key string) string { return c.labels[key] }
