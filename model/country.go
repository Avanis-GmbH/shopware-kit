package model

import "time"

type Country struct {
	Active                         bool                      `json:"active,omitempty"`
	CheckVatIdPattern              bool                      `json:"checkVatIdPattern,omitempty"`
	CompanyTax                     interface{}               `json:"companyTax,omitempty"`
	CompanyTaxFree                 bool                      `json:"companyTaxFree,omitempty"`
	CreatedAt                      time.Time                 `json:"createdAt,omitempty"`
	CurrencyCountryRoundings       []CurrencyCountryRounding `json:"currencyCountryRoundings,omitempty"`
	CustomerAddresses              []CustomerAddress         `json:"customerAddresses,omitempty"`
	CustomerTax                    interface{}               `json:"customerTax,omitempty"`
	CustomFields                   interface{}               `json:"customFields,omitempty"`
	DisplayStateInRegistration     bool                      `json:"displayStateInRegistration,omitempty"`
	ForceStateInRegistration       bool                      `json:"forceStateInRegistration,omitempty"`
	Id                             string                    `json:"id,omitempty"`
	Iso                            string                    `json:"iso,omitempty"`
	Iso3                           string                    `json:"iso3,omitempty"`
	Name                           string                    `json:"name,omitempty"`
	OrderAddresses                 []OrderAddress            `json:"orderAddresses,omitempty"`
	Position                       float64                   `json:"position,omitempty"`
	SalesChannelDefaultAssignments []SalesChannel            `json:"salesChannelDefaultAssignments,omitempty"`
	SalesChannels                  []SalesChannel            `json:"salesChannels,omitempty"`
	ShippingAvailable              bool                      `json:"shippingAvailable,omitempty"`
	States                         []CountryState            `json:"states,omitempty"`
	TaxFree                        bool                      `json:"taxFree,omitempty"`
	TaxRules                       []TaxRule                 `json:"taxRules,omitempty"`
	Translated                     interface{}               `json:"translated,omitempty"`
	Translations                   []CountryTranslation      `json:"translations,omitempty"`
	UpdatedAt                      time.Time                 `json:"updatedAt,omitempty"`
	VatIdPattern                   string                    `json:"vatIdPattern,omitempty"`
	VatIdRequired                  bool                      `json:"vatIdRequired,omitempty"`
}
