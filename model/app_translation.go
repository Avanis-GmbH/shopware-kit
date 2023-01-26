package model

import "time"

type AppTranslation struct {
	App                     *App        `json:"app,omitempty"`
	AppId                   string      `json:"appId,omitempty"`
	CreatedAt               time.Time   `json:"createdAt,omitempty"`
	CustomFields            interface{} `json:"customFields,omitempty"`
	Description             string      `json:"description,omitempty"`
	Label                   string      `json:"label,omitempty"`
	Language                *Language   `json:"language,omitempty"`
	LanguageId              string      `json:"languageId,omitempty"`
	PrivacyPolicyExtensions string      `json:"privacyPolicyExtensions,omitempty"`
	UpdatedAt               time.Time   `json:"updatedAt,omitempty"`
}

type AppTranslationCollection struct {
	EntityCollection

	Data []AppTranslation `json:"data"`
}
