package model

import "time"

type AppTemplate struct {
	Active    bool      `json:"active,omitempty"`
	App       *App      `json:"app,omitempty"`
	AppId     string    `json:"appId,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	Id        string    `json:"id,omitempty"`
	Path      string    `json:"path,omitempty"`
	Template  string    `json:"template,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type AppTemplateCollection struct {
	EntityCollection

	Data []AppTemplate `json:"data"`
}
