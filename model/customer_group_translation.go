package model

import "time"

type CustomerGroupTranslation struct {
	CreatedAt                           *time.Time     `json:"createdAt,omitempty"`
	CustomerGroup                       *CustomerGroup `json:"customerGroup,omitempty"`
	CustomerGroupId                     string         `json:"customerGroupId,omitempty"`
	CustomFields                        interface{}    `json:"customFields,omitempty"`
	Language                            *Language      `json:"language,omitempty"`
	LanguageId                          string         `json:"languageId,omitempty"`
	Name                                string         `json:"name,omitempty"`
	RegistrationIntroduction            string         `json:"registrationIntroduction,omitempty"`
	RegistrationOnlyCompanyRegistration bool           `json:"registrationOnlyCompanyRegistration"`
	RegistrationSeoMetaDescription      string         `json:"registrationSeoMetaDescription,omitempty"`
	RegistrationTitle                   string         `json:"registrationTitle,omitempty"`
	UpdatedAt                           *time.Time     `json:"updatedAt,omitempty"`
}
