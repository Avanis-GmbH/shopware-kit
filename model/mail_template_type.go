package model

import "time"

type MailTemplateType struct {
	AvailableEntities interface{}                   `json:"availableEntities,omitempty"`
	CreatedAt         time.Time                     `json:"createdAt,omitempty"`
	CustomFields      interface{}                   `json:"customFields,omitempty"`
	Id                string                        `json:"id,omitempty"`
	MailTemplates     []MailTemplate                `json:"mailTemplates,omitempty"`
	Name              string                        `json:"name"`          // required
	TechnicalName     string                        `json:"technicalName"` // required
	TemplateData      interface{}                   `json:"templateData,omitempty"`
	Translated        interface{}                   `json:"translated,omitempty"`
	Translations      []MailTemplateTypeTranslation `json:"translations,omitempty"`
	UpdatedAt         time.Time                     `json:"updatedAt,omitempty"`
}
