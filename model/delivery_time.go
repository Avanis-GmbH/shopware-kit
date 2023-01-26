package model

import "time"

type DeliveryTime struct {
	CreatedAt       time.Time                 `json:"createdAt,omitempty"`
	CustomFields    interface{}               `json:"customFields,omitempty"`
	Id              string                    `json:"id,omitempty"`
	Max             float64                   `json:"max,omitempty"`
	Min             float64                   `json:"min,omitempty"`
	Name            string                    `json:"name,omitempty"`
	Products        []Product                 `json:"products,omitempty"`
	ShippingMethods []ShippingMethod          `json:"shippingMethods,omitempty"`
	Translated      interface{}               `json:"translated,omitempty"`
	Translations    []DeliveryTimeTranslation `json:"translations,omitempty"`
	Unit            string                    `json:"unit,omitempty"`
	UpdatedAt       time.Time                 `json:"updatedAt,omitempty"`
}
