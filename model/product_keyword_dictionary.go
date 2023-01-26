package model

type ProductKeywordDictionary struct {
	Id         string    `json:"id,omitempty"`
	Keyword    string    `json:"keyword,omitempty"`
	Language   *Language `json:"language,omitempty"`
	LanguageId string    `json:"languageId,omitempty"`
	Reversed   string    `json:"reversed,omitempty"`
}
