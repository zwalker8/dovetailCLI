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

func (api *API) ListNotes(page string, limit uint8, projects ...string) (*ListNotes, *APIError) {
	var apiResponse ListNotes
	var apiError APIError

	url := api.Routes.Notes + JoinQueryParams(page, limit, projects...)

	res := api.SendRequest(http.MethodGet, url, nil)

	return DecodeResponse(res, &apiResponse, &apiError)
}

func (api *API) GetNote(id string) (*Note, *APIError) {
	var apiResponse Note
	var apiError APIError

	url := fmt.Sprintf("%v/%v", api.Routes.Notes, id)
	res := api.SendRequest(http.MethodGet, url, nil)

	return DecodeResponse(res, &apiResponse, &apiError)
}

func (api *API) PatchNote(noteID string, patch PatchNote) (*Note, *APIError) {
	var apiResponse Note
	var apiError APIError

	url := fmt.Sprintf("%v/%v", api.Routes.Notes, noteID)

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

func (notes *ListNotes) Print() {
	if len(notes.Data) == 0 {
		fmt.Println("No notes for this project")
		return
	}
	for _, note := range notes.Data {
		if note.Title != "" {
			fmt.Println(note.Title)
		} else {
			fmt.Println("Untitled")
		}
	}
}

func (note *Note) Print() {
	fmt.Println(note.Data.Title)
	fmt.Printf("ID: %v\n", note.Data.ID)
	fmt.Println("Files:")

	if len(note.Data.Files) == 0 {
		fmt.Println("None")
		return
	}

	for _, file := range note.Data.Files {
		fmt.Printf("Name: %v, Type: ", file.Name)
		if file.Type != nil {
			fmt.Println(*file.Type)
		} else {
			fmt.Println("Unknown")
		}

	}
}

func (notes *ListNotes) NextPage() Page {
	return notes.Page
}
