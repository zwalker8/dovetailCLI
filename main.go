package main

import (
	//"net/http"
	"log"
	"net/http"

	"github.com/charmbracelet/huh"
)

var (
	choice     string
	singleNote bool
	noteID     string
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
		huh.NewConfirm().
			Title("View a single note?").
			Affirmative("Yes.").
			Negative("No.").
			Value(&singleNote).Run()

		if singleNote == true {
			huh.NewInput().
				Title("What's the ID").
				Prompt("?").
				Value(&noteID).Run()
			api.GetNote(noteID)
			return
		}

		api.ListNotes()

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

	form := huh.NewForm(huh.NewGroup(
		huh.NewSelect[string]().
			Title("What would you like to do?").
			Options(
				huh.NewOption("Get token info", "token"),
				huh.NewOption("Get highlights", "highlights"),
				huh.NewOption("Get insights", "insights"),
				huh.NewOption("Get projects", "projects"),
				huh.NewOption("Get notes", "notes"),
			).
			Value(&choice)))

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	api.GetResponse(choice)
}
