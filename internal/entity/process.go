package entity

import (
	"fmt"
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

func (p Process) ToMap() map[string]any {
	return map[string]any{
		"id":       p.Id.String(),
		"group_id": p.GroupId.String(),
		"user":     p.User,
		"pid":      p.PID,
		"cpu":      p.CPU,
		"mem":      p.Memory,
		"vsz":      p.VSZ,
		"rss":      p.RSS,
		"tty":      p.TTY,
		"stat":     p.Stat,
		"start":    p.Start.UTC().Format(time.RFC3339),
		"cpu_time": p.CPUTime,
		"command":  p.Command,
		"args":     p.Args,
	}
}

type ProcessList struct {
	CreatedAt time.Time
	Processes []*Process
}

func (l ProcessList) ToMap() map[string]any {
	data := make(map[string]any)
	data["created_at"] = l.CreatedAt.UTC().Format(time.RFC3339)
	for i, p := range l.Processes {
		id := fmt.Sprintf("p%d", i)
		data[id] = p.ToMap()
	}
	return data
}
