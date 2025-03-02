package upsert

import (
	"github.com/mbn18/dream/internal/entity"
)

const (
	mapKeyHost      = "host"
	mapKeyUser      = "user"
	mapKeyProcesses = "processes"
)

func extractQueryParams(h *entity.Host) map[string]any {
	return map[string]any{
		mapKeyHost:      h.ToMap(),
		mapKeyUser:      h.User.ToMap(),
		mapKeyProcesses: h.Processes.ToMap(),
	}
}
