package model

type SalesChannelLanguage struct {
	Language       *Language     `json:"language,omitempty"`
	LanguageId     string        `json:"languageId"` // required
	SalesChannel   *SalesChannel `json:"salesChannel,omitempty"`
	SalesChannelId string        `json:"salesChannelId"` // required
}
