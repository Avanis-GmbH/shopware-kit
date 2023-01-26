package model

import "time"

type PaymentMethodTranslation struct {
	CreatedAt           time.Time      `json:"createdAt,omitempty"`
	CustomFields        interface{}    `json:"customFields,omitempty"`
	Description         string         `json:"description,omitempty"`
	DistinguishableName string         `json:"distinguishableName,omitempty"`
	Language            *Language      `json:"language,omitempty"`
	LanguageId          string         `json:"languageId,omitempty"`
	Name                string         `json:"name,omitempty"`
	PaymentMethod       *PaymentMethod `json:"paymentMethod,omitempty"`
	PaymentMethodId     string         `json:"paymentMethodId,omitempty"`
	UpdatedAt           time.Time      `json:"updatedAt,omitempty"`
}

type PaymentMethodTranslationCollection struct {
	EntityCollection

	Data []PaymentMethodTranslation `json:"data"`
}
