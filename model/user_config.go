package model

import "time"

type UserConfig struct {
	CreatedAt *time.Time  `json:"createdAt,omitempty"`
	Id        string      `json:"id,omitempty"`
	Key       string      `json:"key,omitempty"` // required
	UpdatedAt *time.Time  `json:"updatedAt,omitempty"`
	User      *User       `json:"user,omitempty"`
	UserId    string      `json:"userId,omitempty"` // required
	Value     interface{} `json:"value,omitempty"`
}
