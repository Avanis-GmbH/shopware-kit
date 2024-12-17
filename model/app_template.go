package model

import "time"

type AppTemplate struct {
	Active    *bool      `json:"active,omitempty"` // required
	App       *App       `json:"app,omitempty"`
	AppId     string     `json:"appId,omitempty"` // required
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Id        string     `json:"id,omitempty"`
	Path      string     `json:"path,omitempty"`     // required
	Template  string     `json:"template,omitempty"` // required
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}
