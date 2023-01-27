package model

import "time"

type ProductCrossSellingAssignedProducts struct {
	CreatedAt        time.Time            `json:"createdAt,omitempty"`
	CrossSelling     *ProductCrossSelling `json:"crossSelling,omitempty"`
	CrossSellingId   string               `json:"crossSellingId"` // required
	Id               string               `json:"id,omitempty"`
	Position         float64              `json:"position"`
	Product          *Product             `json:"product,omitempty"`
	ProductId        string               `json:"productId"` // required
	ProductVersionId string               `json:"productVersionId,omitempty"`
	UpdatedAt        time.Time            `json:"updatedAt,omitempty"`
}
