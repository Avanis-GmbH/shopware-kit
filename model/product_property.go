package model

type ProductProperty struct {
	Option           *PropertyGroupOption `json:"option,omitempty"`
	OptionId         string               `json:"optionId,omitempty"` // required
	Product          *Product             `json:"product,omitempty"`
	ProductId        string               `json:"productId,omitempty"` // required
	ProductVersionId string               `json:"productVersionId,omitempty"`
}
