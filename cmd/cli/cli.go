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
	var ids []string

	projects := app.FilterProjects()

	huh.NewForm(huh.NewGroup(huh.NewSelect[string]().
		Title("Choose an option").
		Options(
			huh.NewOption("Display all notes", "all"),
			huh.NewOption("Display select notes", "select"),
			huh.NewOption("Delete notes", "delete"),
		).
		Value(&option))).Run()

	notePages, err := GetAllPages(app.API.ListNotes, 10, projects...)
	if err != nil {
		PrettyPrint(err)
		return
	}

	if option == "all" {
		var items []Printable
		for _, page := range notePages {
			items = append(items, page)
		}
		Paginate(items...)
		return
	}

	options := []huh.Option[string]{}

	for _, page := range notePages {
		for _, note := range page.Data {
			if note.Title == "" {
				note.Title = "Untitled"
			}
			if !note.Deleted {
				options = append(options, huh.NewOption(
					note.Title, note.ID))
			}
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
	pages, err := app.GetAllProjects(100)
	if err != nil {
		PrettyPrint(err)
	}

	for _, projects := range pages {
		for _, project := range projects.Data {
			projectOptions = append(projectOptions, huh.NewOption(project.Title, project.ID))
		}
	}

	huh.NewForm(huh.NewGroup(huh.NewMultiSelect[string]().
		Title("Choose Project(s)").
		Options(
			projectOptions...,
		).
		Value(&projectIDS))).Run()

	return projectIDS
}

func (app *Application) DisplayHighlights() {
	projects := app.FilterProjects()
	var items []Printable

	highlightPages, err := GetAllPages(app.API.ListHighlights, 10, projects...)
	if err != nil {
		PrettyPrint(err)
		return
	}
	for _, page := range highlightPages {
		items = append(items, page)
	}

	Paginate(items...)
}

func (app *Application) DisplayProjects() {
	var items []Printable
	projectPages, err := app.GetAllProjects(10)
	if err != nil {
		PrettyPrint(err)
	}

	for _, page := range projectPages {
		items = append(items, page)
	}

	Paginate(items...)
}

func (app *Application) DisplayInsights() {
	projects := app.FilterProjects()
	var items []Printable

	projectPages, err := GetAllPages(app.API.ListInsights, 10, projects...)
	if err != nil {
		PrettyPrint(err)
		return
	}
	for _, page := range projectPages {
		items = append(items, page)
	}

	Paginate(items...)
}

func (app *Application) GetResponse(choice string) {
	switch choice {
	case "token":
		PrettyPrint(app.API.TokenInfo())
	case "highlights":
		app.DisplayHighlights()
	case "insights":
		app.DisplayInsights()
	case "projects":
		app.DisplayProjects()
	case "notes":
		app.ChooseNotes()
	case "files":
		app.ChooseFile()
	case "exit":
		return
	}
}
