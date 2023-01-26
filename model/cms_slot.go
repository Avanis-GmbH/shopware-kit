package model

import "time"

type CmsSlot struct {
	CmsBlockVersionId string               `json:"cmsBlockVersionId,omitempty"`
	VersionId         string               `json:"versionId,omitempty"`
	Slot              string               `json:"slot,omitempty"`
	BlockId           string               `json:"blockId,omitempty"`
	Locked            bool                 `json:"locked,omitempty"`
	Config            interface{}          `json:"config,omitempty"`
	CustomFields      interface{}          `json:"customFields,omitempty"`
	UpdatedAt         time.Time            `json:"updatedAt,omitempty"`
	Translated        interface{}          `json:"translated,omitempty"`
	CreatedAt         time.Time            `json:"createdAt,omitempty"`
	Id                string               `json:"id,omitempty"`
	Type              string               `json:"type,omitempty"`
	Data              interface{}          `json:"data,omitempty"`
	Block             *CmsBlock            `json:"block,omitempty"`
	Translations      []CmsSlotTranslation `json:"translations,omitempty"`
}

type CmsSlotCollection struct {
	EntityCollection

	Data []CmsSlot `json:"data"`
}
