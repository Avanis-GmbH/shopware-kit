package model

import "time"

type MessageQueueStats struct {
	CreatedAt time.Time `json:"createdAt,omitempty"`
	Id        string    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Size      float64   `json:"size"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
