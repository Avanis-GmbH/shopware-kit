package model

import "time"

type MediaThumbnail struct {
	CreatedAt    time.Time   `json:"createdAt,omitempty"`
	CustomFields interface{} `json:"customFields,omitempty"`
	Height       float64     `json:"height"` // required
	Id           string      `json:"id,omitempty"`
	Media        *Media      `json:"media,omitempty"`
	MediaId      string      `json:"mediaId"` // required
	UpdatedAt    time.Time   `json:"updatedAt,omitempty"`
	Url          string      `json:"url,omitempty"`
	Width        float64     `json:"width"` // required
}
