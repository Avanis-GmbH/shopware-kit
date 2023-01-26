package model

import "time"

type OrderLineItem struct {
	Children               []OrderLineItem         `json:"children,omitempty"`
	Cover                  *Media                  `json:"cover,omitempty"`
	CoverId                string                  `json:"coverId,omitempty"`
	CreatedAt              time.Time               `json:"createdAt,omitempty"`
	CustomFields           interface{}             `json:"customFields,omitempty"`
	Description            string                  `json:"description,omitempty"`
	Good                   bool                    `json:"good,omitempty"`
	Id                     string                  `json:"id,omitempty"`
	Identifier             string                  `json:"identifier,omitempty"`
	Label                  string                  `json:"label,omitempty"`
	Order                  *Order                  `json:"order,omitempty"`
	OrderDeliveryPositions []OrderDeliveryPosition `json:"orderDeliveryPositions,omitempty"`
	OrderId                string                  `json:"orderId,omitempty"`
	OrderVersionId         string                  `json:"orderVersionId,omitempty"`
	Parent                 *OrderLineItem          `json:"parent,omitempty"`
	ParentId               string                  `json:"parentId,omitempty"`
	ParentVersionId        string                  `json:"parentVersionId,omitempty"`
	Payload                interface{}             `json:"payload,omitempty"`
	Position               float64                 `json:"position,omitempty"`
	Price                  interface{}             `json:"price,omitempty"`
	PriceDefinition        interface{}             `json:"priceDefinition,omitempty"`
	Product                *Product                `json:"product,omitempty"`
	ProductId              string                  `json:"productId,omitempty"`
	ProductVersionId       string                  `json:"productVersionId,omitempty"`
	Quantity               float64                 `json:"quantity,omitempty"`
	ReferencedId           string                  `json:"referencedId,omitempty"`
	Removable              bool                    `json:"removable,omitempty"`
	Stackable              bool                    `json:"stackable,omitempty"`
	TotalPrice             float64                 `json:"totalPrice,omitempty"`
	Type                   string                  `json:"type,omitempty"`
	UnitPrice              float64                 `json:"unitPrice,omitempty"`
	UpdatedAt              time.Time               `json:"updatedAt,omitempty"`
	VersionId              string                  `json:"versionId,omitempty"`
}

type OrderLineItemCollection struct {
	EntityCollection

	Data []OrderLineItem `json:"data"`
}
