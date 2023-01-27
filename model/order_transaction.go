package model

import "time"

type OrderTransaction struct {
	Amount            Amount             `json:"amount,omitempty"`
	Captures          Capture            `json:"captures,omitempty"`
	CreatedAt         time.Time          `json:"createdAt,omitempty"`
	CustomFields      interface{}        `json:"customFields,omitempty"`
	Id                string             `json:"id,omitempty"`
	Order             *Order             `json:"order,omitempty"`
	OrderId           string             `json:"orderId,omitempty"` // required
	OrderVersionId    string             `json:"orderVersionId,omitempty"`
	PaymentMethod     *PaymentMethod     `json:"paymentMethod,omitempty"`
	PaymentMethodId   string             `json:"paymentMethodId,omitempty"` // required
	StateId           string             `json:"stateId,omitempty"`
	StateMachineState *StateMachineState `json:"stateMachineState,omitempty"`
	UpdatedAt         time.Time          `json:"updatedAt,omitempty"`
	VersionId         string             `json:"versionId,omitempty"`
}

type Amount struct {
	UnitPrice       float64         `json:"unitPrice,omitempty"`  // required
	TotalPrice      float64         `json:"totalPrice,omitempty"` // required
	Quantity        float64         `json:"quantity,omitempty"`   // required
	CalculatedTaxes interface{}     `json:"calculatedTaxes,omitempty"`
	TaxRules        interface{}     `json:"taxRules,omitempty"`
	ReferencePrice  interface{}     `json:"referencePrice,omitempty"`
	ListPrice       ListPrice       `json:"listPrice,omitempty"`
	RegulationPrice RegulationPrice `json:"regulationPrice,omitempty"`
}

type Capture struct {
	Id                        string             `json:"id,omitempty"`
	OrderTransactionId        string             `json:"orderTransactionId,omitempty"` // required
	OrderTransactionVersionId string             `json:"orderTransactionVersionId,omitempty"`
	StateId                   string             `json:"stateId,omitempty"` // required
	ExternalReference         string             `json:"externalReference,omitempty"`
	TotalAmount               float64            `json:"totalAmount,omitempty"` // required
	Amount                    Amount             `json:"amount,omitempty"`      // required
	CustomFields              interface{}        `json:"customFields,omitempty"`
	CreatedAt                 time.Time          `json:"createdAt,omitempty"`
	UpdatedAt                 time.Time          `json:"updatedAt,omitempty"`
	StateMachineState         *StateMachineState `json:"stateMachineState,omitempty"`
	Transaction               interface{}        `json:"transaction,omitempty"`
	Refunds                   Refund             `json:"refunds,omitempty"`
}

type Refund struct {
	Id                 string             `json:"id,omitempty"`
	CaptureId          string             `json:"captureId,omitempty"` // required
	StateId            string             `json:"stateId,omitempty"`   // required
	ExternalReference  string             `json:"externalReference,omitempty"`
	Reason             string             `json:"reason,omitempty"`
	TotalAmount        float64            `json:"totalAmount,omitempty"` // required
	Amount             Amount             `json:"amount,omitempty"`      // required
	CustomFields       interface{}        `json:"customFields,omitempty"`
	CreatedAt          time.Time          `json:"createdAt,omitempty"`
	UpdatedAt          time.Time          `json:"updatedAt,omitempty"`
	StateMachineState  *StateMachineState `json:"stateMachineState,omitempty"`
	TransactionCapture interface{}        `json:"transactionCapture,omitempty"`
	Positions          Position           `json:"positions,omitempty"`
}

type Position struct {
	Id                            string        `json:"id,omitempty"`
	RefundId                      string        `json:"refundId,omitempty"`        // required
	OrderLineItemId               string        `json:"orderLineItemId,omitempty"` // required
	ExternalReference             string        `json:"externalReference,omitempty"`
	Reason                        string        `json:"reason,omitempty"`
	Quantity                      float64       `json:"quantity,omitempty"` // required
	Amount                        Amount        `json:"amount,omitempty"`   // required
	RefundPrice                   float64       `json:"refundPrice,omitempty"`
	CustomFields                  interface{}   `json:"customFields,omitempty"`
	CreatedAt                     time.Time     `json:"createdAt,omitempty"`
	UpdatedAt                     time.Time     `json:"updatedAt,omitempty"`
	OrderLineItem                 OrderLineItem `json:"orderLineItem,omitempty"`
	OrderTransactionCaptureRefund interface{}   `json:"orderTransactionCaptureRefund,omitempty"`
}
