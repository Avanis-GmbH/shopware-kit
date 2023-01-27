package model

import "time"

type TaxRuleType struct {
	CreatedAt     time.Time                `json:"createdAt,omitempty"`
	Id            string                   `json:"id,omitempty"`
	Position      int64                    `json:"position"` // required
	Rules         []TaxRule                `json:"rules,omitempty"`
	TechnicalName string                   `json:"technicalName"` // required
	Translated    interface{}              `json:"translated,omitempty"`
	Translations  []TaxRuleTypeTranslation `json:"translations,omitempty"`
	TypeName      string                   `json:"typeName"` // required
	UpdatedAt     time.Time                `json:"updatedAt,omitempty"`
}
