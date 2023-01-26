package model

import "time"

type Integration struct {
	AccessKey            string         `json:"accessKey,omitempty"`
	AclRoles             []AclRole      `json:"aclRoles,omitempty"`
	Admin                bool           `json:"admin,omitempty"`
	App                  *App           `json:"app,omitempty"`
	CreatedAt            time.Time      `json:"createdAt,omitempty"`
	CreatedNotifications []Notification `json:"createdNotifications,omitempty"`
	CustomFields         interface{}    `json:"customFields,omitempty"`
	DeletedAt            time.Time      `json:"deletedAt,omitempty"`
	Id                   string         `json:"id,omitempty"`
	Label                string         `json:"label,omitempty"`
	LastUsageAt          time.Time      `json:"lastUsageAt,omitempty"`
	SecretAccessKey      interface{}    `json:"secretAccessKey,omitempty"`
	UpdatedAt            time.Time      `json:"updatedAt,omitempty"`
	WriteAccess          bool           `json:"writeAccess,omitempty"`
}
