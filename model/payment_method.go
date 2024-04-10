package model

import "time"

type PaymentMethod struct {
	Active                         bool                       `json:"active"`
	AfterOrderEnabled              bool                       `json:"afterOrderEnabled"`
	AppPaymentMethod               *AppPaymentMethod          `json:"appPaymentMethod,omitempty"`
	Asynchronous                   bool                       `json:"asynchronous"`
	AvailabilityRule               *Rule                      `json:"availabilityRule,omitempty"`
	AvailabilityRuleId             string                     `json:"availabilityRuleId,omitempty"`
	CreatedAt                      *time.Time                 `json:"createdAt,omitempty"`
	Customers                      []Customer                 `json:"customers,omitempty"`
	CustomFields                   interface{}                `json:"customFields,omitempty"`
	Description                    string                     `json:"description,omitempty"`
	DistinguishableName            string                     `json:"distinguishableName,omitempty"`
	FormattedHandlerIdentifier     string                     `json:"formattedHandlerIdentifier,omitempty"`
	HandlerIdentifier              string                     `json:"handlerIdentifier,omitempty"`
	Id                             string                     `json:"id,omitempty"`
	Media                          *Media                     `json:"media,omitempty"`
	MediaId                        string                     `json:"mediaId,omitempty"`
	Name                           string                     `json:"name,omitempty"` // required
	OrderTransactions              []OrderTransaction         `json:"orderTransactions,omitempty"`
	Plugin                         *Plugin                    `json:"plugin,omitempty"`
	PluginId                       string                     `json:"pluginId,omitempty"`
	Position                       float64                    `json:"position,omitempty"`
	Prepared                       bool                       `json:"prepared"`
	SalesChannelDefaultAssignments []SalesChannel             `json:"salesChannelDefaultAssignments,omitempty"`
	SalesChannels                  []SalesChannel             `json:"salesChannels,omitempty"`
	Synchronous                    bool                       `json:"synchronous"`
	Translated                     interface{}                `json:"translated,omitempty"`
	Translations                   []PaymentMethodTranslation `json:"translations,omitempty"`
	UpdatedAt                      *time.Time                 `json:"updatedAt,omitempty"`
}
