package api

import (
	"fmt"
	"net/http"
)

type ListProjects struct {
	Data []struct {
		ID     string
		Author *struct {
			ID   string
			Name string
		}
		Title     string
		Type      string
		CreatedAt string `json:"created_at,omitempty"`
	}
	Page Page
}

func (api *API) ListProjects() (*ListProjects, *APIError) {
	var apiResponse ListProjects
	var apiError APIError

	url := api.Routes.Projects

	res := api.SendRequest(http.MethodGet, url, nil)

	return DecodeResponse(res, &apiResponse, &apiError)
}

func (p *ListProjects) Print() {
	for _, project := range p.Data {
		fmt.Println(project.Title)
	}
}
