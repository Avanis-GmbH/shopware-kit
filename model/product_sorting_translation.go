package model

import "time"

type ProductSortingTranslation struct {
	CreatedAt        *time.Time      `json:"createdAt,omitempty"`
	Label            string          `json:"label,omitempty"`
	Language         *Language       `json:"language,omitempty"`
	LanguageId       string          `json:"languageId,omitempty"`
	ProductSorting   *ProductSorting `json:"productSorting,omitempty"`
	ProductSortingId string          `json:"productSortingId,omitempty"`
	UpdatedAt        *time.Time      `json:"updatedAt,omitempty"`
}
