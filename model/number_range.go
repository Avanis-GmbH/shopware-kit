package model

import "time"

type NumberRange struct {
	CreatedAt                time.Time                 `json:"createdAt,omitempty"`
	CustomFields             interface{}               `json:"customFields,omitempty"`
	Description              string                    `json:"description,omitempty"`
	Global                   bool                      `json:"global,omitempty"`
	Id                       string                    `json:"id,omitempty"`
	Name                     string                    `json:"name,omitempty"`
	NumberRangeSalesChannels []NumberRangeSalesChannel `json:"numberRangeSalesChannels,omitempty"`
	Pattern                  string                    `json:"pattern,omitempty"`
	Start                    float64                   `json:"start,omitempty"`
	State                    *NumberRangeState         `json:"state,omitempty"`
	Translated               interface{}               `json:"translated,omitempty"`
	Translations             []NumberRangeTranslation  `json:"translations,omitempty"`
	Type                     *NumberRangeType          `json:"type,omitempty"`
	TypeId                   string                    `json:"typeId,omitempty"`
	UpdatedAt                time.Time                 `json:"updatedAt,omitempty"`
}

type NumberRangeCollection struct {
	EntityCollection

	Data []NumberRange `json:"data"`
}
