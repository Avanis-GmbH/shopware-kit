package model

import "time"

type PromotionSalesChannel struct {
	CreatedAt      time.Time     `json:"createdAt,omitempty"`
	Id             string        `json:"id,omitempty"`
	Priority       int64         `json:"priority"` // required
	Promotion      *Promotion    `json:"promotion,omitempty"`
	PromotionId    string        `json:"promotionId"` // required
	SalesChannel   *SalesChannel `json:"salesChannel,omitempty"`
	SalesChannelId string        `json:"salesChannelId"` // required
	UpdatedAt      time.Time     `json:"updatedAt,omitempty"`
}
