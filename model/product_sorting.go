package model

import "time"

type ProductSorting struct {
	Active       bool                        `json:"active"` // required
	CreatedAt    time.Time                   `json:"createdAt,omitempty"`
	Fields       interface{}                 `json:"fields"` // required
	Id           string                      `json:"id,omitempty"`
	Key          string                      `json:"key"`   // required
	Label        string                      `json:"label"` // required
	Locked       bool                        `json:"locked,omitempty"`
	Priority     int64                       `json:"priority"` // required
	Translated   interface{}                 `json:"translated,omitempty"`
	Translations []ProductSortingTranslation `json:"translations,omitempty"`
	UpdatedAt    time.Time                   `json:"updatedAt,omitempty"`
}
