package model

import "time"

type PropertyGroupOptionTranslation struct {
	CreatedAt             time.Time            `json:"createdAt,omitempty"`
	CustomFields          interface{}          `json:"customFields,omitempty"`
	Language              *Language            `json:"language,omitempty"`
	LanguageId            string               `json:"languageId,omitempty"`
	Name                  string               `json:"name,omitempty"`
	Position              float64              `json:"position,omitempty"`
	PropertyGroupOption   *PropertyGroupOption `json:"propertyGroupOption,omitempty"`
	PropertyGroupOptionId string               `json:"propertyGroupOptionId,omitempty"`
	UpdatedAt             time.Time            `json:"updatedAt,omitempty"`
}

type PropertyGroupOptionTranslationCollection struct {
	EntityCollection

	Data []PropertyGroupOptionTranslation `json:"data"`
}
