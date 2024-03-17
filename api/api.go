package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/charmbracelet/huh/spinner"
)

type Routes struct {
	TokenInfo  string
	Highlights string
	Insights   string
	Projects   string
	Notes      string
	Files      string
}

type API struct {
	Client *http.Client
	Key    string
	Routes Routes
}

type APIError struct {
	Errors []struct {
		Code    string `json:"code"`
		Title   string `json:"Title"`
		Message string `json:"Message,omitempty"`
		Path    string `json:"Path,omitempty"`
	}
}

type Page struct {
	TotalCount int     `json:"total_count"`
	HasMore    bool    `json:"has_more"`
	NextCursor *string `json:"next_cursor,omitempty"`
}

func (api *API) SendRequest(method string, url string, content any) *http.Response {
	var body bytes.Buffer

	if content != nil {
		json, err := json.Marshal(content)
		if err != nil {
			log.Fatal(err)
		}

		body = *bytes.NewBuffer(json)
	}

	req, err := http.NewRequest(method, url, &body)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", fmt.Sprintf("Bearer %v", api.Key))

	c := make(chan *http.Response, 1)

	load := func() {
		res, err := api.Client.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		c <- res
	}

	spinner.New().Action(load).Run()

	return <-c
}
