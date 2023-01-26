package model

import "time"

type LandingPageTranslation struct {
	CreatedAt            time.Time    `json:"createdAt,omitempty"`
	CustomFields         interface{}  `json:"customFields,omitempty"`
	Keywords             string       `json:"keywords,omitempty"`
	LandingPage          *LandingPage `json:"landingPage,omitempty"`
	LandingPageId        string       `json:"landingPageId,omitempty"`
	LandingPageVersionId string       `json:"landingPageVersionId,omitempty"`
	Language             *Language    `json:"language,omitempty"`
	LanguageId           string       `json:"languageId,omitempty"`
	MetaDescription      string       `json:"metaDescription,omitempty"`
	MetaTitle            string       `json:"metaTitle,omitempty"`
	Name                 string       `json:"name,omitempty"`
	SlotConfig           interface{}  `json:"slotConfig,omitempty"`
	UpdatedAt            time.Time    `json:"updatedAt,omitempty"`
	Url                  string       `json:"url,omitempty"`
}

type LandingPageTranslationCollection struct {
	EntityCollection

	Data []LandingPageTranslation `json:"data"`
}
