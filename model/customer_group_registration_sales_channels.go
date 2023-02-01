package model

import "time"

type CustomerGroupRegistrationSalesChannels struct {
	CreatedAt       *time.Time     `json:"createdAt,omitempty"`
	CustomerGroup   *CustomerGroup `json:"customerGroup,omitempty"`
	CustomerGroupId string         `json:"customerGroupId,omitempty"`
	SalesChannel    *SalesChannel  `json:"salesChannel,omitempty"`
	SalesChannelId  string         `json:"salesChannelId,omitempty"`
}
