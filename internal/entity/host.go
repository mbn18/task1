package entity

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
