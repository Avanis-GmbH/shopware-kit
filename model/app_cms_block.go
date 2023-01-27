package model

import "time"

type AppCmsBlock struct {
	App          *App                     `json:"app,omitempty"`
	AppId        string                   `json:"appId,omitempty"` // required
	Block        interface{}              `json:"block,omitempty"` // required
	CreatedAt    time.Time                `json:"createdAt,omitempty"`
	Id           string                   `json:"id,omitempty"`
	Label        string                   `json:"label,omitempty"`
	Name         string                   `json:"name,omitempty"`     // required
	Styles       string                   `json:"styles,omitempty"`   // required
	Template     string                   `json:"template,omitempty"` // required
	Translated   interface{}              `json:"translated,omitempty"`
	Translations []AppCmsBlockTranslation `json:"translations,omitempty"`
	UpdatedAt    time.Time                `json:"updatedAt,omitempty"`
}
