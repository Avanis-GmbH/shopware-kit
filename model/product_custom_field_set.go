package model

type ProductCustomFieldSet struct {
	CustomFieldSet   *CustomFieldSet `json:"customFieldSet,omitempty"`
	CustomFieldSetId string          `json:"customFieldSetId,omitempty"`
	Product          *Product        `json:"product,omitempty"`
	ProductId        string          `json:"productId,omitempty"`
	ProductVersionId string          `json:"productVersionId,omitempty"`
}
