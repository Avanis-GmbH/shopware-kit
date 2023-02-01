package model

import "time"

type PromotionSalesChannel struct {
	CreatedAt      *time.Time    `json:"createdAt,omitempty"`
	Id             string        `json:"id,omitempty"`
	Priority       int64         `json:"priority,omitempty"` // required
	Promotion      *Promotion    `json:"promotion,omitempty"`
	PromotionId    string        `json:"promotionId,omitempty"` // required
	SalesChannel   *SalesChannel `json:"salesChannel,omitempty"`
	SalesChannelId string        `json:"salesChannelId,omitempty"` // required
	UpdatedAt      *time.Time    `json:"updatedAt,omitempty"`
}
