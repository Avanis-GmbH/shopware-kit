package model

type ShippingMethodTag struct {
	ShippingMethod   *ShippingMethod `json:"shippingMethod,omitempty"`
	ShippingMethodId string          `json:"shippingMethodId"` // required
	Tag              *Tag            `json:"tag,omitempty"`
	TagId            string          `json:"tagId"` // required
}
