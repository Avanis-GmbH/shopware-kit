package com

import "github.com/Avanis-GmbH/shopware-kit/model"

// Collection is the interface for all collections returned by the shopware api
// Implemented by EntityCollection which in turn is embedded by all other collections.
type Collection interface {
	setTotal(int64)
	getTotal() int64
	setAggregations(interface{})
	getAggregations() interface{}
	setData([]interface{})
	getData() []interface{}
}

// EntityCollection is the base collection for all other collections
type EntityCollection struct {
	Total        int64       `json:"total"`
	Aggregations interface{} `json:"aggregations"`

	Data []interface{} `json:"data"`
}

func (c *EntityCollection) setTotal(total int64) {
	c.Total = total
}

func (c EntityCollection) getTotal() int64 {
	return c.Total
}

func (c *EntityCollection) setAggregations(aggregations interface{}) {
	c.Aggregations = aggregations
}

func (c EntityCollection) getAggregations() interface{} {
	return c.Aggregations
}

func (c *EntityCollection) setData(data []interface{}) {
	c.Data = data
}

func (c EntityCollection) getData() []interface{} {
	return c.Data
}

type AclRoleCollection struct {
	EntityCollection

	Data []model.AclRole `json:"data"`
}

type AclUserRoleCollection struct {
	EntityCollection

	Data []model.AclUserRole `json:"data"`
}

type AppActionButtonTranslationCollection struct {
	EntityCollection

	Data []model.AppActionButtonTranslation `json:"data"`
}

type AppActionButtonCollection struct {
	EntityCollection

	Data []model.AppActionButton `json:"data"`
}

type AppCmsBlockTranslationCollection struct {
	EntityCollection

	Data []model.AppCmsBlockTranslation `json:"data"`
}

type AppCmsBlockCollection struct {
	EntityCollection

	Data []model.AppCmsBlock `json:"data"`
}

type AppPaymentMethodCollection struct {
	EntityCollection

	Data []model.AppPaymentMethod `json:"data"`
}

type AppTemplateCollection struct {
	EntityCollection

	Data []model.AppTemplate `json:"data"`
}

type AppTranslationCollection struct {
	EntityCollection

	Data []model.AppTranslation `json:"data"`
}

type AppCollection struct {
	EntityCollection

	Data []model.App `json:"data"`
}

type CategoryTagCollection struct {
	EntityCollection

	Data []model.CategoryTag `json:"data"`
}

type CategoryTranslationCollection struct {
	EntityCollection

	Data []model.CategoryTranslation `json:"data"`
}

type CategoryCollection struct {
	EntityCollection

	Data []model.Category `json:"data"`
}

type CmsBlockCollection struct {
	EntityCollection

	Data []model.CmsBlock `json:"data"`
}

type CmsPageTranslationCollection struct {
	EntityCollection

	Data []model.CmsPageTranslation `json:"data"`
}

type CmsPageCollection struct {
	EntityCollection

	Data []model.CmsPage `json:"data"`
}

type CmsSectionCollection struct {
	EntityCollection

	Data []model.CmsSection `json:"data"`
}

type CmsSlotTranslationCollection struct {
	EntityCollection

	Data []model.CmsSlotTranslation `json:"data"`
}

type CmsSlotCollection struct {
	EntityCollection

	Data []model.CmsSlot `json:"data"`
}

type CountryStateTranslationCollection struct {
	EntityCollection

	Data []model.CountryStateTranslation `json:"data"`
}

type CountryStateCollection struct {
	EntityCollection

	Data []model.CountryState `json:"data"`
}

type CountryTranslationCollection struct {
	EntityCollection

	Data []model.CountryTranslation `json:"data"`
}

type CountryCollection struct {
	EntityCollection

	Data []model.Country `json:"data"`
}

type CurrencyCountryRoundingCollection struct {
	EntityCollection

	Data []model.CurrencyCountryRounding `json:"data"`
}

type CurrencyTranslationCollection struct {
	EntityCollection

	Data []model.CurrencyTranslation `json:"data"`
}

type CurrencyCollection struct {
	EntityCollection

	Data []model.Currency `json:"data"`
}

type CustomFieldSetRelationCollection struct {
	EntityCollection

	Data []model.CustomFieldSetRelation `json:"data"`
}

type CustomFieldSetCollection struct {
	EntityCollection

	Data []model.CustomFieldSet `json:"data"`
}

