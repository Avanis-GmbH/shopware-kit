package model

import "time"

type ProductExport struct {
	AccessKey                string              `json:"accessKey,omitempty"` // required
	BodyTemplate             string              `json:"bodyTemplate,omitempty"`
	CreatedAt                *time.Time          `json:"createdAt,omitempty"`
	Currency                 *Currency           `json:"currency,omitempty"`
	CurrencyId               string              `json:"currencyId,omitempty"` // required
	Encoding                 string              `json:"encoding,omitempty"`   // required
	FileFormat               string              `json:"fileFormat,omitempty"` // required
	FileName                 string              `json:"fileName,omitempty"`   // required
	FooterTemplate           string              `json:"footerTemplate,omitempty"`
	GenerateByCronjob        bool                `json:"generateByCronjob,omitempty"` // required
	GeneratedAt              *time.Time          `json:"generatedAt,omitempty"`
	HeaderTemplate           string              `json:"headerTemplate,omitempty"`
	Id                       string              `json:"id,omitempty"`
	IncludeVariants          bool                `json:"includeVariants,omitempty"`
	Interval                 float64             `json:"interval,omitempty"` // required
	PausedSchedule           bool                `json:"pausedSchedule,omitempty"`
	ProductStream            *ProductStream      `json:"productStream,omitempty"`
	ProductStreamId          string              `json:"productStreamId,omitempty"` // required
	SalesChannel             *SalesChannel       `json:"salesChannel,omitempty"`
	SalesChannelDomain       *SalesChannelDomain `json:"salesChannelDomain,omitempty"`
	SalesChannelDomainId     string              `json:"salesChannelDomainId,omitempty"` // required
	SalesChannelId           string              `json:"salesChannelId,omitempty"`       // required
	StorefrontSalesChannel   *SalesChannel       `json:"storefrontSalesChannel,omitempty"`
	StorefrontSalesChannelId string              `json:"storefrontSalesChannelId,omitempty"` // required
	UpdatedAt                *time.Time          `json:"updatedAt,omitempty"`
}
