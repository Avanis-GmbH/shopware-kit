package model

import "time"

type FlowSequence struct {
	ActionName   string         `json:"actionName,omitempty"`
	Children     []FlowSequence `json:"children,omitempty"`
	Config       interface{}    `json:"config,omitempty"`
	CreatedAt    *time.Time     `json:"createdAt,omitempty"`
	CustomFields interface{}    `json:"customFields,omitempty"`
	DisplayGroup float64        `json:"displayGroup,omitempty"`
	Flow         *Flow          `json:"flow,omitempty"`
	FlowId       string         `json:"flowId,omitempty"` // required
	Id           string         `json:"id,omitempty"`
	Parent       *FlowSequence  `json:"parent,omitempty"`
	ParentId     string         `json:"parentId,omitempty"`
	Position     float64        `json:"position,omitempty"`
	Rule         *Rule          `json:"rule,omitempty"`
	RuleId       string         `json:"ruleId,omitempty"`
	TrueCase     bool           `json:"trueCase,omitempty"`
	UpdatedAt    *time.Time     `json:"updatedAt,omitempty"`
}
