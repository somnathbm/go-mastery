package model

import (
	"day08/health"
	"day08/utils"
)

type Service struct {
	SName        string `json:"name"`
	CPU          int    `json:"cpu"`
	Memory       int    `json:"memory"`
	RestartCount int    `json:"restartcount"`
}

func (s Service) Name() string {
	return s.SName
}

// Determines the severity of a service in terms of CPU or RAM usage
func (s Service) Severity() health.Severity {
	switch {
	case s.CPU >= 90 || s.Memory >= 90:
		return health.SeverityCritical
	case s.CPU >= 80 || s.Memory >= 80:
		return health.SeverityWarning
	default:
		return health.SeverityNormal
	}
}

// Determines if the service is healthy or unhealthy
func (s Service) Healthy() health.HealthStatus {

	severity := s.Severity()
	return health.HealthStatus{
		Name:     s.Name(),
		Healthy:  severity != health.SeverityCritical,
		Severity: severity,
		Reason:   utils.GetReason(severity),
	}
}

// Restart service
// func (s *Service) Restart() error {
// 	s.RestartCount++
// 	s.CPU = 0
// 	s.Memory = 0
// 	return nil
// }
