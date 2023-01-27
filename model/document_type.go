package model

import "time"

type DocumentType struct {
	CreatedAt                       time.Time                        `json:"createdAt,omitempty"`
	CustomFields                    interface{}                      `json:"customFields,omitempty"`
	DocumentBaseConfigs             []DocumentBaseConfig             `json:"documentBaseConfigs,omitempty"`
	DocumentBaseConfigSalesChannels []DocumentBaseConfigSalesChannel `json:"documentBaseConfigSalesChannels,omitempty"`
	Documents                       []Document                       `json:"documents,omitempty"`
	Id                              string                           `json:"id,omitempty"`
	Name                            string                           `json:"name"`          // required
	TechnicalName                   string                           `json:"technicalName"` // required
	Translated                      interface{}                      `json:"translated,omitempty"`
	Translations                    []DocumentTypeTranslation        `json:"translations,omitempty"`
	UpdatedAt                       time.Time                        `json:"updatedAt,omitempty"`
}
