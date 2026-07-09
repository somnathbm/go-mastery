package model

import (
	"day08/health"
	"day08/utils"
	"time"
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

func (d Database) Name() string {
	return d.DBName
}

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
	time.Sleep(5 * time.Second)

	severity := d.Severity()
	return health.HealthStatus{
		Name:     d.Name(),
		Healthy:  severity != health.SeverityCritical,
		Severity: severity,
		Reason:   utils.GetReason(severity),
	}
}
