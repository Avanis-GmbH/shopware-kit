package model

import "time"

type OrderCustomer struct {
	Company        string      `json:"company,omitempty"`
	CreatedAt      time.Time   `json:"createdAt,omitempty"`
	Customer       *Customer   `json:"customer,omitempty"`
	CustomerId     string      `json:"customerId,omitempty"`
	CustomerNumber string      `json:"customerNumber,omitempty"`
	CustomFields   interface{} `json:"customFields,omitempty"`
	Email          string      `json:"email,omitempty"`     // required
	FirstName      string      `json:"firstName,omitempty"` // required
	Id             string      `json:"id,omitempty"`
	LastName       string      `json:"lastName,omitempty"` // required
	Order          *Order      `json:"order,omitempty"`
	OrderId        string      `json:"orderId,omitempty"` // required
	OrderVersionId string      `json:"orderVersionId,omitempty"`
	RemoteAddress  interface{} `json:"remoteAddress,omitempty"`
	Salutation     *Salutation `json:"salutation,omitempty"`
	SalutationId   string      `json:"salutationId,omitempty"` // required
	Title          string      `json:"title,omitempty"`
	UpdatedAt      time.Time   `json:"updatedAt,omitempty"`
	VatIds         interface{} `json:"vatIds,omitempty"`
	VersionId      string      `json:"versionId,omitempty"`
}
