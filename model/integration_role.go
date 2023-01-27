package model

type IntegrationRole struct {
	AclRoleId     string       `json:"aclRoleId,omitempty"` // required
	Integration   *Integration `json:"integration,omitempty"`
	IntegrationId string       `json:"integrationId,omitempty"` // required
	Role          *AclRole     `json:"role,omitempty"`
}