type CustomPriceCollection struct {
	EntityCollection

	Data []model.CustomPrice `json:"data"`
}

type CustomFieldCollection struct {
	EntityCollection

	Data []model.CustomField `json:"data"`
}

type CustomerAddressCollection struct {
	EntityCollection

	Data []model.CustomerAddress `json:"data"`
}

type CustomerGroupRegistrationSalesChannelsCollection struct {
	EntityCollection

	Data []model.CustomerGroupRegistrationSalesChannels `json:"data"`
}

type CustomerGroupTranslationCollection struct {
	EntityCollection

	Data []model.CustomerGroupTranslation `json:"data"`
}

type CustomerGroupCollection struct {
	EntityCollection

	Data []model.CustomerGroup `json:"data"`
}

type CustomerRecoveryCollection struct {
	EntityCollection

	Data []model.CustomerRecovery `json:"data"`
}

type CustomerTagCollection struct {
	EntityCollection

	Data []model.CustomerTag `json:"data"`
}

type CustomerWishlistProductCollection struct {
	EntityCollection

	Data []model.CustomerWishlistProduct `json:"data"`
}

type CustomerWishlistCollection struct {
	EntityCollection

	Data []model.CustomerWishlist `json:"data"`
}

type CustomerCollection struct {
	EntityCollection

	Data []model.Customer `json:"data"`
}

type DeadMessageCollection struct {
	EntityCollection

	Data []model.DeadMessage `json:"data"`
}

type DeliveryTimeTranslationCollection struct {
	EntityCollection

	Data []model.DeliveryTimeTranslation `json:"data"`
}

type DeliveryTimeCollection struct {
	EntityCollection

	Data []model.DeliveryTime `json:"data"`
}

type DocumentBaseConfigSalesChannelCollection struct {
	EntityCollection

	Data []model.DocumentBaseConfigSalesChannel `json:"data"`
}

type DocumentBaseConfigCollection struct {
	EntityCollection

	Data []model.DocumentBaseConfig `json:"data"`
}

type DocumentTypeTranslationCollection struct {
	EntityCollection

	Data []model.DocumentTypeTranslation `json:"data"`
}

type DocumentTypeCollection struct {
	EntityCollection

	Data []model.DocumentType `json:"data"`
}

type DocumentCollection struct {
	EntityCollection

	Data []model.Document `json:"data"`
}

type EventActionRuleCollection struct {
	EntityCollection

	Data []model.EventActionRule `json:"data"`
}

type EventActionSalesChannelCollection struct {
	EntityCollection

	Data []model.EventActionSalesChannel `json:"data"`
}

type EventActionCollection struct {
	EntityCollection

	Data []model.EventAction `json:"data"`
}

type FlowSequenceCollection struct {
	EntityCollection

	Data []model.FlowSequence `json:"data"`
}

type FlowCollection struct {
	EntityCollection

	Data []model.Flow `json:"data"`
}

type ImportExportFileCollection struct {
	EntityCollection

	Data []model.ImportExportFile `json:"data"`
}

type ImportExportLogCollection struct {
	EntityCollection

	Data []model.ImportExportLog `json:"data"`
}

type ImportExportProfileTranslationCollection struct {
	EntityCollection

	Data []model.ImportExportProfileTranslation `json:"data"`
}

type ImportExportProfileCollection struct {
	EntityCollection

	Data []model.ImportExportProfile `json:"data"`
}

type IntegrationRoleCollection struct {
	EntityCollection

	Data []model.IntegrationRole `json:"data"`
}

type IntegrationCollection struct {
	EntityCollection

	Data []model.Integration `json:"data"`
}

type LandingPageSalesChannelCollection struct {
	EntityCollection

	Data []model.LandingPageSalesChannel `json:"data"`
}

type LandingPageTagCollection struct {
	EntityCollection

	Data []model.LandingPageTag `json:"data"`
}

type LandingPageTranslationCollection struct {
	EntityCollection

	Data []model.LandingPageTranslation `json:"data"`
}

type LandingPageCollection struct {
	EntityCollection

	Data []model.LandingPage `json:"data"`
}

type LanguageCollection struct {
	EntityCollection

	Data []model.Language `json:"data"`
}

type LocaleTranslationCollection struct {
	EntityCollection

	Data []model.LocaleTranslation `json:"data"`
}

