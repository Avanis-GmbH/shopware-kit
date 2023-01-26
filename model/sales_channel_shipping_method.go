package model

type SalesChannelShippingMethod struct {
	SalesChannel     *SalesChannel   `json:"salesChannel,omitempty"`
	SalesChannelId   string          `json:"salesChannelId,omitempty"`
	ShippingMethod   *ShippingMethod `json:"shippingMethod,omitempty"`
	ShippingMethodId string          `json:"shippingMethodId,omitempty"`
}

type SalesChannelShippingMethodCollection struct {
	EntityCollection

	Data []SalesChannelShippingMethod `json:"data"`
}
