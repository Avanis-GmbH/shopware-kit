package model

import "time"

type ProductSearchKeyword struct {
	CreatedAt        time.Time `json:"createdAt,omitempty"`
	Id               string    `json:"id,omitempty"`
	Keyword          string    `json:"keyword,omitempty"`
	Language         *Language `json:"language,omitempty"`
	LanguageId       string    `json:"languageId,omitempty"`
	Product          *Product  `json:"product,omitempty"`
	ProductId        string    `json:"productId,omitempty"`
	ProductVersionId string    `json:"productVersionId,omitempty"`
	Ranking          float64   `json:"ranking,omitempty"`
	UpdatedAt        time.Time `json:"updatedAt,omitempty"`
	VersionId        string    `json:"versionId,omitempty"`
}

type ProductSearchKeywordCollection struct {
	EntityCollection

	Data []ProductSearchKeyword `json:"data"`
}
