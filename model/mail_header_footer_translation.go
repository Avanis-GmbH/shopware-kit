package model

import "time"

type MailHeaderFooterTranslation struct {
	CreatedAt          time.Time         `json:"createdAt,omitempty"`
	Description        string            `json:"description,omitempty"`
	FooterHtml         string            `json:"footerHtml,omitempty"`
	FooterPlain        string            `json:"footerPlain,omitempty"`
	HeaderHtml         string            `json:"headerHtml,omitempty"`
	HeaderPlain        string            `json:"headerPlain,omitempty"`
	Language           *Language         `json:"language,omitempty"`
	LanguageId         string            `json:"languageId,omitempty"`
	MailHeaderFooter   *MailHeaderFooter `json:"mailHeaderFooter,omitempty"`
	MailHeaderFooterId string            `json:"mailHeaderFooterId,omitempty"`
	Name               string            `json:"name,omitempty"`
	UpdatedAt          time.Time         `json:"updatedAt,omitempty"`
}
