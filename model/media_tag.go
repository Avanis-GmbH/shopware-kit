package model

type MediaTag struct {
	Media   *Media `json:"media,omitempty"`
	MediaId string `json:"mediaId,omitempty"`
	Tag     *Tag   `json:"tag,omitempty"`
	TagId   string `json:"tagId,omitempty"`
}
