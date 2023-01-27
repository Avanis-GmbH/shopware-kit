package model

import "time"

type ImportExportFile struct {
	AccessToken  string           `json:"accessToken,omitempty"`
	CreatedAt    time.Time        `json:"createdAt,omitempty"`
	ExpireDate   time.Time        `json:"expireDate"` // required
	Id           string           `json:"id,omitempty"`
	Log          *ImportExportLog `json:"log,omitempty"`
	OriginalName string           `json:"originalName"` // required
	Path         string           `json:"path"`         // required
	Size         float64          `json:"size,omitempty"`
	UpdatedAt    time.Time        `json:"updatedAt,omitempty"`
}
