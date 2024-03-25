package api

import (
	"fmt"
	"net/http"
)

type ListHighlights struct {
	Data []struct {
		ID   string
		Tags []struct {
			ID    string
			Title string
		}
		Text      *string `json:"text,omitempty"`
		Type      string
		CreatedAt string `json:"created_at"`
	}
	Page Page
}

func (api *API) ListHighlights(page string, limit uint8, projects ...string) (*ListHighlights, *APIError) {
	var apiResponse ListHighlights
	var apiError APIError

	url := api.Routes.Highlights + JoinQueryParams(page, limit, projects...)

	res := api.SendRequest(http.MethodGet, url, nil)

	return DecodeResponse(res, &apiResponse, &apiError)
}

func (h *ListHighlights) Print() {
	if len(h.Data) == 0 {
		fmt.Println("No highlights for this project")
		return
	}

	for _, highlight := range h.Data {
		if len(highlight.Tags) != 0 {
			fmt.Printf("Title: %v\n", highlight.Tags[0].Title)
		}
		if highlight.Text != nil {
			fmt.Printf("Text: %v\n\n", *highlight.Text)
		}
	}
}

func (h *ListHighlights) NextPage() Page {
	return h.Page
}
