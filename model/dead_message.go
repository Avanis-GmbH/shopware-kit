package model

import "time"

type DeadMessage struct {
	CreatedAt                 *time.Time     `json:"createdAt,omitempty"`
	Encrypted                 bool           `json:"encrypted"`
	ErrorCount                float64        `json:"errorCount,omitempty"`
	Exception                 string         `json:"exception,omitempty"`
	ExceptionFile             string         `json:"exceptionFile,omitempty"`
	ExceptionLine             float64        `json:"exceptionLine,omitempty"`
	ExceptionMessage          string         `json:"exceptionMessage,omitempty"`
	HandlerClass              string         `json:"handlerClass,omitempty"`
	Id                        string         `json:"id,omitempty"`
	NextExecutionTime         *time.Time     `json:"nextExecutionTime,omitempty"`
	OriginalMessageClass      string         `json:"originalMessageClass,omitempty"`
	ScheduledTask             *ScheduledTask `json:"scheduledTask,omitempty"`
	ScheduledTaskId           string         `json:"scheduledTaskId,omitempty"`
	SerializedOriginalMessage interface{}    `json:"serializedOriginalMessage,omitempty"`
	UpdatedAt                 *time.Time     `json:"updatedAt,omitempty"`
}
