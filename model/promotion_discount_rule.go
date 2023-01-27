package model

type PromotionDiscountRule struct {
	Discount   *PromotionDiscount `json:"discount,omitempty"`
	DiscountId string             `json:"discountId"` // required
	Rule       *Rule              `json:"rule,omitempty"`
	RuleId     string             `json:"ruleId"` // required
}
