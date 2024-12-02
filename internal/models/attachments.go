package models

import "time"

type Attachment struct {
	Id             int       `json:"id"`
	Name           string    `json:"file_name"`
	FilePath       string    `json:"file_path"`
	MimeType       string    `json:"mime_type"`
	Size           int       `json:"size"`
	UploadedBy     int       `json:"uploaded_by"`
	UnloadingDurMS int       `json:"unloading_dur_ms"`
	CreatedAt      time.Time `json:"created_at"`
}
