package model

import "time"

type Salutation struct {
	CreatedAt            time.Time               `json:"createdAt,omitempty"`
	CustomerAddresses    []CustomerAddress       `json:"customerAddresses,omitempty"`
	Customers            []Customer              `json:"customers,omitempty"`
	CustomFields         interface{}             `json:"customFields,omitempty"`
	DisplayName          string                  `json:"displayName,omitempty"`
	Id                   string                  `json:"id,omitempty"`
	LetterName           string                  `json:"letterName,omitempty"`
	NewsletterRecipients []NewsletterRecipient   `json:"newsletterRecipients,omitempty"`
	OrderAddresses       []OrderAddress          `json:"orderAddresses,omitempty"`
	OrderCustomers       []OrderCustomer         `json:"orderCustomers,omitempty"`
	SalutationKey        string                  `json:"salutationKey,omitempty"`
	Translated           interface{}             `json:"translated,omitempty"`
	Translations         []SalutationTranslation `json:"translations,omitempty"`
	UpdatedAt            time.Time               `json:"updatedAt,omitempty"`
}

type SalutationCollection struct {
	EntityCollection

	Data []Salutation `json:"data"`
}
