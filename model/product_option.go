package model

type ProductOption struct {
	Option           *PropertyGroupOption `json:"option,omitempty"`
	OptionId         string               `json:"optionId"` // required
	Product          *Product             `json:"product,omitempty"`
	ProductId        string               `json:"productId"` // required
	ProductVersionId string               `json:"productVersionId,omitempty"`
}
