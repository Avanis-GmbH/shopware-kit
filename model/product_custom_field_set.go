package model

type ProductCustomFieldSet struct {
	CustomFieldSet   *CustomFieldSet `json:"customFieldSet,omitempty"`
	CustomFieldSetId string          `json:"customFieldSetId,omitempty"` // required
	Product          *Product        `json:"product,omitempty"`
	ProductId        string          `json:"productId,omitempty"` // required
	ProductVersionId string          `json:"productVersionId,omitempty"`
}
