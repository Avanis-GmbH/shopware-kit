package model

import "time"

type StateMachineTransition struct {
	ActionName            string             `json:"actionName,omitempty"` // required
	CreatedAt             *time.Time         `json:"createdAt,omitempty"`
	CustomFields          interface{}        `json:"customFields,omitempty"`
	FromStateId           string             `json:"fromStateId,omitempty"` // required
	FromStateMachineState *StateMachineState `json:"fromStateMachineState,omitempty"`
	Id                    string             `json:"id,omitempty"`
	StateMachine          *StateMachine      `json:"stateMachine,omitempty"`
	StateMachineId        string             `json:"stateMachineId,omitempty"` // required
	ToStateId             string             `json:"toStateId,omitempty"`      // required
	ToStateMachineState   *StateMachineState `json:"toStateMachineState,omitempty"`
	UpdatedAt             *time.Time         `json:"updatedAt,omitempty"`
}
