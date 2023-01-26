package model

import "time"

type SalesChannelTranslation struct {
	CreatedAt           time.Time     `json:"createdAt,omitempty"`
	CustomFields        interface{}   `json:"customFields,omitempty"`
	HomeEnabled         bool          `json:"homeEnabled,omitempty"`
	HomeKeywords        string        `json:"homeKeywords,omitempty"`
	HomeMetaDescription string        `json:"homeMetaDescription,omitempty"`
	HomeMetaTitle       string        `json:"homeMetaTitle,omitempty"`
	HomeName            string        `json:"homeName,omitempty"`
	HomeSlotConfig      interface{}   `json:"homeSlotConfig,omitempty"`
	Language            *Language     `json:"language,omitempty"`
	LanguageId          string        `json:"languageId,omitempty"`
	Name                string        `json:"name,omitempty"`
	SalesChannel        *SalesChannel `json:"salesChannel,omitempty"`
	SalesChannelId      string        `json:"salesChannelId,omitempty"`
	UpdatedAt           time.Time     `json:"updatedAt,omitempty"`
}
