package model

type Service struct {
	Name         string `json:"name"`
	CPU          int    `json:"cpu"`
	Memory       int    `json:"memory"`
	RestartCount int    `json:"restartcount"`
}

// Determines the severity of a service in terms of CPU or RAM usage
func (s Service) Severity() string {
	switch {
	case s.CPU >= 90 || s.Memory >= 90:
		return "CRITICAL"
	case s.CPU >= 80 || s.Memory >= 80:
		return "WARNING"
	default:
		return "NORMAL"
	}
}

// Determines if the service is healthy or unhealthy
func (s Service) Healthy() bool {
	return s.Severity() != "CRITICAL"
}

// Restart service
func (s *Service) Restart() {
	s.RestartCount++
	s.CPU = 0
	s.Memory = 0
}
