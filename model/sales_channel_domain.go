package model

import "time"

type SalesChannelDomain struct {
	CreatedAt                   time.Time       `json:"createdAt,omitempty"`
	Currency                    *Currency       `json:"currency,omitempty"`
	CurrencyId                  string          `json:"currencyId,omitempty"` // required
	CustomFields                interface{}     `json:"customFields,omitempty"`
	HreflangUseOnlyLocale       bool            `json:"hreflangUseOnlyLocale,omitempty"`
	Id                          string          `json:"id,omitempty"`
	Language                    *Language       `json:"language,omitempty"`
	LanguageId                  string          `json:"languageId,omitempty"` // required
	ProductExports              []ProductExport `json:"productExports,omitempty"`
	SalesChannel                *SalesChannel   `json:"salesChannel,omitempty"`
	SalesChannelDefaultHreflang *SalesChannel   `json:"salesChannelDefaultHreflang,omitempty"`
	SalesChannelId              string          `json:"salesChannelId,omitempty"` // required
	SnippetSet                  *SnippetSet     `json:"snippetSet,omitempty"`
	SnippetSetId                string          `json:"snippetSetId,omitempty"` // required
	UpdatedAt                   time.Time       `json:"updatedAt,omitempty"`
	Url                         string          `json:"url,omitempty"` // required
}
