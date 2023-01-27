package model

import "time"

type AppCmsBlock struct {
	App          *App                     `json:"app,omitempty"`
	AppId        string                   `json:"appId"` // required
	Block        interface{}              `json:"block"` // required
	CreatedAt    time.Time                `json:"createdAt,omitempty"`
	Id           string                   `json:"id,omitempty"`
	Label        string                   `json:"label,omitempty"`
	Name         string                   `json:"name"`     // required
	Styles       string                   `json:"styles"`   // required
	Template     string                   `json:"template"` // required
	Translated   interface{}              `json:"translated,omitempty"`
	Translations []AppCmsBlockTranslation `json:"translations,omitempty"`
	UpdatedAt    time.Time                `json:"updatedAt,omitempty"`
}
