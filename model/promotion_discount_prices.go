package model

import "time"

type PromotionDiscountPrices struct {
	CreatedAt         time.Time          `json:"createdAt,omitempty"`
	Currency          *Currency          `json:"currency,omitempty"`
	CurrencyId        string             `json:"currencyId,omitempty"`
	DiscountId        string             `json:"discountId,omitempty"`
	Id                string             `json:"id,omitempty"`
	Price             float64            `json:"price,omitempty"`
	PromotionDiscount *PromotionDiscount `json:"promotionDiscount,omitempty"`
	UpdatedAt         time.Time          `json:"updatedAt,omitempty"`
}

type PromotionDiscountPricesCollection struct {
	EntityCollection

	Data []PromotionDiscountPrices `json:"data"`
}
