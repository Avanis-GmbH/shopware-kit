package model

import "time"

type CurrencyCountryRounding struct {
	Country       *Country  `json:"country,omitempty"`
	CountryId     string    `json:"countryId"` // required
	CreatedAt     time.Time `json:"createdAt,omitempty"`
	Currency      *Currency `json:"currency,omitempty"`
	CurrencyId    string    `json:"currencyId"` // required
	Id            string    `json:"id,omitempty"`
	ItemRounding  *Rounding `json:"itemRounding,omitempty"`
	TotalRounding *Rounding `json:"totalRounding,omitempty"`
	UpdatedAt     time.Time `json:"updatedAt,omitempty"`
}
