package model

import "time"

type User struct {
	AccessKeys                 []UserAccessKey       `json:"accessKeys,omitempty"`
	AclRoles                   []AclRole             `json:"aclRoles,omitempty"`
	Active                     bool                  `json:"active"`
	Admin                      bool                  `json:"admin"`
	AvatarId                   string                `json:"avatarId,omitempty"`
	AvatarMedia                *Media                `json:"avatarMedia,omitempty"`
	Configs                    []UserConfig          `json:"configs,omitempty"`
	CreatedAt                  *time.Time            `json:"createdAt,omitempty"`
	CreatedNotifications       []Notification        `json:"createdNotifications,omitempty"`
	CreatedOrders              []Order               `json:"createdOrders,omitempty"`
	CustomFields               interface{}           `json:"customFields,omitempty"`
	Email                      string                `json:"email,omitempty"`     // required
	FirstName                  string                `json:"firstName,omitempty"` // required
	Id                         string                `json:"id,omitempty"`
	ImportExportLogEntries     []ImportExportLog     `json:"importExportLogEntries,omitempty"`
	LastName                   string                `json:"lastName,omitempty"` // required
	LastUpdatedPasswordAt      *time.Time            `json:"lastUpdatedPasswordAt,omitempty"`
	Locale                     *Locale               `json:"locale,omitempty"`
	LocaleId                   string                `json:"localeId,omitempty"` // required
	Media                      []Media               `json:"media,omitempty"`
	Password                   interface{}           `json:"password,omitempty"`
	RecoveryUser               *UserRecovery         `json:"recoveryUser,omitempty"`
	StateMachineHistoryEntries []StateMachineHistory `json:"stateMachineHistoryEntries,omitempty"`
	StoreToken                 string                `json:"storeToken,omitempty"`
	TimeZone                   string                `json:"timeZone,omitempty"` // required
	Title                      string                `json:"title,omitempty"`
	UpdatedAt                  *time.Time            `json:"updatedAt,omitempty"`
	UpdatedOrders              []Order               `json:"updatedOrders,omitempty"`
	Username                   string                `json:"username,omitempty"` // required
}
