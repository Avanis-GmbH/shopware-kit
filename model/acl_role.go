package model

import "time"

type AclRole struct {
	App          *App          `json:"app,omitempty"`
	CreatedAt    time.Time     `json:"createdAt,omitempty"`
	DeletedAt    time.Time     `json:"deletedAt,omitempty"`
	Description  string        `json:"description,omitempty"`
	Id           string        `json:"id,omitempty"`
	Integrations []Integration `json:"integrations,omitempty"`
	Name         string        `json:"name,omitempty"`
	Privileges   interface{}   `json:"privileges,omitempty"`
	UpdatedAt    time.Time     `json:"updatedAt,omitempty"`
	Users        []User        `json:"users,omitempty"`
}

type AclRoleCollection struct {
	EntityCollection

	Data []AclRole `json:"data"`
}
