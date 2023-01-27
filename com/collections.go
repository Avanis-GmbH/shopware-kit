package com

import "github.com/Avanis-GmbH/GoSUCK/model"

type Collection interface {
	setTotal(int64)
	getTotal() int64
	setAggregations(interface{})
	getAggregations() interface{}
	setData([]interface{})
	getData() []interface{}
}

type entityCollection struct {
	Total        int64       `json:"total"`
	Aggregations interface{} `json:"aggregations"`

	Data []interface{} `json:"data"`
}

func (c *entityCollection) setTotal(total int64) {
	c.Total = total
}

func (c entityCollection) getTotal() int64 {
	return c.Total
}

func (c *entityCollection) setAggregations(aggregations interface{}) {
	c.Aggregations = aggregations
}

func (c entityCollection) getAggregations() interface{} {
	return c.Aggregations
}

func (c *entityCollection) setData(data []interface{}) {
	c.Data = data
}

func (c entityCollection) getData() []interface{} {
	return c.Data
}

type AclRoleCollection struct {
	entityCollection

	Data []model.AclRole `json:"data"`
}

type AclUserRoleCollection struct {
	entityCollection

	Data []model.AclUserRole `json:"data"`
}

type AppActionButtonTranslationCollection struct {
	entityCollection

	Data []model.AppActionButtonTranslation `json:"data"`
}

type AppActionButtonCollection struct {
	entityCollection

	Data []model.AppActionButton `json:"data"`
}

type AppCmsBlockTranslationCollection struct {
	entityCollection

	Data []model.AppCmsBlockTranslation `json:"data"`
}

type AppCmsBlockCollection struct {
	entityCollection

	Data []model.AppCmsBlock `json:"data"`
}

type AppPaymentMethodCollection struct {
	entityCollection

	Data []model.AppPaymentMethod `json:"data"`
}

type AppTemplateCollection struct {
	entityCollection

	Data []model.AppTemplate `json:"data"`
}

type AppTranslationCollection struct {
	entityCollection

	Data []model.AppTranslation `json:"data"`
}

type AppCollection struct {
	entityCollection

	Data []model.App `json:"data"`
}

type CategoryTagCollection struct {
	entityCollection

	Data []model.CategoryTag `json:"data"`
}

type CategoryTranslationCollection struct {
	entityCollection

	Data []model.CategoryTranslation `json:"data"`
}

type CategoryCollection struct {
	entityCollection

	Data []model.Category `json:"data"`
}

type CmsBlockCollection struct {
	entityCollection

	Data []model.CmsBlock `json:"data"`
}

type CmsPageTranslationCollection struct {
	entityCollection

	Data []model.CmsPageTranslation `json:"data"`
}

type CmsPageCollection struct {
	entityCollection

	Data []model.CmsPage `json:"data"`
}

type CmsSectionCollection struct {
	entityCollection

	Data []model.CmsSection `json:"data"`
}

type CmsSlotTranslationCollection struct {
	entityCollection

	Data []model.CmsSlotTranslation `json:"data"`
}

type CmsSlotCollection struct {
	entityCollection

	Data []model.CmsSlot `json:"data"`
}

type CountryStateTranslationCollection struct {
	entityCollection

	Data []model.CountryStateTranslation `json:"data"`
}

type CountryStateCollection struct {
	entityCollection

	Data []model.CountryState `json:"data"`
}

type CountryTranslationCollection struct {
	entityCollection

	Data []model.CountryTranslation `json:"data"`
}

type CountryCollection struct {
	entityCollection

	Data []model.Country `json:"data"`
}

type CurrencyCountryRoundingCollection struct {
	entityCollection

	Data []model.CurrencyCountryRounding `json:"data"`
}

type CurrencyTranslationCollection struct {
	entityCollection

	Data []model.CurrencyTranslation `json:"data"`
}

type CurrencyCollection struct {
	entityCollection

	Data []model.Currency `json:"data"`
}

type CustomFieldSetRelationCollection struct {
	entityCollection

	Data []model.CustomFieldSetRelation `json:"data"`
}

type CustomFieldSetCollection struct {
	entityCollection

	Data []model.CustomFieldSet `json:"data"`
}

