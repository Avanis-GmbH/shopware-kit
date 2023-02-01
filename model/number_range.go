package model

import "time"

type NumberRange struct {
	CreatedAt                *time.Time                `json:"createdAt,omitempty"`
	CustomFields             interface{}               `json:"customFields,omitempty"`
	Description              string                    `json:"description,omitempty"`
	Global                   bool                      `json:"global,omitempty"`
	Id                       string                    `json:"id,omitempty"`
	Name                     string                    `json:"name,omitempty"`
	NumberRangeSalesChannels []NumberRangeSalesChannel `json:"numberRangeSalesChannels,omitempty"`
	Pattern                  string                    `json:"pattern,omitempty"` // required
	Start                    float64                   `json:"start,omitempty"`   // required
	State                    *NumberRangeState         `json:"state,omitempty"`
	Translated               interface{}               `json:"translated,omitempty"`
	Translations             []NumberRangeTranslation  `json:"translations,omitempty"`
	Type                     *NumberRangeType          `json:"type,omitempty"` // required
	TypeId                   string                    `json:"typeId,omitempty"`
	UpdatedAt                *time.Time                `json:"updatedAt,omitempty"`
}
