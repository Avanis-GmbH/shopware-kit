package model

type ThemeMedia struct {
	Media   *Media `json:"media,omitempty"`
	MediaId string `json:"mediaId,omitempty"` // required
	Theme   *Theme `json:"theme,omitempty"`
	ThemeId string `json:"themeId,omitempty"` // required
}
