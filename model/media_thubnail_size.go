package model

import "time"

type MediaThumbnailSize struct {
	CreatedAt                 time.Time                  `json:"createdAt,omitempty"`
	CustomFields              interface{}                `json:"customFields,omitempty"`
	Height                    float64                    `json:"height,omitempty"`
	Id                        string                     `json:"id,omitempty"`
	MediaFolderConfigurations []MediaFolderConfiguration `json:"mediaFolderConfigurations,omitempty"`
	UpdatedAt                 time.Time                  `json:"updatedAt,omitempty"`
	Width                     float64                    `json:"width,omitempty"`
}
