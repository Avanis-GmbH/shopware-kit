package model

import "time"

type ProductSearchKeyword struct {
	CreatedAt        time.Time `json:"createdAt,omitempty"`
	Id               string    `json:"id,omitempty"`
	Keyword          string    `json:"keyword"` // required
	Language         *Language `json:"language,omitempty"`
	LanguageId       string    `json:"languageId"` // required
	Product          *Product  `json:"product,omitempty"`
	ProductId        string    `json:"productId"` // required
	ProductVersionId string    `json:"productVersionId,omitempty"`
	Ranking          float64   `json:"ranking"` // required
	UpdatedAt        time.Time `json:"updatedAt,omitempty"`
	VersionId        string    `json:"versionId,omitempty"`
}
