package model

import "time"

type DocumentBaseConfigSalesChannel struct {
	CreatedAt            time.Time           `json:"createdAt,omitempty"`
	DocumentBaseConfig   *DocumentBaseConfig `json:"documentBaseConfig,omitempty"`
	DocumentBaseConfigId string              `json:"documentBaseConfigId,omitempty"`
	DocumentType         *DocumentType       `json:"documentType,omitempty"`
	DocumentTypeId       string              `json:"documentTypeId,omitempty"`
	Id                   string              `json:"id,omitempty"`
	SalesChannel         *SalesChannel       `json:"salesChannel,omitempty"`
	SalesChannelId       string              `json:"salesChannelId,omitempty"`
	UpdatedAt            time.Time           `json:"updatedAt,omitempty"`
}

type DocumentBaseConfigSalesChannelCollection struct {
	EntityCollection

	Data []DocumentBaseConfigSalesChannel `json:"data"`
}
