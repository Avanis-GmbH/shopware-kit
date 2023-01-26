package model

import "time"

type Version struct {
	Commits   []VersionCommit `json:"commits,omitempty"`
	CreatedAt time.Time       `json:"createdAt,omitempty"`
	Id        string          `json:"id,omitempty"`
	Name      string          `json:"name,omitempty"`
	UpdatedAt time.Time       `json:"updatedAt,omitempty"`
}

type VersionCollection struct {
	EntityCollection

	Data []Version `json:"data"`
}
