package model

import "time"

type MailTemplateTypeTranslation struct {
	CreatedAt          *time.Time        `json:"createdAt,omitempty"`
	CustomFields       interface{}       `json:"customFields,omitempty"`
	Language           *Language         `json:"language,omitempty"`
	LanguageId         string            `json:"languageId,omitempty"`
	MailTemplateType   *MailTemplateType `json:"mailTemplateType,omitempty"`
	MailTemplateTypeId string            `json:"mailTemplateTypeId,omitempty"`
	Name               string            `json:"name,omitempty"`
	UpdatedAt          *time.Time        `json:"updatedAt,omitempty"`
}
