package model

import "time"

type ProductFeatureSetTranslation struct {
	CreatedAt           time.Time          `json:"createdAt,omitempty"`
	Description         string             `json:"description,omitempty"`
	Language            *Language          `json:"language,omitempty"`
	LanguageId          string             `json:"languageId,omitempty"`
	Name                string             `json:"name,omitempty"`
	ProductFeatureSet   *ProductFeatureSet `json:"productFeatureSet,omitempty"`
	ProductFeatureSetId string             `json:"productFeatureSetId,omitempty"`
	UpdatedAt           time.Time          `json:"updatedAt,omitempty"`
}
