package model

import "time"

type AppTemplate struct {
	Active    bool      `json:"active"` // required
	App       *App      `json:"app,omitempty"`
	AppId     string    `json:"appId"` // required
	CreatedAt time.Time `json:"createdAt,omitempty"`
	Id        string    `json:"id,omitempty"`
	Path      string    `json:"path"`     // required
	Template  string    `json:"template"` // required
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
