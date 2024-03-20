package main

import (
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/charmbracelet/huh"
	"github.com/joho/godotenv"
	"github.com/zwalker8/dovetailCLI/api"
)

type Printable interface {
	Print()
}

type PaginateFunc func(string) (*string, *api.APIError)

func PrettyPrint(inputs ...Printable) {
	// prettyFmt, err := json.MarshalIndent(input, "", "\t")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(prettyFmt))

	for _, input := range inputs {
		nilStruct := reflect.ValueOf(input).IsNil()
		if !nilStruct {
			input.Print()
		}
	}
}

func GetAPIKEY() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	return os.Getenv("API_KEY")
}

func Paginate(page api.Page, paginateFunc PaginateFunc) {
	var next bool
	nextPage := page.NextCursor
	if nextPage == nil {
		return
	}

	currPage := 1
	pages := page.TotalCount / 100
	for nextPage != nil {
		huh.NewConfirm().
			Title(fmt.Sprintf("(%v/%v) Would you like to view the next page of items?", currPage, pages)).
			Affirmative("Yes").
			Negative("No").
			Value(&next).Run()
		if !next {
			return
		}
		page, err := paginateFunc(*nextPage)
		if err != nil {
			return
		}
		nextPage = page
		currPage++
	}
}
