package model

import "time"

type SalutationTranslation struct {
	CreatedAt    *time.Time  `json:"createdAt,omitempty"`
	CustomFields interface{} `json:"customFields,omitempty"`
	DisplayName  string      `json:"displayName,omitempty"`
	Language     *Language   `json:"language,omitempty"`
	LanguageId   string      `json:"languageId,omitempty"`
	LetterName   string      `json:"letterName,omitempty"`
	Salutation   *Salutation `json:"salutation,omitempty"`
	SalutationId string      `json:"salutationId,omitempty"`
	UpdatedAt    *time.Time  `json:"updatedAt,omitempty"`
}
