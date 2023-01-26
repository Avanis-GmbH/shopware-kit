package model

type MailTemplateMedia struct {
	Id             string        `json:"id,omitempty"`
	LanguageId     string        `json:"languageId,omitempty"`
	MailTemplate   *MailTemplate `json:"mailTemplate,omitempty"`
	MailTemplateId string        `json:"mailTemplateId,omitempty"`
	Media          *Media        `json:"media,omitempty"`
	MediaId        string        `json:"mediaId,omitempty"`
	Position       float64       `json:"position,omitempty"`
}
