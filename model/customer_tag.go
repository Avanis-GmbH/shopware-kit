package model

type CustomerTag struct {
	Customer   *Customer `json:"customer,omitempty"`
	CustomerId string    `json:"customerId,omitempty"`
	Tag        *Tag      `json:"tag,omitempty"`
	TagId      string    `json:"tagId,omitempty"`
}
