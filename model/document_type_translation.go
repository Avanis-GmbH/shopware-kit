package model

import "time"

type DocumentTypeTranslation struct {
	CreatedAt      time.Time     `json:"createdAt,omitempty"`
	CustomFields   interface{}   `json:"customFields,omitempty"`
	DocumentType   *DocumentType `json:"documentType,omitempty"`
	DocumentTypeId string        `json:"documentTypeId,omitempty"`
	Language       *Language     `json:"language,omitempty"`
	LanguageId     string        `json:"languageId,omitempty"`
	Name           string        `json:"name,omitempty"`
	UpdatedAt      time.Time     `json:"updatedAt,omitempty"`
}

type DocumentTypeTranslationCollection struct {
	EntityCollection

	Data []DocumentTypeTranslation `json:"data"`
}
