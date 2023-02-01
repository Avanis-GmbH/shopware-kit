package model

import "time"

type AppActionButton struct {
	Action       string                       `json:"action,omitempty"` // required
	App          *App                         `json:"app,omitempty"`
	AppId        string                       `json:"appId,omitempty"` // required
	CreatedAt    *time.Time                   `json:"createdAt,omitempty"`
	Entity       string                       `json:"entity,omitempty"` // required
	Id           string                       `json:"id,omitempty"`
	Label        string                       `json:"label,omitempty"` // required
	OpenNewTab   bool                         `json:"openNewTab,omitempty"`
	Translated   interface{}                  `json:"translated,omitempty"`
	Translations []AppActionButtonTranslation `json:"translations,omitempty"`
	UpdatedAt    *time.Time                   `json:"updatedAt,omitempty"`
	Url          string                       `json:"url,omitempty"`  // required
	View         string                       `json:"view,omitempty"` // required
}
