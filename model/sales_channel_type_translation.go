package model

import "time"

type SalesChannelTypeTranslation struct {
	CreatedAt          *time.Time        `json:"createdAt,omitempty"`
	CustomFields       interface{}       `json:"customFields,omitempty"`
	Description        string            `json:"description,omitempty"`
	DescriptionLong    string            `json:"descriptionLong,omitempty"`
	Language           *Language         `json:"language,omitempty"`
	LanguageId         string            `json:"languageId,omitempty"`
	Manufacturer       string            `json:"manufacturer,omitempty"`
	Name               string            `json:"name,omitempty"`
	SalesChannelType   *SalesChannelType `json:"salesChannelType,omitempty"`
	SalesChannelTypeId string            `json:"salesChannelTypeId,omitempty"`
	UpdatedAt          *time.Time        `json:"updatedAt,omitempty"`
}
