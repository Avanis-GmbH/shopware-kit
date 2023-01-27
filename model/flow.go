package model

import "time"

type Flow struct {
	Active       bool           `json:"active,omitempty"`
	CreatedAt    time.Time      `json:"createdAt,omitempty"`
	CustomFields interface{}    `json:"customFields,omitempty"`
	Description  string         `json:"description,omitempty"`
	EventName    string         `json:"eventName"` // required
	Id           string         `json:"id,omitempty"`
	Invalid      bool           `json:"invalid,omitempty"`
	Name         string         `json:"name"` // required
	Payload      interface{}    `json:"payload,omitempty"`
	Priority     float64        `json:"priority"`
	Sequences    []FlowSequence `json:"sequences,omitempty"`
	UpdatedAt    time.Time      `json:"updatedAt,omitempty"`
}
