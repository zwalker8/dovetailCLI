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

func (api *API) ListProjects(page string, limit uint8) (*ListProjects, *APIError) {
	var apiResponse ListProjects
	var apiError APIError

	url := api.Routes.Projects + JoinQueryParams(page, limit)

	res := api.SendRequest(http.MethodGet, url, nil)

	return DecodeResponse(res, &apiResponse, &apiError)
}

func (p *ListProjects) Print() {
	if len(p.Data) == 0 {
		fmt.Println("No projects to display")
		return
	}
	for _, project := range p.Data {
		fmt.Println(project.Title)
	}
}

func (p *ListProjects) NextPage() Page {
	return p.Page
}
