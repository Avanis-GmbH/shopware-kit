package model

type PromotionSetgroupRule struct {
	Rule       *Rule              `json:"rule,omitempty"`
	RuleId     string             `json:"ruleId"` // required
	Setgroup   *PromotionSetgroup `json:"setgroup,omitempty"`
	SetgroupId string             `json:"setgroupId"` // required
}
