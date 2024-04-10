package model

import "time"

type MailHeaderFooter struct {
	CreatedAt     *time.Time                    `json:"createdAt,omitempty"`
	Description   string                        `json:"description,omitempty"`
	FooterHtml    string                        `json:"footerHtml,omitempty"`
	FooterPlain   string                        `json:"footerPlain,omitempty"`
	HeaderHtml    string                        `json:"headerHtml,omitempty"`
	HeaderPlain   string                        `json:"headerPlain,omitempty"`
	Id            string                        `json:"id,omitempty"`
	Name          string                        `json:"name,omitempty"` // required
	SalesChannels []SalesChannel                `json:"salesChannels,omitempty"`
	SystemDefault bool                          `json:"systemDefault"`
	Translated    interface{}                   `json:"translated,omitempty"`
	Translations  []MailHeaderFooterTranslation `json:"translations,omitempty"`
	UpdatedAt     *time.Time                    `json:"updatedAt,omitempty"`
}
