package model

type ProductCustomFieldSet struct {
	CustomFieldSet   *CustomFieldSet `json:"customFieldSet,omitempty"`
	CustomFieldSetId string          `json:"customFieldSetId"` // required
	Product          *Product        `json:"product,omitempty"`
	ProductId        string          `json:"productId"` // required
	ProductVersionId string          `json:"productVersionId,omitempty"`
}
