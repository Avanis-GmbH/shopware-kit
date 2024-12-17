package model

import "time"

type ShippingMethod struct {
	Active                         *bool                       `json:"active,omitempty"`
	AvailabilityRule               *Rule                       `json:"availabilityRule,omitempty"`
	AvailabilityRuleId             string                      `json:"availabilityRuleId,omitempty"` // required
	CreatedAt                      *time.Time                  `json:"createdAt,omitempty"`
	CustomFields                   interface{}                 `json:"customFields,omitempty"`
	DeliveryTime                   *DeliveryTime               `json:"deliveryTime,omitempty"`
	DeliveryTimeId                 string                      `json:"deliveryTimeId,omitempty"` // required
	Description                    string                      `json:"description,omitempty"`
	Id                             string                      `json:"id,omitempty"`
	Media                          *Media                      `json:"media,omitempty"`
	MediaId                        string                      `json:"mediaId,omitempty"`
	Name                           string                      `json:"name,omitempty"` // required
	OrderDeliveries                []OrderDelivery             `json:"orderDeliveries,omitempty"`
	Prices                         []ShippingMethodPrice       `json:"prices,omitempty"`
	SalesChannelDefaultAssignments []SalesChannel              `json:"salesChannelDefaultAssignments,omitempty"`
	SalesChannels                  []SalesChannel              `json:"salesChannels,omitempty"`
	Tags                           []Tag                       `json:"tags,omitempty"`
	Tax                            *Tax                        `json:"tax,omitempty"`
	TaxId                          string                      `json:"taxId,omitempty"`
	TaxType                        string                      `json:"taxType,omitempty"` // required
	TrackingUrl                    string                      `json:"trackingUrl,omitempty"`
	Translated                     interface{}                 `json:"translated,omitempty"`
	Translations                   []ShippingMethodTranslation `json:"translations,omitempty"`
	UpdatedAt                      *time.Time                  `json:"updatedAt,omitempty"`
}
