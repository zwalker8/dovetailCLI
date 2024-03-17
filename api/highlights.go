package api

import "net/http"

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

func (api *API) ListHighlights() (*ListHighlights, *APIError) {
	var apiResponse ListHighlights
	var apiError APIError
	res := api.SendRequest(http.MethodGet, api.Routes.Highlights, nil)

	return DecodeResponse(res, &apiResponse, &apiError)
}
