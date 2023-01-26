package model

type IntegrationRole struct {
	AclRoleId     string       `json:"aclRoleId,omitempty"`
	Integration   *Integration `json:"integration,omitempty"`
	IntegrationId string       `json:"integrationId,omitempty"`
	Role          *AclRole     `json:"role,omitempty"`
}
