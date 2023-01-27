package model

import "time"

type ProductMedia struct {
	CreatedAt        time.Time   `json:"createdAt,omitempty"`
	CustomFields     interface{} `json:"customFields,omitempty"`
	Id               string      `json:"id,omitempty"`
	Media            *Media      `json:"media,omitempty"`
	MediaId          string      `json:"mediaId"` // required
	Position         float64     `json:"position"`
	Product          *Product    `json:"product,omitempty"`
	ProductId        string      `json:"productId"` // required
	ProductVersionId string      `json:"productVersionId,omitempty"`
	UpdatedAt        time.Time   `json:"updatedAt,omitempty"`
	VersionId        string      `json:"versionId,omitempty"`
}
