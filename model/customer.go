package model

import "time"

type Customer struct {
	Active                    bool               `json:"active,omitempty"`
	Addresses                 []CustomerAddress  `json:"addresses,omitempty"`
	AffiliateCode             string             `json:"affiliateCode,omitempty"`
	AutoIncrement             float64            `json:"autoIncrement,omitempty"`
	Birthday                  time.Time          `json:"birthday,omitempty"`
	BoundSalesChannel         *SalesChannel      `json:"boundSalesChannel,omitempty"`
	BoundSalesChannelId       string             `json:"boundSalesChannelId,omitempty"`
	CampaignCode              string             `json:"campaignCode,omitempty"`
	Company                   string             `json:"company,omitempty"`
	CreatedAt                 time.Time          `json:"createdAt,omitempty"`
	CustomerNumber            string             `json:"customerNumber,omitempty"`
	CustomFields              interface{}        `json:"customFields,omitempty"`
	DefaultBillingAddress     *CustomerAddress   `json:"defaultBillingAddress,omitempty"`
	DefaultBillingAddressId   string             `json:"defaultBillingAddressId,omitempty"`
	DefaultPaymentMethod      *PaymentMethod     `json:"defaultPaymentMethod,omitempty"`
	DefaultPaymentMethodId    string             `json:"defaultPaymentMethodId,omitempty"`
	DefaultShippingAddress    *CustomerAddress   `json:"defaultShippingAddress,omitempty"`
	DefaultShippingAddressId  string             `json:"defaultShippingAddressId,omitempty"`
	DoubleOptInConfirmDate    time.Time          `json:"doubleOptInConfirmDate,omitempty"`
	DoubleOptInEmailSentDate  time.Time          `json:"doubleOptInEmailSentDate,omitempty"`
	DoubleOptInRegistration   bool               `json:"doubleOptInRegistration,omitempty"`
	Email                     string             `json:"email,omitempty"`
	FirstLogin                time.Time          `json:"firstLogin,omitempty"`
	FirstName                 string             `json:"firstName,omitempty"`
	Group                     *CustomerGroup     `json:"group,omitempty"`
	GroupId                   string             `json:"groupId,omitempty"`
	Guest                     bool               `json:"guest,omitempty"`
	Hash                      string             `json:"hash,omitempty"`
	Id                        string             `json:"id,omitempty"`
	Language                  *Language          `json:"language,omitempty"`
	LanguageId                string             `json:"languageId,omitempty"`
	LastLogin                 time.Time          `json:"lastLogin,omitempty"`
	LastName                  string             `json:"lastName,omitempty"`
	LastOrderDate             time.Time          `json:"lastOrderDate,omitempty"`
	LastPaymentMethod         *PaymentMethod     `json:"lastPaymentMethod,omitempty"`
	LastPaymentMethodId       string             `json:"lastPaymentMethodId,omitempty"`
	LegacyEncoder             string             `json:"legacyEncoder,omitempty"`
	LegacyPassword            string             `json:"legacyPassword,omitempty"`
	Newsletter                bool               `json:"newsletter,omitempty"`
	NewsletterSalesChannelIds interface{}        `json:"newsletterSalesChannelIds,omitempty"`
	OrderCount                float64            `json:"orderCount,omitempty"`
	OrderCustomers            []OrderCustomer    `json:"orderCustomers,omitempty"`
	OrderTotalAmount          float64            `json:"orderTotalAmount,omitempty"`
	Password                  interface{}        `json:"password,omitempty"`
	ProductReviews            []ProductReview    `json:"productReviews,omitempty"`
	Promotions                []Promotion        `json:"promotions,omitempty"`
	RecoveryCustomer          *CustomerRecovery  `json:"recoveryCustomer,omitempty"`
	RemoteAddress             interface{}        `json:"remoteAddress,omitempty"`
	RequestedGroup            *CustomerGroup     `json:"requestedGroup,omitempty"`
	RequestedGroupId          string             `json:"requestedGroupId,omitempty"`
	SalesChannel              *SalesChannel      `json:"salesChannel,omitempty"`
	SalesChannelId            string             `json:"salesChannelId,omitempty"`
	Salutation                *Salutation        `json:"salutation,omitempty"`
	SalutationId              string             `json:"salutationId,omitempty"`
	TagIds                    interface{}        `json:"tagIds,omitempty"`
	Tags                      []Tag              `json:"tags,omitempty"`
	Title                     string             `json:"title,omitempty"`
	UpdatedAt                 time.Time          `json:"updatedAt,omitempty"`
	VatIds                    interface{}        `json:"vatIds,omitempty"`
	Wishlists                 []CustomerWishlist `json:"wishlists,omitempty"`
}

type CustomerCollection struct {
	EntityCollection

	Data []Customer `json:"data"`
}
