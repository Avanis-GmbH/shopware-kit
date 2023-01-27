package model

import "time"

type DeliveryTime struct {
	CreatedAt       time.Time                 `json:"createdAt,omitempty"`
	CustomFields    interface{}               `json:"customFields,omitempty"`
	Id              string                    `json:"id,omitempty"`
	Max             int64                     `json:"max"`  // required
	Min             int64                     `json:"min"`  // required
	Name            string                    `json:"name"` // required
	Products        []Product                 `json:"products,omitempty"`
	ShippingMethods []ShippingMethod          `json:"shippingMethods,omitempty"`
	Translated      interface{}               `json:"translated,omitempty"`
	Translations    []DeliveryTimeTranslation `json:"translations,omitempty"`
	Unit            string                    `json:"unit"` // required
	UpdatedAt       time.Time                 `json:"updatedAt,omitempty"`
}
