package model

import "time"

type TaxRuleType struct {
	CreatedAt     time.Time                `json:"createdAt,omitempty"`
	Id            string                   `json:"id,omitempty"`
	Position      int64                    `json:"position,omitempty"` // required
	Rules         []TaxRule                `json:"rules,omitempty"`
	TechnicalName string                   `json:"technicalName,omitempty"` // required
	Translated    interface{}              `json:"translated,omitempty"`
	Translations  []TaxRuleTypeTranslation `json:"translations,omitempty"`
	TypeName      string                   `json:"typeName,omitempty"` // required
	UpdatedAt     time.Time                `json:"updatedAt,omitempty"`
}
