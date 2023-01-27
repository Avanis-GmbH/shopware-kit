package model

type ProductStreamMapping struct {
	Product          *Product       `json:"product,omitempty"`
	ProductId        string         `json:"productId"` // required
	ProductStream    *ProductStream `json:"productStream,omitempty"`
	ProductStreamId  string         `json:"productStreamId"` // required
	ProductVersionId string         `json:"productVersionId,omitempty"`
}
