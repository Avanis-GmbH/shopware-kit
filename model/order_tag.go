package model

type OrderTag struct {
	Order          *Order `json:"order,omitempty"`
	OrderId        string `json:"orderId"` // required
	OrderVersionId string `json:"orderVersionId,omitempty"`
	Tag            *Tag   `json:"tag,omitempty"`
	TagId          string `json:"tagId"` // required
}
