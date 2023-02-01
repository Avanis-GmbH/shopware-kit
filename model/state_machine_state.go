package model

import "time"

type StateMachineState struct {
	CreatedAt                      *time.Time                     `json:"createdAt,omitempty"`
	CustomFields                   interface{}                    `json:"customFields,omitempty"`
	FromStateMachineHistoryEntries []StateMachineHistory          `json:"fromStateMachineHistoryEntries,omitempty"`
	FromStateMachineTransitions    []StateMachineTransition       `json:"fromStateMachineTransitions,omitempty"`
	Id                             string                         `json:"id,omitempty"`
	Name                           string                         `json:"name,omitempty"` // required
	OrderDeliveries                []OrderDelivery                `json:"orderDeliveries,omitempty"`
	Orders                         []Order                        `json:"orders,omitempty"`
	OrderTransactions              []OrderTransaction             `json:"orderTransactions,omitempty"`
	StateMachine                   *StateMachine                  `json:"stateMachine,omitempty"`
	StateMachineId                 string                         `json:"stateMachineId,omitempty"` // required
	TechnicalName                  string                         `json:"technicalName,omitempty"`  // required
	ToStateMachineHistoryEntries   []StateMachineHistory          `json:"toStateMachineHistoryEntries,omitempty"`
	ToStateMachineTransitions      []StateMachineTransition       `json:"toStateMachineTransitions,omitempty"`
	Translated                     interface{}                    `json:"translated,omitempty"`
	Translations                   []StateMachineStateTranslation `json:"translations,omitempty"`
	UpdatedAt                      *time.Time                     `json:"updatedAt,omitempty"`
}
