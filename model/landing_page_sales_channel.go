package model

type LandingPageSalesChannel struct {
	LandingPage          *LandingPage  `json:"landingPage,omitempty"`
	LandingPageId        string        `json:"landingPageId,omitempty"` // required
	LandingPageVersionId string        `json:"landingPageVersionId,omitempty"`
	SalesChannel         *SalesChannel `json:"salesChannel,omitempty"`
	SalesChannelId       string        `json:"salesChannelId,omitempty"` // required
}
