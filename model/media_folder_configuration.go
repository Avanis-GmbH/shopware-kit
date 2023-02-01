package model

import "time"

type MediaFolderConfiguration struct {
	CreatedAt             *time.Time           `json:"createdAt,omitempty"`
	CreateThumbnails      bool                 `json:"createThumbnails,omitempty"`
	CustomFields          interface{}          `json:"customFields,omitempty"`
	Id                    string               `json:"id,omitempty"`
	KeepAspectRatio       bool                 `json:"keepAspectRatio,omitempty"`
	MediaFolders          []MediaFolder        `json:"mediaFolders,omitempty"`
	MediaThumbnailSizes   []MediaThumbnailSize `json:"mediaThumbnailSizes,omitempty"`
	MediaThumbnailSizesRo interface{}          `json:"mediaThumbnailSizesRo,omitempty"`
	NoAssociation         bool                 `json:"noAssociation,omitempty"`
	Private               bool                 `json:"private,omitempty"`
	ThumbnailQuality      float64              `json:"thumbnailQuality,omitempty"`
	UpdatedAt             *time.Time           `json:"updatedAt,omitempty"`
}
