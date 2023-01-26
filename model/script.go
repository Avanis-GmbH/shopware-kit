package model

import "time"

type Script struct {
	Active    bool      `json:"active,omitempty"`
	App       *App      `json:"app,omitempty"`
	AppId     string    `json:"appId,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	Hook      string    `json:"hook,omitempty"`
	Id        string    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Script    string    `json:"script,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
