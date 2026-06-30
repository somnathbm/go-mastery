package contracts

type HealthChecker interface {
	Name() string
	Info() map[string]any
	Healthy() bool
	Severity() string
	Restart()
}
