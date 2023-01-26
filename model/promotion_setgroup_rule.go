package model

type PromotionSetgroupRule struct {
	Rule       *Rule              `json:"rule,omitempty"`
	RuleId     string             `json:"ruleId,omitempty"`
	Setgroup   *PromotionSetgroup `json:"setgroup,omitempty"`
	SetgroupId string             `json:"setgroupId,omitempty"`
}

type PromotionSetgroupRuleCollection struct {
	EntityCollection

	Data []PromotionSetgroupRule `json:"data"`
}
