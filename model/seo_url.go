package model

import "time"

type SeoUrl struct {
	CreatedAt      *time.Time    `json:"createdAt,omitempty"`
	CustomFields   interface{}   `json:"customFields,omitempty"`
	ForeignKey     string        `json:"foreignKey,omitempty"` // required
	Id             string        `json:"id,omitempty"`
	IsCanonical    *bool         `json:"isCanonical,omitempty"`
	IsDeleted      *bool         `json:"isDeleted,omitempty"`
	IsModified     *bool         `json:"isModified,omitempty"`
	Language       *Language     `json:"language,omitempty"`
	LanguageId     string        `json:"languageId,omitempty"` // required
	PathInfo       string        `json:"pathInfo,omitempty"`   // required
	RouteName      string        `json:"routeName,omitempty"`  // required
	SalesChannel   *SalesChannel `json:"salesChannel,omitempty"`
	SalesChannelId string        `json:"salesChannelId,omitempty"`
	SeoPathInfo    string        `json:"seoPathInfo,omitempty"` // required
	UpdatedAt      *time.Time    `json:"updatedAt,omitempty"`
	Url            string        `json:"url,omitempty"`
}
