package model

import "time"

type ProductSearchConfig struct {
	AndLogic        bool                       `json:"andLogic"` // required
	ConfigFields    []ProductSearchConfigField `json:"configFields,omitempty"`
	CreatedAt       time.Time                  `json:"createdAt,omitempty"`
	ExcludedTerms   interface{}                `json:"excludedTerms,omitempty"`
	Id              string                     `json:"id,omitempty"`
	Language        *Language                  `json:"language,omitempty"`
	LanguageId      string                     `json:"languageId"`      // required
	MinSearchLength int64                      `json:"minSearchLength"` // required
	UpdatedAt       time.Time                  `json:"updatedAt,omitempty"`
}