type LocaleCollection struct {
	EntityCollection

	Data []model.Locale `json:"data"`
}

type LogEntryCollection struct {
	EntityCollection

	Data []model.LogEntry `json:"data"`
}

type MailHeaderFooterTranslationCollection struct {
	EntityCollection

	Data []model.MailHeaderFooterTranslation `json:"data"`
}

type MailHeaderFooterCollection struct {
	EntityCollection

	Data []model.MailHeaderFooter `json:"data"`
}

type MailTemplateMediaCollection struct {
	EntityCollection

	Data []model.MailTemplateMedia `json:"data"`
}

type MailTemplateTranslationCollection struct {
	EntityCollection

	Data []model.MailTemplateTranslation `json:"data"`
}

type MailTemplateTypeTranslationCollection struct {
	EntityCollection

	Data []model.MailTemplateTypeTranslation `json:"data"`
}

type MailTemplateTypeCollection struct {
	EntityCollection

	Data []model.MailTemplateType `json:"data"`
}

type MailTemplateCollection struct {
	EntityCollection

	Data []model.MailTemplate `json:"data"`
}

type MainCategoryCollection struct {
	EntityCollection

	Data []model.MainCategory `json:"data"`
}

type MediaDefaultFolderCollection struct {
	EntityCollection

	Data []model.MediaDefaultFolder `json:"data"`
}

type MediaFolderConfigurationMediaThumbnailSizeCollection struct {
	EntityCollection

	Data []model.MediaFolderConfigurationMediaThumbnailSize `json:"data"`
}

type MediaFolderConfigurationCollection struct {
	EntityCollection

	Data []model.MediaFolderConfiguration `json:"data"`
}

type MediaFolderCollection struct {
	EntityCollection

	Data []model.MediaFolder `json:"data"`
}

type MediaTagCollection struct {
	EntityCollection

	Data []model.MediaTag `json:"data"`
}

type MediaThumbnailSizeCollection struct {
	EntityCollection

	Data []model.MediaThumbnailSize `json:"data"`
}

type MediaThumbnailCollection struct {
	EntityCollection

	Data []model.MediaThumbnail `json:"data"`
}

type MediaTranslationCollection struct {
	EntityCollection

	Data []model.MediaTranslation `json:"data"`
}

type MediaCollection struct {
	EntityCollection

	Data []model.Media `json:"data"`
}

type MessageQueueStatsCollection struct {
	EntityCollection

	Data []model.MessageQueueStats `json:"data"`
}

type NewsletterRecipientTagCollection struct {
	EntityCollection

	Data []model.NewsletterRecipientTag `json:"data"`
}

type NewsletterRecipientCollection struct {
	EntityCollection

	Data []model.NewsletterRecipient `json:"data"`
}

type NotificationCollection struct {
	EntityCollection

	Data []model.Notification `json:"data"`
}

type NumberRangeSalesChannelCollection struct {
	EntityCollection

	Data []model.NumberRangeSalesChannel `json:"data"`
}

type NumberRangeStateCollection struct {
	EntityCollection

	Data []model.NumberRangeState `json:"data"`
}

type NumberRangeTranslationCollection struct {
	EntityCollection

	Data []model.NumberRangeTranslation `json:"data"`
}

type NumberRangeTypeTranslationCollection struct {
	EntityCollection

	Data []model.NumberRangeTypeTranslation `json:"data"`
}

type NumberRangeTypeCollection struct {
	EntityCollection

	Data []model.NumberRangeType `json:"data"`
}

type NumberRangeCollection struct {
	EntityCollection

	Data []model.NumberRange `json:"data"`
}

type OrderAddressCollection struct {
	EntityCollection

	Data []model.OrderAddress `json:"data"`
}

type OrderCustomerCollection struct {
	EntityCollection

	Data []model.OrderCustomer `json:"data"`
}

type OrderDeliveryPositionCollection struct {
	EntityCollection

	Data []model.OrderDeliveryPosition `json:"data"`
}

type OrderDeliveryCollection struct {
	EntityCollection

	Data []model.OrderDelivery `json:"data"`
}

type OrderLineItemCollection struct {
	EntityCollection

	Data []model.OrderLineItem `json:"data"`
}

type OrderTagCollection struct {
	EntityCollection

	Data []model.OrderTag `json:"data"`
}

