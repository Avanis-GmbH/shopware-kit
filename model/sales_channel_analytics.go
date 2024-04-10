package model

import "time"

type SalesChannelAnalytics struct {
	Active       bool          `json:"active"`
	AnonymizeIp  bool          `json:"anonymizeIp"`
	CreatedAt    *time.Time    `json:"createdAt,omitempty"`
	Id           string        `json:"id,omitempty"`
	SalesChannel *SalesChannel `json:"salesChannel,omitempty"`
	TrackingId   string        `json:"trackingId,omitempty"`
	TrackOrders  bool          `json:"trackOrders"`
	UpdatedAt    *time.Time    `json:"updatedAt,omitempty"`
}
