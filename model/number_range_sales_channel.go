package model

import "time"

type NumberRangeSalesChannel struct {
	CreatedAt         time.Time        `json:"createdAt,omitempty"`
	Id                string           `json:"id,omitempty"`
	NumberRange       *NumberRange     `json:"numberRange,omitempty"`
	NumberRangeId     string           `json:"numberRangeId,omitempty"`
	NumberRangeType   *NumberRangeType `json:"numberRangeType,omitempty"`
	NumberRangeTypeId string           `json:"numberRangeTypeId,omitempty"`
	SalesChannel      *SalesChannel    `json:"salesChannel,omitempty"`
	SalesChannelId    string           `json:"salesChannelId,omitempty"`
	UpdatedAt         time.Time        `json:"updatedAt,omitempty"`
}

type NumberRangeSalesChannelCollection struct {
	EntityCollection

	Data []NumberRangeSalesChannel `json:"data"`
}
