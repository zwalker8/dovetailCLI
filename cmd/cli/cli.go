package main

import (
	"fmt"
	"log"

	"github.com/charmbracelet/huh"
	"github.com/zwalker8/dovetailCLI/api"
)

func (app *Application) MainMenu() {
	var choice string

	for choice != "exit" {
		form := huh.NewForm(huh.NewGroup(
			huh.NewSelect[string]().
				Title("What would you like to access?").
				Options(
					huh.NewOption("Token Info", "token"),
					huh.NewOption("Highlights", "highlights"),
					huh.NewOption("Insights", "insights"),
					huh.NewOption("Projects", "projects"),
					huh.NewOption("Notes", "notes"),
					huh.NewOption("Files", "files"),
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

	projects := app.FilterProjects()

	huh.NewForm(huh.NewGroup(huh.NewSelect[string]().
		Title("Choose an option").
		Options(
			huh.NewOption("List all notes", "all"),
			huh.NewOption("Select indvidual notes", "select"),
			huh.NewOption("Delete note", "delete"),
		).
		Value(&option))).Run()

	if option == "all" {
		PrettyPrint(app.API.ListNotes("", projects...))
		return
	}

	var ids []string

	notes, err := app.API.ListNotes("", projects...)
	if err != nil {
		log.Fatal(err)
	}

	options := []huh.Option[string]{}

	for _, note := range notes.Data {
		if note.Title == "" {
			note.Title = "Untitled"
		}
		if !note.Deleted {
			options = append(options, huh.NewOption(
				note.Title, note.ID))
		}
	}

	huh.NewForm(huh.NewGroup(huh.NewMultiSelect[string]().
		Options(options...).
		Title("Notes").
		Value(&ids))).Run()

	for _, id := range ids {
		switch option {
		case "select":
			PrettyPrint(app.API.GetNote(id))
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

	PrettyPrint(app.API.GetFile(fileID))
}

func (app *Application) FilterProjects() []string {
	var projectIDS []string

	projectOptions := []huh.Option[string]{}
	projects, _ := app.API.ListProjects()
	for _, project := range projects.Data {
		projectOptions = append(projectOptions, huh.NewOption(project.Title, project.ID))
	}

	huh.NewForm(huh.NewGroup(huh.NewMultiSelect[string]().
		Title("Choose Project(s) or none for all Projects").
		Options(
			projectOptions...,
		).
		Value(&projectIDS))).Run()

	return projectIDS
}

func (app *Application) ChooseHighlights() {
	projects := app.FilterProjects()
	h, err := app.API.ListHighlights("", projects...)

	PrettyPrint(h, err)

	paginateFunc := func(page string) (*string, *api.APIError) {
		h, err := app.API.ListHighlights(page, projects...)
		PrettyPrint(h, err)

		return h.Page.NextCursor, err
	}
	// fmt.Println(h.Page.TotalCount)
	Paginate(h.Page, paginateFunc)
}

func (app *Application) GetResponse(choice string) {
	switch choice {
	case "token":
		PrettyPrint(app.API.TokenInfo())
	case "highlights":
		app.ChooseHighlights()
	case "insights":
		PrettyPrint(app.API.ListInsights(""))
	case "projects":
		PrettyPrint(app.API.ListProjects())
	case "notes":
		app.ChooseNotes()
	case "files":
		app.ChooseFile()
	case "exit":
		return
	}
}
