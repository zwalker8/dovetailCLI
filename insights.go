package main

type ListInsights struct {
	Data []struct {
		ID        string
		Title     string
		Type      string
		CreatedAt string `json:"created_at"`
	}
	Page Page
}
