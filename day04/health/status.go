package health

type Severity string

const (
	SeverityCritical Severity = "CRITICAL"
	SeverityWarning  Severity = "WARNING"
	SeverityNormal   Severity = "NORMAL"
)

type HealthStatus struct {
	Healthy  bool
	Severity Severity
	Reason   string
}

type HealthChecker interface {
	Healthy() HealthStatus
}

type Restartable interface {
	Restart() error
}
