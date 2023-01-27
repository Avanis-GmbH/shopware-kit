package model

import "time"

type CustomerRecovery struct {
	CreatedAt  time.Time `json:"createdAt,omitempty"`
	Customer   *Customer `json:"customer,omitempty"`
	CustomerId string    `json:"customerId,omitempty"` // required
	Hash       string    `json:"hash,omitempty"`       // required
	Id         string    `json:"id,omitempty"`
	UpdatedAt  time.Time `json:"updatedAt,omitempty"`
}
