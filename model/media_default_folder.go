package model

import "time"

type MediaDefaultFolder struct {
	AssociationFields []string     `json:"associationFields,omitempty"` // required
	CreatedAt         *time.Time   `json:"createdAt,omitempty"`
	CustomFields      interface{}  `json:"customFields,omitempty"`
	Entity            string       `json:"entity,omitempty"` // required
	Folder            *MediaFolder `json:"folder,omitempty"`
	Id                string       `json:"id,omitempty"`
	UpdatedAt         *time.Time   `json:"updatedAt,omitempty"`
}
