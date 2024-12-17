package model

import "time"

type ProductCrossSelling struct {
	Active           *bool                                 `json:"active,omitempty"`
	AssignedProducts []ProductCrossSellingAssignedProducts `json:"assignedProducts,omitempty"`
	CreatedAt        *time.Time                            `json:"createdAt,omitempty"`
	Id               string                                `json:"id,omitempty"`
	Limit            float64                               `json:"limit,omitempty"`
	Name             string                                `json:"name,omitempty"`     // required
	Position         float64                               `json:"position,omitempty"` // required
	Product          *Product                              `json:"product,omitempty"`
	ProductId        string                                `json:"productId,omitempty"` // required
	ProductStream    *ProductStream                        `json:"productStream,omitempty"`
	ProductStreamId  string                                `json:"productStreamId,omitempty"`
	ProductVersionId string                                `json:"productVersionId,omitempty"`
	SortBy           string                                `json:"sortBy,omitempty"`
	SortDirection    string                                `json:"sortDirection,omitempty"`
	Translated       interface{}                           `json:"translated,omitempty"`
	Translations     []ProductCrossSellingTranslation      `json:"translations,omitempty"`
	Type             string                                `json:"type,omitempty"` // required
	UpdatedAt        *time.Time                            `json:"updatedAt,omitempty"`
}
