package model

import "time"

type Category struct {
	Active                  bool                  `json:"active"`
	AfterCategoryId         string                `json:"afterCategoryId,omitempty"`
	AfterCategoryVersionId  string                `json:"afterCategoryVersionId,omitempty"`
	AutoIncrement           float64               `json:"autoIncrement,omitempty"`
	Breadcrumb              interface{}           `json:"breadcrumb,omitempty"`
	ChildCount              float64               `json:"childCount,omitempty"`
	Children                []Category            `json:"children,omitempty"`
	CmsPage                 *CmsPage              `json:"cmsPage,omitempty"`
	CmsPageId               string                `json:"cmsPageId,omitempty"`
	CmsPageVersionId        string                `json:"cmsPageVersionId,omitempty"`
	CreatedAt               *time.Time            `json:"createdAt,omitempty"`
	CustomFields            interface{}           `json:"customFields,omitempty"`
	Description             string                `json:"description"`
	DisplayNestedProducts   bool                  `json:"displayNestedProducts"` // required
	ExternalLink            string                `json:"externalLink,omitempty"`
	FooterSalesChannels     []SalesChannel        `json:"footerSalesChannels,omitempty"`
	Id                      string                `json:"id,omitempty"`
	InternalLink            string                `json:"internalLink,omitempty"`
	Keywords                string                `json:"keywords,omitempty"`
	Level                   float64               `json:"level,omitempty"`
	LinkNewTab              bool                  `json:"linkNewTab"`
	LinkType                string                `json:"linkType,omitempty"`
	MainCategories          []MainCategory        `json:"mainCategories,omitempty"`
	Media                   *Media                `json:"media,omitempty"`
	MediaId                 *string               `json:"mediaId"`
	MetaDescription         string                `json:"metaDescription"`
	MetaTitle               string                `json:"metaTitle"`
	Name                    string                `json:"name,omitempty"` // required
	NavigationSalesChannels []SalesChannel        `json:"navigationSalesChannels,omitempty"`
	NestedProducts          []Product             `json:"nestedProducts,omitempty"`
	Parent                  *Category             `json:"parent,omitempty"`
	ParentId                string                `json:"parentId,omitempty"`
	ParentVersionId         string                `json:"parentVersionId,omitempty"`
	Path                    string                `json:"path,omitempty"`
	ProductAssignmentType   string                `json:"productAssignmentType,omitempty"` // required
	Products                []Product             `json:"products,omitempty"`
	ProductStream           *ProductStream        `json:"productStream,omitempty"`
	ProductStreamId         string                `json:"productStreamId,omitempty"`
	SeoUrls                 []SeoUrl              `json:"seoUrls,omitempty"`
	ServiceSalesChannels    []SalesChannel        `json:"serviceSalesChannels,omitempty"`
	SlotConfig              interface{}           `json:"slotConfig,omitempty"`
	Tags                    []Tag                 `json:"tags,omitempty"`
	Translated              interface{}           `json:"translated,omitempty"`
	Translations            []CategoryTranslation `json:"translations,omitempty"`
	Type                    string                `json:"type,omitempty"` // required
	UpdatedAt               *time.Time            `json:"updatedAt,omitempty"`
	VersionId               string                `json:"versionId,omitempty"`
	Visible                 bool                  `json:"visible"`
}