type CustomFieldCollection struct {
	entityCollection

	Data []model.CustomField `json:"data"`
}

type CustomerAddressCollection struct {
	entityCollection

	Data []model.CustomerAddress `json:"data"`
}

type CustomerGroupRegistrationSalesChannelsCollection struct {
	entityCollection

	Data []model.CustomerGroupRegistrationSalesChannels `json:"data"`
}

type CustomerGroupTranslationCollection struct {
	entityCollection

	Data []model.CustomerGroupTranslation `json:"data"`
}

type CustomerGroupCollection struct {
	entityCollection

	Data []model.CustomerGroup `json:"data"`
}

type CustomerRecoveryCollection struct {
	entityCollection

	Data []model.CustomerRecovery `json:"data"`
}

type CustomerTagCollection struct {
	entityCollection

	Data []model.CustomerTag `json:"data"`
}

type CustomerWishlistProductCollection struct {
	entityCollection

	Data []model.CustomerWishlistProduct `json:"data"`
}

type CustomerWishlistCollection struct {
	entityCollection

	Data []model.CustomerWishlist `json:"data"`
}

type CustomerCollection struct {
	entityCollection

	Data []model.Customer `json:"data"`
}

type DeadMessageCollection struct {
	entityCollection

	Data []model.DeadMessage `json:"data"`
}

type DeliveryTimeTranslationCollection struct {
	entityCollection

	Data []model.DeliveryTimeTranslation `json:"data"`
}

type DeliveryTimeCollection struct {
	entityCollection

	Data []model.DeliveryTime `json:"data"`
}

type DocumentBaseConfigSalesChannelCollection struct {
	entityCollection

	Data []model.DocumentBaseConfigSalesChannel `json:"data"`
}

type DocumentBaseConfigCollection struct {
	entityCollection

	Data []model.DocumentBaseConfig `json:"data"`
}

type DocumentTypeTranslationCollection struct {
	entityCollection

	Data []model.DocumentTypeTranslation `json:"data"`
}

type DocumentTypeCollection struct {
	entityCollection

	Data []model.DocumentType `json:"data"`
}

type DocumentCollection struct {
	entityCollection

	Data []model.Document `json:"data"`
}

type EventActionRuleCollection struct {
	entityCollection

	Data []model.EventActionRule `json:"data"`
}

type EventActionSalesChannelCollection struct {
	entityCollection

	Data []model.EventActionSalesChannel `json:"data"`
}

type EventActionCollection struct {
	entityCollection

	Data []model.EventAction `json:"data"`
}

type FlowSequenceCollection struct {
	entityCollection

	Data []model.FlowSequence `json:"data"`
}

type FlowCollection struct {
	entityCollection

	Data []model.Flow `json:"data"`
}

type ImportExportFileCollection struct {
	entityCollection

	Data []model.ImportExportFile `json:"data"`
}

type ImportExportLogCollection struct {
	entityCollection

	Data []model.ImportExportLog `json:"data"`
}

type ImportExportProfileTranslationCollection struct {
	entityCollection

	Data []model.ImportExportProfileTranslation `json:"data"`
}

type ImportExportProfileCollection struct {
	entityCollection

	Data []model.ImportExportProfile `json:"data"`
}

type IntegrationRoleCollection struct {
	entityCollection

	Data []model.IntegrationRole `json:"data"`
}

type IntegrationCollection struct {
	entityCollection

	Data []model.Integration `json:"data"`
}

type LandingPageSalesChannelCollection struct {
	entityCollection

	Data []model.LandingPageSalesChannel `json:"data"`
}

type LandingPageTagCollection struct {
	entityCollection

	Data []model.LandingPageTag `json:"data"`
}

type LandingPageTranslationCollection struct {
	entityCollection

	Data []model.LandingPageTranslation `json:"data"`
}

type LandingPageCollection struct {
	entityCollection

	Data []model.LandingPage `json:"data"`
}

type LanguageCollection struct {
	entityCollection

	Data []model.Language `json:"data"`
}

type LocaleTranslationCollection struct {
	entityCollection

	Data []model.LocaleTranslation `json:"data"`
}

type LocaleCollection struct {
	entityCollection

	Data []model.Locale `json:"data"`
}

