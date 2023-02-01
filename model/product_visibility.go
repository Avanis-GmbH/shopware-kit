package model

import "time"

type ProductVisibility struct {
	CreatedAt        *time.Time    `json:"createdAt,omitempty"`
	Id               string        `json:"id,omitempty"`
	Product          *Product      `json:"product,omitempty"`
	ProductId        string        `json:"productId,omitempty"` // required
	ProductVersionId string        `json:"productVersionId,omitempty"`
	SalesChannel     *SalesChannel `json:"salesChannel,omitempty"`
	SalesChannelId   string        `json:"salesChannelId,omitempty"` // required
	UpdatedAt        *time.Time    `json:"updatedAt,omitempty"`
	Visibility       int64         `json:"visibility,omitempty"` // required
}
