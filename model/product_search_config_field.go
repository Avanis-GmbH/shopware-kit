package model

import "time"

type ProductSearchConfigField struct {
	CreatedAt      time.Time            `json:"createdAt,omitempty"`
	CustomField    *CustomField         `json:"customField,omitempty"`
	CustomFieldId  string               `json:"customFieldId,omitempty"`
	Field          string               `json:"field,omitempty"`
	Id             string               `json:"id,omitempty"`
	Ranking        float64              `json:"ranking,omitempty"`
	Searchable     bool                 `json:"searchable,omitempty"`
	SearchConfig   *ProductSearchConfig `json:"searchConfig,omitempty"`
	SearchConfigId string               `json:"searchConfigId,omitempty"`
	Tokenize       bool                 `json:"tokenize,omitempty"`
	UpdatedAt      time.Time            `json:"updatedAt,omitempty"`
}

type ProductSearchConfigFieldCollection struct {
	EntityCollection

	Data []ProductSearchConfigField `json:"data"`
}