type LogEntryCollection struct {
	entityCollection

	Data []model.LogEntry `json:"data"`
}

type MailHeaderFooterTranslationCollection struct {
	entityCollection

	Data []model.MailHeaderFooterTranslation `json:"data"`
}

type MailHeaderFooterCollection struct {
	entityCollection

	Data []model.MailHeaderFooter `json:"data"`
}

type MailTemplateMediaCollection struct {
	entityCollection

	Data []model.MailTemplateMedia `json:"data"`
}

type MailTemplateTranslationCollection struct {
	entityCollection

	Data []model.MailTemplateTranslation `json:"data"`
}

type MailTemplateTypeTranslationCollection struct {
	entityCollection

	Data []model.MailTemplateTypeTranslation `json:"data"`
}

type MailTemplateTypeCollection struct {
	entityCollection

	Data []model.MailTemplateType `json:"data"`
}

type MailTemplateCollection struct {
	entityCollection

	Data []model.MailTemplate `json:"data"`
}

type MainCategoryCollection struct {
	entityCollection

	Data []model.MainCategory `json:"data"`
}

type MediaDefaultFolderCollection struct {
	entityCollection

	Data []model.MediaDefaultFolder `json:"data"`
}

type MediaFolderConfigurationMediaThumbnailSizeCollection struct {
	entityCollection

	Data []model.MediaFolderConfigurationMediaThumbnailSize `json:"data"`
}

type MediaFolderConfigurationCollection struct {
	entityCollection

	Data []model.MediaFolderConfiguration `json:"data"`
}

type MediaFolderCollection struct {
	entityCollection

	Data []model.MediaFolder `json:"data"`
}

type MediaTagCollection struct {
	entityCollection

	Data []model.MediaTag `json:"data"`
}

type MediaThumbnailSizeCollection struct {
	entityCollection

	Data []model.MediaThumbnailSize `json:"data"`
}

type MediaThumbnailCollection struct {
	entityCollection

	Data []model.MediaThumbnail `json:"data"`
}

type MediaTranslationCollection struct {
	entityCollection

	Data []model.MediaTranslation `json:"data"`
}

type MediaCollection struct {
	entityCollection

	Data []model.Media `json:"data"`
}

type MessageQueueStatsCollection struct {
	entityCollection

	Data []model.MessageQueueStats `json:"data"`
}

type NewsletterRecipientTagCollection struct {
	entityCollection

	Data []model.NewsletterRecipientTag `json:"data"`
}

type NewsletterRecipientCollection struct {
	entityCollection

	Data []model.NewsletterRecipient `json:"data"`
}

type NotificationCollection struct {
	entityCollection

	Data []model.Notification `json:"data"`
}

type NumberRangeSalesChannelCollection struct {
	entityCollection

	Data []model.NumberRangeSalesChannel `json:"data"`
}

type NumberRangeStateCollection struct {
	entityCollection

	Data []model.NumberRangeState `json:"data"`
}

type NumberRangeTranslationCollection struct {
	entityCollection

	Data []model.NumberRangeTranslation `json:"data"`
}

type NumberRangeTypeTranslationCollection struct {
	entityCollection

	Data []model.NumberRangeTypeTranslation `json:"data"`
}

type NumberRangeTypeCollection struct {
	entityCollection

	Data []model.NumberRangeType `json:"data"`
}

type NumberRangeCollection struct {
	entityCollection

	Data []model.NumberRange `json:"data"`
}

type OrderAddressCollection struct {
	entityCollection

	Data []model.OrderAddress `json:"data"`
}

type OrderCustomerCollection struct {
	entityCollection

	Data []model.OrderCustomer `json:"data"`
}

type OrderDeliveryPositionCollection struct {
	entityCollection

	Data []model.OrderDeliveryPosition `json:"data"`
}

type OrderDeliveryCollection struct {
	entityCollection

	Data []model.OrderDelivery `json:"data"`
}

type OrderLineItemCollection struct {
	entityCollection

	Data []model.OrderLineItem `json:"data"`
}

type OrderTagCollection struct {
	entityCollection

	Data []model.OrderTag `json:"data"`
}

