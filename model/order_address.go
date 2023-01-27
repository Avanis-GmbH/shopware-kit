package model

import "time"

type OrderAddress struct {
	AdditionalAddressLine1 string          `json:"additionalAddressLine1,omitempty"`
	AdditionalAddressLine2 string          `json:"additionalAddressLine2,omitempty"`
	City                   string          `json:"city"` // required
	Company                string          `json:"company,omitempty"`
	Country                *Country        `json:"country,omitempty"`
	CountryId              string          `json:"countryId"` // required
	CountryState           *CountryState   `json:"countryState,omitempty"`
	CountryStateId         string          `json:"countryStateId,omitempty"`
	CreatedAt              time.Time       `json:"createdAt,omitempty"`
	CustomFields           interface{}     `json:"customFields,omitempty"`
	Department             string          `json:"department,omitempty"`
	FirstName              string          `json:"firstName"` // required
	Id                     string          `json:"id,omitempty"`
	LastName               string          `json:"lastName"` // required
	Order                  *Order          `json:"order,omitempty"`
	OrderDeliveries        []OrderDelivery `json:"orderDeliveries,omitempty"`
	OrderId                string          `json:"orderId"` // required
	OrderVersionId         string          `json:"orderVersionId,omitempty"`
	PhoneNumber            string          `json:"phoneNumber,omitempty"`
	Salutation             *Salutation     `json:"salutation,omitempty"`
	SalutationId           string          `json:"salutationId"` // required
	Street                 string          `json:"street"`       // required
	Title                  string          `json:"title,omitempty"`
	UpdatedAt              time.Time       `json:"updatedAt,omitempty"`
	VatId                  string          `json:"vatId,omitempty"`
	VersionId              string          `json:"versionId,omitempty"`
	Zipcode                string          `json:"zipcode,omitempty"`
}
