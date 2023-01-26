package model

import "time"

type CmsPageTranslation struct {
	CmsPage          *CmsPage    `json:"cmsPage,omitempty"`
	CmsPageId        string      `json:"cmsPageId,omitempty"`
	CmsPageVersionId string      `json:"cmsPageVersionId,omitempty"`
	CreatedAt        time.Time   `json:"createdAt,omitempty"`
	CustomFields     interface{} `json:"customFields,omitempty"`
	Language         *Language   `json:"language,omitempty"`
	LanguageId       string      `json:"languageId,omitempty"`
	Name             string      `json:"name,omitempty"`
	UpdatedAt        time.Time   `json:"updatedAt,omitempty"`
}

type CmsPageTranslationCollection struct {
	EntityCollection

	Data []CmsPageTranslation `json:"data"`
}
