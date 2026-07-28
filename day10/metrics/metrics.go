package metrics

import (
	"day10/health"
	"sync"
)

// Metrics
type Metrics struct {
	mu sync.Mutex

	Healthy  int
	Warning  int
	Critical int
}

type Snapshot struct {
	Healthy  int
	Warning  int
	Critical int
}

func (m *Metrics) Snapshot() Snapshot {
	m.mu.Lock()
	defer m.mu.Unlock()
	return Snapshot{
		Healthy:  m.Healthy,
		Warning:  m.Warning,
		Critical: m.Critical,
	}
}

func (m *Metrics) Increment(severity health.Severity) {
	m.mu.Lock()
	defer m.mu.Unlock()

	switch severity {
	case health.SeverityNormal:
		m.Healthy++
	case health.SeverityWarning:
		m.Warning++
	case health.SeverityCritical:
		m.Critical++
	}
}
