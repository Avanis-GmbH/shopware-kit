package model

type SalesChannelCurrency struct {
	Currency       *Currency     `json:"currency,omitempty"`
	CurrencyId     string        `json:"currencyId"` // required
	SalesChannel   *SalesChannel `json:"salesChannel,omitempty"`
	SalesChannelId string        `json:"salesChannelId"` // required
}
