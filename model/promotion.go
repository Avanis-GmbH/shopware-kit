package model

import "time"

type Promotion struct {
	Active                    bool                      `json:"active,omitempty"`
	CartRules                 []Rule                    `json:"cartRules,omitempty"`
	Code                      string                    `json:"code,omitempty"`
	CreatedAt                 time.Time                 `json:"createdAt,omitempty"`
	CustomerRestriction       bool                      `json:"customerRestriction,omitempty"`
	CustomFields              interface{}               `json:"customFields,omitempty"`
	Discounts                 []PromotionDiscount       `json:"discounts,omitempty"`
	ExclusionIds              interface{}               `json:"exclusionIds,omitempty"`
	Exclusive                 bool                      `json:"exclusive,omitempty"`
	Id                        string                    `json:"id,omitempty"`
	IndividualCodePattern     string                    `json:"individualCodePattern,omitempty"`
	IndividualCodes           []PromotionIndividualCode `json:"individualCodes,omitempty"`
	MaxRedemptionsGlobal      float64                   `json:"maxRedemptionsGlobal,omitempty"`
	MaxRedemptionsPerCustomer float64                   `json:"maxRedemptionsPerCustomer,omitempty"`
	Name                      string                    `json:"name,omitempty"`
	OrderCount                float64                   `json:"orderCount,omitempty"`
	OrderRules                []Rule                    `json:"orderRules,omitempty"`
	OrdersPerCustomerCount    interface{}               `json:"ordersPerCustomerCount,omitempty"`
	PersonaCustomers          []Customer                `json:"personaCustomers,omitempty"`
	PersonaRules              []Rule                    `json:"personaRules,omitempty"`
	PreventCombination        bool                      `json:"preventCombination,omitempty"`
	Priority                  float64                   `json:"priority,omitempty"`
	SalesChannels             []PromotionSalesChannel   `json:"salesChannels,omitempty"`
	Setgroups                 []PromotionSetgroup       `json:"setgroups,omitempty"`
	Translated                interface{}               `json:"translated,omitempty"`
	Translations              []PromotionTranslation    `json:"translations,omitempty"`
	UpdatedAt                 time.Time                 `json:"updatedAt,omitempty"`
	UseCodes                  bool                      `json:"useCodes,omitempty"`
	UseIndividualCodes        bool                      `json:"useIndividualCodes,omitempty"`
	UseSetGroups              bool                      `json:"useSetGroups,omitempty"`
	ValidFrom                 time.Time                 `json:"validFrom,omitempty"`
	ValidUntil                time.Time                 `json:"validUntil,omitempty"`
}

type PromotionCollection struct {
	EntityCollection

	Data []Promotion `json:"data"`
}
