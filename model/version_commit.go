package model

import "time"

type VersionCommit struct {
	AutoIncrement float64             `json:"autoIncrement,omitempty"`
	CreatedAt     *time.Time          `json:"createdAt,omitempty"`
	Data          []VersionCommitData `json:"data,omitempty"`
	Id            string              `json:"id,omitempty"`
	IntegrationId string              `json:"integrationId,omitempty"`
	IsMerge       bool                `json:"isMerge"`
	Message       string              `json:"message,omitempty"`
	UpdatedAt     *time.Time          `json:"updatedAt,omitempty"`
	UserId        string              `json:"userId,omitempty"`
	Version       *Version            `json:"version,omitempty"`
	VersionId     string              `json:"versionId,omitempty"`
}
