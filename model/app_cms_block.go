package model

import "time"

type AppCmsBlock struct {
	App          *App                     `json:"app,omitempty"`
	AppId        string                   `json:"appId,omitempty"`
	Block        interface{}              `json:"block,omitempty"`
	CreatedAt    time.Time                `json:"createdAt,omitempty"`
	Id           string                   `json:"id,omitempty"`
	Label        string                   `json:"label,omitempty"`
	Name         string                   `json:"name,omitempty"`
	Styles       string                   `json:"styles,omitempty"`
	Template     string                   `json:"template,omitempty"`
	Translated   interface{}              `json:"translated,omitempty"`
	Translations []AppCmsBlockTranslation `json:"translations,omitempty"`
	UpdatedAt    time.Time                `json:"updatedAt,omitempty"`
}

type AppCmsBlockCollection struct {
	EntityCollection

	Data []AppCmsBlock `json:"data"`
}
