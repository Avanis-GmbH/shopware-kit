package model

type ProductTag struct {
	Product          *Product `json:"product,omitempty"`
	ProductId        string   `json:"productId,omitempty"`
	ProductVersionId string   `json:"productVersionId,omitempty"`
	Tag              *Tag     `json:"tag,omitempty"`
	TagId            string   `json:"tagId,omitempty"`
}

type ProductTagCollection struct {
	EntityCollection

	Data []ProductTag `json:"data"`
}
