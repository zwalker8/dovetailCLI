package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

func (api *API) GetRequest(url string) *http.Response {
	client := api.Client
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", api.Key))

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	return res
}

func (api *API) TokenInfo() (*TokenInfo, *APIError) {
	res := api.GetRequest(api.Routes.TokenInfo)
	defer res.Body.Close()

	var apiResponse TokenInfo
	var apiError APIError

	return DecodeResponse(res, &apiResponse, &apiError)
}

func (api *API) ListHighlights() (*ListHighlights, *APIError) {
	var apiResponse ListHighlights
	var apiError APIError
	res := api.GetRequest(api.Routes.Highlights)

	return DecodeResponse(res, &apiResponse, &apiError)
}

func (api *API) ListInsights() (*ListInsights, *APIError) {
	var apiResponse ListInsights
	var apiError APIError
	res := api.GetRequest(api.Routes.Insights)

	return DecodeResponse(res, &apiResponse, &apiError)
}

func (api *API) ListProjects() (*ListProjects, *APIError) {
	var apiResponse ListProjects
	var apiError APIError

	res := api.GetRequest(api.Routes.Projects)

	return DecodeResponse(res, &apiResponse, &apiError)
}

func (api *API) ListNotes() (*ListNotes, *APIError) {
	var apiResponse ListNotes
	var apiError APIError

	res := api.GetRequest(api.Routes.Notes)

	return DecodeResponse(res, &apiResponse, &apiError)
}

func (api *API) GetNote(id string) (*Note, *APIError) {
	var apiResponse Note
	var apiError APIError

	res := api.GetRequest(fmt.Sprintf("%v/%v", api.Routes.Notes, id))

	return DecodeResponse(res, &apiResponse, &apiError)
}

func (api *API) GetFile(id string) (*File, *APIError) {
	var apiResponse File
	var apiError APIError

	res := api.GetRequest(fmt.Sprintf("%v/%v", api.Routes.Files, id))

	return DecodeResponse(res, &apiResponse, &apiError)
}

func (api *API) PatchRequest(url string, patch any) *http.Response {
	json, err := json.Marshal(patch)
	if err != nil {
		log.Fatal(err)
	}

	body := bytes.NewBuffer(json)

	req, err := http.NewRequest(http.MethodPatch, url, body)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", fmt.Sprintf("Bearer %v", api.Key))

	res, err := api.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	return res
}

func (api *API) PatchNote(id string, title string, fields Fields) (*Note, *APIError) {
	var apiResponse Note
	var apiError APIError

	patch := PatchNote{
		Title:  title,
		Fields: fields,
	}

	res := api.PatchRequest(fmt.Sprintf("%v/%v", api.Routes.Notes, id), patch)

	return DecodeResponse(res, &apiResponse, &apiError)
}
