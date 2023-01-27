package model

import "time"

type PromotionDiscount struct {
	ApplierKey              string                    `json:"applierKey,omitempty"`
	ConsiderAdvancedRules   bool                      `json:"considerAdvancedRules"` // required
	CreatedAt               time.Time                 `json:"createdAt,omitempty"`
	DiscountRules           []Rule                    `json:"discountRules,omitempty"`
	Id                      string                    `json:"id,omitempty"`
	MaxValue                float64                   `json:"maxValue"`
	PickerKey               string                    `json:"pickerKey,omitempty"`
	Promotion               *Promotion                `json:"promotion,omitempty"`
	PromotionDiscountPrices []PromotionDiscountPrices `json:"promotionDiscountPrices,omitempty"`
	PromotionId             string                    `json:"promotionId"` // required
	Scope                   string                    `json:"scope"`       // required
	SorterKey               string                    `json:"sorterKey,omitempty"`
	Type                    string                    `json:"type"` // required
	UpdatedAt               time.Time                 `json:"updatedAt,omitempty"`
	UsageKey                string                    `json:"usageKey,omitempty"`
	Value                   float64                   `json:"value"` // required
}
