package model

import "time"

type NumberRangeType struct {
	CreatedAt                time.Time                    `json:"createdAt,omitempty"`
	CustomFields             interface{}                  `json:"customFields,omitempty"`
	Global                   bool                         `json:"global"` // required
	Id                       string                       `json:"id,omitempty"`
	NumberRanges             []NumberRange                `json:"numberRanges,omitempty"`
	NumberRangeSalesChannels []NumberRangeSalesChannel    `json:"numberRangeSalesChannels,omitempty"`
	TechnicalName            string                       `json:"technicalName,omitempty"`
	Translated               interface{}                  `json:"translated,omitempty"`
	Translations             []NumberRangeTypeTranslation `json:"translations,omitempty"`
	TypeName                 string                       `json:"typeName"` // required
	UpdatedAt                time.Time                    `json:"updatedAt,omitempty"`
}
