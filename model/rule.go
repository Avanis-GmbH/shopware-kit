package model

import "time"

type Rule struct {
	CartPromotions                  []Promotion           `json:"cartPromotions,omitempty"`
	Conditions                      []RuleCondition       `json:"conditions,omitempty"`
	CreatedAt                       *time.Time            `json:"createdAt,omitempty"`
	CustomFields                    interface{}           `json:"customFields,omitempty"`
	Description                     string                `json:"description,omitempty"`
	EventActions                    []EventAction         `json:"eventActions,omitempty"`
	FlowSequences                   []FlowSequence        `json:"flowSequences,omitempty"`
	Id                              string                `json:"id,omitempty"`
	Invalid                         bool                  `json:"invalid"`
	ModuleTypes                     interface{}           `json:"moduleTypes,omitempty"`
	Name                            string                `json:"name,omitempty"` // required
	OrderPromotions                 []Promotion           `json:"orderPromotions,omitempty"`
	Payload                         interface{}           `json:"payload,omitempty"`
	PaymentMethods                  []PaymentMethod       `json:"paymentMethods,omitempty"`
	PersonaPromotions               []Promotion           `json:"personaPromotions,omitempty"`
	Priority                        int64                 `json:"priority,omitempty"` // required
	ProductPrices                   []ProductPrice        `json:"productPrices,omitempty"`
	PromotionDiscounts              []PromotionDiscount   `json:"promotionDiscounts,omitempty"`
	PromotionSetGroups              []PromotionSetgroup   `json:"promotionSetGroups,omitempty"`
	ShippingMethodPriceCalculations []ShippingMethodPrice `json:"shippingMethodPriceCalculations,omitempty"`
	ShippingMethodPrices            []ShippingMethodPrice `json:"shippingMethodPrices,omitempty"`
	ShippingMethods                 []ShippingMethod      `json:"shippingMethods,omitempty"`
	Tags                            RuleTag               `json:"tags,omitempty"`
	UpdatedAt                       *time.Time            `json:"updatedAt,omitempty"`
}
