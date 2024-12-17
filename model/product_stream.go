package model

import "time"

type ProductStream struct {
	ApiFilter            interface{}                `json:"apiFilter,omitempty"`
	Categories           []Category                 `json:"categories,omitempty"`
	CreatedAt            *time.Time                 `json:"createdAt,omitempty"`
	CustomFields         interface{}                `json:"customFields,omitempty"`
	Description          string                     `json:"description,omitempty"`
	Filters              []ProductStreamFilter      `json:"filters,omitempty"`
	Id                   string                     `json:"id,omitempty"`
	Invalid              *bool                      `json:"invalid,omitempty"`
	Name                 string                     `json:"name,omitempty"` // required
	ProductCrossSellings []ProductCrossSelling      `json:"productCrossSellings,omitempty"`
	ProductExports       []ProductExport            `json:"productExports,omitempty"`
	Translated           interface{}                `json:"translated,omitempty"`
	Translations         []ProductStreamTranslation `json:"translations,omitempty"`
	UpdatedAt            *time.Time                 `json:"updatedAt,omitempty"`
}
