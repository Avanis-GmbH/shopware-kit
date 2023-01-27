package model

type ShippingMethodTag struct {
	ShippingMethod   *ShippingMethod `json:"shippingMethod,omitempty"`
	ShippingMethodId string          `json:"shippingMethodId,omitempty"` // required
	Tag              *Tag            `json:"tag,omitempty"`
	TagId            string          `json:"tagId,omitempty"` // required
}
