package model

import "time"

type ProductReview struct {
	Comment          string        `json:"comment,omitempty"`
	Content          string        `json:"content"` // required
	CreatedAt        time.Time     `json:"createdAt,omitempty"`
	Customer         *Customer     `json:"customer,omitempty"`
	CustomerId       string        `json:"customerId,omitempty"`
	CustomFields     interface{}   `json:"customFields,omitempty"`
	ExternalEmail    string        `json:"externalEmail,omitempty"`
	ExternalUser     string        `json:"externalUser,omitempty"`
	Id               string        `json:"id,omitempty"`
	Language         *Language     `json:"language,omitempty"`
	LanguageId       string        `json:"languageId"` // required
	Points           float64       `json:"points"`
	Product          *Product      `json:"product,omitempty"`
	ProductId        string        `json:"productId"` // required
	ProductVersionId string        `json:"productVersionId,omitempty"`
	SalesChannel     *SalesChannel `json:"salesChannel,omitempty"`
	SalesChannelId   string        `json:"salesChannelId"` // required
	Status           bool          `json:"status,omitempty"`
	Title            string        `json:"title"` // required
	UpdatedAt        time.Time     `json:"updatedAt,omitempty"`
}
