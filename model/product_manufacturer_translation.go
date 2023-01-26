package model

import "time"

type ProductManufacturerTranslation struct {
	CreatedAt                    time.Time            `json:"createdAt,omitempty"`
	CustomFields                 interface{}          `json:"customFields,omitempty"`
	Description                  string               `json:"description,omitempty"`
	Language                     *Language            `json:"language,omitempty"`
	LanguageId                   string               `json:"languageId,omitempty"`
	Name                         string               `json:"name,omitempty"`
	ProductManufacturer          *ProductManufacturer `json:"productManufacturer,omitempty"`
	ProductManufacturerId        string               `json:"productManufacturerId,omitempty"`
	ProductManufacturerVersionId string               `json:"productManufacturerVersionId,omitempty"`
	UpdatedAt                    time.Time            `json:"updatedAt,omitempty"`
}

type ProductManufacturerTranslationCollection struct {
	EntityCollection

	Data []ProductManufacturerTranslation `json:"data"`
}
