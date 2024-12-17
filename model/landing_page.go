package model

import "time"

type LandingPage struct {
	Active           *bool                    `json:"active,omitempty"`
	CmsPage          *CmsPage                 `json:"cmsPage,omitempty"`
	CmsPageId        string                   `json:"cmsPageId,omitempty"`
	CmsPageVersionId string                   `json:"cmsPageVersionId,omitempty"`
	CreatedAt        *time.Time               `json:"createdAt,omitempty"`
	CustomFields     interface{}              `json:"customFields,omitempty"`
	Id               string                   `json:"id,omitempty"`
	Keywords         string                   `json:"keywords,omitempty"`
	MetaDescription  string                   `json:"metaDescription,omitempty"`
	MetaTitle        string                   `json:"metaTitle,omitempty"`
	Name             string                   `json:"name,omitempty"` // required
	SalesChannels    []SalesChannel           `json:"salesChannels,omitempty"`
	SeoUrls          []SeoUrl                 `json:"seoUrls,omitempty"`
	SlotConfig       interface{}              `json:"slotConfig,omitempty"`
	Tags             []Tag                    `json:"tags,omitempty"`
	Translated       interface{}              `json:"translated,omitempty"`
	Translations     []LandingPageTranslation `json:"translations,omitempty"`
	UpdatedAt        *time.Time               `json:"updatedAt,omitempty"`
	Url              string                   `json:"url,omitempty"` // required
	VersionId        string                   `json:"versionId,omitempty"`
}
