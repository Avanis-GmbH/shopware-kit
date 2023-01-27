package model

import "time"

type CmsSlotTranslation struct {
	CmsSlot          *CmsSlot    `json:"cmsSlot,omitempty"`
	CmsSlotId        string      `json:"cmsSlotId,omitempty"`
	CmsSlotVersionId string      `json:"cmsSlotVersionId,omitempty"`
	Config           interface{} `json:"config,omitempty"`
	CreatedAt        time.Time   `json:"createdAt,omitempty"`
	CustomFields     interface{} `json:"customFields,omitempty"`
	Language         *Language   `json:"language,omitempty"`
	LanguageId       string      `json:"languageId,omitempty"`
	UpdatedAt        time.Time   `json:"updatedAt,omitempty"`
}
