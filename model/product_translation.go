package model

import "time"

type ProductTranslation struct {
	CreatedAt            time.Time   `json:"createdAt,omitempty"`
	CustomFields         interface{} `json:"customFields,omitempty"`
	CustomSearchKeywords interface{} `json:"customSearchKeywords,omitempty"`
	Description          string      `json:"description,omitempty"`
	Keywords             string      `json:"keywords,omitempty"`
	Language             *Language   `json:"language,omitempty"`
	LanguageId           string      `json:"languageId,omitempty"`
	MetaDescription      string      `json:"metaDescription,omitempty"`
	MetaTitle            string      `json:"metaTitle,omitempty"`
	Name                 string      `json:"name,omitempty"`
	PackUnit             string      `json:"packUnit,omitempty"`
	PackUnitPlural       string      `json:"packUnitPlural,omitempty"`
	Product              *Product    `json:"product,omitempty"`
	ProductId            string      `json:"productId,omitempty"`
	ProductVersionId     string      `json:"productVersionId,omitempty"`
	SlotConfig           interface{} `json:"slotConfig,omitempty"`
	UpdatedAt            time.Time   `json:"updatedAt,omitempty"`
}

type ProductTranslationCollection struct {
	EntityCollection

	Data []ProductTranslation `json:"data"`
}
