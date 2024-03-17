package api

import "net/http"

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

	res := api.SendRequest(http.MethodGet, api.Routes.Projects, nil)

	return DecodeResponse(res, &apiResponse, &apiError)
}
