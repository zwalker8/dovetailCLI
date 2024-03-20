package main

import (
	"github.com/charmbracelet/huh/spinner"
	"github.com/zwalker8/dovetailCLI/api"
)

func Load(app *Application) {
	c := make(chan *api.ListNotes, 1)

	action := func() {
		notes, _ := app.API.ListNotes("")
		c <- notes

		// time.Sleep(2 * time.Second)
	}

	spinner.New().Action(action).Run()
	list := <-c

	PrettyPrint(list)
}
