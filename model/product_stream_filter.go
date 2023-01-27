package model

import "time"

type ProductStreamFilter struct {
	CreatedAt       time.Time             `json:"createdAt,omitempty"`
	CustomFields    interface{}           `json:"customFields,omitempty"`
	Field           string                `json:"field,omitempty"`
	Id              string                `json:"id,omitempty"`
	Operator        string                `json:"operator,omitempty"`
	Parameters      interface{}           `json:"parameters,omitempty"`
	Parent          *ProductStreamFilter  `json:"parent,omitempty"`
	ParentId        string                `json:"parentId,omitempty"`
	Position        float64               `json:"position"`
	ProductStream   *ProductStream        `json:"productStream,omitempty"`
	ProductStreamId string                `json:"productStreamId"` // required
	Queries         []ProductStreamFilter `json:"queries,omitempty"`
	Type            string                `json:"type"` // required
	UpdatedAt       time.Time             `json:"updatedAt,omitempty"`
	Value           string                `json:"value,omitempty"`
}
