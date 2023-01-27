package model

import "time"

type CustomerRecovery struct {
	CreatedAt  time.Time `json:"createdAt,omitempty"`
	Customer   *Customer `json:"customer,omitempty"`
	CustomerId string    `json:"customerId"` // required
	Hash       string    `json:"hash"`       // required
	Id         string    `json:"id,omitempty"`
	UpdatedAt  time.Time `json:"updatedAt,omitempty"`
}
