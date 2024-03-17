package api

import (
	"fmt"
	"net/http"
)

type ListNotes struct {
	Data []struct {
		ID        string
		Title     string
		Type      string
		CreatedAt string `json:"created_at"`
		Deleted   bool
	}

	Page Page
}

type Fields []struct {
	Label string
	Value any
}

type Note struct {
	Data struct {
		ID      string
		Type    string
		Title   string
		Project struct {
			ID string
		}
		Fields Fields
		Files  []struct {
			ID     string
			Name   string
			Type   *string
			Status string
		}
		CreatedAt string `json:"created_at"`
		Deleted   bool
	}
}

type ExportNote struct {
	Data struct {
		Content   string
		Type      string
		Title     string
		CreatedAt string `json:"created_at"`
	}
}

type ImportFile struct {
	Fields     Fields `json:"fields,omitempty"`
	MimeType   string `json:"mime_type"`
	ProjectID  string `json:"project_id"`
	Title      string `json:"title"`
	Transcribe bool   `json:"transcribe"`
	Url        string `json:"url"`
}

type PatchNote struct {
	Title  string `json:"title"`
	Fields Fields `json:"fields"`
}

func (api *API) ListNotes() (*ListNotes, *APIError) {
	var apiResponse ListNotes
	var apiError APIError

	res := api.SendRequest(http.MethodGet, api.Routes.Notes, nil)

	return DecodeResponse(res, &apiResponse, &apiError)
}

func (api *API) GetNote(id string) (*Note, *APIError) {
	var apiResponse Note
	var apiError APIError

	url := fmt.Sprintf("%v/%v", api.Routes.Notes, id)
	res := api.SendRequest(http.MethodGet, url, nil)

	return DecodeResponse(res, &apiResponse, &apiError)
}

func (api *API) PatchNote(id string, title string, fields Fields) (*Note, *APIError) {
	var apiResponse Note
	var apiError APIError

	url := fmt.Sprintf("%v/%v", api.Routes.Notes, id)

	patch := PatchNote{
		Title:  title,
		Fields: fields,
	}

	res := api.SendRequest(http.MethodPatch, url, patch)

	return DecodeResponse(res, &apiResponse, &apiError)
}

func (api *API) DeleteNote(id string) (*Note, *APIError) {
	var apiResponse Note
	var apiError APIError

	url := fmt.Sprintf("%v/%v", api.Routes.Notes, id)

	res := api.SendRequest(http.MethodDelete, url, nil)

	return DecodeResponse(res, &apiResponse, &apiError)
}

func (api *API) FileToNote(fileUrl string, fields Fields, mimeType string, projectID string, title string, transcribe bool) (*Note, *APIError) {
	var apiResponse Note
	var apiError APIError

	subdomain := "import/file"
	url := fmt.Sprintf("%v/%v", api.Routes.Notes, subdomain)

	file := ImportFile{
		Fields:     nil,
		MimeType:   mimeType,
		ProjectID:  projectID,
		Title:      title,
		Transcribe: transcribe,
		Url:        fileUrl,
	}

	res := api.SendRequest(http.MethodPost, url, file)

	return DecodeResponse(res, &apiResponse, &apiError)
}
