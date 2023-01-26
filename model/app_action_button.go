package model

import "time"

type AppActionButton struct {
	Action       string                       `json:"action,omitempty"`
	App          *App                         `json:"app,omitempty"`
	AppId        string                       `json:"appId,omitempty"`
	CreatedAt    time.Time                    `json:"createdAt,omitempty"`
	Entity       string                       `json:"entity,omitempty"`
	Id           string                       `json:"id,omitempty"`
	Label        string                       `json:"label,omitempty"`
	OpenNewTab   bool                         `json:"openNewTab,omitempty"`
	Translated   interface{}                  `json:"translated,omitempty"`
	Translations []AppActionButtonTranslation `json:"translations,omitempty"`
	UpdatedAt    time.Time                    `json:"updatedAt,omitempty"`
	Url          string                       `json:"url,omitempty"`
	View         string                       `json:"view,omitempty"`
}

type AppActionButtonCollection struct {
	EntityCollection

	Data []AppActionButton `json:"data"`
}
