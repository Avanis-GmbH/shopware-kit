package model

type ProductKeywordDictionary struct {
	Id         string    `json:"id,omitempty"`
	Keyword    string    `json:"keyword,omitempty"` // required
	Language   *Language `json:"language,omitempty"`
	LanguageId string    `json:"languageId,omitempty"` // required
	Reversed   string    `json:"reversed,omitempty"`
}
