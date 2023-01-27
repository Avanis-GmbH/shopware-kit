package model

type ProductKeywordDictionary struct {
	Id         string    `json:"id,omitempty"`
	Keyword    string    `json:"keyword"` // required
	Language   *Language `json:"language,omitempty"`
	LanguageId string    `json:"languageId"` // required
	Reversed   string    `json:"reversed,omitempty"`
}
