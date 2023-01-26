package model

type EventActionSalesChannel struct {
	EventAction    *EventAction  `json:"eventAction,omitempty"`
	EventActionId  string        `json:"eventActionId,omitempty"`
	SalesChannel   *SalesChannel `json:"salesChannel,omitempty"`
	SalesChannelId string        `json:"salesChannelId,omitempty"`
}
