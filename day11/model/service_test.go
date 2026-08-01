package model

import (
	"day11/health"
	"testing"
)

func TestServiceSeverity(t *testing.T) {
	tests := []struct {
		name         string
		cpu, memory  int
		wantSeverity health.Severity
	}{
		{"auth-service", 80, 90, health.SeverityCritical},
		{"payment-service", 30, 20, health.SeverityNormal},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			myService := Service{SName: tc.name, CPU: tc.cpu, Memory: tc.memory}
			gotSeverity := myService.Severity()

			if gotSeverity != tc.wantSeverity {
				t.Errorf("TestServiceSeverity expects: %v, but received: %v", tc.wantSeverity, gotSeverity)
			}
		})
	}
}

func TestServiceHealthy(t *testing.T) {
	tests := []struct {
		name        string
		cpu, memory int
		severity    health.Severity
		isHealthy   bool
		reason      string
	}{
		{"auth-service", 80, 92, health.SeverityCritical, false, "CPU or memory exceeded critical threshold"},
		{"payment-service", 30, 22, health.SeverityNormal, true, "Healthy"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			myService := Service{SName: tc.name, CPU: tc.cpu, Memory: tc.memory}
			gotHealthStatus := myService.Healthy()

			if gotHealthStatus.Healthy != tc.isHealthy || gotHealthStatus.Reason != tc.reason {
				t.Errorf("TestServiceHealthy wanted %v, got %v", tc.isHealthy, gotHealthStatus.Healthy)
			}
		})
	}
}
