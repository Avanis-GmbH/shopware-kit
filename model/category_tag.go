package model

type CategoryTag struct {
	Category          *Category `json:"category,omitempty"`
	CategoryId        string    `json:"categoryId,omitempty"`
	CategoryVersionId string    `json:"categoryVersionId,omitempty"`
	Tag               *Tag      `json:"tag,omitempty"`
	TagId             string    `json:"tagId,omitempty"`
}
