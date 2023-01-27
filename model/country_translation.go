package model

import "time"

type CountryTranslation struct {
	Country      *Country    `json:"country,omitempty"`
	CountryId    string      `json:"countryId,omitempty"`
	CreatedAt    time.Time   `json:"createdAt,omitempty"`
	CustomFields interface{} `json:"customFields,omitempty"`
	Language     *Language   `json:"language,omitempty"`
	LanguageId   string      `json:"languageId,omitempty"`
	Name         string      `json:"name,omitempty"`
	UpdatedAt    time.Time   `json:"updatedAt,omitempty"`
}
