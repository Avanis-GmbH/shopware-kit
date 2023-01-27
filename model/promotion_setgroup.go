package model

import "time"

type PromotionSetgroup struct {
	CreatedAt     time.Time  `json:"createdAt,omitempty"`
	Id            string     `json:"id,omitempty"`
	PackagerKey   string     `json:"packagerKey"` // required
	Promotion     *Promotion `json:"promotion,omitempty"`
	PromotionId   string     `json:"promotionId"` // required
	SetGroupRules []Rule     `json:"setGroupRules,omitempty"`
	SorterKey     string     `json:"sorterKey"` // required
	UpdatedAt     time.Time  `json:"updatedAt,omitempty"`
	Value         float64    `json:"value"` // required
}
