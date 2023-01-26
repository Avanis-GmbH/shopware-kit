package model

type PromotionPersonaRule struct {
	PromotionId string `json:"promotionId,omitempty"`

	RuleId string `json:"ruleId,omitempty"`

	Promotion *Promotion `json:"promotion,omitempty"`

	Rule *Rule `json:"rule,omitempty"`
}
