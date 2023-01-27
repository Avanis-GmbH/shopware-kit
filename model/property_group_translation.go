package model

import "time"

type PropertyGroupTranslation struct {
	CreatedAt       time.Time      `json:"createdAt,omitempty"`
	CustomFields    interface{}    `json:"customFields,omitempty"`
	Description     string         `json:"description,omitempty"`
	Language        *Language      `json:"language,omitempty"`
	LanguageId      string         `json:"languageId,omitempty"`
	Name            string         `json:"name,omitempty"`
	Position        float64        `json:"position"`
	PropertyGroup   *PropertyGroup `json:"propertyGroup,omitempty"`
	PropertyGroupId string         `json:"propertyGroupId,omitempty"`
	UpdatedAt       time.Time      `json:"updatedAt,omitempty"`
}
