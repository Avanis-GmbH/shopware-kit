package model

import "time"

type DeliveryTimeTranslation struct {
	CreatedAt      *time.Time    `json:"createdAt,omitempty"`
	CustomFields   interface{}   `json:"customFields,omitempty"`
	DeliveryTime   *DeliveryTime `json:"deliveryTime,omitempty"`
	DeliveryTimeId string        `json:"deliveryTimeId,omitempty"`
	Language       *Language     `json:"language,omitempty"`
	LanguageId     string        `json:"languageId,omitempty"`
	Name           string        `json:"name,omitempty"`
	UpdatedAt      *time.Time    `json:"updatedAt,omitempty"`
}
