package model

type LandingPageSalesChannel struct {
	LandingPage          *LandingPage  `json:"landingPage,omitempty"`
	LandingPageId        string        `json:"landingPageId"` // required
	LandingPageVersionId string        `json:"landingPageVersionId,omitempty"`
	SalesChannel         *SalesChannel `json:"salesChannel,omitempty"`
	SalesChannelId       string        `json:"salesChannelId"` // required
}
