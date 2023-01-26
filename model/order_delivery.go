package model

import "time"

type OrderDelivery struct {
	CreatedAt                     time.Time               `json:"createdAt,omitempty"`
	CustomFields                  interface{}             `json:"customFields,omitempty"`
	Id                            string                  `json:"id,omitempty"`
	Order                         *Order                  `json:"order,omitempty"`
	OrderId                       string                  `json:"orderId,omitempty"`
	OrderVersionId                string                  `json:"orderVersionId,omitempty"`
	Positions                     []OrderDeliveryPosition `json:"positions,omitempty"`
	ShippingCosts                 interface{}             `json:"shippingCosts,omitempty"`
	ShippingDateEarliest          time.Time               `json:"shippingDateEarliest,omitempty"`
	ShippingDateLatest            time.Time               `json:"shippingDateLatest,omitempty"`
	ShippingMethod                *ShippingMethod         `json:"shippingMethod,omitempty"`
	ShippingMethodId              string                  `json:"shippingMethodId,omitempty"`
	ShippingOrderAddress          *OrderAddress           `json:"shippingOrderAddress,omitempty"`
	ShippingOrderAddressId        string                  `json:"shippingOrderAddressId,omitempty"`
	ShippingOrderAddressVersionId string                  `json:"shippingOrderAddressVersionId,omitempty"`
	StateId                       string                  `json:"stateId,omitempty"`
	StateMachineState             *StateMachineState      `json:"stateMachineState,omitempty"`
	TrackingCodes                 interface{}             `json:"trackingCodes,omitempty"`
	UpdatedAt                     time.Time               `json:"updatedAt,omitempty"`
	VersionId                     string                  `json:"versionId,omitempty"`
}

type OrderDeliveryCollection struct {
	EntityCollection

	Data []OrderDelivery `json:"data"`
}
