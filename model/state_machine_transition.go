package model

import "time"

type StateMachineTransition struct {
	ActionName            string             `json:"actionName,omitempty"`
	CreatedAt             time.Time          `json:"createdAt,omitempty"`
	CustomFields          interface{}        `json:"customFields,omitempty"`
	FromStateId           string             `json:"fromStateId,omitempty"`
	FromStateMachineState *StateMachineState `json:"fromStateMachineState,omitempty"`
	Id                    string             `json:"id,omitempty"`
	StateMachine          *StateMachine      `json:"stateMachine,omitempty"`
	StateMachineId        string             `json:"stateMachineId,omitempty"`
	ToStateId             string             `json:"toStateId,omitempty"`
	ToStateMachineState   *StateMachineState `json:"toStateMachineState,omitempty"`
	UpdatedAt             time.Time          `json:"updatedAt,omitempty"`
}
