package model

import "time"

type ImportExportLog struct {
	Activity            string               `json:"activity,omitempty"`
	Config              interface{}          `json:"config,omitempty"`
	CreatedAt           time.Time            `json:"createdAt,omitempty"`
	FailedImportLog     *ImportExportLog     `json:"failedImportLog,omitempty"`
	File                *ImportExportFile    `json:"file,omitempty"`
	FileId              string               `json:"fileId,omitempty"`
	Id                  string               `json:"id,omitempty"`
	InvalidRecordsLog   *ImportExportLog     `json:"invalidRecordsLog,omitempty"`
	InvalidRecordsLogId string               `json:"invalidRecordsLogId,omitempty"`
	Profile             *ImportExportProfile `json:"profile,omitempty"`
	ProfileId           string               `json:"profileId,omitempty"`
	ProfileName         string               `json:"profileName,omitempty"`
	Records             float64              `json:"records,omitempty"`
	Result              interface{}          `json:"result,omitempty"`
	State               string               `json:"state,omitempty"`
	UpdatedAt           time.Time            `json:"updatedAt,omitempty"`
	User                *User                `json:"user,omitempty"`
	UserId              string               `json:"userId,omitempty"`
	Username            string               `json:"username,omitempty"`
}

type ImportExportLogCollection struct {
	EntityCollection

	Data []ImportExportLog `json:"data"`
}
