package model

import "time"

type VersionCommitData struct {
	Action          string         `json:"action,omitempty"`
	AutoIncrement   float64        `json:"autoIncrement,omitempty"`
	Commit          *VersionCommit `json:"commit,omitempty"`
	CreatedAt       time.Time      `json:"createdAt,omitempty"`
	EntityId        interface{}    `json:"entityId,omitempty"`
	EntityName      string         `json:"entityName,omitempty"`
	Id              string         `json:"id,omitempty"`
	IntegrationId   string         `json:"integrationId,omitempty"`
	Payload         interface{}    `json:"payload,omitempty"`
	UpdatedAt       time.Time      `json:"updatedAt,omitempty"`
	UserId          string         `json:"userId,omitempty"`
	VersionCommitId string         `json:"versionCommitId,omitempty"`
}
