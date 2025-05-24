package models

import "time"

type MediaFile struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	URL        string    `json:"url"`
	UploadedBy *int      `json:"uploaded_by"`
	CreatedAt  time.Time `json:"created_at"`
}
