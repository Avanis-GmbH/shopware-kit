package model

import "time"

type StateMachineStateTranslation struct {
	CreatedAt           time.Time          `json:"createdAt,omitempty"`
	CustomFields        interface{}        `json:"customFields,omitempty"`
	Language            *Language          `json:"language,omitempty"`
	LanguageId          string             `json:"languageId,omitempty"`
	Name                string             `json:"name,omitempty"`
	StateMachineState   *StateMachineState `json:"stateMachineState,omitempty"`
	StateMachineStateId string             `json:"stateMachineStateId,omitempty"`
	UpdatedAt           time.Time          `json:"updatedAt,omitempty"`
}

type StateMachineStateTranslationCollection struct {
	EntityCollection

	Data []StateMachineStateTranslation `json:"data"`
}
