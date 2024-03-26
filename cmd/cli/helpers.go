package main

import (
	"fmt"
	"os"
	"reflect"

	"github.com/charmbracelet/huh"
	"github.com/zwalker8/dovetailCLI/api"
)

type Printable interface {
	Print()
}

type MultiPage interface {
	NextPage() api.Page
}

func PrettyPrint(inputs ...Printable) {
	for _, input := range inputs {
		nilStruct := reflect.ValueOf(input).IsNil()
		if !nilStruct {
			input.Print()
		}
	}
}

func GetAPIKEY() string {
	var key string

	key, exists := os.LookupEnv("DOVETAIL_API_KEY")
	if exists && key != "" {
		return key
	}

	for key == "" {
		huh.NewInput().
			Title("What is your API key?").
			Prompt(">").
			Value(&key).Run()
	}

	os.Setenv("DOVETAIL_API_KEY", key)

	return key
}

func Paginate(pages ...Printable) {
	var choice string
	currPage := 0
	pageCount := len(pages)

	if pageCount == 0 {
		fmt.Println("No items to display.")
		return
	}

	if pageCount == 1 {
		PrettyPrint(pages[currPage])
		return
	}

	for {
		fmt.Println("------------------------")
		PrettyPrint(pages[currPage])
		fmt.Println("------------------------")

		if currPage == 0 {
			huh.NewForm(huh.NewGroup(huh.NewSelect[string]().Title(fmt.Sprintf("(%v/%v)", currPage+1, pageCount)).Options(
				huh.NewOption("Next", "next"),
				huh.NewOption("Menu", "menu"),
			).Value(&choice))).Run()
		} else if currPage == pageCount-1 {
			huh.NewForm(huh.NewGroup(huh.NewSelect[string]().Title(fmt.Sprintf("(%v/%v)", currPage+1, pageCount)).Options(
				huh.NewOption("Prev", "prev"),
				huh.NewOption("Menu", "menu"),
			).Value(&choice))).Run()
		} else if currPage > 0 && currPage < pageCount-1 {
			huh.NewForm(huh.NewGroup(huh.NewSelect[string]().Title(fmt.Sprintf("(%v/%v)", currPage+1, pageCount)).Options(
				huh.NewOption("Next", "next"),
				huh.NewOption("Prev", "prev"),
				huh.NewOption("Menu", "menu"),
			).Value(&choice))).Run()
		}

		switch choice {
		case "next":
			currPage++
		case "prev":
			currPage--
		case "menu":
			return
		}
	}
}

func GetAllPages[M MultiPage](getFunc func(string, uint8, ...string) (M, *api.APIError), limit uint8, projects ...string) ([]M, *api.APIError) {
	var data []M
	curr, err := getFunc("", limit, projects...)
	if err != nil {
		return nil, err
	}

	data = append(data, curr)

	if !curr.NextPage().HasMore {
		return data, nil
	}

	nextCursor := curr.NextPage().NextCursor

	for nextCursor != nil {
		curr, err = getFunc(*nextCursor, limit, projects...)
		if err != nil {
			return nil, err
		}
		data = append(data, curr)
		nextCursor = curr.NextPage().NextCursor
	}
	return data, nil
}

func (app *Application) GetAllProjects(limit uint8) ([]*api.ListProjects, *api.APIError) {
	var data []*api.ListProjects
	curr, err := app.API.ListProjects("", limit)
	if err != nil {
		return nil, err
	}

	data = append(data, curr)

	if !curr.NextPage().HasMore {
		return data, nil
	}

	nextCursor := curr.NextPage().NextCursor

	for nextCursor != nil {
		curr, err = app.API.ListProjects(*nextCursor, limit)
		if err != nil {
			return nil, err
		}
		data = append(data, curr)
		nextCursor = curr.NextPage().NextCursor
	}
	return data, nil
}
