package model

type RuleTag struct {
	Id     string `json:"id,omitempty"`
	Rule   *Rule  `json:"rule,omitempty"`
	RuleId string `json:"ruleId"` // required
	Tag    *Tag   `json:"tag,omitempty"`
	TagId  string `json:"tagId"` // required
}
