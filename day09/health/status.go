package health

import "time"

type Severity string

const (
	SeverityCritical Severity = "CRITICAL"
	SeverityWarning  Severity = "WARNING"
	SeverityNormal   Severity = "NORMAL"
)

type HealthStatus struct {
	Name         string
	Healthy      bool
	Severity     Severity
	Reason       string
	ResponseTime time.Duration
}

type HealthChecker interface {
	Name() string
	CheckHealth() (HealthStatus, error)
}

type Restartable interface {
	Restart() error
}
