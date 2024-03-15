package main

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

type Note struct {
	Data struct {
		ID      string
		Type    string
		Title   string
		Project struct {
			ID string
		}
		Fields []struct {
			Label string
			Value any
		}
		Files []struct {
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
	Fields []struct {
		Label string
		Value any
	}
	MimeType   string `json:"mime_type"`
	ProjectID  string `json:"project_id"`
	Title      string
	Transcribe bool
	Url        string
}
