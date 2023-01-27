package model

import "time"

type PropertyGroupOption struct {
	ColorHexCode                string                           `json:"colorHexCode,omitempty"`
	CreatedAt                   time.Time                        `json:"createdAt,omitempty"`
	CustomFields                interface{}                      `json:"customFields,omitempty"`
	Group                       *PropertyGroup                   `json:"group,omitempty"`
	GroupId                     string                           `json:"groupId"` // required
	Id                          string                           `json:"id,omitempty"`
	Media                       *Media                           `json:"media,omitempty"`
	MediaId                     string                           `json:"mediaId,omitempty"`
	Name                        string                           `json:"name"` // required
	Position                    float64                          `json:"position"`
	ProductConfiguratorSettings []ProductConfiguratorSetting     `json:"productConfiguratorSettings,omitempty"`
	ProductOptions              []Product                        `json:"productOptions,omitempty"`
	ProductProperties           []Product                        `json:"productProperties,omitempty"`
	Translated                  interface{}                      `json:"translated,omitempty"`
	Translations                []PropertyGroupOptionTranslation `json:"translations,omitempty"`
	UpdatedAt                   time.Time                        `json:"updatedAt,omitempty"`
}
