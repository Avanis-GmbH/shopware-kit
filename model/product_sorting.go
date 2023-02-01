package model

import "time"

type ProductSorting struct {
	Active       bool                        `json:"active,omitempty"` // required
	CreatedAt    *time.Time                  `json:"createdAt,omitempty"`
	Fields       interface{}                 `json:"fields,omitempty"` // required
	Id           string                      `json:"id,omitempty"`
	Key          string                      `json:"key,omitempty"`   // required
	Label        string                      `json:"label,omitempty"` // required
	Locked       bool                        `json:"locked,omitempty"`
	Priority     int64                       `json:"priority,omitempty"` // required
	Translated   interface{}                 `json:"translated,omitempty"`
	Translations []ProductSortingTranslation `json:"translations,omitempty"`
	UpdatedAt    *time.Time                  `json:"updatedAt,omitempty"`
}
