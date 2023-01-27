package model

import "time"

type CustomFieldSet struct {
	Active       bool                     `json:"active,omitempty"`
	App          *App                     `json:"app,omitempty"`
	AppId        string                   `json:"appId,omitempty"`
	Config       interface{}              `json:"config,omitempty"`
	CreatedAt    time.Time                `json:"createdAt,omitempty"`
	CustomFields []CustomField            `json:"customFields,omitempty"`
	Global       bool                     `json:"global,omitempty"`
	Id           string                   `json:"id,omitempty"`
	Name         string                   `json:"name"` // required
	Position     float64                  `json:"position,omitempty"`
	Products     []Product                `json:"products,omitempty"`
	Relations    []CustomFieldSetRelation `json:"relations,omitempty"`
	UpdatedAt    time.Time                `json:"updatedAt,omitempty"`
}
