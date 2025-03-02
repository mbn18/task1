package entity

import (
	"github.com/google/uuid"
	"time"
)

type Process struct {
	Id      uuid.UUID
	GroupId uuid.UUID
	User    string
	PID     int
	CPU     float64
	Memory  float64
	VSZ     int
	RSS     int
	TTY     string
	Stat    string
	Start   time.Time
	CPUTime time.Duration
	Command string
	Args    string
}

type ProcessList struct {
	CreatedAt time.Time
	Processes []*Process
}
