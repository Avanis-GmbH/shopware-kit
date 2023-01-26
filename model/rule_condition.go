package model

import "time"

type RuleCondition struct {
	Children     []RuleCondition `json:"children,omitempty"`
	CreatedAt    time.Time       `json:"createdAt,omitempty"`
	CustomFields interface{}     `json:"customFields,omitempty"`
	Id           string          `json:"id,omitempty"`
	Parent       *RuleCondition  `json:"parent,omitempty"`
	ParentId     string          `json:"parentId,omitempty"`
	Position     float64         `json:"position,omitempty"`
	Rule         *Rule           `json:"rule,omitempty"`
	RuleId       string          `json:"ruleId,omitempty"`
	Type         string          `json:"type,omitempty"`
	UpdatedAt    time.Time       `json:"updatedAt,omitempty"`
	Value        interface{}     `json:"value,omitempty"`
}
