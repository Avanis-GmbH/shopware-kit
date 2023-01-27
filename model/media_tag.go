package model

type MediaTag struct {
	Media   *Media `json:"media,omitempty"`
	MediaId string `json:"mediaId,omitempty"` // required
	Tag     *Tag   `json:"tag,omitempty"`
	TagId   string `json:"tagId,omitempty"` // required
}
