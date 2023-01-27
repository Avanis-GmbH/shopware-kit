package model

import "time"

type ProductPrice struct {
	CreatedAt        time.Time   `json:"createdAt,omitempty"`
	CustomFields     interface{} `json:"customFields,omitempty"`
	Id               string      `json:"id,omitempty"`
	Price            interface{} `json:"price,omitempty"` // required
	Product          *Product    `json:"product,omitempty"`
	ProductId        string      `json:"productId,omitempty"` // required
	ProductVersionId string      `json:"productVersionId,omitempty"`
	QuantityEnd      int64       `json:"quantityEnd,omitempty"`
	QuantityStart    int64       `json:"quantityStart,omitempty"` // required
	Rule             *Rule       `json:"rule,omitempty"`
	RuleId           string      `json:"ruleId,omitempty"` // required
	UpdatedAt        time.Time   `json:"updatedAt,omitempty"`
	VersionId        string      `json:"versionId,omitempty"`
}

type Price struct {
	RuleID     string  `json:"ruleId,omitempty"`
	CurrencyID string  `json:"currencyId,omitempty"`
	Gross      float32 `json:"gross"` // required
	Net        float32 `json:"net"`   // required
	Linked     bool    `json:"linked,omitempty"`
}

type OrderPrice struct {
	NetPrice        float32         `json:"netPrice,omitempty"`   // required
	TotalPrice      float32         `json:"totalPrice,omitempty"` // required
	UnitPrice       float32         `json:"unitPrice,omitempty"`  // required
	Quantity        float32         `json:"quantity,omitempty"`   // required
	CalculatedTaxes interface{}     `json:"calculatedTaxes,omitempty"`
	TaxRules        interface{}     `json:"taxRules,omitempty"`
	ReferencePrice  interface{}     `json:"referencePrice,omitempty"`
	PositionPrice   float32         `json:"positionPrice,omitempty"` // required
	RawTotal        float32         `json:"rawTotal,omitempty"`      // required
	TaxStatus       string          `json:"taxStatus,omitempty"`     // required
	ListPrice       ListPrice       `json:"listPrice,omitempty"`
	RegulationPrice RegulationPrice `json:"regulationPrice,omitempty"`
}

type ListPrice struct {
	Price      float64 `json:"price,omitempty"`      // required
	Discount   float64 `json:"discount,omitempty"`   // required
	Percentage float64 `json:"percentage,omitempty"` // required
}

type RegulationPrice struct {
	Price float64 `json:"price,omitempty"` // required
}
