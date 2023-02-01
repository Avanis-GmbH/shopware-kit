package model

import "time"

type UserRecovery struct {
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Hash      string     `json:"hash,omitempty"` // required
	Id        string     `json:"id,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	User      *User      `json:"user,omitempty"`
	UserId    string     `json:"userId,omitempty"` // required
}
