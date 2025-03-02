package entity

type User struct {
	// Should have an ID, maybe UUID and alike
	Meta map[string]any `json:"meta"`
}

// This method is used by the mapper, Should it be a fuction in the mapper or stay here as method?
// If we extend the system, then the ToMap() will serve diffrent cli and by that might violate the single responsebility principle.
func (u User) ToMap() map[string]any {
	data := make(map[string]any)
	for k, v := range u.Meta {
		data[k] = v
	}
	return data
}
