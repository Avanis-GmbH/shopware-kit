package model

import "time"

type ProductVisibility struct {
	CreatedAt        time.Time     `json:"createdAt,omitempty"`
	Id               string        `json:"id,omitempty"`
	Product          *Product      `json:"product,omitempty"`
	ProductId        string        `json:"productId,omitempty"`
	ProductVersionId string        `json:"productVersionId,omitempty"`
	SalesChannel     *SalesChannel `json:"salesChannel,omitempty"`
	SalesChannelId   string        `json:"salesChannelId,omitempty"`
	UpdatedAt        time.Time     `json:"updatedAt,omitempty"`
	Visibility       float64       `json:"visibility,omitempty"`
}

type ProductVisibilityCollection struct {
	EntityCollection

	Data []ProductVisibility `json:"data"`
}
