package model

import "time"

type PromotionIndividualCode struct {
	Code        string      `json:"code"` // required
	CreatedAt   time.Time   `json:"createdAt,omitempty"`
	Id          string      `json:"id,omitempty"`
	Payload     interface{} `json:"payload,omitempty"`
	Promotion   *Promotion  `json:"promotion,omitempty"`
	PromotionId string      `json:"promotionId"` // required
	UpdatedAt   time.Time   `json:"updatedAt,omitempty"`
}
