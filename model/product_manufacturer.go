package model

import "time"

type ProductManufacturer struct {
	CreatedAt    *time.Time                       `json:"createdAt,omitempty"`
	CustomFields interface{}                      `json:"customFields,omitempty"`
	Description  string                           `json:"description"`
	Id           string                           `json:"id,omitempty"`
	Link         string                           `json:"link,omitempty"`
	Media        *Media                           `json:"media,omitempty"`
	MediaId      *string                          `json:"mediaId"`
	Name         string                           `json:"name,omitempty"` // required
	Products     []Product                        `json:"products,omitempty"`
	Translated   interface{}                      `json:"translated,omitempty"`
	Translations []ProductManufacturerTranslation `json:"translations,omitempty"`
	UpdatedAt    *time.Time                       `json:"updatedAt,omitempty"`
	VersionId    string                           `json:"versionId,omitempty"`
}
