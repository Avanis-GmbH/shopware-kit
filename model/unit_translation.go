package model

import "time"

type UnitTranslation struct {
	CreatedAt    *time.Time  `json:"createdAt,omitempty"`
	CustomFields interface{} `json:"customFields,omitempty"`
	Language     *Language   `json:"language,omitempty"`
	LanguageId   string      `json:"languageId,omitempty"`
	Name         string      `json:"name,omitempty"`
	ShortCode    string      `json:"shortCode,omitempty"`
	Unit         *Unit       `json:"unit,omitempty"`
	UnitId       string      `json:"unitId,omitempty"`
	UpdatedAt    *time.Time  `json:"updatedAt,omitempty"`
}
