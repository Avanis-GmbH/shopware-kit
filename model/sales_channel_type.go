package model

import "time"

type SalesChannelType struct {
	CoverUrl        string                        `json:"coverUrl,omitempty"`
	CreatedAt       time.Time                     `json:"createdAt,omitempty"`
	CustomFields    interface{}                   `json:"customFields,omitempty"`
	Description     string                        `json:"description,omitempty"`
	DescriptionLong string                        `json:"descriptionLong,omitempty"`
	IconName        string                        `json:"iconName,omitempty"`
	Id              string                        `json:"id,omitempty"`
	Manufacturer    string                        `json:"manufacturer,omitempty"`
	Name            string                        `json:"name,omitempty"`
	SalesChannels   []SalesChannel                `json:"salesChannels,omitempty"`
	ScreenshotUrls  interface{}                   `json:"screenshotUrls,omitempty"`
	Translated      interface{}                   `json:"translated,omitempty"`
	Translations    []SalesChannelTypeTranslation `json:"translations,omitempty"`
	UpdatedAt       time.Time                     `json:"updatedAt,omitempty"`
}

type SalesChannelTypeCollection struct {
	EntityCollection

	Data []SalesChannelType `json:"data"`
}
