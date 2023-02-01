package model

import "time"

type CustomFieldSetRelation struct {
	CreatedAt        *time.Time      `json:"createdAt,omitempty"`
	CustomFieldSet   *CustomFieldSet `json:"customFieldSet,omitempty"`
	CustomFieldSetId string          `json:"customFieldSetId,omitempty"` // required
	EntityName       string          `json:"entityName,omitempty"`       // required
	Id               string          `json:"id,omitempty"`
	UpdatedAt        *time.Time      `json:"updatedAt,omitempty"`
}
