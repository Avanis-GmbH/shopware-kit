package model

import "time"

type CurrencyTranslation struct {
	CreatedAt    time.Time   `json:"createdAt,omitempty"`
	Currency     *Currency   `json:"currency,omitempty"`
	CurrencyId   string      `json:"currencyId,omitempty"`
	CustomFields interface{} `json:"customFields,omitempty"`
	Language     *Language   `json:"language,omitempty"`
	LanguageId   string      `json:"languageId,omitempty"`
	Name         string      `json:"name,omitempty"`
	ShortName    string      `json:"shortName,omitempty"`
	UpdatedAt    time.Time   `json:"updatedAt,omitempty"`
}

type CurrencyTranslationCollection struct {
	EntityCollection

	Data []CurrencyTranslation `json:"data"`
}
