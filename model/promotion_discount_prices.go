package model

import "time"

type PromotionDiscountPrices struct {
	CreatedAt         time.Time          `json:"createdAt,omitempty"`
	Currency          *Currency          `json:"currency,omitempty"`
	CurrencyId        string             `json:"currencyId"` // required
	DiscountId        string             `json:"discountId"` // required
	Id                string             `json:"id,omitempty"`
	Price             float64            `json:"price"` // required
	PromotionDiscount *PromotionDiscount `json:"promotionDiscount,omitempty"`
	UpdatedAt         time.Time          `json:"updatedAt,omitempty"`
}
