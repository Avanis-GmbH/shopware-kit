package model

type ProductStreamMapping struct {
	Product          *Product       `json:"product,omitempty"`
	ProductId        string         `json:"productId,omitempty"` // required
	ProductStream    *ProductStream `json:"productStream,omitempty"`
	ProductStreamId  string         `json:"productStreamId,omitempty"` // required
	ProductVersionId string         `json:"productVersionId,omitempty"`
}
