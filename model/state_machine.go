package model

import "time"

type StateMachine struct {
	CreatedAt      time.Time                 `json:"createdAt,omitempty"`
	CustomFields   interface{}               `json:"customFields,omitempty"`
	HistoryEntries []StateMachineHistory     `json:"historyEntries,omitempty"`
	Id             string                    `json:"id,omitempty"`
	InitialStateId string                    `json:"initialStateId,omitempty"`
	Name           string                    `json:"name"` // required
	States         []StateMachineState       `json:"states,omitempty"`
	TechnicalName  string                    `json:"technicalName"` // required
	Transitions    []StateMachineTransition  `json:"transitions,omitempty"`
	Translated     interface{}               `json:"translated,omitempty"`
	Translations   []StateMachineTranslation `json:"translations,omitempty"`
	UpdatedAt      time.Time                 `json:"updatedAt,omitempty"`
}
