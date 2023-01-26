package model

import "time"

type ProductCrossSellingAssignedProducts struct {
	CreatedAt        time.Time            `json:"createdAt,omitempty"`
	CrossSelling     *ProductCrossSelling `json:"crossSelling,omitempty"`
	CrossSellingId   string               `json:"crossSellingId,omitempty"`
	Id               string               `json:"id,omitempty"`
	Position         float64              `json:"position,omitempty"`
	Product          *Product             `json:"product,omitempty"`
	ProductId        string               `json:"productId,omitempty"`
	ProductVersionId string               `json:"productVersionId,omitempty"`
	UpdatedAt        time.Time            `json:"updatedAt,omitempty"`
}
