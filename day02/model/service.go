package model

type Service struct {
	Name   string `json:"name"`
	CPU    int    `json:"cpu"`
	Memory int    `json:"memory"`
}

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
