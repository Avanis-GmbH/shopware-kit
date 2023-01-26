package model

type ShippingMethodTag struct {
	ShippingMethod   *ShippingMethod `json:"shippingMethod,omitempty"`
	ShippingMethodId string          `json:"shippingMethodId,omitempty"`
	Tag              *Tag            `json:"tag,omitempty"`
	TagId            string          `json:"tagId,omitempty"`
}

type ShippingMethodTagCollection struct {
	EntityCollection

	Data []ShippingMethodTag `json:"data"`
}
