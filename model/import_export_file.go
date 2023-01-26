package model

import "time"

type ImportExportFile struct {
	AccessToken  string           `json:"accessToken,omitempty"`
	CreatedAt    time.Time        `json:"createdAt,omitempty"`
	ExpireDate   time.Time        `json:"expireDate,omitempty"`
	Id           string           `json:"id,omitempty"`
	Log          *ImportExportLog `json:"log,omitempty"`
	OriginalName string           `json:"originalName,omitempty"`
	Path         string           `json:"path,omitempty"`
	Size         float64          `json:"size,omitempty"`
	UpdatedAt    time.Time        `json:"updatedAt,omitempty"`
}

type ImportExportFileCollection struct {
	EntityCollection

	Data []ImportExportFile `json:"data"`
}
