package model

import "time"

type TaxRuleType struct {
	CreatedAt     time.Time                `json:"createdAt,omitempty"`
	Id            string                   `json:"id,omitempty"`
	Position      float64                  `json:"position,omitempty"`
	Rules         []TaxRule                `json:"rules,omitempty"`
	TechnicalName string                   `json:"technicalName,omitempty"`
	Translated    interface{}              `json:"translated,omitempty"`
	Translations  []TaxRuleTypeTranslation `json:"translations,omitempty"`
	TypeName      string                   `json:"typeName,omitempty"`
	UpdatedAt     time.Time                `json:"updatedAt,omitempty"`
}
