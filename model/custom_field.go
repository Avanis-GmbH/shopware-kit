package model

import "time"

type CustomField struct {
	Active                    bool                       `json:"active,omitempty"`
	Config                    interface{}                `json:"config,omitempty"`
	CreatedAt                 time.Time                  `json:"createdAt,omitempty"`
	CustomFieldSet            *CustomFieldSet            `json:"customFieldSet,omitempty"`
	CustomFieldSetId          string                     `json:"customFieldSetId,omitempty"`
	Id                        string                     `json:"id,omitempty"`
	Name                      string                     `json:"name,omitempty"`
	ProductSearchConfigFields []ProductSearchConfigField `json:"productSearchConfigFields,omitempty"`
	Type                      string                     `json:"type,omitempty"`
	UpdatedAt                 time.Time                  `json:"updatedAt,omitempty"`
}

type CustomFieldCollection struct {
	EntityCollection

	Data []CustomField `json:"data"`
}
