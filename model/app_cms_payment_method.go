package model

import "time"

type AppPaymentMethod struct {
	App             *App           `json:"app,omitempty"`
	AppId           string         `json:"appId,omitempty"`
	AppName         string         `json:"appName,omitempty"` // required
	CreatedAt       time.Time      `json:"createdAt,omitempty"`
	FinalizeUrl     string         `json:"finalizeUrl,omitempty"`
	Id              string         `json:"id,omitempty"`
	Identifier      string         `json:"identifier,omitempty"` // required
	OriginalMedia   *Media         `json:"originalMedia,omitempty"`
	OriginalMediaId string         `json:"originalMediaId,omitempty"`
	PaymentMethod   *PaymentMethod `json:"paymentMethod,omitempty"`
	PaymentMethodId string         `json:"paymentMethodId,omitempty"` // required
	PayUrl          string         `json:"payUrl,omitempty"`
	UpdatedAt       time.Time      `json:"updatedAt,omitempty"`
}
