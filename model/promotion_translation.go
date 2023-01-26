package model

import "time"

type PromotionTranslation struct {
	CreatedAt    time.Time   `json:"createdAt,omitempty"`
	CustomFields interface{} `json:"customFields,omitempty"`
	Language     *Language   `json:"language,omitempty"`
	LanguageId   string      `json:"languageId,omitempty"`
	Name         string      `json:"name,omitempty"`
	Promotion    *Promotion  `json:"promotion,omitempty"`
	PromotionId  string      `json:"promotionId,omitempty"`
	UpdatedAt    time.Time   `json:"updatedAt,omitempty"`
}

type PromotionTranslationCollection struct {
	EntityCollection

	Data []PromotionTranslation `json:"data"`
}
