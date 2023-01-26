package model

type PromotionCartRule struct {
	Promotion   *Promotion `json:"promotion,omitempty"`
	PromotionId string     `json:"promotionId,omitempty"`
	Rule        *Rule      `json:"rule,omitempty"`
	RuleId      string     `json:"ruleId,omitempty"`
}

type PromotionCartRuleCollection struct {
	EntityCollection

	Data []PromotionCartRule `json:"data"`
}
