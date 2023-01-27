package model

import "time"

type MediaFolder struct {
	ChildCount             float64                   `json:"childCount,omitempty"`
	Children               []MediaFolder             `json:"children,omitempty"`
	Configuration          *MediaFolderConfiguration `json:"configuration,omitempty"`
	ConfigurationId        string                    `json:"configurationId,omitempty"` // required
	CreatedAt              time.Time                 `json:"createdAt,omitempty"`
	CustomFields           interface{}               `json:"customFields,omitempty"`
	DefaultFolder          *MediaDefaultFolder       `json:"defaultFolder,omitempty"`
	DefaultFolderId        string                    `json:"defaultFolderId,omitempty"`
	Id                     string                    `json:"id,omitempty"`
	Media                  []Media                   `json:"media,omitempty"`
	Name                   string                    `json:"name,omitempty"` // required
	Parent                 *MediaFolder              `json:"parent,omitempty"`
	ParentId               string                    `json:"parentId,omitempty"`
	UpdatedAt              time.Time                 `json:"updatedAt,omitempty"`
	UseParentConfiguration bool                      `json:"useParentConfiguration,omitempty"`
}
