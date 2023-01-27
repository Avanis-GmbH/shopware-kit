package model

type ProductCategoryTree struct {
	Category          *Category `json:"category,omitempty"`
	CategoryId        string    `json:"categoryId"`
	CategoryVersionId string    `json:"categoryVersionId,omitempty"`
	Product           *Product  `json:"product,omitempty"`
	ProductId         string    `json:"productId"`
	ProductVersionId  string    `json:"productVersionId,omitempty"`
}