type OrderTransactionCollection struct {
	entityCollection

	Data []model.OrderTransaction `json:"data"`
}

type OrderCollection struct {
	entityCollection

	Data []model.Order `json:"data"`
}

type PaymentMethodTranslationCollection struct {
	entityCollection

	Data []model.PaymentMethodTranslation `json:"data"`
}

type PaymentMethodCollection struct {
	entityCollection

	Data []model.PaymentMethod `json:"data"`
}

type PluginTranslationCollection struct {
	entityCollection

	Data []model.PluginTranslation `json:"data"`
}

type PluginCollection struct {
	entityCollection

	Data []model.Plugin `json:"data"`
}

type ProductCategoryTreeCollection struct {
	entityCollection

	Data []model.ProductCategoryTree `json:"data"`
}

type ProductCategoryCollection struct {
	entityCollection

	Data []model.ProductCategory `json:"data"`
}

type ProductConfiguratorSettingCollection struct {
	entityCollection

	Data []model.ProductConfiguratorSetting `json:"data"`
}

type ProductCrossSellingAssignedProductsCollection struct {
	entityCollection

	Data []model.ProductCrossSellingAssignedProducts `json:"data"`
}

type ProductCrossSellingTranslationCollection struct {
	entityCollection

	Data []model.ProductCrossSellingTranslation `json:"data"`
}

type ProductCrossSellingCollection struct {
	entityCollection

	Data []model.ProductCrossSelling `json:"data"`
}

type ProductCustomFieldSetCollection struct {
	entityCollection

	Data []model.ProductCustomFieldSet `json:"data"`
}

type ProductExportCollection struct {
	entityCollection

	Data []model.ProductExport `json:"data"`
}

type ProductFeatureSetTranslationCollection struct {
	entityCollection

	Data []model.ProductFeatureSetTranslation `json:"data"`
}

type ProductFeatureSetCollection struct {
	entityCollection

	Data []model.ProductFeatureSet `json:"data"`
}

type ProductKeywordDictionaryCollection struct {
	entityCollection

	Data []model.ProductKeywordDictionary `json:"data"`
}

type ProductManufacturerTranslationCollection struct {
	entityCollection

	Data []model.ProductManufacturerTranslation `json:"data"`
}

type ProductManufacturerCollection struct {
	entityCollection

	Data []model.ProductManufacturer `json:"data"`
}

type ProductMediaCollection struct {
	entityCollection

	Data []model.ProductMedia `json:"data"`
}

type ProductOptionCollection struct {
	entityCollection

	Data []model.ProductOption `json:"data"`
}

type ProductPriceCollection struct {
	entityCollection

	Data []model.ProductPrice `json:"data"`
}

type ProductPropertyCollection struct {
	entityCollection

	Data []model.ProductProperty `json:"data"`
}

type ProductReviewCollection struct {
	entityCollection

	Data []model.ProductReview `json:"data"`
}

type ProductSearchConfigFieldCollection struct {
	entityCollection

	Data []model.ProductSearchConfigField `json:"data"`
}

type ProductSearchConfigCollection struct {
	entityCollection

	Data []model.ProductSearchConfig `json:"data"`
}

type ProductSearchKeywordCollection struct {
	entityCollection

	Data []model.ProductSearchKeyword `json:"data"`
}

type ProductSortingTranslationCollection struct {
	entityCollection

	Data []model.ProductSortingTranslation `json:"data"`
}

type ProductSortingCollection struct {
	entityCollection

	Data []model.ProductSorting `json:"data"`
}

type ProductStreamFilterCollection struct {
	entityCollection

	Data []model.ProductStreamFilter `json:"data"`
}

type ProductStreamMappingCollection struct {
	entityCollection

	Data []model.ProductStreamMapping `json:"data"`
}

type ProductStreamTranslationCollection struct {
	entityCollection

	Data []model.ProductStreamTranslation `json:"data"`
}

type ProductStreamCollection struct {
	entityCollection

	Data []model.ProductStream `json:"data"`
}

type ProductTagCollection struct {
	entityCollection

	Data []model.ProductTag `json:"data"`
}

type ProductTranslationCollection struct {
	entityCollection

	Data []model.ProductTranslation `json:"data"`
}

