package model

type ProductCategory struct {
	Category          *Category `json:"category,omitempty"`
	CategoryId        string    `json:"categoryId,omitempty"` // required
	CategoryVersionId string    `json:"categoryVersionId,omitempty"`
	Product           *Product  `json:"product,omitempty"`
	ProductId         string    `json:"productId,omitempty"` // required
	ProductVersionId  string    `json:"productVersionId,omitempty"`
}
