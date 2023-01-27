package model

import "time"

type ScheduledTask struct {
	CreatedAt          time.Time     `json:"createdAt,omitempty"`
	DeadMessages       []DeadMessage `json:"deadMessages,omitempty"`
	Id                 string        `json:"id,omitempty"`
	LastExecutionTime  time.Time     `json:"lastExecutionTime,omitempty"`
	Name               string        `json:"name"` // required
	NextExecutionTime  time.Time     `json:"nextExecutionTime,omitempty"`
	RunInterval        int64         `json:"runInterval"`        // required
	ScheduledTaskClass string        `json:"scheduledTaskClass"` // required
	Status             string        `json:"status"`             // required
	UpdatedAt          time.Time     `json:"updatedAt,omitempty"`
}
