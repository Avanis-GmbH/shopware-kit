package model

import "time"

type ShippingMethodPrice struct {
	Calculation       float64         `json:"calculation,omitempty"`
	CalculationRule   *Rule           `json:"calculationRule,omitempty"`
	CalculationRuleId string          `json:"calculationRuleId,omitempty"`
	CreatedAt         time.Time       `json:"createdAt,omitempty"`
	CurrencyPrice     interface{}     `json:"currencyPrice,omitempty"`
	CustomFields      interface{}     `json:"customFields,omitempty"`
	Id                string          `json:"id,omitempty"`
	QuantityEnd       float64         `json:"quantityEnd,omitempty"`
	QuantityStart     float64         `json:"quantityStart,omitempty"`
	Rule              *Rule           `json:"rule,omitempty"`
	RuleId            string          `json:"ruleId,omitempty"`
	ShippingMethod    *ShippingMethod `json:"shippingMethod,omitempty"`
	ShippingMethodId  string          `json:"shippingMethodId,omitempty"`
	UpdatedAt         time.Time       `json:"updatedAt,omitempty"`
}
