package model

import "time"

type MediaThumbnail struct {
	CreatedAt    time.Time   `json:"createdAt,omitempty"`
	CustomFields interface{} `json:"customFields,omitempty"`
	Height       float64     `json:"height,omitempty"` // required
	Id           string      `json:"id,omitempty"`
	Media        *Media      `json:"media,omitempty"`
	MediaId      string      `json:"mediaId,omitempty"` // required
	UpdatedAt    time.Time   `json:"updatedAt,omitempty"`
	Url          string      `json:"url,omitempty"`
	Width        float64     `json:"width,omitempty"` // required
}
