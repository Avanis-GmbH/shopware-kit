package model

import "time"

type Script struct {
	Active    bool       `json:"active,omitempty"` // required
	App       *App       `json:"app,omitempty"`
	AppId     string     `json:"appId,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Hook      string     `json:"hook,omitempty"` // required
	Id        string     `json:"id,omitempty"`
	Name      string     `json:"name,omitempty"`   // required
	Script    string     `json:"script,omitempty"` // required
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}
