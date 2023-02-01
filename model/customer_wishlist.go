package model

import "time"

type CustomerWishlist struct {
	CreatedAt      *time.Time                `json:"createdAt,omitempty"`
	Customer       *Customer                 `json:"customer,omitempty"`
	CustomerId     string                    `json:"customerId,omitempty"` // required
	CustomFields   interface{}               `json:"customFields,omitempty"`
	Id             string                    `json:"id,omitempty"`
	Products       []CustomerWishlistProduct `json:"products,omitempty"`
	SalesChannel   *SalesChannel             `json:"salesChannel,omitempty"`
	SalesChannelId string                    `json:"salesChannelId,omitempty"` // required
	UpdatedAt      *time.Time                `json:"updatedAt,omitempty"`
}