type ProductVisibilityCollection struct {
	entityCollection

	Data []model.ProductVisibility `json:"data"`
}

type ProductCollection struct {
	entityCollection

	Data []model.Product `json:"data"`
}

type PromotionCartRuleCollection struct {
	entityCollection

	Data []model.PromotionCartRule `json:"data"`
}

type PromotionDiscountPricesCollection struct {
	entityCollection

	Data []model.PromotionDiscountPrices `json:"data"`
}

type PromotionDiscountRuleCollection struct {
	entityCollection

	Data []model.PromotionDiscountRule `json:"data"`
}

type PromotionDiscountCollection struct {
	entityCollection

	Data []model.PromotionDiscount `json:"data"`
}

type PromotionIndividualCodeCollection struct {
	entityCollection

	Data []model.PromotionIndividualCode `json:"data"`
}

type PromotionOrderRuleCollection struct {
	entityCollection

	Data []model.PromotionOrderRule `json:"data"`
}

type PromotionPersonaCustomerCollection struct {
	entityCollection

	Data []model.PromotionPersonaCustomer `json:"data"`
}

type PromotionPersonaRuleCollection struct {
	entityCollection

	Data []model.PromotionPersonaRule `json:"data"`
}

type PromotionSalesChannelCollection struct {
	entityCollection

	Data []model.PromotionSalesChannel `json:"data"`
}

type PromotionSetgroupRuleCollection struct {
	entityCollection

	Data []model.PromotionSetgroupRule `json:"data"`
}

type PromotionSetgroupCollection struct {
	entityCollection

	Data []model.PromotionSetgroup `json:"data"`
}

type PromotionTranslationCollection struct {
	entityCollection

	Data []model.PromotionTranslation `json:"data"`
}

type PromotionCollection struct {
	entityCollection

	Data []model.Promotion `json:"data"`
}

type PropertyGroupOptionTranslationCollection struct {
	entityCollection

	Data []model.PropertyGroupOptionTranslation `json:"data"`
}

type PropertyGroupOptionCollection struct {
	entityCollection

	Data []model.PropertyGroupOption `json:"data"`
}

type PropertyGroupTranslationCollection struct {
	entityCollection

	Data []model.PropertyGroupTranslation `json:"data"`
}

type PropertyGroupCollection struct {
	entityCollection

	Data []model.PropertyGroup `json:"data"`
}

type RuleConditionCollection struct {
	entityCollection

	Data []model.RuleCondition `json:"data"`
}

type RuleCollection struct {
	entityCollection

	Data []model.Rule `json:"data"`
}

type SalesChannelAnalyticsCollection struct {
	entityCollection

	Data []model.SalesChannelAnalytics `json:"data"`
}

type SalesChannelCountryCollection struct {
	entityCollection

	Data []model.SalesChannelCountry `json:"data"`
}

type SalesChannelCurrencyCollection struct {
	entityCollection

	Data []model.SalesChannelCurrency `json:"data"`
}

type SalesChannelDomainCollection struct {
	entityCollection

	Data []model.SalesChannelDomain `json:"data"`
}

type SalesChannelLanguageCollection struct {
	entityCollection

	Data []model.SalesChannelLanguage `json:"data"`
}

type SalesChannelPaymentMethodCollection struct {
	entityCollection

	Data []model.SalesChannelPaymentMethod `json:"data"`
}

type SalesChannelShippingMethodCollection struct {
	entityCollection

	Data []model.SalesChannelShippingMethod `json:"data"`
}

type SalesChannelTranslationCollection struct {
	entityCollection

	Data []model.SalesChannelTranslation `json:"data"`
}

type SalesChannelTypeTranslationCollection struct {
	entityCollection

	Data []model.SalesChannelTypeTranslation `json:"data"`
}

type SalesChannelTypeCollection struct {
	entityCollection

	Data []model.SalesChannelType `json:"data"`
}

type SalesChannelCollection struct {
	entityCollection

	Data []model.SalesChannel `json:"data"`
}

type SalutationTranslationCollection struct {
	entityCollection

	Data []model.SalutationTranslation `json:"data"`
}

type SalutationCollection struct {
	entityCollection

	Data []model.Salutation `json:"data"`
}

