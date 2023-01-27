package model

import "time"

type Language struct {
	ActionButtonTranslations        []AppActionButtonTranslation     `json:"actionButtonTranslations,omitempty"`
	AppCmsBlockTranslations         []AppCmsBlockTranslation         `json:"appCmsBlockTranslations,omitempty"`
	AppTranslations                 []AppTranslation                 `json:"appTranslations,omitempty"`
	CategoryTranslations            []CategoryTranslation            `json:"categoryTranslations,omitempty"`
	Children                        []Language                       `json:"children,omitempty"`
	CmsPageTranslations             []CmsPageTranslation             `json:"cmsPageTranslations,omitempty"`
	CmsSlotTranslations             []CmsSlotTranslation             `json:"cmsSlotTranslations,omitempty"`
	CountryStateTranslations        []CountryStateTranslation        `json:"countryStateTranslations,omitempty"`
	CountryTranslations             []CountryTranslation             `json:"countryTranslations,omitempty"`
	CreatedAt                       time.Time                        `json:"createdAt,omitempty"`
	CurrencyTranslations            []CurrencyTranslation            `json:"currencyTranslations,omitempty"`
	CustomerGroupTranslations       []CustomerGroupTranslation       `json:"customerGroupTranslations,omitempty"`
	Customers                       []Customer                       `json:"customers,omitempty"`
	CustomFields                    interface{}                      `json:"customFields,omitempty"`
	DeliveryTimeTranslations        []DeliveryTimeTranslation        `json:"deliveryTimeTranslations,omitempty"`
	DocumentTypeTranslations        []DocumentTypeTranslation        `json:"documentTypeTranslations,omitempty"`
	Id                              string                           `json:"id,omitempty"`
	ImportExportProfileTranslations []ImportExportProfileTranslation `json:"importExportProfileTranslations,omitempty"`
	LandingPageTranslations         []LandingPageTranslation         `json:"landingPageTranslations,omitempty"`
	Locale                          *Locale                          `json:"locale,omitempty"`
	LocaleId                        string                           `json:"localeId"` // required
	LocaleTranslations              []LocaleTranslation              `json:"localeTranslations,omitempty"`
	MailHeaderFooterTranslations    []MailHeaderFooterTranslation    `json:"mailHeaderFooterTranslations,omitempty"`
	MailTemplateTranslations        []MailTemplateTranslation        `json:"mailTemplateTranslations,omitempty"`
	MailTemplateTypeTranslations    []MailTemplateTypeTranslation    `json:"mailTemplateTypeTranslations,omitempty"`
	MediaTranslations               []MediaTranslation               `json:"mediaTranslations,omitempty"`
	Name                            string                           `json:"name"` // required
	NewsletterRecipients            []NewsletterRecipient            `json:"newsletterRecipients,omitempty"`
	NumberRangeTranslations         []NumberRangeTranslation         `json:"numberRangeTranslations,omitempty"`
	NumberRangeTypeTranslations     []NumberRangeTypeTranslation     `json:"numberRangeTypeTranslations,omitempty"`
	Orders                          []Order                          `json:"orders,omitempty"`
	Parent                          *Language                        `json:"parent,omitempty"`
	ParentId                        string                           `json:"parentId,omitempty"`
	PaymentMethodTranslations       []PaymentMethodTranslation       `json:"paymentMethodTranslations,omitempty"`
	PluginTranslations              []PluginTranslation              `json:"pluginTranslations,omitempty"`
	ProductCrossSellingTranslations []ProductCrossSellingTranslation `json:"productCrossSellingTranslations,omitempty"`
	ProductFeatureSetTranslations   []ProductFeatureSetTranslation   `json:"productFeatureSetTranslations,omitempty"`
	ProductKeywordDictionaries      []ProductKeywordDictionary       `json:"productKeywordDictionaries,omitempty"`
	ProductManufacturerTranslations []ProductManufacturerTranslation `json:"productManufacturerTranslations,omitempty"`
	ProductReviews                  []ProductReview                  `json:"productReviews,omitempty"`
	ProductSearchConfig             *ProductSearchConfig             `json:"productSearchConfig,omitempty"`
	ProductSearchKeywords           []ProductSearchKeyword           `json:"productSearchKeywords,omitempty"`
	ProductSortingTranslations      []ProductSortingTranslation      `json:"productSortingTranslations,omitempty"`
	ProductStreamTranslations       []ProductStreamTranslation       `json:"productStreamTranslations,omitempty"`
	ProductTranslations             []ProductTranslation             `json:"productTranslations,omitempty"`
	PromotionTranslations           []PromotionTranslation           `json:"promotionTranslations,omitempty"`
	PropertyGroupOptionTranslations []PropertyGroupOptionTranslation `json:"propertyGroupOptionTranslations,omitempty"`
	PropertyGroupTranslations       []PropertyGroupTranslation       `json:"propertyGroupTranslations,omitempty"`
	SalesChannelDefaultAssignments  []SalesChannel                   `json:"salesChannelDefaultAssignments,omitempty"`
	SalesChannelDomains             []SalesChannelDomain             `json:"salesChannelDomains,omitempty"`
	SalesChannels                   []SalesChannel                   `json:"salesChannels,omitempty"`
	SalesChannelTranslations        []SalesChannelTranslation        `json:"salesChannelTranslations,omitempty"`
	SalesChannelTypeTranslations    []SalesChannelTypeTranslation    `json:"salesChannelTypeTranslations,omitempty"`
	SalutationTranslations          []SalutationTranslation          `json:"salutationTranslations,omitempty"`
	SeoUrlTranslations              []SeoUrl                         `json:"seoUrlTranslations,omitempty"`
	ShippingMethodTranslations      []ShippingMethodTranslation      `json:"shippingMethodTranslations,omitempty"`
	StateMachineStateTranslations   []StateMachineStateTranslation   `json:"stateMachineStateTranslations,omitempty"`
	StateMachineTranslations        []StateMachineTranslation        `json:"stateMachineTranslations,omitempty"`
	TaxRuleTypeTranslations         []TaxRuleTypeTranslation         `json:"taxRuleTypeTranslations,omitempty"`
	ThemeTranslations               []ThemeTranslation               `json:"themeTranslations,omitempty"`
	TranslationCode                 *Locale                          `json:"translationCode,omitempty"`
	TranslationCodeId               string                           `json:"translationCodeId,omitempty"`
	UnitTranslations                []UnitTranslation                `json:"unitTranslations,omitempty"`
	UpdatedAt                       time.Time                        `json:"updatedAt,omitempty"`
}
