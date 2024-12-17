package model

import "time"

type DocumentBaseConfig struct {
	Config         interface{}                      `json:"config,omitempty"`
	CreatedAt      *time.Time                       `json:"createdAt,omitempty"`
	CustomFields   interface{}                      `json:"customFields,omitempty"`
	DocumentNumber string                           `json:"documentNumber,omitempty"`
	DocumentType   *DocumentType                    `json:"documentType,omitempty"`
	DocumentTypeId string                           `json:"documentTypeId,omitempty"` // required
	FilenamePrefix string                           `json:"filenamePrefix,omitempty"`
	FilenameSuffix string                           `json:"filenameSuffix,omitempty"`
	Global         *bool                            `json:"global,omitempty"` // required
	Id             string                           `json:"id,omitempty"`
	Logo           *Media                           `json:"logo,omitempty"`
	LogoId         string                           `json:"logoId,omitempty"`
	Name           string                           `json:"name,omitempty"` // required
	SalesChannels  []DocumentBaseConfigSalesChannel `json:"salesChannels,omitempty"`
	UpdatedAt      *time.Time                       `json:"updatedAt,omitempty"`
}
