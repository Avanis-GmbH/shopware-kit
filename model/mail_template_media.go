package model

type MailTemplateMedia struct {
	Id             string        `json:"id,omitempty"`
	LanguageId     string        `json:"languageId,omitempty"` // required
	MailTemplate   *MailTemplate `json:"mailTemplate,omitempty"`
	MailTemplateId string        `json:"mailTemplateId,omitempty"` // required
	Media          *Media        `json:"media,omitempty"`
	MediaId        string        `json:"mediaId,omitempty"` // required
	Position       float64       `json:"position,omitempty"`
}
