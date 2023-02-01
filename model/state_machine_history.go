package model

import "time"

type StateMachineHistory struct {
	CreatedAt             *time.Time         `json:"createdAt,omitempty"`
	EntityId              interface{}        `json:"entityId,omitempty"`    // required
	EntityName            string             `json:"entityName,omitempty"`  // required
	FromStateId           string             `json:"fromStateId,omitempty"` // required
	FromStateMachineState *StateMachineState `json:"fromStateMachineState,omitempty"`
	Id                    string             `json:"id,omitempty"`
	StateMachine          *StateMachine      `json:"stateMachine,omitempty"`
	StateMachineId        string             `json:"stateMachineId,omitempty"` // required
	ToStateId             string             `json:"toStateId,omitempty"`      // required
	ToStateMachineState   *StateMachineState `json:"toStateMachineState,omitempty"`
	TransitionActionName  string             `json:"transitionActionName,omitempty"`
	UpdatedAt             *time.Time         `json:"updatedAt,omitempty"`
	User                  *User              `json:"user,omitempty"`
	UserId                string             `json:"userId,omitempty"`
}
