package model

import "time"

type StateMachineTranslation struct {
	CreatedAt      time.Time     `json:"createdAt,omitempty"`
	CustomFields   interface{}   `json:"customFields,omitempty"`
	Language       *Language     `json:"language,omitempty"`
	LanguageId     string        `json:"languageId,omitempty"`
	Name           string        `json:"name,omitempty"`
	StateMachine   *StateMachine `json:"stateMachine,omitempty"`
	StateMachineId string        `json:"stateMachineId,omitempty"`
	UpdatedAt      time.Time     `json:"updatedAt,omitempty"`
}

type StateMachineTranslationCollection struct {
	EntityCollection

	Data []StateMachineTranslation `json:"data"`
}
