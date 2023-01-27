package model

import "time"

type Locale struct {
	Code         string              `json:"code"` // required
	CreatedAt    time.Time           `json:"createdAt,omitempty"`
	CustomFields interface{}         `json:"customFields,omitempty"`
	Id           string              `json:"id,omitempty"`
	Languages    []Language          `json:"languages,omitempty"`
	Name         string              `json:"name"`      // required
	Territory    string              `json:"territory"` // required
	Translated   interface{}         `json:"translated,omitempty"`
	Translations []LocaleTranslation `json:"translations,omitempty"`
	UpdatedAt    time.Time           `json:"updatedAt,omitempty"`
	Users        []User              `json:"users,omitempty"`
}
