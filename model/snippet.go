package model

import "time"

type Snippet struct {
	Author         string      `json:"author"` // required
	CreatedAt      time.Time   `json:"createdAt,omitempty"`
	CustomFields   interface{} `json:"customFields,omitempty"`
	Id             string      `json:"id,omitempty"`
	Set            *SnippetSet `json:"set,omitempty"`
	SetId          string      `json:"setId"`          // required
	TranslationKey string      `json:"translationKey"` // required
	UpdatedAt      time.Time   `json:"updatedAt,omitempty"`
	Value          string      `json:"value"` // required
}
