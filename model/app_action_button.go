package model

import "time"

type AppActionButton struct {
	Action       string                       `json:"action"` // required
	App          *App                         `json:"app,omitempty"`
	AppId        string                       `json:"appId"` // required
	CreatedAt    time.Time                    `json:"createdAt,omitempty"`
	Entity       string                       `json:"entity"` // required
	Id           string                       `json:"id,omitempty"`
	Label        string                       `json:"label"` // required
	OpenNewTab   bool                         `json:"openNewTab,omitempty"`
	Translated   interface{}                  `json:"translated,omitempty"`
	Translations []AppActionButtonTranslation `json:"translations,omitempty"`
	UpdatedAt    time.Time                    `json:"updatedAt,omitempty"`
	Url          string                       `json:"url"`  // required
	View         string                       `json:"view"` // required
}
