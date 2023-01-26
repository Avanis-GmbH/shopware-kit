package model

import "time"

type StateMachineHistory struct {
	CreatedAt             time.Time          `json:"createdAt,omitempty"`
	EntityId              interface{}        `json:"entityId,omitempty"`
	EntityName            string             `json:"entityName,omitempty"`
	FromStateId           string             `json:"fromStateId,omitempty"`
	FromStateMachineState *StateMachineState `json:"fromStateMachineState,omitempty"`
	Id                    string             `json:"id,omitempty"`
	StateMachine          *StateMachine      `json:"stateMachine,omitempty"`
	StateMachineId        string             `json:"stateMachineId,omitempty"`
	ToStateId             string             `json:"toStateId,omitempty"`
	ToStateMachineState   *StateMachineState `json:"toStateMachineState,omitempty"`
	TransitionActionName  string             `json:"transitionActionName,omitempty"`
	UpdatedAt             time.Time          `json:"updatedAt,omitempty"`
	User                  *User              `json:"user,omitempty"`
	UserId                string             `json:"userId,omitempty"`
}
