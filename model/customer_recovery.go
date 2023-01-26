package model

import "time"

type CustomerRecovery struct {
	CreatedAt  time.Time `json:"createdAt,omitempty"`
	Customer   *Customer `json:"customer,omitempty"`
	CustomerId string    `json:"customerId,omitempty"`
	Hash       string    `json:"hash,omitempty"`
	Id         string    `json:"id,omitempty"`
	UpdatedAt  time.Time `json:"updatedAt,omitempty"`
}

type CustomerRecoveryCollection struct {
	EntityCollection

	Data []CustomerRecovery `json:"data"`
}
