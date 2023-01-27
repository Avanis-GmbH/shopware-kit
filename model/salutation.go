package model

import "time"

type Salutation struct {
	CreatedAt            time.Time               `json:"createdAt,omitempty"`
	CustomerAddresses    []CustomerAddress       `json:"customerAddresses,omitempty"`
	Customers            []Customer              `json:"customers,omitempty"`
	CustomFields         interface{}             `json:"customFields,omitempty"`
	DisplayName          string                  `json:"displayName,omitempty"` // required
	Id                   string                  `json:"id,omitempty"`
	LetterName           string                  `json:"letterName,omitempty"` // required
	NewsletterRecipients []NewsletterRecipient   `json:"newsletterRecipients,omitempty"`
	OrderAddresses       []OrderAddress          `json:"orderAddresses,omitempty"`
	OrderCustomers       []OrderCustomer         `json:"orderCustomers,omitempty"`
	SalutationKey        string                  `json:"salutationKey,omitempty"` // required
	Translated           interface{}             `json:"translated,omitempty"`
	Translations         []SalutationTranslation `json:"translations,omitempty"`
	UpdatedAt            time.Time               `json:"updatedAt,omitempty"`
}
