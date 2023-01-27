package model

import "time"

type Order struct {
	Addresses               []OrderAddress     `json:"addresses,omitempty"`
	AffiliateCode           string             `json:"affiliateCode,omitempty"`
	AmountNet               float64            `json:"amountNet,omitempty"`
	AmountTotal             float64            `json:"amountTotal,omitempty"`
	AutoIncrement           float64            `json:"autoIncrement,omitempty"`
	BillingAddress          *OrderAddress      `json:"billingAddress,omitempty"`
	BillingAddressId        string             `json:"billingAddressId,omitempty"` // required
	BillingAddressVersionId string             `json:"billingAddressVersionId,omitempty"`
	CampaignCode            string             `json:"campaignCode,omitempty"`
	CreatedAt               time.Time          `json:"createdAt,omitempty"`
	CreatedBy               *User              `json:"createdBy,omitempty"`
	CreatedById             string             `json:"createdById,omitempty"`
	Currency                *Currency          `json:"currency,omitempty"`
	CurrencyFactor          float64            `json:"currencyFactor,omitempty"` // required
	CurrencyId              string             `json:"currencyId,omitempty"`     // required
	CustomerComment         string             `json:"customerComment,omitempty"`
	CustomFields            interface{}        `json:"customFields,omitempty"`
	DeepLinkCode            string             `json:"deepLinkCode,omitempty"`
	Deliveries              []OrderDelivery    `json:"deliveries,omitempty"`
	Documents               []Document         `json:"documents,omitempty"`
	Id                      string             `json:"id,omitempty"`
	ItemRounding            interface{}        `json:"itemRounding,omitempty"`
	Language                *Language          `json:"language,omitempty"`
	LanguageId              string             `json:"languageId,omitempty"` // required
	LineItems               []OrderLineItem    `json:"lineItems,omitempty"`
	OrderCustomer           *OrderCustomer     `json:"orderCustomer,omitempty"`
	OrderDate               time.Time          `json:"orderDate,omitempty"`
	OrderDateTime           time.Time          `json:"orderDateTime,omitempty"` // required
	OrderNumber             string             `json:"orderNumber,omitempty"`
	PositionPrice           float64            `json:"positionPrice,omitempty"`
	Price                   OrderPrice         `json:"price,omitempty"`
	RuleIds                 interface{}        `json:"ruleIds,omitempty"`
	SalesChannel            *SalesChannel      `json:"salesChannel,omitempty"`
	SalesChannelId          string             `json:"salesChannelId,omitempty"` // required
	ShippingCosts           interface{}        `json:"shippingCosts,omitempty"`
	ShippingTotal           float64            `json:"shippingTotal,omitempty"`
	StateId                 string             `json:"stateId,omitempty"`
	StateMachineState       *StateMachineState `json:"stateMachineState,omitempty"`
	Tags                    []Tag              `json:"tags,omitempty"`
	TaxStatus               string             `json:"taxStatus,omitempty"`
	TotalRounding           interface{}        `json:"totalRounding,omitempty"`
	Transactions            []OrderTransaction `json:"transactions,omitempty"`
	UpdatedAt               time.Time          `json:"updatedAt,omitempty"`
	UpdatedBy               *User              `json:"updatedBy,omitempty"`
	UpdatedById             string             `json:"updatedById,omitempty"`
	VersionId               string             `json:"versionId,omitempty"`
}
