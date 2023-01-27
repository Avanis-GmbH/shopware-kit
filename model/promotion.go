package model

import "time"

type Promotion struct {
	Active                    bool                      `json:"active"` // required
	CartRules                 []Rule                    `json:"cartRules,omitempty"`
	Code                      string                    `json:"code,omitempty"`
	CreatedAt                 time.Time                 `json:"createdAt,omitempty"`
	CustomerRestriction       bool                      `json:"customerRestriction,omitempty"`
	CustomFields              interface{}               `json:"customFields,omitempty"`
	Discounts                 []PromotionDiscount       `json:"discounts,omitempty"`
	ExclusionIds              interface{}               `json:"exclusionIds,omitempty"`
	Exclusive                 bool                      `json:"exclusive"` // required
	Id                        string                    `json:"id,omitempty"`
	IndividualCodePattern     string                    `json:"individualCodePattern,omitempty"`
	IndividualCodes           []PromotionIndividualCode `json:"individualCodes,omitempty"`
	MaxRedemptionsGlobal      float64                   `json:"maxRedemptionsGlobal"`
	MaxRedemptionsPerCustomer float64                   `json:"maxRedemptionsPerCustomer"`
	Name                      string                    `json:"name"` // required
	OrderCount                float64                   `json:"orderCount"`
	OrderRules                []Rule                    `json:"orderRules,omitempty"`
	OrdersPerCustomerCount    interface{}               `json:"ordersPerCustomerCount,omitempty"`
	PersonaCustomers          []Customer                `json:"personaCustomers,omitempty"`
	PersonaRules              []Rule                    `json:"personaRules,omitempty"`
	PreventCombination        bool                      `json:"preventCombination"` // required
	Priority                  int64                     `json:"priority"`           // required
	SalesChannels             []PromotionSalesChannel   `json:"salesChannels,omitempty"`
	Setgroups                 []PromotionSetgroup       `json:"setgroups,omitempty"`
	Translated                interface{}               `json:"translated,omitempty"`
	Translations              []PromotionTranslation    `json:"translations,omitempty"`
	UpdatedAt                 time.Time                 `json:"updatedAt,omitempty"`
	UseCodes                  bool                      `json:"useCodes"`           // required
	UseIndividualCodes        bool                      `json:"useIndividualCodes"` // required
	UseSetGroups              bool                      `json:"useSetGroups"`       // required
	ValidFrom                 time.Time                 `json:"validFrom,omitempty"`
	ValidUntil                time.Time                 `json:"validUntil,omitempty"`
}