type OrderTransactionCollection struct {
	EntityCollection

	Data []model.OrderTransaction `json:"data"`
}

type OrderCollection struct {
	EntityCollection

	Data []model.Order `json:"data"`
}

type PaymentMethodTranslationCollection struct {
	EntityCollection

	Data []model.PaymentMethodTranslation `json:"data"`
}

type PaymentMethodCollection struct {
	EntityCollection

	Data []model.PaymentMethod `json:"data"`
}

type PluginTranslationCollection struct {
	EntityCollection

	Data []model.PluginTranslation `json:"data"`
}

type PluginCollection struct {
	EntityCollection

	Data []model.Plugin `json:"data"`
}

type ProductCategoryTreeCollection struct {
	EntityCollection

	Data []model.ProductCategoryTree `json:"data"`
}

type ProductCategoryCollection struct {
	EntityCollection

	Data []model.ProductCategory `json:"data"`
}

type ProductConfiguratorSettingCollection struct {
	EntityCollection

	Data []model.ProductConfiguratorSetting `json:"data"`
}

type ProductCrossSellingAssignedProductsCollection struct {
	EntityCollection

	Data []model.ProductCrossSellingAssignedProducts `json:"data"`
}

type ProductCrossSellingTranslationCollection struct {
	EntityCollection

	Data []model.ProductCrossSellingTranslation `json:"data"`
}

type ProductCrossSellingCollection struct {
	EntityCollection

	Data []model.ProductCrossSelling `json:"data"`
}

type ProductCustomFieldSetCollection struct {
	EntityCollection

	Data []model.ProductCustomFieldSet `json:"data"`
}

type ProductExportCollection struct {
	EntityCollection

	Data []model.ProductExport `json:"data"`
}

type ProductFeatureSetTranslationCollection struct {
	EntityCollection

	Data []model.ProductFeatureSetTranslation `json:"data"`
}

type ProductFeatureSetCollection struct {
	EntityCollection

	Data []model.ProductFeatureSet `json:"data"`
}

type ProductKeywordDictionaryCollection struct {
	EntityCollection

	Data []model.ProductKeywordDictionary `json:"data"`
}

type ProductManufacturerTranslationCollection struct {
	EntityCollection

	Data []model.ProductManufacturerTranslation `json:"data"`
}

type ProductManufacturerCollection struct {
	EntityCollection

	Data []model.ProductManufacturer `json:"data"`
}

type ProductMediaCollection struct {
	EntityCollection

	Data []model.ProductMedia `json:"data"`
}

type ProductOptionCollection struct {
	EntityCollection

	Data []model.ProductOption `json:"data"`
}

type ProductPriceCollection struct {
	EntityCollection

	Data []model.ProductPrice `json:"data"`
}

type ProductPropertyCollection struct {
	EntityCollection

	Data []model.ProductProperty `json:"data"`
}

type ProductReviewCollection struct {
	EntityCollection

	Data []model.ProductReview `json:"data"`
}

type ProductSearchConfigFieldCollection struct {
	EntityCollection

	Data []model.ProductSearchConfigField `json:"data"`
}

type ProductSearchConfigCollection struct {
	EntityCollection

	Data []model.ProductSearchConfig `json:"data"`
}

type ProductSearchKeywordCollection struct {
	EntityCollection

	Data []model.ProductSearchKeyword `json:"data"`
}

type ProductSortingTranslationCollection struct {
	EntityCollection

	Data []model.ProductSortingTranslation `json:"data"`
}

type ProductSortingCollection struct {
	EntityCollection

	Data []model.ProductSorting `json:"data"`
}

type ProductStreamFilterCollection struct {
	EntityCollection

	Data []model.ProductStreamFilter `json:"data"`
}

type ProductStreamMappingCollection struct {
	EntityCollection

	Data []model.ProductStreamMapping `json:"data"`
}

type ProductStreamTranslationCollection struct {
	EntityCollection

	Data []model.ProductStreamTranslation `json:"data"`
}

type ProductStreamCollection struct {
	EntityCollection

	Data []model.ProductStream `json:"data"`
}

type ProductTagCollection struct {
	EntityCollection

	Data []model.ProductTag `json:"data"`
}

type ProductTranslationCollection struct {
	EntityCollection

	Data []model.ProductTranslation `json:"data"`
}

