package model

type ProductStreamMapping struct {
	Product          *Product       `json:"product,omitempty"`
	ProductId        string         `json:"productId,omitempty"`
	ProductStream    *ProductStream `json:"productStream,omitempty"`
	ProductStreamId  string         `json:"productStreamId,omitempty"`
	ProductVersionId string         `json:"productVersionId,omitempty"`
}

type ProductStreamMappingCollection struct {
	EntityCollection

	Data []ProductStreamMapping `json:"data"`
}
