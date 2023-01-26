package model

import "time"

type LocaleTranslation struct {
	Language     *Language   `json:"language,omitempty"`
	Name         string      `json:"name,omitempty"`
	LocaleId     string      `json:"localeId,omitempty"`
	LanguageId   string      `json:"languageId,omitempty"`
	UpdatedAt    time.Time   `json:"updatedAt,omitempty"`
	Locale       *Locale     `json:"locale,omitempty"`
	Territory    string      `json:"territory,omitempty"`
	CustomFields interface{} `json:"customFields,omitempty"`
	CreatedAt    time.Time   `json:"createdAt,omitempty"`
}

type LocaleTranslationCollection struct {
	EntityCollection

	Data []LocaleTranslation `json:"data"`
}