type ProductVisibilityCollection struct {
	EntityCollection

	Data []model.ProductVisibility `json:"data"`
}

type ProductCollection struct {
	EntityCollection

	Data []model.Product `json:"data"`
}

type PromotionCartRuleCollection struct {
	EntityCollection

	Data []model.PromotionCartRule `json:"data"`
}

type PromotionDiscountPricesCollection struct {
	EntityCollection

	Data []model.PromotionDiscountPrices `json:"data"`
}

type PromotionDiscountRuleCollection struct {
	EntityCollection

	Data []model.PromotionDiscountRule `json:"data"`
}

type PromotionDiscountCollection struct {
	EntityCollection

	Data []model.PromotionDiscount `json:"data"`
}

type PromotionIndividualCodeCollection struct {
	EntityCollection

	Data []model.PromotionIndividualCode `json:"data"`
}

type PromotionOrderRuleCollection struct {
	EntityCollection

	Data []model.PromotionOrderRule `json:"data"`
}

type PromotionPersonaCustomerCollection struct {
	EntityCollection

	Data []model.PromotionPersonaCustomer `json:"data"`
}

type PromotionPersonaRuleCollection struct {
	EntityCollection

	Data []model.PromotionPersonaRule `json:"data"`
}

type PromotionSalesChannelCollection struct {
	EntityCollection

	Data []model.PromotionSalesChannel `json:"data"`
}

type PromotionSetgroupRuleCollection struct {
	EntityCollection

	Data []model.PromotionSetgroupRule `json:"data"`
}

type PromotionSetgroupCollection struct {
	EntityCollection

	Data []model.PromotionSetgroup `json:"data"`
}

type PromotionTranslationCollection struct {
	EntityCollection

	Data []model.PromotionTranslation `json:"data"`
}

type PromotionCollection struct {
	EntityCollection

	Data []model.Promotion `json:"data"`
}

type PropertyGroupOptionTranslationCollection struct {
	EntityCollection

	Data []model.PropertyGroupOptionTranslation `json:"data"`
}

type PropertyGroupOptionCollection struct {
	EntityCollection

	Data []model.PropertyGroupOption `json:"data"`
}

type PropertyGroupTranslationCollection struct {
	EntityCollection

	Data []model.PropertyGroupTranslation `json:"data"`
}

type PropertyGroupCollection struct {
	EntityCollection

	Data []model.PropertyGroup `json:"data"`
}

type RuleConditionCollection struct {
	EntityCollection

	Data []model.RuleCondition `json:"data"`
}

type RuleTagCollection struct {
	EntityCollection

	Data []model.RuleTag `json:"data"`
}

type RuleCollection struct {
	EntityCollection

	Data []model.Rule `json:"data"`
}

type SalesChannelAnalyticsCollection struct {
	EntityCollection

	Data []model.SalesChannelAnalytics `json:"data"`
}

type SalesChannelCountryCollection struct {
	EntityCollection

	Data []model.SalesChannelCountry `json:"data"`
}

type SalesChannelCurrencyCollection struct {
	EntityCollection

	Data []model.SalesChannelCurrency `json:"data"`
}

type SalesChannelDomainCollection struct {
	EntityCollection

	Data []model.SalesChannelDomain `json:"data"`
}

type SalesChannelLanguageCollection struct {
	EntityCollection

	Data []model.SalesChannelLanguage `json:"data"`
}

type SalesChannelPaymentMethodCollection struct {
	EntityCollection

	Data []model.SalesChannelPaymentMethod `json:"data"`
}

type SalesChannelShippingMethodCollection struct {
	EntityCollection

	Data []model.SalesChannelShippingMethod `json:"data"`
}

type SalesChannelTranslationCollection struct {
	EntityCollection

	Data []model.SalesChannelTranslation `json:"data"`
}

type SalesChannelTypeTranslationCollection struct {
	EntityCollection

	Data []model.SalesChannelTypeTranslation `json:"data"`
}

type SalesChannelTypeCollection struct {
	EntityCollection

	Data []model.SalesChannelType `json:"data"`
}

type SalesChannelCollection struct {
	EntityCollection

	Data []model.SalesChannel `json:"data"`
}

type SalutationTranslationCollection struct {
	EntityCollection

	Data []model.SalutationTranslation `json:"data"`
}

type SalutationCollection struct {
	EntityCollection

	Data []model.Salutation `json:"data"`
}

