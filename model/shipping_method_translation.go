package model

import "time"

type ShippingMethodTranslation struct {
	CreatedAt        *time.Time      `json:"createdAt,omitempty"`
	CustomFields     interface{}     `json:"customFields,omitempty"`
	Description      string          `json:"description,omitempty"`
	Language         *Language       `json:"language,omitempty"`
	LanguageId       string          `json:"languageId,omitempty"`
	Name             string          `json:"name,omitempty"`
	ShippingMethod   *ShippingMethod `json:"shippingMethod,omitempty"`
	ShippingMethodId string          `json:"shippingMethodId,omitempty"`
	TrackingUrl      string          `json:"trackingUrl,omitempty"`
	UpdatedAt        *time.Time      `json:"updatedAt,omitempty"`
}
