package model

type CustomerTag struct {
	Customer   *Customer `json:"customer,omitempty"`
	CustomerId string    `json:"customerId,omitempty"` // required
	Tag        *Tag      `json:"tag,omitempty"`
	TagId      string    `json:"tagId,omitempty"` // required
}
