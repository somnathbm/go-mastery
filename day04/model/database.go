package model

import (
	"day04/health"
)

type Database struct {
	DBName      string
	Connections int
	DiskUsage   int
}

const (
	maxConnections     = 100
	warningConnections = 80
)

func (d Database) Severity() health.Severity {
	switch {
	case d.DiskUsage >= 90 &&
		d.Connections >= maxConnections:
		return health.SeverityCritical
	case d.DiskUsage >= 80 ||
		d.Connections >= warningConnections:
		return health.SeverityWarning
	default:
		return health.SeverityNormal
	}
}

// Determines if the service is healthy or unhealthy
func (d Database) Healthy() health.HealthStatus {
	severity := d.Severity()
	return health.HealthStatus{
		Healthy:  severity != health.SeverityCritical,
		Severity: severity,
		Reason:   getReason(severity),
	}
}

// func (d Database) Healthy() bool {
// 	return d.Severity() != "CRITICAL"
// }
