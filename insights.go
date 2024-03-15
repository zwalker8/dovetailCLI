package main

import "net/http"

type ListInsights struct {
	Data []struct {
		ID        string
		Title     string
		Type      string
		CreatedAt string `json:"created_at"`
	}
	Page Page
}

func (api *API) ListInsights() (*ListInsights, *APIError) {
	var apiResponse ListInsights
	var apiError APIError
	res := api.SendRequest(http.MethodGet, api.Routes.Insights, nil)

	return DecodeResponse(res, &apiResponse, &apiError)
}
