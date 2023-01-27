package model

import "time"

type Currency struct {
	CountryRoundings               []CurrencyCountryRounding `json:"countryRoundings,omitempty"`
	CreatedAt                      time.Time                 `json:"createdAt,omitempty"`
	CustomFields                   interface{}               `json:"customFields,omitempty"`
	Factor                         float64                   `json:"factor"` // required
	Id                             string                    `json:"id,omitempty"`
	IsoCode                        string                    `json:"isoCode"` // required
	IsSystemDefault                bool                      `json:"isSystemDefault,omitempty"`
	ItemRounding                   Rounding                  `json:"itemRounding"` // required
	Name                           string                    `json:"name"`         // required
	Orders                         []Order                   `json:"orders,omitempty"`
	Position                       float64                   `json:"position,omitempty"`
	ProductExports                 []ProductExport           `json:"productExports,omitempty"`
	PromotionDiscountPrices        []PromotionDiscountPrices `json:"promotionDiscountPrices,omitempty"`
	SalesChannelDefaultAssignments []SalesChannel            `json:"salesChannelDefaultAssignments,omitempty"`
	SalesChannelDomains            []SalesChannelDomain      `json:"salesChannelDomains,omitempty"`
	SalesChannels                  []SalesChannel            `json:"salesChannels,omitempty"`
	ShortName                      string                    `json:"shortName"` // required
	Symbol                         string                    `json:"symbol"`    // required
	TaxFreeFrom                    float64                   `json:"taxFreeFrom"`
	TotalRounding                  Rounding                  `json:"totalRounding"` // required
	Translated                     interface{}               `json:"translated,omitempty"`
	Translations                   []CurrencyTranslation     `json:"translations,omitempty"`
	UpdatedAt                      time.Time                 `json:"updatedAt,omitempty"`
}

type Rounding struct {
	Decimals    int64 `json:"decimals"`    // required
	Intervals   int64 `json:"intervals"`   // required
	RoundForNet bool  `json:"roundForNet"` // required
}
