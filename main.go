package main

import (
	"net/http"
)

func main() {
	client := &http.Client{}
	api := &API{
		Client: client,
		Routes: Routes{
			TokenInfo:  "https://dovetail.com/api/v1/token/info",
			Highlights: "https://dovetail.com/api/v1/highlights",
			Insights:   "https://dovetail.com/api/v1/insights",
			Projects:   "https://dovetail.com/api/v1/projects",
			Notes:      "https://dovetail.com/api/v1/notes",
			Files:      "https://dovetail.com/api/v1/files",
		},
	}

	api.MainMenu()
}
