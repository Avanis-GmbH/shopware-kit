package model

import "time"

type AppCmsBlockTranslation struct {
	AppCmsBlock   *AppCmsBlock `json:"appCmsBlock,omitempty"`
	AppCmsBlockId string       `json:"appCmsBlockId,omitempty"`
	CreatedAt     time.Time    `json:"createdAt,omitempty"`
	Label         string       `json:"label,omitempty"`
	Language      *Language    `json:"language,omitempty"`
	LanguageId    string       `json:"languageId,omitempty"`
	UpdatedAt     time.Time    `json:"updatedAt,omitempty"`
}

type AppCmsBlockTranslationCollection struct {
	EntityCollection

	Data []AppCmsBlockTranslation `json:"data"`
}
