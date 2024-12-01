package models

import "time"

type Attachment struct {
	ID             string    `json:"id"`
	Name           string    `json:"file_name"`
	MimeType       string    `json:"mime_type"`
	Size           int       `json:"size"`
	UploadedBy     string    `json:"uploaded_by"`
	UnloadingDurMS int       `json:"unloading_dur_ms"`
	CreatedAt      time.Time `json:"created_at"`
}
