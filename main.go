package main

import (
	"net/http"
)

func (api *API) GetResponse(choice string) {
	switch choice {
	case "token":
		api.TokenInfo()
	case "highlights":
		api.ListHighlights()
	case "insights":
		api.ListInsights()
	case "projects":
		api.ListProjects()
	case "notes":
		api.ChooseNotes()
	}
}

func main() {
	client := &http.Client{}
	apiKey := "api.5SiJujQaspZZcrywzdrUpQ.7biOUmU9mafWazA2SVQWEf"
	api := &API{
		Client: client,
		Key:    apiKey,
		Routes: Routes{
			TokenInfo:  "https://dovetail.com/api/v1/token/info",
			Highlights: "https://dovetail.com/api/v1/highlights",
			Insights:   "https://dovetail.com/api/v1/insights",
			Projects:   "https://dovetail.com/api/v1/projects",
			Notes:      "https://dovetail.com/api/v1/notes",
			Files:      "https://dovetail.com/api/v1/files",
		},
	}

	// api.TokenInfo()
	// api.ListInsights()
	// api.ListProjects()
	// api.ListNotes()
	// api.GetNote("6FtiqmKwLt9jvEKpnmTPzC")

	api.MainMenu()
}
