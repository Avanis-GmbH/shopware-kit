package model

import "time"

type CustomField struct {
	Active                    bool                       `json:"active"`
	Config                    interface{}                `json:"config,omitempty"`
	CreatedAt                 *time.Time                 `json:"createdAt,omitempty"`
	CustomFieldSet            *CustomFieldSet            `json:"customFieldSet,omitempty"`
	CustomFieldSetId          string                     `json:"customFieldSetId,omitempty"`
	Id                        string                     `json:"id,omitempty"`
	Name                      string                     `json:"name,omitempty"` // required
	ProductSearchConfigFields []ProductSearchConfigField `json:"productSearchConfigFields,omitempty"`
	Type                      string                     `json:"type,omitempty"` // required
	UpdatedAt                 *time.Time                 `json:"updatedAt,omitempty"`
}
