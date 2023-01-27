package model

type ThemeSalesChannel struct {
	SalesChannel   *SalesChannel `json:"salesChannel,omitempty"`
	SalesChannelId string        `json:"salesChannelId,omitempty"` // required
	Theme          *Theme        `json:"theme,omitempty"`
	ThemeId        string        `json:"themeId,omitempty"` // required
}
