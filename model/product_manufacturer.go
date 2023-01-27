package model

import "time"

type ProductManufacturer struct {
	CreatedAt    time.Time                        `json:"createdAt,omitempty"`
	CustomFields interface{}                      `json:"customFields,omitempty"`
	Description  string                           `json:"description,omitempty"`
	Id           string                           `json:"id,omitempty"`
	Link         string                           `json:"link,omitempty"`
	Media        *Media                           `json:"media,omitempty"`
	MediaId      string                           `json:"mediaId,omitempty"`
	Name         string                           `json:"name"` // required
	Products     []Product                        `json:"products,omitempty"`
	Translated   interface{}                      `json:"translated,omitempty"`
	Translations []ProductManufacturerTranslation `json:"translations,omitempty"`
	UpdatedAt    time.Time                        `json:"updatedAt,omitempty"`
	VersionId    string                           `json:"versionId,omitempty"`
}