type ScheduledTaskCollection struct {
	EntityCollection

	Data []model.ScheduledTask `json:"data"`
}

type ScriptCollection struct {
	EntityCollection

	Data []model.Script `json:"data"`
}

type SeoUrlTemplateCollection struct {
	EntityCollection

	Data []model.SeoUrlTemplate `json:"data"`
}

type SeoUrlCollection struct {
	EntityCollection

	Data []model.SeoUrl `json:"data"`
}

type ShippingMethodPriceCollection struct {
	EntityCollection

	Data []model.ShippingMethodPrice `json:"data"`
}

type ShippingMethodTagCollection struct {
	EntityCollection

	Data []model.ShippingMethodTag `json:"data"`
}

type ShippingMethodTranslationCollection struct {
	EntityCollection

	Data []model.ShippingMethodTranslation `json:"data"`
}

type ShippingMethodCollection struct {
	EntityCollection

	Data []model.ShippingMethod `json:"data"`
}

type SnippetSetCollection struct {
	EntityCollection

	Data []model.SnippetSet `json:"data"`
}

type SnippetCollection struct {
	EntityCollection

	Data []model.Snippet `json:"data"`
}

type StateMachineHistoryCollection struct {
	EntityCollection

	Data []model.StateMachineHistory `json:"data"`
}

type StateMachineStateTranslationCollection struct {
	EntityCollection

	Data []model.StateMachineStateTranslation `json:"data"`
}

type StateMachineStateCollection struct {
	EntityCollection

	Data []model.StateMachineState `json:"data"`
}

type StateMachineTransitionCollection struct {
	EntityCollection

	Data []model.StateMachineTransition `json:"data"`
}

type StateMachineTranslationCollection struct {
	EntityCollection

	Data []model.StateMachineTranslation `json:"data"`
}

type StateMachineCollection struct {
	EntityCollection

	Data []model.StateMachine `json:"data"`
}

type SystemConfigCollection struct {
	EntityCollection

	Data []model.SystemConfig `json:"data"`
}

type TagCollection struct {
	EntityCollection

	Data []model.Tag `json:"data"`
}

type TaxRuleTypeTranslationCollection struct {
	EntityCollection

	Data []model.TaxRuleTypeTranslation `json:"data"`
}

type TaxRuleTypeCollection struct {
	EntityCollection

	Data []model.TaxRuleType `json:"data"`
}

type TaxRuleCollection struct {
	EntityCollection

	Data []model.TaxRule `json:"data"`
}

type TaxCollection struct {
	EntityCollection

	Data []model.Tax `json:"data"`
}

type ThemeChildCollection struct {
	EntityCollection

	Data []model.ThemeChild `json:"data"`
}

type ThemeMediaCollection struct {
	EntityCollection

	Data []model.ThemeMedia `json:"data"`
}

type ThemeSalesChannelCollection struct {
	EntityCollection

	Data []model.ThemeSalesChannel `json:"data"`
}

type ThemeTranslationCollection struct {
	EntityCollection

	Data []model.ThemeTranslation `json:"data"`
}

type ThemeCollection struct {
	EntityCollection

	Data []model.Theme `json:"data"`
}

type UnitTranslationCollection struct {
	EntityCollection

	Data []model.UnitTranslation `json:"data"`
}

type UnitCollection struct {
	EntityCollection

	Data []model.Unit `json:"data"`
}

type UserAccessKeyCollection struct {
	EntityCollection

	Data []model.UserAccessKey `json:"data"`
}

type UserConfigCollection struct {
	EntityCollection

	Data []model.UserConfig `json:"data"`
}

type UserRecoveryCollection struct {
	EntityCollection

	Data []model.UserRecovery `json:"data"`
}

type UserCollection struct {
	EntityCollection

	Data []model.User `json:"data"`
}

type VersionCommitDataCollection struct {
	EntityCollection

	Data []model.VersionCommitData `json:"data"`
}

type VersionCommitCollection struct {
	EntityCollection

	Data []model.VersionCommit `json:"data"`
}

type VersionCollection struct {
	EntityCollection

	Data []model.Version `json:"data"`
}

type WebhookEventLogCollection struct {
	EntityCollection

	Data []model.WebhookEventLog `json:"data"`
}

type WebhookCollection struct {
	EntityCollection

	Data []model.Webhook `json:"data"`
}
