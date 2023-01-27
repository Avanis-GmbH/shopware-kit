package model

import "time"

type PromotionDiscountPrices struct {
	CreatedAt         time.Time          `json:"createdAt,omitempty"`
	Currency          *Currency          `json:"currency,omitempty"`
	CurrencyId        string             `json:"currencyId,omitempty"` // required
	DiscountId        string             `json:"discountId,omitempty"` // required
	Id                string             `json:"id,omitempty"`
	Price             float64            `json:"price,omitempty"` // required
	PromotionDiscount *PromotionDiscount `json:"promotionDiscount,omitempty"`
	UpdatedAt         time.Time          `json:"updatedAt,omitempty"`
}
