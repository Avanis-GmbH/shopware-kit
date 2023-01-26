package model

import "time"

type OrderTransaction struct {
	Amount            interface{}        `json:"amount,omitempty"`
	CreatedAt         time.Time          `json:"createdAt,omitempty"`
	CustomFields      interface{}        `json:"customFields,omitempty"`
	Id                string             `json:"id,omitempty"`
	Order             *Order             `json:"order,omitempty"`
	OrderId           string             `json:"orderId,omitempty"`
	OrderVersionId    string             `json:"orderVersionId,omitempty"`
	PaymentMethod     *PaymentMethod     `json:"paymentMethod,omitempty"`
	PaymentMethodId   string             `json:"paymentMethodId,omitempty"`
	StateId           string             `json:"stateId,omitempty"`
	StateMachineState *StateMachineState `json:"stateMachineState,omitempty"`
	UpdatedAt         time.Time          `json:"updatedAt,omitempty"`
	VersionId         string             `json:"versionId,omitempty"`
}
