package model

import "time"

type Media struct {
	Alt                         string                       `json:"alt,omitempty"`
	AppPaymentMethods           []AppPaymentMethod           `json:"appPaymentMethods,omitempty"`
	AvatarUser                  *User                        `json:"avatarUser,omitempty"`
	Categories                  []Category                   `json:"categories,omitempty"`
	CmsBlocks                   []CmsBlock                   `json:"cmsBlocks,omitempty"`
	CmsPages                    []CmsPage                    `json:"cmsPages,omitempty"`
	CmsSections                 []CmsSection                 `json:"cmsSections,omitempty"`
	CreatedAt                   *time.Time                   `json:"createdAt,omitempty"`
	CustomFields                interface{}                  `json:"customFields,omitempty"`
	DocumentBaseConfigs         []DocumentBaseConfig         `json:"documentBaseConfigs,omitempty"`
	Documents                   []Document                   `json:"documents,omitempty"`
	FileExtension               string                       `json:"fileExtension,omitempty"`
	FileName                    string                       `json:"fileName,omitempty"`
	FileSize                    float64                      `json:"fileSize,omitempty"`
	HasFile                     bool                         `json:"hasFile"`
	Id                          string                       `json:"id,omitempty"`
	MailTemplateMedia           []MailTemplateMedia          `json:"mailTemplateMedia,omitempty"`
	MediaFolder                 *MediaFolder                 `json:"mediaFolder,omitempty"`
	MediaFolderId               string                       `json:"mediaFolderId,omitempty"`
	MediaType                   interface{}                  `json:"mediaType,omitempty"`
	MediaTypeRaw                interface{}                  `json:"mediaTypeRaw,omitempty"`
	MetaData                    interface{}                  `json:"metaData,omitempty"`
	MimeType                    string                       `json:"mimeType,omitempty"`
	OrderLineItems              []OrderLineItem              `json:"orderLineItems,omitempty"`
	PaymentMethods              []PaymentMethod              `json:"paymentMethods,omitempty"`
	Private                     bool                         `json:"private"`
	ProductConfiguratorSettings []ProductConfiguratorSetting `json:"productConfiguratorSettings,omitempty"`
	ProductManufacturers        []ProductManufacturer        `json:"productManufacturers,omitempty"`
	ProductMedia                []ProductMedia               `json:"productMedia,omitempty"`
	PropertyGroupOptions        []PropertyGroupOption        `json:"propertyGroupOptions,omitempty"`
	ShippingMethods             []ShippingMethod             `json:"shippingMethods,omitempty"`
	Tags                        []Tag                        `json:"tags,omitempty"`
	ThemeMedia                  []Theme                      `json:"themeMedia,omitempty"`
	Themes                      []Theme                      `json:"themes,omitempty"`
	Thumbnails                  []MediaThumbnail             `json:"thumbnails,omitempty"`
	ThumbnailsRo                interface{}                  `json:"thumbnailsRo,omitempty"`
	Title                       string                       `json:"title,omitempty"`
	Translated                  interface{}                  `json:"translated,omitempty"`
	Translations                []MediaTranslation           `json:"translations,omitempty"`
	UpdatedAt                   *time.Time                   `json:"updatedAt,omitempty"`
	UploadedAt                  *time.Time                   `json:"uploadedAt,omitempty"`
	Url                         string                       `json:"url,omitempty"`
	User                        *User                        `json:"user,omitempty"`
	UserId                      string                       `json:"userId,omitempty"`
}

type MediaURLUpload struct {
	Url string `json:"url"`
}
