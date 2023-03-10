package model

import "time"

type Unit struct {
	CreatedAt    *time.Time        `json:"createdAt,omitempty"`
	CustomFields interface{}       `json:"customFields,omitempty"`
	Id           string            `json:"id,omitempty"`
	Name         string            `json:"name,omitempty"` // required
	Products     []Product         `json:"products,omitempty"`
	ShortCode    string            `json:"shortCode,omitempty"` // required
	Translated   interface{}       `json:"translated,omitempty"`
	Translations []UnitTranslation `json:"translations,omitempty"`
	UpdatedAt    *time.Time        `json:"updatedAt,omitempty"`
}
