package model

import "time"

type MailTemplate struct {
	ContentHtml        string                    `json:"contentHtml,omitempty"`  // required
	ContentPlain       string                    `json:"contentPlain,omitempty"` // required
	CreatedAt          *time.Time                `json:"createdAt,omitempty"`
	CustomFields       interface{}               `json:"customFields,omitempty"`
	Description        string                    `json:"description,omitempty"`
	Id                 string                    `json:"id,omitempty"`
	MailTemplateType   *MailTemplateType         `json:"mailTemplateType,omitempty"`
	MailTemplateTypeId string                    `json:"mailTemplateTypeId,omitempty"` // required
	Media              []MailTemplateMedia       `json:"media,omitempty"`
	SenderName         string                    `json:"senderName,omitempty"`
	Subject            string                    `json:"subject,omitempty"` // required
	SystemDefault      bool                      `json:"systemDefault"`
	Translated         interface{}               `json:"translated,omitempty"`
	Translations       []MailTemplateTranslation `json:"translations,omitempty"`
	UpdatedAt          *time.Time                `json:"updatedAt,omitempty"`
}
