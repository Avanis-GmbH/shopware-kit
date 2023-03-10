package model

import "time"

type NewsletterRecipient struct {
	City           string        `json:"city,omitempty"`
	ConfirmedAt    *time.Time    `json:"confirmedAt,omitempty"`
	CreatedAt      *time.Time    `json:"createdAt,omitempty"`
	CustomFields   interface{}   `json:"customFields,omitempty"`
	Email          string        `json:"email,omitempty"` // required
	FirstName      string        `json:"firstName,omitempty"`
	Hash           string        `json:"hash,omitempty"` // required
	Id             string        `json:"id,omitempty"`
	Language       *Language     `json:"language,omitempty"`
	LanguageId     string        `json:"languageId,omitempty"` // required
	LastName       string        `json:"lastName,omitempty"`
	SalesChannel   *SalesChannel `json:"salesChannel,omitempty"`
	SalesChannelId string        `json:"salesChannelId,omitempty"` // required
	Salutation     *Salutation   `json:"salutation,omitempty"`
	SalutationId   string        `json:"salutationId,omitempty"`
	Status         string        `json:"status,omitempty"` // required
	Street         string        `json:"street,omitempty"`
	Tags           []Tag         `json:"tags,omitempty"`
	Title          string        `json:"title,omitempty"`
	UpdatedAt      *time.Time    `json:"updatedAt,omitempty"`
	ZipCode        string        `json:"zipCode,omitempty"`
}
