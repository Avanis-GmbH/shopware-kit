package model

import "time"

type MediaDefaultFolder struct {
	AssociationFields interface{}  `json:"associationFields,omitempty"`
	CreatedAt         time.Time    `json:"createdAt,omitempty"`
	CustomFields      interface{}  `json:"customFields,omitempty"`
	Entity            string       `json:"entity,omitempty"`
	Folder            *MediaFolder `json:"folder,omitempty"`
	Id                string       `json:"id,omitempty"`
	UpdatedAt         time.Time    `json:"updatedAt,omitempty"`
}

type MediaDefaultFolderCollection struct {
	EntityCollection

	Data []MediaDefaultFolder `json:"data"`
}
