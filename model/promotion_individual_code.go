package model

import "time"

type PromotionIndividualCode struct {
	Payload     interface{} `json:"payload,omitempty"`
	Promotion   *Promotion  `json:"promotion,omitempty"`
	CreatedAt   time.Time   `json:"createdAt,omitempty"`
	UpdatedAt   time.Time   `json:"updatedAt,omitempty"`
	Id          string      `json:"id,omitempty"`
	PromotionId string      `json:"promotionId,omitempty"`
	Code        string      `json:"code,omitempty"`
}
