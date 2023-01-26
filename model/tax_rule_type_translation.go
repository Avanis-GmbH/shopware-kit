package model

import "time"

type TaxRuleTypeTranslation struct {
	CreatedAt     time.Time    `json:"createdAt,omitempty"`
	Language      *Language    `json:"language,omitempty"`
	LanguageId    string       `json:"languageId,omitempty"`
	TaxRuleType   *TaxRuleType `json:"taxRuleType,omitempty"`
	TaxRuleTypeId string       `json:"taxRuleTypeId,omitempty"`
	TypeName      string       `json:"typeName,omitempty"`
	UpdatedAt     time.Time    `json:"updatedAt,omitempty"`
}
