package model

import "time"

type ProductSearchConfig struct {
	AndLogic        *bool                      `json:"andLogic,omitempty"` // required
	ConfigFields    []ProductSearchConfigField `json:"configFields,omitempty"`
	CreatedAt       *time.Time                 `json:"createdAt,omitempty"`
	ExcludedTerms   interface{}                `json:"excludedTerms,omitempty"`
	Id              string                     `json:"id,omitempty"`
	Language        *Language                  `json:"language,omitempty"`
	LanguageId      string                     `json:"languageId,omitempty"`      // required
	MinSearchLength int64                      `json:"minSearchLength,omitempty"` // required
	UpdatedAt       *time.Time                 `json:"updatedAt,omitempty"`
}
