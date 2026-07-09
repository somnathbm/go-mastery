package health

type Severity string

const (
	SeverityCritical Severity = "CRITICAL"
	SeverityWarning  Severity = "WARNING"
	SeverityNormal   Severity = "NORMAL"
)

type HealthStatus struct {
	Name     string
	Healthy  bool
	Severity Severity
	Reason   string
}

type HealthChecker interface {
	Name() string
	Healthy() HealthStatus
}

type Restartable interface {
	Restart() error
}
