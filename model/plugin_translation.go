package model

import "time"

type PluginTranslation struct {
	Changelog        interface{} `json:"changelog,omitempty"`
	CreatedAt        time.Time   `json:"createdAt,omitempty"`
	CustomFields     interface{} `json:"customFields,omitempty"`
	Description      string      `json:"description,omitempty"`
	Label            string      `json:"label,omitempty"`
	Language         *Language   `json:"language,omitempty"`
	LanguageId       string      `json:"languageId,omitempty"`
	ManufacturerLink string      `json:"manufacturerLink,omitempty"`
	Plugin           *Plugin     `json:"plugin,omitempty"`
	PluginId         string      `json:"pluginId,omitempty"`
	SupportLink      string      `json:"supportLink,omitempty"`
	UpdatedAt        time.Time   `json:"updatedAt,omitempty"`
}

type PluginTranslationCollection struct {
	EntityCollection

	Data []PluginTranslation `json:"data"`
}
