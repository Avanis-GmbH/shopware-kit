package model

import "time"

type SnippetSet struct {
	BaseFile            string               `json:"baseFile,omitempty"`
	CreatedAt           time.Time            `json:"createdAt,omitempty"`
	CustomFields        interface{}          `json:"customFields,omitempty"`
	Id                  string               `json:"id,omitempty"`
	Iso                 string               `json:"iso,omitempty"`
	Name                string               `json:"name,omitempty"`
	SalesChannelDomains []SalesChannelDomain `json:"salesChannelDomains,omitempty"`
	Snippets            []Snippet            `json:"snippets,omitempty"`
	UpdatedAt           time.Time            `json:"updatedAt,omitempty"`
}
