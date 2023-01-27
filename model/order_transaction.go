package model

import "time"

type OrderTransaction struct {
	Amount            Amount             `json:"amount,omitempty"`
	Captures          Capture            `json:"captures,omitempty"`
	CreatedAt         time.Time          `json:"createdAt,omitempty"`
	CustomFields      interface{}        `json:"customFields,omitempty"`
	Id                string             `json:"id,omitempty"`
	Order             *Order             `json:"order,omitempty"`
	OrderId           string             `json:"orderId"` // required
	OrderVersionId    string             `json:"orderVersionId,omitempty"`
	PaymentMethod     *PaymentMethod     `json:"paymentMethod,omitempty"`
	PaymentMethodId   string             `json:"paymentMethodId"` // required
	StateId           string             `json:"stateId,omitempty"`
	StateMachineState *StateMachineState `json:"stateMachineState,omitempty"`
	UpdatedAt         time.Time          `json:"updatedAt,omitempty"`
	VersionId         string             `json:"versionId,omitempty"`
}

type Amount struct {
	UnitPrice       float64         `json:"unitPrice"`  // required
	TotalPrice      float64         `json:"totalPrice"` // required
	Quantity        float64         `json:"quantity"`   // required
	CalculatedTaxes interface{}     `json:"calculatedTaxes,omitempty"`
	TaxRules        interface{}     `json:"taxRules,omitempty"`
	ReferencePrice  interface{}     `json:"referencePrice,omitempty"`
	ListPrice       ListPrice       `json:"listPrice,omitempty"`
	RegulationPrice RegulationPrice `json:"regulationPrice,omitempty"`
}

type Capture struct {
	Id                        string             `json:"id,omitempty"`
	OrderTransactionId        string             `json:"orderTransactionId"` // required
	OrderTransactionVersionId string             `json:"orderTransactionVersionId,omitempty"`
	StateId                   string             `json:"stateId"` // required
	ExternalReference         string             `json:"externalReference,omitempty"`
	TotalAmount               float64            `json:"totalAmount"` // required
	Amount                    Amount             `json:"amount"`      // required
	CustomFields              interface{}        `json:"customFields,omitempty"`
	CreatedAt                 time.Time          `json:"createdAt,omitempty"`
	UpdatedAt                 time.Time          `json:"updatedAt,omitempty"`
	StateMachineState         *StateMachineState `json:"stateMachineState,omitempty"`
	Transaction               interface{}        `json:"transaction,omitempty"`
	Refunds                   Refund             `json:"refunds,omitempty"`
}

type Refund struct {
	Id                 string             `json:"id,omitempty"`
	CaptureId          string             `json:"captureId"` // required
	StateId            string             `json:"stateId"`   // required
	ExternalReference  string             `json:"externalReference,omitempty"`
	Reason             string             `json:"reason,omitempty"`
	TotalAmount        float64            `json:"totalAmount"` // required
	Amount             Amount             `json:"amount"`      // required
	CustomFields       interface{}        `json:"customFields,omitempty"`
	CreatedAt          time.Time          `json:"createdAt,omitempty"`
	UpdatedAt          time.Time          `json:"updatedAt,omitempty"`
	StateMachineState  *StateMachineState `json:"stateMachineState,omitempty"`
	TransactionCapture interface{}        `json:"transactionCapture,omitempty"`
	Positions          Position           `json:"positions,omitempty"`
}

type Position struct {
	Id                            string        `json:"id,omitempty"`
	RefundId                      string        `json:"refundId"`        // required
	OrderLineItemId               string        `json:"orderLineItemId"` // required
	ExternalReference             string        `json:"externalReference,omitempty"`
	Reason                        string        `json:"reason,omitempty"`
	Quantity                      float64       `json:"quantity"` // required
	Amount                        Amount        `json:"amount"`   // required
	RefundPrice                   float64       `json:"refundPrice,omitempty"`
	CustomFields                  interface{}   `json:"customFields,omitempty"`
	CreatedAt                     time.Time     `json:"createdAt,omitempty"`
	UpdatedAt                     time.Time     `json:"updatedAt,omitempty"`
	OrderLineItem                 OrderLineItem `json:"orderLineItem,omitempty"`
	OrderTransactionCaptureRefund interface{}   `json:"orderTransactionCaptureRefund,omitempty"`
}
