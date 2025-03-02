package entity

import "time"

type OS string

const Windows = OS("windows")
const Linux = OS("linux")
const Darwin = OS("darwin")

type Host struct {
	// Should consider something like UUID
	ID        int
	User      User
	OS        OS
	Meta      map[string]any
	Processes ProcessList
}

// This method is used by the mapper, Should it be a fuction in the mapper or stay here as method?
// If we extend the system, then the ToMap() will serve diffrent cli and by that might violate the single responsebility principle.
func (h Host) ToMap() map[string]any {
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
