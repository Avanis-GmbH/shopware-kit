package model

import "time"

type OrderDeliveryPosition struct {
	CreatedAt              *time.Time     `json:"createdAt,omitempty"`
	CustomFields           interface{}    `json:"customFields,omitempty"`
	Id                     string         `json:"id,omitempty"`
	OrderDelivery          *OrderDelivery `json:"orderDelivery,omitempty"`
	OrderDeliveryId        string         `json:"orderDeliveryId,omitempty"` // required
	OrderDeliveryVersionId string         `json:"orderDeliveryVersionId,omitempty"`
	OrderLineItem          *OrderLineItem `json:"orderLineItem,omitempty"`
	OrderLineItemId        string         `json:"orderLineItemId,omitempty"` // required
	OrderLineItemVersionId string         `json:"orderLineItemVersionId,omitempty"`
	Price                  OrderPrice     `json:"price,omitempty"`
	Quantity               float64        `json:"quantity,omitempty"`
	TotalPrice             float64        `json:"totalPrice,omitempty"`
	UnitPrice              float64        `json:"unitPrice,omitempty"`
	UpdatedAt              *time.Time     `json:"updatedAt,omitempty"`
	VersionId              string         `json:"versionId,omitempty"`
}
