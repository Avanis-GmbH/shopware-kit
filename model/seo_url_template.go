package model

import "time"

type SeoUrlTemplate struct {
	CreatedAt      time.Time     `json:"createdAt,omitempty"`
	CustomFields   interface{}   `json:"customFields,omitempty"`
	EntityName     string        `json:"entityName"` // required
	Id             string        `json:"id,omitempty"`
	IsValid        bool          `json:"isValid,omitempty"`
	RouteName      string        `json:"routeName"` // required
	SalesChannel   *SalesChannel `json:"salesChannel,omitempty"`
	SalesChannelId string        `json:"salesChannelId,omitempty"`
	Template       string        `json:"template,omitempty"`
	UpdatedAt      time.Time     `json:"updatedAt,omitempty"`
}
