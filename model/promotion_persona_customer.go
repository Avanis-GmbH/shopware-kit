package model

type PromotionPersonaCustomer struct {
	Customer    *Customer  `json:"customer,omitempty"`
	CustomerId  string     `json:"customerId,omitempty"`
	Promotion   *Promotion `json:"promotion,omitempty"`
	PromotionId string     `json:"promotionId,omitempty"`
}

type PromotionPersonaCustomerCollection struct {
	EntityCollection

	Data []PromotionPersonaCustomer `json:"data"`
}
