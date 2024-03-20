package api

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

func (file *File) Print() {
	fmt.Printf("%v \n Type: ", file.Data.Name)
	if file.Data.Type != nil {
		fmt.Println(*file.Data.Type)
		return
	}
	fmt.Println("Unknown")
}
