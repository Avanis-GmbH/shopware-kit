package model

import "time"

type CustomerAddress struct {
	AdditionalAddressLine1 string        `json:"additionalAddressLine1,omitempty"`
	AdditionalAddressLine2 string        `json:"additionalAddressLine2,omitempty"`
	City                   string        `json:"city,omitempty"` // required
	Company                string        `json:"company,omitempty"`
	Country                *Country      `json:"country,omitempty"`
	CountryId              string        `json:"countryId,omitempty"` // required
	CountryState           *CountryState `json:"countryState,omitempty"`
	CountryStateId         string        `json:"countryStateId,omitempty"`
	CreatedAt              time.Time     `json:"createdAt,omitempty"`
	Customer               *Customer     `json:"customer,omitempty"`
	CustomerId             string        `json:"customerId,omitempty"` // required
	CustomFields           interface{}   `json:"customFields,omitempty"`
	Department             string        `json:"department,omitempty"`
	FirstName              string        `json:"firstName,omitempty"` // required
	Id                     string        `json:"id,omitempty"`
	LastName               string        `json:"lastName,omitempty"` // required
	PhoneNumber            string        `json:"phoneNumber,omitempty"`
	Salutation             *Salutation   `json:"salutation,omitempty"`
	SalutationId           string        `json:"salutationId,omitempty"`
	Street                 string        `json:"street,omitempty"` // required
	Title                  string        `json:"title,omitempty"`
	UpdatedAt              time.Time     `json:"updatedAt,omitempty"`
	Zipcode                string        `json:"zipcode,omitempty"`
}
