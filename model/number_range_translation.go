package model

import "time"

type NumberRangeTranslation struct {
	CreatedAt     time.Time    `json:"createdAt,omitempty"`
	CustomFields  interface{}  `json:"customFields,omitempty"`
	Description   string       `json:"description,omitempty"`
	Language      *Language    `json:"language,omitempty"`
	LanguageId    string       `json:"languageId,omitempty"`
	Name          string       `json:"name,omitempty"`
	NumberRange   *NumberRange `json:"numberRange,omitempty"`
	NumberRangeId string       `json:"numberRangeId,omitempty"`
	UpdatedAt     time.Time    `json:"updatedAt,omitempty"`
}

type NumberRangeTranslationCollection struct {
	EntityCollection

	Data []NumberRangeTranslation `json:"data"`
}
