package model

import "time"

type ProductSearchConfigField struct {
	CreatedAt      time.Time            `json:"createdAt,omitempty"`
	CustomField    *CustomField         `json:"customField,omitempty"`
	CustomFieldId  string               `json:"customFieldId,omitempty"`
	Field          string               `json:"field"` // required
	Id             string               `json:"id,omitempty"`
	Ranking        int64                `json:"ranking"`    // required
	Searchable     bool                 `json:"searchable"` // required
	SearchConfig   *ProductSearchConfig `json:"searchConfig,omitempty"`
	SearchConfigId string               `json:"searchConfigId"` // required
	Tokenize       bool                 `json:"tokenize"`       // required
	UpdatedAt      time.Time            `json:"updatedAt,omitempty"`
}
