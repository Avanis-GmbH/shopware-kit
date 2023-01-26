package model

import "time"

type NumberRangeTypeTranslation struct {
	CreatedAt         time.Time        `json:"createdAt,omitempty"`
	CustomFields      interface{}      `json:"customFields,omitempty"`
	Language          *Language        `json:"language,omitempty"`
	LanguageId        string           `json:"languageId,omitempty"`
	NumberRangeType   *NumberRangeType `json:"numberRangeType,omitempty"`
	NumberRangeTypeId string           `json:"numberRangeTypeId,omitempty"`
	TypeName          string           `json:"typeName,omitempty"`
	UpdatedAt         time.Time        `json:"updatedAt,omitempty"`
}

type NumberRangeTypeTranslationCollection struct {
	EntityCollection

	Data []NumberRangeTypeTranslation `json:"data"`
}
