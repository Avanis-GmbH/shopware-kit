package model

type PromotionPersonaCustomer struct {
	Customer    *Customer  `json:"customer,omitempty"`
	CustomerId  string     `json:"customerId,omitempty"` // required
	Promotion   *Promotion `json:"promotion,omitempty"`
	PromotionId string     `json:"promotionId,omitempty"` // required
}
