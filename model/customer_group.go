package model

import "time"

type CustomerGroup struct {
	CreatedAt                           time.Time                  `json:"createdAt,omitempty"`
	Customers                           []Customer                 `json:"customers,omitempty"`
	CustomFields                        interface{}                `json:"customFields,omitempty"`
	DisplayGross                        bool                       `json:"displayGross,omitempty"`
	Id                                  string                     `json:"id,omitempty"`
	Name                                string                     `json:"name,omitempty"` // required
	RegistrationActive                  bool                       `json:"registrationActive,omitempty"`
	RegistrationIntroduction            string                     `json:"registrationIntroduction,omitempty"`
	RegistrationOnlyCompanyRegistration bool                       `json:"registrationOnlyCompanyRegistration,omitempty"`
	RegistrationSalesChannels           []SalesChannel             `json:"registrationSalesChannels,omitempty"`
	RegistrationSeoMetaDescription      string                     `json:"registrationSeoMetaDescription,omitempty"`
	RegistrationTitle                   string                     `json:"registrationTitle,omitempty"`
	SalesChannels                       []SalesChannel             `json:"salesChannels,omitempty"`
	Translated                          interface{}                `json:"translated,omitempty"`
	Translations                        []CustomerGroupTranslation `json:"translations,omitempty"`
	UpdatedAt                           time.Time                  `json:"updatedAt,omitempty"`
}
