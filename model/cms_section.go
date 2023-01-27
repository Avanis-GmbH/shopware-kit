package model

import "time"

type CmsSection struct {
	BackgroundColor     string      `json:"backgroundColor,omitempty"`
	BackgroundMedia     *Media      `json:"backgroundMedia,omitempty"`
	BackgroundMediaId   string      `json:"backgroundMediaId,omitempty"`
	BackgroundMediaMode string      `json:"backgroundMediaMode,omitempty"`
	Blocks              []CmsBlock  `json:"blocks,omitempty"`
	CmsPageVersionId    string      `json:"cmsPageVersionId,omitempty"`
	CreatedAt           time.Time   `json:"createdAt,omitempty"`
	CssClass            string      `json:"cssClass,omitempty"`
	CustomFields        interface{} `json:"customFields,omitempty"`
	Id                  string      `json:"id,omitempty"`
	Locked              bool        `json:"locked,omitempty"`
	MobileBehavior      string      `json:"mobileBehavior,omitempty"`
	Name                string      `json:"name,omitempty"`
	Page                *CmsPage    `json:"page,omitempty"`
	PageId              string      `json:"pageId,omitempty"`   // required
	Position            float64     `json:"position,omitempty"` // required
	SizingMode          string      `json:"sizingMode,omitempty"`
	Type                string      `json:"type,omitempty"` // required
	UpdatedAt           time.Time   `json:"updatedAt,omitempty"`
	VersionId           string      `json:"versionId,omitempty"`
}
