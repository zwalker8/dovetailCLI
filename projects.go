package main

type ListProjects struct {
	Data []struct {
		ID     string
		Author *struct {
			ID   string
			Name string
		}
		Title     string
		Type      string
		CreatedAt string `json:"created_at,omitempty"`
	}
	Page Page
}
