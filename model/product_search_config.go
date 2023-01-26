package model

import "time"

type ProductSearchConfig struct {
	AndLogic        bool                       `json:"andLogic,omitempty"`
	ConfigFields    []ProductSearchConfigField `json:"configFields,omitempty"`
	CreatedAt       time.Time                  `json:"createdAt,omitempty"`
	ExcludedTerms   interface{}                `json:"excludedTerms,omitempty"`
	Id              string                     `json:"id,omitempty"`
	Language        *Language                  `json:"language,omitempty"`
	LanguageId      string                     `json:"languageId,omitempty"`
	MinSearchLength float64                    `json:"minSearchLength,omitempty"`
	UpdatedAt       time.Time                  `json:"updatedAt,omitempty"`
}

type ProductSearchConfigCollection struct {
	EntityCollection

	Data []ProductSearchConfig `json:"data"`
}
