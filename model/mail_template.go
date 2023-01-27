package model

import "time"

type MailTemplate struct {
	ContentHtml        string                    `json:"contentHtml"`  // required
	ContentPlain       string                    `json:"contentPlain"` // required
	CreatedAt          time.Time                 `json:"createdAt,omitempty"`
	CustomFields       interface{}               `json:"customFields,omitempty"`
	Description        string                    `json:"description,omitempty"`
	Id                 string                    `json:"id,omitempty"`
	MailTemplateType   *MailTemplateType         `json:"mailTemplateType,omitempty"`
	MailTemplateTypeId string                    `json:"mailTemplateTypeId"` // required
	Media              []MailTemplateMedia       `json:"media,omitempty"`
	SenderName         string                    `json:"senderName,omitempty"`
	Subject            string                    `json:"subject"` // required
	SystemDefault      bool                      `json:"systemDefault,omitempty"`
	Translated         interface{}               `json:"translated,omitempty"`
	Translations       []MailTemplateTranslation `json:"translations,omitempty"`
	UpdatedAt          time.Time                 `json:"updatedAt,omitempty"`
}
