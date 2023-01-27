package model

type SalesChannelPaymentMethod struct {
	PaymentMethod   *PaymentMethod `json:"paymentMethod,omitempty"`
	PaymentMethodId string         `json:"paymentMethodId,omitempty"` // required
	SalesChannel    *SalesChannel  `json:"salesChannel,omitempty"`
	SalesChannelId  string         `json:"salesChannelId,omitempty"` // required
}
