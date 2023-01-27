package model

type ProductTag struct {
	Product          *Product `json:"product,omitempty"`
	ProductId        string   `json:"productId,omitempty"` // required
	ProductVersionId string   `json:"productVersionId,omitempty"`
	Tag              *Tag     `json:"tag,omitempty"`
	TagId            string   `json:"tagId,omitempty"` // required
}
