package model

import "time"

type EventAction struct {
	ActionName    string         `json:"actionName,omitempty"`
	Active        bool           `json:"active,omitempty"`
	Config        interface{}    `json:"config,omitempty"`
	CreatedAt     time.Time      `json:"createdAt,omitempty"`
	CustomFields  interface{}    `json:"customFields,omitempty"`
	EventName     string         `json:"eventName,omitempty"`
	Id            string         `json:"id,omitempty"`
	Rules         []Rule         `json:"rules,omitempty"`
	SalesChannels []SalesChannel `json:"salesChannels,omitempty"`
	Title         string         `json:"title,omitempty"`
	UpdatedAt     time.Time      `json:"updatedAt,omitempty"`
}

type EventActionCollection struct {
	EntityCollection

	Data []EventAction `json:"data"`
}
