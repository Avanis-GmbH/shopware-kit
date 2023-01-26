package model

import "time"

type UserConfig struct {
	CreatedAt time.Time   `json:"createdAt,omitempty"`
	Id        string      `json:"id,omitempty"`
	Key       string      `json:"key,omitempty"`
	UpdatedAt time.Time   `json:"updatedAt,omitempty"`
	User      *User       `json:"user,omitempty"`
	UserId    string      `json:"userId,omitempty"`
	Value     interface{} `json:"value,omitempty"`
}

type UserConfigCollection struct {
	EntityCollection

	Data []UserConfig `json:"data"`
}
