package model

import "time"

type CmsSection struct {
	Page                *CmsPage    `json:"page,omitempty"`
	Type                string      `json:"type,omitempty"`
	MobileBehavior      string      `json:"mobileBehavior,omitempty"`
	BackgroundMediaId   string      `json:"backgroundMediaId,omitempty"`
	CssClass            string      `json:"cssClass,omitempty"`
	PageId              string      `json:"pageId,omitempty"`
	BackgroundMedia     *Media      `json:"backgroundMedia,omitempty"`
	Blocks              []CmsBlock  `json:"blocks,omitempty"`
	VersionId           string      `json:"versionId,omitempty"`
	Name                string      `json:"name,omitempty"`
	CreatedAt           time.Time   `json:"createdAt,omitempty"`
	CmsPageVersionId    string      `json:"cmsPageVersionId,omitempty"`
	SizingMode          string      `json:"sizingMode,omitempty"`
	BackgroundColor     string      `json:"backgroundColor,omitempty"`
	CustomFields        interface{} `json:"customFields,omitempty"`
	Id                  string      `json:"id,omitempty"`
	Locked              bool        `json:"locked,omitempty"`
	BackgroundMediaMode string      `json:"backgroundMediaMode,omitempty"`
	UpdatedAt           time.Time   `json:"updatedAt,omitempty"`
	Position            float64     `json:"position,omitempty"`
}
