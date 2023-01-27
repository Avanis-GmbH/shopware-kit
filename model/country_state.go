package model

import "time"

type CountryState struct {
	Active            bool                      `json:"active,omitempty"`
	Country           *Country                  `json:"country,omitempty"`
	CountryId         string                    `json:"countryId,omitempty"` // required
	CreatedAt         time.Time                 `json:"createdAt,omitempty"`
	CustomerAddresses []CustomerAddress         `json:"customerAddresses,omitempty"`
	CustomFields      interface{}               `json:"customFields,omitempty"`
	Id                string                    `json:"id,omitempty"`
	Name              string                    `json:"name,omitempty"` // required
	OrderAddresses    []OrderAddress            `json:"orderAddresses,omitempty"`
	Position          float64                   `json:"position,omitempty"`
	ShortCode         string                    `json:"shortCode,omitempty"` // required
	Translated        interface{}               `json:"translated,omitempty"`
	Translations      []CountryStateTranslation `json:"translations,omitempty"`
	UpdatedAt         time.Time                 `json:"updatedAt,omitempty"`
}
