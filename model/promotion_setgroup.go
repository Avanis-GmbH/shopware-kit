package model

import "time"

type PromotionSetgroup struct {
	CreatedAt     time.Time  `json:"createdAt,omitempty"`
	Id            string     `json:"id,omitempty"`
	PackagerKey   string     `json:"packagerKey,omitempty"` // required
	Promotion     *Promotion `json:"promotion,omitempty"`
	PromotionId   string     `json:"promotionId,omitempty"` // required
	SetGroupRules []Rule     `json:"setGroupRules,omitempty"`
	SorterKey     string     `json:"sorterKey,omitempty"` // required
	UpdatedAt     time.Time  `json:"updatedAt,omitempty"`
	Value         float64    `json:"value,omitempty"` // required
}
