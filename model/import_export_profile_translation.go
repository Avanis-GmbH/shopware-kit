package model

import "time"

type ImportExportProfileTranslation struct {
	CreatedAt             *time.Time           `json:"createdAt,omitempty"`
	ImportExportProfile   *ImportExportProfile `json:"importExportProfile,omitempty"`
	ImportExportProfileId string               `json:"importExportProfileId,omitempty"`
	Label                 string               `json:"label,omitempty"`
	Language              *Language            `json:"language,omitempty"`
	LanguageId            string               `json:"languageId,omitempty"`
	UpdatedAt             *time.Time           `json:"updatedAt,omitempty"`
}
