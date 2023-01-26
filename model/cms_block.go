package model

import "time"

type CmsBlock struct {
	BackgroundColor     string      `json:"backgroundColor,omitempty"`
	BackgroundMedia     *Media      `json:"backgroundMedia,omitempty"`
	BackgroundMediaId   string      `json:"backgroundMediaId,omitempty"`
	BackgroundMediaMode string      `json:"backgroundMediaMode,omitempty"`
	CmsSectionVersionId string      `json:"cmsSectionVersionId,omitempty"`
	CreatedAt           time.Time   `json:"createdAt,omitempty"`
	CssClass            string      `json:"cssClass,omitempty"`
	CustomFields        interface{} `json:"customFields,omitempty"`
	Id                  string      `json:"id,omitempty"`
	Locked              bool        `json:"locked,omitempty"`
	MarginBottom        string      `json:"marginBottom,omitempty"`
	MarginLeft          string      `json:"marginLeft,omitempty"`
	MarginRight         string      `json:"marginRight,omitempty"`
	MarginTop           string      `json:"marginTop,omitempty"`
	Name                string      `json:"name,omitempty"`
	Position            float64     `json:"position,omitempty"`
	Section             *CmsSection `json:"section,omitempty"`
	SectionId           string      `json:"sectionId,omitempty"`
	SectionPosition     string      `json:"sectionPosition,omitempty"`
	Slots               []CmsSlot   `json:"slots,omitempty"`
	Type                string      `json:"type,omitempty"`
	UpdatedAt           time.Time   `json:"updatedAt,omitempty"`
	VersionId           string      `json:"versionId,omitempty"`
}

type CmsBlockCollection struct {
	EntityCollection

	Data []CmsBlock `json:"data"`
}
