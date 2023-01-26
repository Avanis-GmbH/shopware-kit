package model

import "time"

type Tax struct {
	CreatedAt       time.Time        `json:"createdAt,omitempty"`
	CustomFields    interface{}      `json:"customFields,omitempty"`
	Id              string           `json:"id,omitempty"`
	Name            string           `json:"name,omitempty"`
	Position        float64          `json:"position,omitempty"`
	Products        []Product        `json:"products,omitempty"`
	Rules           []TaxRule        `json:"rules,omitempty"`
	ShippingMethods []ShippingMethod `json:"shippingMethods,omitempty"`
	TaxRate         float64          `json:"taxRate,omitempty"`
	UpdatedAt       time.Time        `json:"updatedAt,omitempty"`
}

type TaxCollection struct {
	EntityCollection

	Data []Tax `json:"data"`
}
