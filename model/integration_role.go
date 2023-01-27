package model

type IntegrationRole struct {
	AclRoleId     string       `json:"aclRoleId"` // required
	Integration   *Integration `json:"integration,omitempty"`
	IntegrationId string       `json:"integrationId"` // required
	Role          *AclRole     `json:"role,omitempty"`
}
