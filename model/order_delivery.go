package model

import "time"

type OrderDelivery struct {
	CreatedAt                     time.Time               `json:"createdAt,omitempty"`
	CustomFields                  interface{}             `json:"customFields,omitempty"`
	Id                            string                  `json:"id,omitempty"`
	Order                         *Order                  `json:"order,omitempty"`
	OrderId                       string                  `json:"orderId,omitempty"` // required
	OrderVersionId                string                  `json:"orderVersionId,omitempty"`
	Positions                     []OrderDeliveryPosition `json:"positions,omitempty"`
	ShippingCosts                 ShippingCosts           `json:"shippingCosts,omitempty"`
	ShippingDateEarliest          time.Time               `json:"shippingDateEarliest,omitempty"` // required
	ShippingDateLatest            time.Time               `json:"shippingDateLatest,omitempty"`   // required
	ShippingMethod                *ShippingMethod         `json:"shippingMethod,omitempty"`
	ShippingMethodId              string                  `json:"shippingMethodId,omitempty"` // required
	ShippingOrderAddress          *OrderAddress           `json:"shippingOrderAddress,omitempty"`
	ShippingOrderAddressId        string                  `json:"shippingOrderAddressId,omitempty"` // required
	ShippingOrderAddressVersionId string                  `json:"shippingOrderAddressVersionId,omitempty"`
	StateId                       string                  `json:"stateId,omitempty"` // required
	StateMachineState             *StateMachineState      `json:"stateMachineState,omitempty"`
	TrackingCodes                 []string                `json:"trackingCodes,omitempty"` // required
	UpdatedAt                     time.Time               `json:"updatedAt,omitempty"`
	VersionId                     string                  `json:"versionId,omitempty"`
}

type ShippingCosts struct {
	UnitPrice       float64         `json:"unitPrice,omitempty"`  // required
	TotalPrice      float64         `json:"totalPrice,omitempty"` // required
	Quantity        float64         `json:"quantity,omitempty"`   // required
	CalculatedTaxes interface{}     `json:"calculatedTaxes,omitempty"`
	TaxRules        interface{}     `json:"taxRules,omitempty"`
	ReferencePrice  float64         `json:"referencePrice,omitempty"` // required
	ListPrice       ListPrice       `json:"listPrice,omitempty"`
	RegulationPrice RegulationPrice `json:"regulationPrice,omitempty"`
}
