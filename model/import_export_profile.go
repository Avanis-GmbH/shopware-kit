package model

import "time"

type ImportExportProfile struct {
	Config           interface{}                      `json:"config,omitempty"`
	CreatedAt        time.Time                        `json:"createdAt,omitempty"`
	Delimiter        string                           `json:"delimiter,omitempty"`
	Enclosure        string                           `json:"enclosure,omitempty"`
	FileType         string                           `json:"fileType,omitempty"`
	Id               string                           `json:"id,omitempty"`
	ImportExportLogs []ImportExportLog                `json:"importExportLogs,omitempty"`
	Label            string                           `json:"label,omitempty"`
	Mapping          interface{}                      `json:"mapping,omitempty"`
	Name             string                           `json:"name,omitempty"`
	SourceEntity     string                           `json:"sourceEntity,omitempty"`
	SystemDefault    bool                             `json:"systemDefault,omitempty"`
	Translated       interface{}                      `json:"translated,omitempty"`
	Translations     []ImportExportProfileTranslation `json:"translations,omitempty"`
	Type             string                           `json:"type,omitempty"`
	UpdateBy         interface{}                      `json:"updateBy,omitempty"`
	UpdatedAt        time.Time                        `json:"updatedAt,omitempty"`
}

type ImportExportProfileCollection struct {
	EntityCollection

	Data []ImportExportProfile `json:"data"`
}
