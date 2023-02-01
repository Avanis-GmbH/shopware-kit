package model

import "time"

type NumberRangeState struct {
	CreatedAt     *time.Time   `json:"createdAt,omitempty"`
	Id            string       `json:"id,omitempty"`
	LastValue     float64      `json:"lastValue,omitempty"` // required
	NumberRange   *NumberRange `json:"numberRange,omitempty"`
	NumberRangeId string       `json:"numberRangeId,omitempty"` // required
	UpdatedAt     *time.Time   `json:"updatedAt,omitempty"`
}
