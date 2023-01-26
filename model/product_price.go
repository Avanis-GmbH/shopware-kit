package model

import "time"

type ProductPrice struct {
	VersionId        string      `json:"versionId,omitempty"`
	RuleId           string      `json:"ruleId,omitempty"`
	Product          *Product    `json:"product,omitempty"`
	CustomFields     interface{} `json:"customFields,omitempty"`
	UpdatedAt        time.Time   `json:"updatedAt,omitempty"`
	CreatedAt        time.Time   `json:"createdAt,omitempty"`
	Id               string      `json:"id,omitempty"`
	ProductId        string      `json:"productId,omitempty"`
	ProductVersionId string      `json:"productVersionId,omitempty"`
	Price            interface{} `json:"price,omitempty"`
	QuantityStart    float64     `json:"quantityStart,omitempty"`
	QuantityEnd      float64     `json:"quantityEnd,omitempty"`
	Rule             *Rule       `json:"rule,omitempty"`
}
