package main

type ListHighlights struct {
	Data []struct {
		ID   string
		Tags []struct {
			ID    string
			Title string
		}
		Text      *string `json:"text,omitempty"`
		Type      string
		CreatedAt string `json:"created_at"`
	}
	Page Page
}
