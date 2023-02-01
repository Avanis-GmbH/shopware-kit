package model

import "time"

type Currency struct {
	CountryRoundings               []CurrencyCountryRounding `json:"countryRoundings,omitempty"`
	CreatedAt                      *time.Time                `json:"createdAt,omitempty"`
	CustomFields                   interface{}               `json:"customFields,omitempty"`
	Factor                         float64                   `json:"factor,omitempty"` // required
	Id                             string                    `json:"id,omitempty"`
	IsoCode                        string                    `json:"isoCode,omitempty"` // required
	IsSystemDefault                bool                      `json:"isSystemDefault,omitempty"`
	ItemRounding                   Rounding                  `json:"itemRounding,omitempty"` // required
	Name                           string                    `json:"name,omitempty"`         // required
	Orders                         []Order                   `json:"orders,omitempty"`
	Position                       float64                   `json:"position,omitempty"`
	ProductExports                 []ProductExport           `json:"productExports,omitempty"`
	PromotionDiscountPrices        []PromotionDiscountPrices `json:"promotionDiscountPrices,omitempty"`
	SalesChannelDefaultAssignments []SalesChannel            `json:"salesChannelDefaultAssignments,omitempty"`
	SalesChannelDomains            []SalesChannelDomain      `json:"salesChannelDomains,omitempty"`
	SalesChannels                  []SalesChannel            `json:"salesChannels,omitempty"`
	ShortName                      string                    `json:"shortName,omitempty"` // required
	Symbol                         string                    `json:"symbol,omitempty"`    // required
	TaxFreeFrom                    float64                   `json:"taxFreeFrom,omitempty"`
	TotalRounding                  Rounding                  `json:"totalRounding,omitempty"` // required
	Translated                     interface{}               `json:"translated,omitempty"`
	Translations                   []CurrencyTranslation     `json:"translations,omitempty"`
	UpdatedAt                      *time.Time                `json:"updatedAt,omitempty"`
}

type Rounding struct {
	Decimals    int64 `json:"decimals,omitempty"`    // required
	Intervals   int64 `json:"intervals,omitempty"`   // required
	RoundForNet bool  `json:"roundForNet,omitempty"` // required
}
