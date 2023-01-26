package model

import "time"

type PromotionSalesChannel struct {
	CreatedAt      time.Time     `json:"createdAt,omitempty"`
	Id             string        `json:"id,omitempty"`
	Priority       float64       `json:"priority,omitempty"`
	Promotion      *Promotion    `json:"promotion,omitempty"`
	PromotionId    string        `json:"promotionId,omitempty"`
	SalesChannel   *SalesChannel `json:"salesChannel,omitempty"`
	SalesChannelId string        `json:"salesChannelId,omitempty"`
	UpdatedAt      time.Time     `json:"updatedAt,omitempty"`
}
