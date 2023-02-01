package model

import "time"

type CmsSlot struct {
	Block             *CmsBlock            `json:"block,omitempty"`
	BlockId           string               `json:"blockId,omitempty"` // required
	CmsBlockVersionId string               `json:"cmsBlockVersionId,omitempty"`
	Config            interface{}          `json:"config,omitempty"`
	CreatedAt         *time.Time           `json:"createdAt,omitempty"`
	CustomFields      interface{}          `json:"customFields,omitempty"`
	Data              interface{}          `json:"data,omitempty"`
	Id                string               `json:"id,omitempty"`
	Locked            bool                 `json:"locked,omitempty"`
	Slot              string               `json:"slot,omitempty"` // required
	Translated        interface{}          `json:"translated,omitempty"`
	Translations      []CmsSlotTranslation `json:"translations,omitempty"`
	Type              string               `json:"type,omitempty"` // required
	UpdatedAt         *time.Time           `json:"updatedAt,omitempty"`
	VersionId         string               `json:"versionId,omitempty"`
}
