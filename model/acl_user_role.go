package model

import "time"

type AclUserRole struct {
	AclRole   *AclRole  `json:"aclRole,omitempty"`
	AclRoleId string    `json:"aclRoleId"` // required
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	User      *User     `json:"user,omitempty"`
	UserId    string    `json:"userId"` // required
}
