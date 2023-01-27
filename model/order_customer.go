package model

import "time"

type OrderCustomer struct {
	Company        string      `json:"company,omitempty"`
	CreatedAt      time.Time   `json:"createdAt,omitempty"`
	Customer       *Customer   `json:"customer,omitempty"`
	CustomerId     string      `json:"customerId,omitempty"`
	CustomerNumber string      `json:"customerNumber,omitempty"`
	CustomFields   interface{} `json:"customFields,omitempty"`
	Email          string      `json:"email"`     // required
	FirstName      string      `json:"firstName"` // required
	Id             string      `json:"id,omitempty"`
	LastName       string      `json:"lastName"` // required
	Order          *Order      `json:"order,omitempty"`
	OrderId        string      `json:"orderId"` // required
	OrderVersionId string      `json:"orderVersionId,omitempty"`
	RemoteAddress  interface{} `json:"remoteAddress,omitempty"`
	Salutation     *Salutation `json:"salutation,omitempty"`
	SalutationId   string      `json:"salutationId"` // required
	Title          string      `json:"title,omitempty"`
	UpdatedAt      time.Time   `json:"updatedAt,omitempty"`
	VatIds         interface{} `json:"vatIds,omitempty"`
	VersionId      string      `json:"versionId,omitempty"`
}
