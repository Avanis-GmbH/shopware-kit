package model

import "time"

type SalesChannel struct {
	AccessKey                       string                           `json:"accessKey,omitempty"` // required
	Active                          *bool                            `json:"active,omitempty"`
	Analytics                       *SalesChannelAnalytics           `json:"analytics,omitempty"`
	AnalyticsId                     string                           `json:"analyticsId,omitempty"`
	BoundCustomers                  []Customer                       `json:"boundCustomers,omitempty"`
	Configuration                   interface{}                      `json:"configuration,omitempty"`
	Countries                       []Country                        `json:"countries,omitempty"`
	Country                         *Country                         `json:"country,omitempty"`
	CountryId                       string                           `json:"countryId,omitempty"` // required
	CreatedAt                       *time.Time                       `json:"createdAt,omitempty"`
	Currencies                      []Currency                       `json:"currencies,omitempty"`
	Currency                        *Currency                        `json:"currency,omitempty"`
	CurrencyId                      string                           `json:"currencyId,omitempty"` // required
	CustomerGroup                   *CustomerGroup                   `json:"customerGroup,omitempty"`
	CustomerGroupId                 string                           `json:"customerGroupId,omitempty"` // required
	CustomerGroupsRegistrations     []CustomerGroup                  `json:"customerGroupsRegistrations,omitempty"`
	Customers                       []Customer                       `json:"customers,omitempty"`
	CustomFields                    interface{}                      `json:"customFields,omitempty"`
	DocumentBaseConfigSalesChannels []DocumentBaseConfigSalesChannel `json:"documentBaseConfigSalesChannels,omitempty"`
	Domains                         []SalesChannelDomain             `json:"domains,omitempty"`
	EventActions                    []EventAction                    `json:"eventActions,omitempty"`
	FooterCategory                  *Category                        `json:"footerCategory,omitempty"`
	FooterCategoryId                string                           `json:"footerCategoryId,omitempty"`
	FooterCategoryVersionId         string                           `json:"footerCategoryVersionId,omitempty"`
	HomeCmsPage                     *CmsPage                         `json:"homeCmsPage,omitempty"`
	HomeCmsPageId                   string                           `json:"homeCmsPageId,omitempty"`
	HomeCmsPageVersionId            string                           `json:"homeCmsPageVersionId,omitempty"`
	HomeEnabled                     *bool                            `json:"homeEnabled,omitempty"` // required
	HomeKeywords                    string                           `json:"homeKeywords,omitempty"`
	HomeMetaDescription             string                           `json:"homeMetaDescription,omitempty"`
	HomeMetaTitle                   string                           `json:"homeMetaTitle,omitempty"`
	HomeName                        string                           `json:"homeName,omitempty"`
	HomeSlotConfig                  interface{}                      `json:"homeSlotConfig,omitempty"`
	HreflangActive                  *bool                            `json:"hreflangActive,omitempty"`
	HreflangDefaultDomain           *SalesChannelDomain              `json:"hreflangDefaultDomain,omitempty"`
	HreflangDefaultDomainId         string                           `json:"hreflangDefaultDomainId,omitempty"`
	Id                              string                           `json:"id,omitempty"`
	LandingPages                    []LandingPage                    `json:"landingPages,omitempty"`
	Language                        *Language                        `json:"language,omitempty"`
	LanguageId                      string                           `json:"languageId,omitempty"`
	Languages                       []Language                       `json:"languages,omitempty"`
	MailHeaderFooter                *MailHeaderFooter                `json:"mailHeaderFooter,omitempty"`
	MailHeaderFooterId              string                           `json:"mailHeaderFooterId,omitempty"`
	MainCategories                  []MainCategory                   `json:"mainCategories,omitempty"`
	Maintenance                     *bool                            `json:"maintenance,omitempty"`
	MaintenanceIpWhitelist          interface{}                      `json:"maintenanceIpWhitelist,omitempty"`
	Name                            string                           `json:"name,omitempty"` // required
	NavigationCategory              *Category                        `json:"navigationCategory,omitempty"`
	NavigationCategoryDepth         float64                          `json:"navigationCategoryDepth,omitempty"`
	NavigationCategoryId            string                           `json:"navigationCategoryId,omitempty"` // required
	NavigationCategoryVersionId     string                           `json:"navigationCategoryVersionId,omitempty"`
	NewsletterRecipients            []NewsletterRecipient            `json:"newsletterRecipients,omitempty"`
	NumberRangeSalesChannels        []NumberRangeSalesChannel        `json:"numberRangeSalesChannels,omitempty"`
	Orders                          []Order                          `json:"orders,omitempty"`
	PaymentMethod                   *PaymentMethod                   `json:"paymentMethod,omitempty"`
	PaymentMethodId                 string                           `json:"paymentMethodId,omitempty"` // required
	PaymentMethodIds                interface{}                      `json:"paymentMethodIds,omitempty"`
	PaymentMethods                  []PaymentMethod                  `json:"paymentMethods,omitempty"`
	ProductExports                  []ProductExport                  `json:"productExports,omitempty"`
	ProductReviews                  []ProductReview                  `json:"productReviews,omitempty"`
	ProductVisibilities             []ProductVisibility              `json:"productVisibilities,omitempty"`
	PromotionSalesChannels          []PromotionSalesChannel          `json:"promotionSalesChannels,omitempty"`
	SeoUrls                         []SeoUrl                         `json:"seoUrls,omitempty"`
	SeoUrlTemplates                 []SeoUrlTemplate                 `json:"seoUrlTemplates,omitempty"`
	ServiceCategory                 *Category                        `json:"serviceCategory,omitempty"`
	ServiceCategoryId               string                           `json:"serviceCategoryId,omitempty"`
	ServiceCategoryVersionId        string                           `json:"serviceCategoryVersionId,omitempty"`
	ShippingMethod                  *ShippingMethod                  `json:"shippingMethod,omitempty"`
	ShippingMethodId                string                           `json:"shippingMethodId,omitempty"` // required
	ShippingMethods                 []ShippingMethod                 `json:"shippingMethods,omitempty"`
	ShortName                       string                           `json:"shortName,omitempty"`
	SystemConfigs                   []SystemConfig                   `json:"systemConfigs,omitempty"`
	TaxCalculationType              string                           `json:"taxCalculationType,omitempty"`
	Themes                          []Theme                          `json:"themes,omitempty"`
	Translated                      interface{}                      `json:"translated,omitempty"`
	Translations                    []SalesChannelTranslation        `json:"translations,omitempty"`
	Type                            *SalesChannelType                `json:"type,omitempty"`
	TypeId                          string                           `json:"typeId,omitempty"` // required
	UpdatedAt                       *time.Time                       `json:"updatedAt,omitempty"`
	Wishlists                       []CustomerWishlist               `json:"wishlists,omitempty"`
}
