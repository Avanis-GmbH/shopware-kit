package model

import "time"

type Webhook struct {
	Active     *bool      `json:"active,omitempty"`
	App        *App       `json:"app,omitempty"`
	AppId      string     `json:"appId,omitempty"`
	CreatedAt  *time.Time `json:"createdAt,omitempty"`
	ErrorCount int64      `json:"errorCount,omitempty"` // required
	EventName  string     `json:"eventName,omitempty"`  // required
	Id         string     `json:"id,omitempty"`
	Name       string     `json:"name,omitempty"` // required
	UpdatedAt  *time.Time `json:"updatedAt,omitempty"`
	Url        string     `json:"url,omitempty"` // required
}
