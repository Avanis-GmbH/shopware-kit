package model

import "time"

type TaxRule struct {
	Country       *Country     `json:"country,omitempty"`
	CountryId     string       `json:"countryId"` // required
	CreatedAt     time.Time    `json:"createdAt,omitempty"`
	Data          interface{}  `json:"data,omitempty"`
	Id            string       `json:"id,omitempty"`
	Tax           *Tax         `json:"tax,omitempty"`
	TaxId         string       `json:"taxId,omitempty"`
	TaxRate       float64      `json:"taxRate"`       // required
	TaxRuleTypeId string       `json:"taxRuleTypeId"` // required
	Type          *TaxRuleType `json:"type,omitempty"`
	UpdatedAt     time.Time    `json:"updatedAt,omitempty"`
}
