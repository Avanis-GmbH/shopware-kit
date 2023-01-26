package model

import "time"

type PropertyGroup struct {
	CreatedAt                  time.Time                  `json:"createdAt,omitempty"`
	CustomFields               interface{}                `json:"customFields,omitempty"`
	Description                string                     `json:"description,omitempty"`
	DisplayType                string                     `json:"displayType,omitempty"`
	Filterable                 bool                       `json:"filterable,omitempty"`
	Id                         string                     `json:"id,omitempty"`
	Name                       string                     `json:"name,omitempty"`
	Options                    []PropertyGroupOption      `json:"options,omitempty"`
	Position                   float64                    `json:"position,omitempty"`
	SortingType                string                     `json:"sortingType,omitempty"`
	Translated                 interface{}                `json:"translated,omitempty"`
	Translations               []PropertyGroupTranslation `json:"translations,omitempty"`
	UpdatedAt                  time.Time                  `json:"updatedAt,omitempty"`
	VisibleOnProductDetailPage bool                       `json:"visibleOnProductDetailPage,omitempty"`
}
