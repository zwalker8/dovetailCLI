package main

import (
	"net/http"

	"github.com/zwalker8/dovetailCLI/api"
)

type Application struct {
	API *api.API
}

func main() {
	client := &http.Client{}
	api := &api.API{
		Client: client,
		Key:    GetAPIKEY(),
		Routes: api.Routes{
			TokenInfo:  "https://dovetail.com/api/v1/token/info",
			Highlights: "https://dovetail.com/api/v1/highlights",
			Insights:   "https://dovetail.com/api/v1/insights",
			Projects:   "https://dovetail.com/api/v1/projects",
			Notes:      "https://dovetail.com/api/v1/notes",
			Files:      "https://dovetail.com/api/v1/files",
		},
	}

	app := &Application{
		API: api,
	}

	app.MainMenu()
	// fmt.Println(app.API.FilterProjects())

	// app.API.ListNotes()
}
