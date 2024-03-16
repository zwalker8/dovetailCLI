package main

import (
	"fmt"
	"log"

	"github.com/charmbracelet/huh"
)

func (api *API) MainMenu() {
	var choice string

	for choice != "exit" {
		form := huh.NewForm(huh.NewGroup(
			huh.NewSelect[string]().
				Title("What would you like to do?").
				Options(
					huh.NewOption("Get Token Info", "token"),
					huh.NewOption("List Highlights", "highlights"),
					huh.NewOption("List Insights", "insights"),
					huh.NewOption("List projects", "projects"),
					huh.NewOption("Notes", "notes"),
					huh.NewOption("Get File", "file"),
					huh.NewOption("Exit", "exit"),
				).
				Value(&choice)))

		err := form.Run()
		if err != nil {
			log.Fatal(err)
		}

		api.GetResponse(choice)
	}
}

func (api *API) ChooseNotes() {
	var option string

	huh.NewSelect[string]().
		Title("Choose an option").
		Options(
			huh.NewOption("List all notes", "all"),
			huh.NewOption("Get single note", "get"),
			huh.NewOption("Delete note", "delete"),
		).
		Value(&option).Run()

	if option == "all" {
		n, _ := api.ListNotes()
		PrettyPrint(n)
		return
	}

	var ids []string

	notes, err := api.ListNotes()
	if err != nil {
		log.Fatal(err)
	}

	options := []huh.Option[string]{}

	for _, note := range notes.Data {
		options = append(options, huh.NewOption(
			note.Title, note.ID))
	}

	huh.NewMultiSelect[string]().
		Options(options...).
		Title("Notes").
		Value(&ids).Run()

	for _, id := range ids {
		switch option {
		case "get":
			n, _ := api.GetNote(id)
			PrettyPrint(n)
		case "delete":
			var confirmed bool
			huh.NewConfirm().
				Title("Are you sure?").
				Affirmative("Yes").
				Negative("No").
				Value(&confirmed).Run()

			if confirmed {
				n, _ := api.DeleteNote(id)
				fmt.Printf("Deleted: %v\n", n.Data.Title)
			}
		}
	}
}

func (api *API) ChooseFile() {
	var fileID string

	huh.NewInput().
		Title("What's the ID").
		Prompt("?").
		Value(&fileID).Run()

	api.GetFile(fileID)
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
	case "file":
		api.ChooseFile()
	}
}
