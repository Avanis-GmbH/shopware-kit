package model

import "time"

type ImportExportProfile struct {
	Config           interface{}                      `json:"config,omitempty"`
	CreatedAt        time.Time                        `json:"createdAt,omitempty"`
	Delimiter        string                           `json:"delimiter"` // required
	Enclosure        string                           `json:"enclosure"` // required
	FileType         string                           `json:"fileType"`  // required
	Id               string                           `json:"id,omitempty"`
	ImportExportLogs []ImportExportLog                `json:"importExportLogs,omitempty"`
	Label            string                           `json:"label"` // required
	Mapping          interface{}                      `json:"mapping,omitempty"`
	Name             string                           `json:"name,omitempty"`
	SourceEntity     string                           `json:"sourceEntity"` // required
	SystemDefault    bool                             `json:"systemDefault,omitempty"`
	Translated       interface{}                      `json:"translated,omitempty"`
	Translations     []ImportExportProfileTranslation `json:"translations,omitempty"`
	Type             string                           `json:"type,omitempty"`
	UpdateBy         interface{}                      `json:"updateBy,omitempty"`
	UpdatedAt        time.Time                        `json:"updatedAt,omitempty"`
}
