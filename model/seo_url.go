package model

import "time"

type SeoUrl struct {
	CreatedAt      time.Time     `json:"createdAt,omitempty"`
	CustomFields   interface{}   `json:"customFields,omitempty"`
	ForeignKey     string        `json:"foreignKey,omitempty"`
	Id             string        `json:"id,omitempty"`
	IsCanonical    bool          `json:"isCanonical,omitempty"`
	IsDeleted      bool          `json:"isDeleted,omitempty"`
	IsModified     bool          `json:"isModified,omitempty"`
	Language       *Language     `json:"language,omitempty"`
	LanguageId     string        `json:"languageId,omitempty"`
	PathInfo       string        `json:"pathInfo,omitempty"`
	RouteName      string        `json:"routeName,omitempty"`
	SalesChannel   *SalesChannel `json:"salesChannel,omitempty"`
	SalesChannelId string        `json:"salesChannelId,omitempty"`
	SeoPathInfo    string        `json:"seoPathInfo,omitempty"`
	UpdatedAt      time.Time     `json:"updatedAt,omitempty"`
	Url            string        `json:"url,omitempty"`
}

type SeoUrlCollection struct {
	EntityCollection

	Data []SeoUrl `json:"data"`
}
