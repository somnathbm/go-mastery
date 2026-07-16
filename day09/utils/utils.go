package utils

import (
	"day09/health"
)

func GetReason(s health.Severity) string {
	switch s {
	case health.SeverityCritical:
		return "CPU or memory exceeded critical threshold"
	case health.SeverityWarning:
		return "CPU or memory exceeded warning threshold"
	default:
		return "Healthy"
	}
}
