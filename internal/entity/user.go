package entity

type User struct {
	// Should have an ID, maybe UUID and alike
	Meta map[string]any `json:"meta"`
}

func (u User) ToMap() map[string]any {
	data := make(map[string]any)
	for k, v := range u.Meta {
		data[k] = v
	}
	return data
}
