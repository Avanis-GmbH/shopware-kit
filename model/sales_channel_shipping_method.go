package model

type SalesChannelShippingMethod struct {
	SalesChannel     *SalesChannel   `json:"salesChannel,omitempty"`
	SalesChannelId   string          `json:"salesChannelId"` // required
	ShippingMethod   *ShippingMethod `json:"shippingMethod,omitempty"`
	ShippingMethodId string          `json:"shippingMethodId"` // required
}
