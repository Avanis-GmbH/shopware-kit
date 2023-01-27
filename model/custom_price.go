package model

import "time"

type CustomPrice struct {
	CreatedAt        time.Time      `json:"createdAt,omitempty"`
	Customer         *Customer      `json:"customer,omitempty"`
	CustomerGroup    *CustomerGroup `json:"customerGroup,omitempty"`
	CustomerGroupId  string         `json:"customerGroupId,omitempty"`
	CustomerId       string         `json:"customerId,omitempty"`
	Id               string         `json:"id,omitempty"`
	Price            interface{}    `json:"price"` // required
	Product          *Product       `json:"product,omitempty"`
	ProductId        string         `json:"productId"` // required
	ProductVersionId string         `json:"productVersionId,omitempty"`
	UpdatedAt        time.Time      `json:"updatedAt,omitempty"`
}
