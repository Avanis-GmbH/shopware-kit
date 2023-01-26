package model

import "time"

type ProductStreamTranslation struct {
	Name            string         `json:"name,omitempty"`
	CreatedAt       time.Time      `json:"createdAt,omitempty"`
	UpdatedAt       time.Time      `json:"updatedAt,omitempty"`
	ProductStreamId string         `json:"productStreamId,omitempty"`
	LanguageId      string         `json:"languageId,omitempty"`
	Description     string         `json:"description,omitempty"`
	CustomFields    interface{}    `json:"customFields,omitempty"`
	ProductStream   *ProductStream `json:"productStream,omitempty"`
	Language        *Language      `json:"language,omitempty"`
}

type ProductStreamTranslationCollection struct {
	EntityCollection

	Data []ProductStreamTranslation `json:"data"`
}
