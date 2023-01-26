package model

import "time"

type ProductSorting struct {
	Active       bool                        `json:"active,omitempty"`
	CreatedAt    time.Time                   `json:"createdAt,omitempty"`
	Fields       interface{}                 `json:"fields,omitempty"`
	Id           string                      `json:"id,omitempty"`
	Key          string                      `json:"key,omitempty"`
	Label        string                      `json:"label,omitempty"`
	Locked       bool                        `json:"locked,omitempty"`
	Priority     float64                     `json:"priority,omitempty"`
	Translated   interface{}                 `json:"translated,omitempty"`
	Translations []ProductSortingTranslation `json:"translations,omitempty"`
	UpdatedAt    time.Time                   `json:"updatedAt,omitempty"`
}
