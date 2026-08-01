package model

import (
	"day11/health"
	"day11/utils"
	"errors"
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
func (s Service) CheckHealth() (health.HealthStatus, error) {
	severity := s.Severity()

	// simulate a health check failure
	if s.Name() == "order-service" {
		return health.HealthStatus{
			Name:     s.Name(),
			Healthy:  false,
			Severity: severity,
		}, errors.New("order-service transient failure")
	}

	return health.HealthStatus{
		Name:     s.Name(),
		Healthy:  severity != health.SeverityCritical,
		Severity: severity,
		Reason:   utils.GetReason(severity),
	}, nil
}

// Restart service
// func (s *Service) Restart() error {
// 	s.RestartCount++
// 	s.CPU = 0
// 	s.Memory = 0
// 	return nil
// }
