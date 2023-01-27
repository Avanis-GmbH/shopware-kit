package model

import "time"

type OrderDeliveryPosition struct {
	CreatedAt              time.Time      `json:"createdAt,omitempty"`
	CustomFields           interface{}    `json:"customFields,omitempty"`
	Id                     string         `json:"id,omitempty"`
	OrderDelivery          *OrderDelivery `json:"orderDelivery,omitempty"`
	OrderDeliveryId        string         `json:"orderDeliveryId"` // required
	OrderDeliveryVersionId string         `json:"orderDeliveryVersionId,omitempty"`
	OrderLineItem          *OrderLineItem `json:"orderLineItem,omitempty"`
	OrderLineItemId        string         `json:"orderLineItemId"` // required
	OrderLineItemVersionId string         `json:"orderLineItemVersionId,omitempty"`
	Price                  OrderPrice     `json:"price,omitempty"`
	Quantity               float64        `json:"quantity"`
	TotalPrice             float64        `json:"totalPrice"`
	UnitPrice              float64        `json:"unitPrice"`
	UpdatedAt              time.Time      `json:"updatedAt,omitempty"`
	VersionId              string         `json:"versionId,omitempty"`
}
