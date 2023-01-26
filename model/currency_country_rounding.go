package model

import "time"

type CurrencyCountryRounding struct {
	Country       *Country    `json:"country,omitempty"`
	CountryId     string      `json:"countryId,omitempty"`
	CreatedAt     time.Time   `json:"createdAt,omitempty"`
	Currency      *Currency   `json:"currency,omitempty"`
	CurrencyId    string      `json:"currencyId,omitempty"`
	Id            string      `json:"id,omitempty"`
	ItemRounding  interface{} `json:"itemRounding,omitempty"`
	TotalRounding interface{} `json:"totalRounding,omitempty"`
	UpdatedAt     time.Time   `json:"updatedAt,omitempty"`
}

type CurrencyCountryRoundingCollection struct {
	EntityCollection

	Data []CurrencyCountryRounding `json:"data"`
}
