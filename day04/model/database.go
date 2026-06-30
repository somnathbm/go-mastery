package model

type Database struct {
	DBName       string
	Connections  int
	DiskUsage    int
	RestartCount int
}

const (
	maxConnections     = 100
	warningConnections = 80
)

func (d Database) Name() string {
	return d.DBName
}

func (d Database) Info() map[string]any {
	return map[string]any{
		"DBName":    d.Name(),
		"DiskUsage": d.DiskUsage,
	}
}

func (d Database) Severity() string {
	switch {
	case d.DiskUsage >= 90 &&
		d.Connections >= maxConnections:
		return "CRITICAL"
	case d.DiskUsage >= 80 ||
		d.Connections >= warningConnections:
		return "WARNING"
	default:
		return "NORMAL"
	}
}

func (d Database) Healthy() bool {
	return d.Severity() != "CRITICAL"
}

func (d *Database) Restart() {
	d.RestartCount++
	d.DiskUsage = 0
}
