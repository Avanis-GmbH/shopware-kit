package model

type LandingPageSalesChannel struct {
	LandingPage          *LandingPage  `json:"landingPage,omitempty"`
	LandingPageId        string        `json:"landingPageId,omitempty"`
	LandingPageVersionId string        `json:"landingPageVersionId,omitempty"`
	SalesChannel         *SalesChannel `json:"salesChannel,omitempty"`
	SalesChannelId       string        `json:"salesChannelId,omitempty"`
}

type LandingPageSalesChannelCollection struct {
	EntityCollection

	Data []LandingPageSalesChannel `json:"data"`
}
