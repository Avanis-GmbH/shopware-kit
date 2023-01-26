package model

type ProductOption struct {
	Option           *PropertyGroupOption `json:"option,omitempty"`
	OptionId         string               `json:"optionId,omitempty"`
	Product          *Product             `json:"product,omitempty"`
	ProductId        string               `json:"productId,omitempty"`
	ProductVersionId string               `json:"productVersionId,omitempty"`
}
