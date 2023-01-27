package model

import "time"

type MainCategory struct {
	CategoryId        string        `json:"categoryId"` // required
	CategoryVersionId string        `json:"categoryVersionId,omitempty"`
	SalesChannelId    string        `json:"salesChannelId"` // required
	Product           *Product      `json:"product,omitempty"`
	Category          *Category     `json:"category,omitempty"`
	SalesChannel      *SalesChannel `json:"salesChannel,omitempty"`
	CreatedAt         time.Time     `json:"createdAt,omitempty"`
	ProductVersionId  string        `json:"productVersionId,omitempty"`
	UpdatedAt         time.Time     `json:"updatedAt,omitempty"`
	ProductId         string        `json:"productId"` // required
	Id                string        `json:"id,omitempty"`
}
