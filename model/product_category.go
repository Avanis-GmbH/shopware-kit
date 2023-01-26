package model

type ProductCategory struct {
	Category          *Category `json:"category,omitempty"`
	CategoryId        string    `json:"categoryId,omitempty"`
	CategoryVersionId string    `json:"categoryVersionId,omitempty"`
	Product           *Product  `json:"product,omitempty"`
	ProductId         string    `json:"productId,omitempty"`
	ProductVersionId  string    `json:"productVersionId,omitempty"`
}

type ProductCategoryCollection struct {
	EntityCollection

	Data []ProductCategory `json:"data"`
}
