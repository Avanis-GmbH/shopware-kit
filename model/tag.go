package model

import "time"

type Tag struct {
	Categories           []Category            `json:"categories,omitempty"`
	CreatedAt            *time.Time            `json:"createdAt,omitempty"`
	Customers            []Customer            `json:"customers,omitempty"`
	Id                   string                `json:"id,omitempty"`
	LandingPages         []LandingPage         `json:"landingPages,omitempty"`
	Media                []Media               `json:"media,omitempty"`
	Name                 string                `json:"name,omitempty"` // required
	NewsletterRecipients []NewsletterRecipient `json:"newsletterRecipients,omitempty"`
	Orders               []Order               `json:"orders,omitempty"`
	Products             []Product             `json:"products,omitempty"`
	ShippingMethods      []ShippingMethod      `json:"shippingMethods,omitempty"`
	UpdatedAt            *time.Time            `json:"updatedAt,omitempty"`
}
