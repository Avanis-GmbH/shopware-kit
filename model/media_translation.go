package model

import "time"

type MediaTranslation struct {
	Alt          string      `json:"alt,omitempty"`
	CreatedAt    *time.Time  `json:"createdAt,omitempty"`
	CustomFields interface{} `json:"customFields,omitempty"`
	Language     *Language   `json:"language,omitempty"`
	LanguageId   string      `json:"languageId,omitempty"`
	Media        *Media      `json:"media,omitempty"`
	MediaId      string      `json:"mediaId,omitempty"`
	Title        string      `json:"title,omitempty"`
	UpdatedAt    *time.Time  `json:"updatedAt,omitempty"`
}
