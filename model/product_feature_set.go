package model

import "time"

type ProductFeatureSet struct {
	CreatedAt    *time.Time                     `json:"createdAt,omitempty"`
	Description  string                         `json:"description,omitempty"`
	Features     interface{}                    `json:"features,omitempty"`
	Id           string                         `json:"id,omitempty"`
	Name         string                         `json:"name,omitempty"` // required
	Products     []Product                      `json:"products,omitempty"`
	Translated   interface{}                    `json:"translated,omitempty"`
	Translations []ProductFeatureSetTranslation `json:"translations,omitempty"`
	UpdatedAt    *time.Time                     `json:"updatedAt,omitempty"`
}
