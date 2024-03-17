package main

import (
	"fmt"
	"log"

	"github.com/charmbracelet/huh"
)

func (app *Application) MainMenu() {
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

		app.GetResponse(choice)

	}
}

func (app *Application) ChooseNotes() {
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
		n, _ := app.API.ListNotes()
		PrettyPrint(n)
		return
	}

	var ids []string

	notes, err := app.API.ListNotes()
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
			n, _ := app.API.GetNote(id)
			PrettyPrint(n)
		case "delete":
			var confirmed bool
			huh.NewConfirm().
				Title("Are you sure?").
				Affirmative("Yes").
				Negative("No").
				Value(&confirmed).Run()

			if confirmed {
				n, _ := app.API.DeleteNote(id)
				fmt.Printf("Deleted: %v\n", n.Data.Title)
			}
		}
	}
}

func (app *Application) ChooseFile() {
	var fileID string

	huh.NewInput().
		Title("What's the ID").
		Prompt("?").
		Value(&fileID).Run()

	app.API.GetFile(fileID)
}

func (app *Application) GetResponse(choice string) {
	switch choice {
	case "token":
		info, _ := app.API.TokenInfo()
		PrettyPrint(info)
	case "highlights":
		app.API.ListHighlights()
	case "insights":
		app.API.ListInsights()
	case "projects":
		app.API.ListProjects()
	case "notes":
		app.ChooseNotes()
	case "file":
		app.ChooseFile()
	case "exit":
		return
	}
}
