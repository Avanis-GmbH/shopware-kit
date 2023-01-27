package model

import "time"

type CustomFieldSetRelation struct {
	CreatedAt        time.Time       `json:"createdAt,omitempty"`
	CustomFieldSet   *CustomFieldSet `json:"customFieldSet,omitempty"`
	CustomFieldSetId string          `json:"customFieldSetId"` // required
	EntityName       string          `json:"entityName"`       // required
	Id               string          `json:"id,omitempty"`
	UpdatedAt        time.Time       `json:"updatedAt,omitempty"`
}
