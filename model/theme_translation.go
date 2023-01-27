package model

import "time"

type ThemeTranslation struct {
	CreatedAt    time.Time   `json:"createdAt,omitempty"`
	CustomFields interface{} `json:"customFields,omitempty"`
	Description  string      `json:"description,omitempty"`
	HelpTexts    interface{} `json:"helpTexts,omitempty"`
	Labels       interface{} `json:"labels,omitempty"`
	Language     *Language   `json:"language,omitempty"`
	LanguageId   string      `json:"languageId,omitempty"`
	Theme        *Theme      `json:"theme,omitempty"`
	ThemeId      string      `json:"themeId,omitempty"`
	UpdatedAt    time.Time   `json:"updatedAt,omitempty"`
}
