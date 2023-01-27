package model

import "time"

type StateMachineHistory struct {
	CreatedAt             time.Time          `json:"createdAt,omitempty"`
	EntityId              interface{}        `json:"entityId"`    // required
	EntityName            string             `json:"entityName"`  // required
	FromStateId           string             `json:"fromStateId"` // required
	FromStateMachineState *StateMachineState `json:"fromStateMachineState,omitempty"`
	Id                    string             `json:"id,omitempty"`
	StateMachine          *StateMachine      `json:"stateMachine,omitempty"`
	StateMachineId        string             `json:"stateMachineId"` // required
	ToStateId             string             `json:"toStateId"`      // required
	ToStateMachineState   *StateMachineState `json:"toStateMachineState,omitempty"`
	TransitionActionName  string             `json:"transitionActionName,omitempty"`
	UpdatedAt             time.Time          `json:"updatedAt,omitempty"`
	User                  *User              `json:"user,omitempty"`
	UserId                string             `json:"userId,omitempty"`
}
