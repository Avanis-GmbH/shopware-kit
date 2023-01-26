package model

import "time"

type Currency struct {
	CountryRoundings               []CurrencyCountryRounding `json:"countryRoundings,omitempty"`
	CreatedAt                      time.Time                 `json:"createdAt,omitempty"`
	CustomFields                   interface{}               `json:"customFields,omitempty"`
	Factor                         float64                   `json:"factor,omitempty"`
	Id                             string                    `json:"id,omitempty"`
	IsoCode                        string                    `json:"isoCode,omitempty"`
	IsSystemDefault                bool                      `json:"isSystemDefault,omitempty"`
	ItemRounding                   interface{}               `json:"itemRounding,omitempty"`
	Name                           string                    `json:"name,omitempty"`
	Orders                         []Order                   `json:"orders,omitempty"`
	Position                       float64                   `json:"position,omitempty"`
	ProductExports                 []ProductExport           `json:"productExports,omitempty"`
	PromotionDiscountPrices        []PromotionDiscountPrices `json:"promotionDiscountPrices,omitempty"`
	SalesChannelDefaultAssignments []SalesChannel            `json:"salesChannelDefaultAssignments,omitempty"`
	SalesChannelDomains            []SalesChannelDomain      `json:"salesChannelDomains,omitempty"`
	SalesChannels                  []SalesChannel            `json:"salesChannels,omitempty"`
	ShortName                      string                    `json:"shortName,omitempty"`
	Symbol                         string                    `json:"symbol,omitempty"`
	TaxFreeFrom                    float64                   `json:"taxFreeFrom,omitempty"`
	TotalRounding                  interface{}               `json:"totalRounding,omitempty"`
	Translated                     interface{}               `json:"translated,omitempty"`
	Translations                   []CurrencyTranslation     `json:"translations,omitempty"`
	UpdatedAt                      time.Time                 `json:"updatedAt,omitempty"`
}
