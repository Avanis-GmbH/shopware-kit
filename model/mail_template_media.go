package model

type MailTemplateMedia struct {
	Id             string        `json:"id,omitempty"`
	LanguageId     string        `json:"languageId"` // required
	MailTemplate   *MailTemplate `json:"mailTemplate,omitempty"`
	MailTemplateId string        `json:"mailTemplateId"` // required
	Media          *Media        `json:"media,omitempty"`
	MediaId        string        `json:"mediaId"` // required
	Position       float64       `json:"position"`
}
