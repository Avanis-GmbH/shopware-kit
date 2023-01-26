package model

type ProductProperty struct {
	Option           *PropertyGroupOption `json:"option,omitempty"`
	OptionId         string               `json:"optionId,omitempty"`
	Product          *Product             `json:"product,omitempty"`
	ProductId        string               `json:"productId,omitempty"`
	ProductVersionId string               `json:"productVersionId,omitempty"`
}

type ProductPropertyCollection struct {
	EntityCollection

	Data []ProductProperty `json:"data"`
}
