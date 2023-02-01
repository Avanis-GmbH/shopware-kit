package model

import "time"

type CountryStateTranslation struct {
	CountryState   *CountryState `json:"countryState,omitempty"`
	CountryStateId string        `json:"countryStateId,omitempty"`
	CreatedAt      *time.Time    `json:"createdAt,omitempty"`
	CustomFields   interface{}   `json:"customFields,omitempty"`
	Language       *Language     `json:"language,omitempty"`
	LanguageId     string        `json:"languageId,omitempty"`
	Name           string        `json:"name,omitempty"`
	UpdatedAt      *time.Time    `json:"updatedAt,omitempty"`
}
