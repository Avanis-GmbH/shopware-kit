package model

import "time"

type PromotionDiscount struct {
	ApplierKey              string                    `json:"applierKey,omitempty"`
	ConsiderAdvancedRules   bool                      `json:"considerAdvancedRules,omitempty"`
	CreatedAt               time.Time                 `json:"createdAt,omitempty"`
	DiscountRules           []Rule                    `json:"discountRules,omitempty"`
	Id                      string                    `json:"id,omitempty"`
	MaxValue                float64                   `json:"maxValue,omitempty"`
	PickerKey               string                    `json:"pickerKey,omitempty"`
	Promotion               *Promotion                `json:"promotion,omitempty"`
	PromotionDiscountPrices []PromotionDiscountPrices `json:"promotionDiscountPrices,omitempty"`
	PromotionId             string                    `json:"promotionId,omitempty"`
	Scope                   string                    `json:"scope,omitempty"`
	SorterKey               string                    `json:"sorterKey,omitempty"`
	Type                    string                    `json:"type,omitempty"`
	UpdatedAt               time.Time                 `json:"updatedAt,omitempty"`
	UsageKey                string                    `json:"usageKey,omitempty"`
	Value                   float64                   `json:"value,omitempty"`
}

type PromotionDiscountCollection struct {
	EntityCollection

	Data []PromotionDiscount `json:"data"`
}
