package model

type EventActionRule struct {
	EventAction   *EventAction `json:"eventAction,omitempty"`
	EventActionId string       `json:"eventActionId,omitempty"`
	Rule          *Rule        `json:"rule,omitempty"`
	RuleId        string       `json:"ruleId,omitempty"`
}
