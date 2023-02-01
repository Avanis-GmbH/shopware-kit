package model

import "time"

type ProductSearchConfigField struct {
	CreatedAt      *time.Time           `json:"createdAt,omitempty"`
	CustomField    *CustomField         `json:"customField,omitempty"`
	CustomFieldId  string               `json:"customFieldId,omitempty"`
	Field          string               `json:"field,omitempty"` // required
	Id             string               `json:"id,omitempty"`
	Ranking        int64                `json:"ranking,omitempty"`    // required
	Searchable     bool                 `json:"searchable,omitempty"` // required
	SearchConfig   *ProductSearchConfig `json:"searchConfig,omitempty"`
	SearchConfigId string               `json:"searchConfigId,omitempty"` // required
	Tokenize       bool                 `json:"tokenize,omitempty"`       // required
	UpdatedAt      *time.Time           `json:"updatedAt,omitempty"`
}
