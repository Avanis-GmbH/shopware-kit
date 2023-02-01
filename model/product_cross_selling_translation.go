package model

import "time"

type ProductCrossSellingTranslation struct {
	CreatedAt             *time.Time           `json:"createdAt,omitempty"`
	Language              *Language            `json:"language,omitempty"`
	LanguageId            string               `json:"languageId,omitempty"`
	Name                  string               `json:"name,omitempty"`
	ProductCrossSelling   *ProductCrossSelling `json:"productCrossSelling,omitempty"`
	ProductCrossSellingId string               `json:"productCrossSellingId,omitempty"`
	UpdatedAt             *time.Time           `json:"updatedAt,omitempty"`
}
