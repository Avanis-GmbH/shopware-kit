package model

import "time"

type MailTemplateTranslation struct {
	ContentHtml    string        `json:"contentHtml,omitempty"`
	ContentPlain   string        `json:"contentPlain,omitempty"`
	CreatedAt      time.Time     `json:"createdAt,omitempty"`
	CustomFields   interface{}   `json:"customFields,omitempty"`
	Description    string        `json:"description,omitempty"`
	Language       *Language     `json:"language,omitempty"`
	LanguageId     string        `json:"languageId,omitempty"`
	MailTemplate   *MailTemplate `json:"mailTemplate,omitempty"`
	MailTemplateId string        `json:"mailTemplateId,omitempty"`
	SenderName     string        `json:"senderName,omitempty"`
	Subject        string        `json:"subject,omitempty"`
	UpdatedAt      time.Time     `json:"updatedAt,omitempty"`
}
