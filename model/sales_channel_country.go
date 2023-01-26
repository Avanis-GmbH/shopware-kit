package model

type SalesChannelCountry struct {
	Country        *Country      `json:"country,omitempty"`
	CountryId      string        `json:"countryId,omitempty"`
	SalesChannel   *SalesChannel `json:"salesChannel,omitempty"`
	SalesChannelId string        `json:"salesChannelId,omitempty"`
}

type SalesChannelCountryCollection struct {
	EntityCollection

	Data []SalesChannelCountry `json:"data"`
}
