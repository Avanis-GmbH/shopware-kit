package model

import "time"

type App struct {
	AclRole                 *AclRole           `json:"aclRole,omitempty"`
	AclRoleId               string             `json:"aclRoleId,omitempty"`
	ActionButtons           []AppActionButton  `json:"actionButtons,omitempty"`
	Active                  bool               `json:"active,omitempty"`
	AppSecret               string             `json:"appSecret,omitempty"`
	Author                  string             `json:"author,omitempty"`
	CmsBlocks               []AppCmsBlock      `json:"cmsBlocks,omitempty"`
	Configurable            bool               `json:"configurable,omitempty"`
	Cookies                 interface{}        `json:"cookies,omitempty"`
	Copyright               string             `json:"copyright,omitempty"`
	CreatedAt               time.Time          `json:"createdAt,omitempty"`
	CustomFields            interface{}        `json:"customFields,omitempty"`
	CustomFieldSets         []CustomFieldSet   `json:"customFieldSets,omitempty"`
	Description             string             `json:"description,omitempty"`
	Icon                    string             `json:"icon,omitempty"`
	IconRaw                 interface{}        `json:"iconRaw,omitempty"`
	Id                      string             `json:"id,omitempty"`
	Integration             *Integration       `json:"integration,omitempty"`
	IntegrationId           string             `json:"integrationId,omitempty"`
	Label                   string             `json:"label,omitempty"`
	License                 string             `json:"license,omitempty"`
	MainModule              interface{}        `json:"mainModule,omitempty"`
	Modules                 interface{}        `json:"modules,omitempty"`
	Name                    string             `json:"name,omitempty"`
	Path                    string             `json:"path,omitempty"`
	PaymentMethods          []AppPaymentMethod `json:"paymentMethods,omitempty"`
	Privacy                 string             `json:"privacy,omitempty"`
	PrivacyPolicyExtensions string             `json:"privacyPolicyExtensions,omitempty"`
	Scripts                 []Script           `json:"scripts,omitempty"`
	Templates               []AppTemplate      `json:"templates,omitempty"`
	Translated              interface{}        `json:"translated,omitempty"`
	Translations            []AppTranslation   `json:"translations,omitempty"`
	UpdatedAt               time.Time          `json:"updatedAt,omitempty"`
	Version                 string             `json:"version,omitempty"`
	Webhooks                []Webhook          `json:"webhooks,omitempty"`
}
