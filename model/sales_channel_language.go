package model

type SalesChannelLanguage struct {
	Language       *Language     `json:"language,omitempty"`
	LanguageId     string        `json:"languageId,omitempty"` // required
	SalesChannel   *SalesChannel `json:"salesChannel,omitempty"`
	SalesChannelId string        `json:"salesChannelId,omitempty"` // required
}
