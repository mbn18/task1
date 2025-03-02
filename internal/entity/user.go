package entity

type User struct {
	// Should have an ID, maybe UUID and alike
	Meta map[string]any `json:"meta"`
}
