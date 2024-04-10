package model

import "time"

type Country struct {
	Active                         bool                      `json:"active"`
	AddressFormat                  string                    `json:"addressFormat,omitempty"` // required
	CheckVatIdPattern              bool                      `json:"checkVatIdPattern"`
	CompanyTax                     interface{}               `json:"companyTax,omitempty"`
	CompanyTaxFree                 bool                      `json:"companyTaxFree"`
	CreatedAt                      *time.Time                `json:"createdAt,omitempty"`
	CurrencyCountryRoundings       []CurrencyCountryRounding `json:"currencyCountryRoundings,omitempty"`
	CustomerAddresses              []CustomerAddress         `json:"customerAddresses,omitempty"`
	CustomerTax                    *CTax                     `json:"customerTax,omitempty"`
	CustomFields                   *CTax                     `json:"customFields,omitempty"`
	DisplayStateInRegistration     bool                      `json:"displayStateInRegistration"`
	ForceStateInRegistration       bool                      `json:"forceStateInRegistration"`
	Id                             string                    `json:"id,omitempty"`
	Iso                            string                    `json:"iso,omitempty"`
	Iso3                           string                    `json:"iso3,omitempty"`
	Name                           string                    `json:"name,omitempty"` // required
	OrderAddresses                 []OrderAddress            `json:"orderAddresses,omitempty"`
	Position                       float64                   `json:"position,omitempty"`
	SalesChannelDefaultAssignments []SalesChannel            `json:"salesChannelDefaultAssignments,omitempty"`
	SalesChannels                  []SalesChannel            `json:"salesChannels,omitempty"`
	ShippingAvailable              bool                      `json:"shippingAvailable"`
	States                         []CountryState            `json:"states,omitempty"`
	TaxFree                        bool                      `json:"taxFree"`
	TaxRules                       []TaxRule                 `json:"taxRules,omitempty"`
	Translated                     interface{}               `json:"translated,omitempty"`
	Translations                   []CountryTranslation      `json:"translations,omitempty"`
	UpdatedAt                      *time.Time                `json:"updatedAt,omitempty"`
	VatIdPattern                   string                    `json:"vatIdPattern,omitempty"`
	VatIdRequired                  bool                      `json:"vatIdRequired"`
}
