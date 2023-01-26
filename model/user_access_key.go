package model

import "time"

type UserAccessKey struct {
	AccessKey       string      `json:"accessKey,omitempty"`
	CreatedAt       time.Time   `json:"createdAt,omitempty"`
	CustomFields    interface{} `json:"customFields,omitempty"`
	Id              string      `json:"id,omitempty"`
	LastUsageAt     time.Time   `json:"lastUsageAt,omitempty"`
	SecretAccessKey interface{} `json:"secretAccessKey,omitempty"`
	UpdatedAt       time.Time   `json:"updatedAt,omitempty"`
	User            *User       `json:"user,omitempty"`
	UserId          string      `json:"userId,omitempty"`
	WriteAccess     bool        `json:"writeAccess,omitempty"`
}

type UserAccessKeyCollection struct {
	EntityCollection

	Data []UserAccessKey `json:"data"`
}
