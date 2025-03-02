package upsert

import (
	"fmt"
	"github.com/mbn18/dream/internal/entity"
	"time"
)

const (
	mapKeyHost      = "host"
	mapKeyUser      = "user"
	mapKeyProcesses = "processes"
)

func toMap(h *entity.Host) map[string]any {
	return map[string]any{
		mapKeyHost:      toMapHost(h),
		mapKeyUser:      toMapUser(h.User),
		mapKeyProcesses: toMapProcesses(h.Processes),
	}
}

func toMapHost(h *entity.Host) map[string]any {
	data := map[string]any{
		"id": h.ID,
		"os": h.OS,
	}
	for k, v := range h.Meta {
		switch v.(type) {
		case time.Time:
			v = v.(time.Time).UTC().Format(time.RFC3339)
		}
		data[k] = v
	}
	return data

}

func toMapUser(u entity.User) map[string]any {
	data := make(map[string]any)
	for k, v := range u.Meta {
		data[k] = v
	}
	return data
}

func toMapProcesses(l entity.ProcessList) map[string]any {
	data := make(map[string]any)
	for i, p := range l.Processes {
		id := fmt.Sprintf("p%d", i)
		data[id] = ToMapProcess(p)
	}
	return data
}

func ToMapProcess(p *entity.Process) map[string]any {
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
