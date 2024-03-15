package main

import (
	"log"

	"github.com/charmbracelet/huh"
)

func (api *API) MainMenu() {
	api.LoadToken()
	var choice string

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

func (api *API) ChooseNotes() {
	var singleNote bool
	var noteID string

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
