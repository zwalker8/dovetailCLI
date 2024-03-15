package main

import (
	"fmt"
	"net/http"
)

type File struct {
	Data struct {
		ID     string
		Name   string
		Type   *string
		Status string
	}
}

func (api *API) GetFile(id string) (*File, *APIError) {
	var apiResponse File
	var apiError APIError

	url := fmt.Sprintf("%v/%v", api.Routes.Files, id)

	res := api.SendRequest(http.MethodGet, url, nil)

	return DecodeResponse(res, &apiResponse, &apiError)
}
