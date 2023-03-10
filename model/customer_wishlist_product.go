package model

import "time"

type CustomerWishlistProduct struct {
	CreatedAt        *time.Time        `json:"createdAt,omitempty"`
	Id               string            `json:"id,omitempty"`
	Product          *Product          `json:"product,omitempty"`
	ProductId        string            `json:"productId,omitempty"` // required
	ProductVersionId string            `json:"productVersionId,omitempty"`
	UpdatedAt        *time.Time        `json:"updatedAt,omitempty"`
	Wishlist         *CustomerWishlist `json:"wishlist,omitempty"`
	WishlistId       string            `json:"wishlistId,omitempty"` // required
}
