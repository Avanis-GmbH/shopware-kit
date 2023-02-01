package model

import "time"

type SystemConfig struct {
	ConfigurationKey   string        `json:"configurationKey,omitempty"`   // required
	ConfigurationValue interface{}   `json:"configurationValue,omitempty"` // required
	CreatedAt          *time.Time    `json:"createdAt,omitempty"`
	Id                 string        `json:"id,omitempty"`
	SalesChannel       *SalesChannel `json:"salesChannel,omitempty"`
	SalesChannelId     string        `json:"salesChannelId,omitempty"`
	UpdatedAt          *time.Time    `json:"updatedAt,omitempty"`
}
