package model

import "time"

type PromotionSetgroup struct {
	CreatedAt     time.Time  `json:"createdAt,omitempty"`
	Id            string     `json:"id,omitempty"`
	PackagerKey   string     `json:"packagerKey,omitempty"`
	Promotion     *Promotion `json:"promotion,omitempty"`
	PromotionId   string     `json:"promotionId,omitempty"`
	SetGroupRules []Rule     `json:"setGroupRules,omitempty"`
	SorterKey     string     `json:"sorterKey,omitempty"`
	UpdatedAt     time.Time  `json:"updatedAt,omitempty"`
	Value         float64    `json:"value,omitempty"`
}

type PromotionSetgroupCollection struct {
	EntityCollection

	Data []PromotionSetgroup `json:"data"`
}
