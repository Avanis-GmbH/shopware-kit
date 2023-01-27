package model

import "time"

type CmsPage struct {
	Categories        []Category           `json:"categories,omitempty"`
	Config            interface{}          `json:"config,omitempty"`
	CreatedAt         time.Time            `json:"createdAt,omitempty"`
	CustomFields      interface{}          `json:"customFields,omitempty"`
	Entity            string               `json:"entity,omitempty"`
	HomeSalesChannels []SalesChannel       `json:"homeSalesChannels,omitempty"`
	Id                string               `json:"id,omitempty"`
	LandingPages      []LandingPage        `json:"landingPages,omitempty"`
	Locked            bool                 `json:"locked,omitempty"`
	Name              string               `json:"name,omitempty"`
	PreviewMedia      *Media               `json:"previewMedia,omitempty"`
	PreviewMediaId    string               `json:"previewMediaId,omitempty"`
	Products          []Product            `json:"products,omitempty"`
	Sections          []CmsSection         `json:"sections,omitempty"`
	Translated        interface{}          `json:"translated,omitempty"`
	Translations      []CmsPageTranslation `json:"translations,omitempty"`
	Type              string               `json:"type"` // required
	UpdatedAt         time.Time            `json:"updatedAt,omitempty"`
	VersionId         string               `json:"versionId,omitempty"`
}
