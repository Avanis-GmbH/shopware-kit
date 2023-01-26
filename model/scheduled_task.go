package model

import "time"

type ScheduledTask struct {
	CreatedAt          time.Time     `json:"createdAt,omitempty"`
	DeadMessages       []DeadMessage `json:"deadMessages,omitempty"`
	Id                 string        `json:"id,omitempty"`
	LastExecutionTime  time.Time     `json:"lastExecutionTime,omitempty"`
	Name               string        `json:"name,omitempty"`
	NextExecutionTime  time.Time     `json:"nextExecutionTime,omitempty"`
	RunInterval        float64       `json:"runInterval,omitempty"`
	ScheduledTaskClass string        `json:"scheduledTaskClass,omitempty"`
	Status             string        `json:"status,omitempty"`
	UpdatedAt          time.Time     `json:"updatedAt,omitempty"`
}

type ScheduledTaskCollection struct {
	EntityCollection

	Data []ScheduledTask `json:"data"`
}
