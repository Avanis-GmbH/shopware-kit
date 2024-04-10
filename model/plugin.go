package model

import "time"

type Plugin struct {
	Active            bool                `json:"active"`
	Author            string              `json:"author,omitempty"`
	Autoload          interface{}         `json:"autoload,omitempty"`  // required
	BaseClass         string              `json:"baseClass,omitempty"` // required
	Changelog         interface{}         `json:"changelog,omitempty"`
	ComposerName      string              `json:"composerName,omitempty"`
	Copyright         string              `json:"copyright,omitempty"`
	CreatedAt         *time.Time          `json:"createdAt,omitempty"`
	CustomFields      interface{}         `json:"customFields,omitempty"`
	Description       string              `json:"description,omitempty"`
	Icon              string              `json:"icon,omitempty"`
	IconRaw           interface{}         `json:"iconRaw,omitempty"`
	Id                string              `json:"id,omitempty"`
	InstalledAt       *time.Time          `json:"installedAt,omitempty"`
	Label             string              `json:"label,omitempty"` // required
	License           string              `json:"license,omitempty"`
	ManagedByComposer bool                `json:"managedByComposer"`
	ManufacturerLink  string              `json:"manufacturerLink,omitempty"`
	Name              string              `json:"name,omitempty"` // required
	Path              string              `json:"path,omitempty"`
	PaymentMethods    []PaymentMethod     `json:"paymentMethods,omitempty"`
	SupportLink       string              `json:"supportLink,omitempty"`
	Translated        interface{}         `json:"translated,omitempty"`
	Translations      []PluginTranslation `json:"translations,omitempty"`
	UpdatedAt         *time.Time          `json:"updatedAt,omitempty"`
	UpgradedAt        *time.Time          `json:"upgradedAt,omitempty"`
	UpgradeVersion    string              `json:"upgradeVersion,omitempty"`
	Version           string              `json:"version,omitempty"` // required
}
