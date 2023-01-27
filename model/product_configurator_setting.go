package model

import "time"

type ProductConfiguratorSetting struct {
	CreatedAt        time.Time            `json:"createdAt,omitempty"`
	CustomFields     interface{}          `json:"customFields,omitempty"`
	Id               string               `json:"id,omitempty"`
	Media            *Media               `json:"media,omitempty"`
	MediaId          string               `json:"mediaId,omitempty"`
	Option           *PropertyGroupOption `json:"option,omitempty"`
	OptionId         string               `json:"optionId"` // required
	Position         float64              `json:"position"`
	Price            interface{}          `json:"price,omitempty"`
	Product          *Product             `json:"product,omitempty"`
	ProductId        string               `json:"productId"` // required
	ProductVersionId string               `json:"productVersionId,omitempty"`
	UpdatedAt        time.Time            `json:"updatedAt,omitempty"`
	VersionId        string               `json:"versionId,omitempty"`
}
