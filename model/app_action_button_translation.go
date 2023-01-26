package model

import "time"

type AppActionButtonTranslation struct {
	AppActionButton   *AppActionButton `json:"appActionButton,omitempty"`
	AppActionButtonId string           `json:"appActionButtonId,omitempty"`
	CreatedAt         time.Time        `json:"createdAt,omitempty"`
	Label             string           `json:"label,omitempty"`
	Language          *Language        `json:"language,omitempty"`
	LanguageId        string           `json:"languageId,omitempty"`
	UpdatedAt         time.Time        `json:"updatedAt,omitempty"`
}
