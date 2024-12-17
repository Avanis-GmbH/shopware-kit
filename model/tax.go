package model

import "time"

type Tax struct {
	CreatedAt       *time.Time       `json:"createdAt,omitempty"`
	CustomFields    interface{}      `json:"customFields,omitempty"`
	Id              string           `json:"id,omitempty"`
	Name            string           `json:"name,omitempty"`     // required
	Position        float64          `json:"position,omitempty"` // required
	Products        []Product        `json:"products,omitempty"`
	Rules           []TaxRule        `json:"rules,omitempty"`
	ShippingMethods []ShippingMethod `json:"shippingMethods,omitempty"`
	TaxRate         float64          `json:"taxRate,omitempty"` // required
	UpdatedAt       *time.Time       `json:"updatedAt,omitempty"`
}

type CTax struct {
	Enabled    *bool   `json:"enabled,omitempty"`    // required
	CurrencyId string  `json:"currencyId,omitempty"` // required
	Amount     float64 `json:"amount,omitempty"`     // required
}
