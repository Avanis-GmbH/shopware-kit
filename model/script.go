package model

import "time"

type Script struct {
	Active    bool      `json:"active"` // required
	App       *App      `json:"app,omitempty"`
	AppId     string    `json:"appId,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	Hook      string    `json:"hook"` // required
	Id        string    `json:"id,omitempty"`
	Name      string    `json:"name"`   // required
	Script    string    `json:"script"` // required
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
