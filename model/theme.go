package model

import "time"

type Theme struct {
	Active         bool               `json:"active,omitempty"`
	Author         string             `json:"author,omitempty"`
	BaseConfig     interface{}        `json:"baseConfig,omitempty"`
	ChildThemes    []Theme            `json:"childThemes,omitempty"`
	ConfigValues   interface{}        `json:"configValues,omitempty"`
	CreatedAt      time.Time          `json:"createdAt,omitempty"`
	CustomFields   interface{}        `json:"customFields,omitempty"`
	Description    string             `json:"description,omitempty"`
	HelpTexts      interface{}        `json:"helpTexts,omitempty"`
	Id             string             `json:"id,omitempty"`
	Labels         interface{}        `json:"labels,omitempty"`
	Media          []Media            `json:"media,omitempty"`
	Name           string             `json:"name,omitempty"`
	ParentThemeId  string             `json:"parentThemeId,omitempty"`
	PreviewMedia   *Media             `json:"previewMedia,omitempty"`
	PreviewMediaId string             `json:"previewMediaId,omitempty"`
	SalesChannels  []SalesChannel     `json:"salesChannels,omitempty"`
	TechnicalName  string             `json:"technicalName,omitempty"`
	Translated     interface{}        `json:"translated,omitempty"`
	Translations   []ThemeTranslation `json:"translations,omitempty"`
	UpdatedAt      time.Time          `json:"updatedAt,omitempty"`
}
