package main

type TokenInfo struct {
	Data struct {
		ID        string `json:"id"`
		Subdomain string `json:"subdomain"`
	}
}