type ScheduledTaskCollection struct {
	entityCollection

	Data []model.ScheduledTask `json:"data"`
}

type ScriptCollection struct {
	entityCollection

	Data []model.Script `json:"data"`
}

type SeoUrlTemplateCollection struct {
	entityCollection

	Data []model.SeoUrlTemplate `json:"data"`
}

type SeoUrlCollection struct {
	entityCollection

	Data []model.SeoUrl `json:"data"`
}

type ShippingMethodPriceCollection struct {
	entityCollection

	Data []model.ShippingMethodPrice `json:"data"`
}

type ShippingMethodTagCollection struct {
	entityCollection

	Data []model.ShippingMethodTag `json:"data"`
}

type ShippingMethodTranslationCollection struct {
	entityCollection

	Data []model.ShippingMethodTranslation `json:"data"`
}

type ShippingMethodCollection struct {
	entityCollection

	Data []model.ShippingMethod `json:"data"`
}

type SnippetSetCollection struct {
	entityCollection

	Data []model.SnippetSet `json:"data"`
}

type SnippetCollection struct {
	entityCollection

	Data []model.Snippet `json:"data"`
}

type StateMachineHistoryCollection struct {
	entityCollection

	Data []model.StateMachineHistory `json:"data"`
}

type StateMachineStateTranslationCollection struct {
	entityCollection

	Data []model.StateMachineStateTranslation `json:"data"`
}

type StateMachineStateCollection struct {
	entityCollection

	Data []model.StateMachineState `json:"data"`
}

type StateMachineTransitionCollection struct {
	entityCollection

	Data []model.StateMachineTransition `json:"data"`
}

type StateMachineTranslationCollection struct {
	entityCollection

	Data []model.StateMachineTranslation `json:"data"`
}

type StateMachineCollection struct {
	entityCollection

	Data []model.StateMachine `json:"data"`
}

type SystemConfigCollection struct {
	entityCollection

	Data []model.SystemConfig `json:"data"`
}

type TagCollection struct {
	entityCollection

	Data []model.Tag `json:"data"`
}

type TaxRuleTypeTranslationCollection struct {
	entityCollection

	Data []model.TaxRuleTypeTranslation `json:"data"`
}

type TaxRuleTypeCollection struct {
	entityCollection

	Data []model.TaxRuleType `json:"data"`
}

type TaxRuleCollection struct {
	entityCollection

	Data []model.TaxRule `json:"data"`
}

type TaxCollection struct {
	entityCollection

	Data []model.Tax `json:"data"`
}

type ThemeMediaCollection struct {
	entityCollection

	Data []model.ThemeMedia `json:"data"`
}

type ThemeSalesChannelCollection struct {
	entityCollection

	Data []model.ThemeSalesChannel `json:"data"`
}

type ThemeTranslationCollection struct {
	entityCollection

	Data []model.ThemeTranslation `json:"data"`
}

type ThemeCollection struct {
	entityCollection

	Data []model.Theme `json:"data"`
}

type UnitTranslationCollection struct {
	entityCollection

	Data []model.UnitTranslation `json:"data"`
}

type UnitCollection struct {
	entityCollection

	Data []model.Unit `json:"data"`
}

type UserAccessKeyCollection struct {
	entityCollection

	Data []model.UserAccessKey `json:"data"`
}

type UserConfigCollection struct {
	entityCollection

	Data []model.UserConfig `json:"data"`
}

type UserRecoveryCollection struct {
	entityCollection

	Data []model.UserRecovery `json:"data"`
}

type UserCollection struct {
	entityCollection

	Data []model.User `json:"data"`
}

type VersionCommitDataCollection struct {
	entityCollection

	Data []model.VersionCommitData `json:"data"`
}

type VersionCommitCollection struct {
	entityCollection

	Data []model.VersionCommit `json:"data"`
}

type VersionCollection struct {
	entityCollection

	Data []model.Version `json:"data"`
}

type WebhookEventLogCollection struct {
	entityCollection

	Data []model.WebhookEventLog `json:"data"`
}

type WebhookCollection struct {
	entityCollection

	Data []model.Webhook `json:"data"`
}
