package model

import "time"

type ProductExport struct {
	AccessKey                string              `json:"accessKey"` // required
	BodyTemplate             string              `json:"bodyTemplate,omitempty"`
	CreatedAt                time.Time           `json:"createdAt,omitempty"`
	Currency                 *Currency           `json:"currency,omitempty"`
	CurrencyId               string              `json:"currencyId"` // required
	Encoding                 string              `json:"encoding"`   // required
	FileFormat               string              `json:"fileFormat"` // required
	FileName                 string              `json:"fileName"`   // required
	FooterTemplate           string              `json:"footerTemplate,omitempty"`
	GenerateByCronjob        bool                `json:"generateByCronjob"` // required
	GeneratedAt              time.Time           `json:"generatedAt,omitempty"`
	HeaderTemplate           string              `json:"headerTemplate,omitempty"`
	Id                       string              `json:"id,omitempty"`
	IncludeVariants          bool                `json:"includeVariants,omitempty"`
	Interval                 float64             `json:"interval"` // required
	PausedSchedule           bool                `json:"pausedSchedule,omitempty"`
	ProductStream            *ProductStream      `json:"productStream,omitempty"`
	ProductStreamId          string              `json:"productStreamId"` // required
	SalesChannel             *SalesChannel       `json:"salesChannel,omitempty"`
	SalesChannelDomain       *SalesChannelDomain `json:"salesChannelDomain,omitempty"`
	SalesChannelDomainId     string              `json:"salesChannelDomainId"` // required
	SalesChannelId           string              `json:"salesChannelId"`       // required
	StorefrontSalesChannel   *SalesChannel       `json:"storefrontSalesChannel,omitempty"`
	StorefrontSalesChannelId string              `json:"storefrontSalesChannelId"` // required
	UpdatedAt                time.Time           `json:"updatedAt,omitempty"`
}
