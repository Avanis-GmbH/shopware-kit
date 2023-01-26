package model

import "time"

type ProductExport struct {
	AccessKey                string              `json:"accessKey,omitempty"`
	BodyTemplate             string              `json:"bodyTemplate,omitempty"`
	CreatedAt                time.Time           `json:"createdAt,omitempty"`
	Currency                 *Currency           `json:"currency,omitempty"`
	CurrencyId               string              `json:"currencyId,omitempty"`
	Encoding                 string              `json:"encoding,omitempty"`
	FileFormat               string              `json:"fileFormat,omitempty"`
	FileName                 string              `json:"fileName,omitempty"`
	FooterTemplate           string              `json:"footerTemplate,omitempty"`
	GenerateByCronjob        bool                `json:"generateByCronjob,omitempty"`
	GeneratedAt              time.Time           `json:"generatedAt,omitempty"`
	HeaderTemplate           string              `json:"headerTemplate,omitempty"`
	Id                       string              `json:"id,omitempty"`
	IncludeVariants          bool                `json:"includeVariants,omitempty"`
	Interval                 float64             `json:"interval,omitempty"`
	PausedSchedule           bool                `json:"pausedSchedule,omitempty"`
	ProductStream            *ProductStream      `json:"productStream,omitempty"`
	ProductStreamId          string              `json:"productStreamId,omitempty"`
	SalesChannel             *SalesChannel       `json:"salesChannel,omitempty"`
	SalesChannelDomain       *SalesChannelDomain `json:"salesChannelDomain,omitempty"`
	SalesChannelDomainId     string              `json:"salesChannelDomainId,omitempty"`
	SalesChannelId           string              `json:"salesChannelId,omitempty"`
	StorefrontSalesChannel   *SalesChannel       `json:"storefrontSalesChannel,omitempty"`
	StorefrontSalesChannelId string              `json:"storefrontSalesChannelId,omitempty"`
	UpdatedAt                time.Time           `json:"updatedAt,omitempty"`
}
