package model

type PromotionCartRule struct {
	Promotion   *Promotion `json:"promotion,omitempty"`
	PromotionId string     `json:"promotionId,omitempty"` // required
	Rule        *Rule      `json:"rule,omitempty"`
	RuleId      string     `json:"ruleId,omitempty"` // required
}
