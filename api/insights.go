package api

import (
	"fmt"
	"net/http"
)

type ListInsights struct {
	Data []struct {
		ID        string
		Title     string
		Type      string
		CreatedAt string `json:"created_at"`
	}
	Page Page
}

func (api *API) ListInsights(page string, limit uint8, projects ...string) (*ListInsights, *APIError) {
	var apiResponse ListInsights
	var apiError APIError

	url := api.Routes.Insights + JoinQueryParams(page, limit, projects...)

	res := api.SendRequest(http.MethodGet, url, nil)

	return DecodeResponse(res, &apiResponse, &apiError)
}

func (i *ListInsights) Print() {
	if len(i.Data) == 0 {
		fmt.Println("No insights for this project")
		return
	}

	for _, insight := range i.Data {
		if insight.Title != "" {
			fmt.Printf("%v\n\n", insight.Title)
		}
	}
}

func (i *ListInsights) NextPage() Page {
	return i.Page
}
