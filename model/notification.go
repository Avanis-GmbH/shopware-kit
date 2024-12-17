package model

import "time"

type Notification struct {
	AdminOnly              *bool        `json:"adminOnly,omitempty"`
	CreatedAt              *time.Time   `json:"createdAt,omitempty"`
	CreatedByIntegration   *Integration `json:"createdByIntegration,omitempty"`
	CreatedByIntegrationId string       `json:"createdByIntegrationId,omitempty"`
	CreatedByUser          *User        `json:"createdByUser,omitempty"`
	CreatedByUserId        string       `json:"createdByUserId,omitempty"`
	Id                     string       `json:"id,omitempty"`
	Message                string       `json:"message,omitempty"` // required
	RequiredPrivileges     interface{}  `json:"requiredPrivileges,omitempty"`
	Status                 string       `json:"status,omitempty"` // required
	UpdatedAt              *time.Time   `json:"updatedAt,omitempty"`
}
