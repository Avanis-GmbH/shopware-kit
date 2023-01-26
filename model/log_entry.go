package model

import "time"

type LogEntry struct {
	Channel   string      `json:"channel,omitempty"`
	Context   interface{} `json:"context,omitempty"`
	CreatedAt time.Time   `json:"createdAt,omitempty"`
	Extra     interface{} `json:"extra,omitempty"`
	Id        string      `json:"id,omitempty"`
	Level     float64     `json:"level,omitempty"`
	Message   string      `json:"message,omitempty"`
	UpdatedAt time.Time   `json:"updatedAt,omitempty"`
}

type LogEntryCollection struct {
	EntityCollection

	Data []LogEntry `json:"data"`
}
