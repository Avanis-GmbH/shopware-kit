package model

type ThemeMedia struct {
	Media   *Media `json:"media,omitempty"`
	MediaId string `json:"mediaId"` // required
	Theme   *Theme `json:"theme,omitempty"`
	ThemeId string `json:"themeId"` // required
}
