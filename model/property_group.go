package model

import "time"

type PropertyGroup struct {
	CreatedAt                  time.Time                  `json:"createdAt,omitempty"`
	CustomFields               interface{}                `json:"customFields,omitempty"`
	Description                string                     `json:"description,omitempty"`
	DisplayType                string                     `json:"displayType"` // required
	Filterable                 bool                       `json:"filterable,omitempty"`
	Id                         string                     `json:"id,omitempty"`
	Name                       string                     `json:"name"` // required
	Options                    []PropertyGroupOption      `json:"options,omitempty"`
	Position                   float64                    `json:"position"`
	SortingType                string                     `json:"sortingType"` // required
	Translated                 interface{}                `json:"translated,omitempty"`
	Translations               []PropertyGroupTranslation `json:"translations,omitempty"`
	UpdatedAt                  time.Time                  `json:"updatedAt,omitempty"`
	VisibleOnProductDetailPage bool                       `json:"visibleOnProductDetailPage,omitempty"`
}
