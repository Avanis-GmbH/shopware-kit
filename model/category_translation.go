package model

import "time"

type CategoryTranslation struct {
	Breadcrumb        interface{} `json:"breadcrumb,omitempty"`
	Category          *Category   `json:"category,omitempty"`
	CategoryId        string      `json:"categoryId,omitempty"`
	CategoryVersionId string      `json:"categoryVersionId,omitempty"`
	CreatedAt         time.Time   `json:"createdAt,omitempty"`
	CustomFields      interface{} `json:"customFields,omitempty"`
	Description       string      `json:"description,omitempty"`
	ExternalLink      string      `json:"externalLink,omitempty"`
	InternalLink      string      `json:"internalLink,omitempty"`
	Keywords          string      `json:"keywords,omitempty"`
	Language          *Language   `json:"language,omitempty"`
	LanguageId        string      `json:"languageId,omitempty"`
	LinkNewTab        bool        `json:"linkNewTab,omitempty"`
	LinkType          string      `json:"linkType,omitempty"`
	MetaDescription   string      `json:"metaDescription,omitempty"`
	MetaTitle         string      `json:"metaTitle,omitempty"`
	Name              string      `json:"name,omitempty"`
	SlotConfig        interface{} `json:"slotConfig,omitempty"`
	UpdatedAt         time.Time   `json:"updatedAt,omitempty"`
}
