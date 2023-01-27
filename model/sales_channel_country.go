package model

type SalesChannelCountry struct {
	Country        *Country      `json:"country,omitempty"`
	CountryId      string        `json:"countryId"` // required
	SalesChannel   *SalesChannel `json:"salesChannel,omitempty"`
	SalesChannelId string        `json:"salesChannelId"` // required
}
