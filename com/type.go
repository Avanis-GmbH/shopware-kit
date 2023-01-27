package com

import (
	"reflect"

	"github.com/Avanis-GmbH/GoSUCK/model"
)

// Returns the segment of the path that is used to identify the resource.
func (c *Client) getSegment(v interface{}) string {
	// If v is a pointer, get the value it points to.
	if reflect.TypeOf(v).Kind() == reflect.Ptr {
		v = reflect.ValueOf(v).Elem().Interface()
	}

	// If v is a slice, get the type of the first element.
	if reflect.TypeOf(v).Kind() == reflect.Slice {
		if reflect.ValueOf(v).Len() == 0 {
			return "unknown"
		}
		v = reflect.ValueOf(v).Index(0).Interface()
	}

	switch v.(type) {
	case model.AclRole, AclRoleCollection:
		return "acl-role"
	case model.AclUserRole, AclUserRoleCollection:
		return "acl-user-role"
	case model.AppActionButtonTranslation, AppActionButtonTranslationCollection:
		return "app-action-button-translation"
	case model.AppActionButton, AppActionButtonCollection:
		return "app-action-button"
	case model.AppCmsBlockTranslation, AppCmsBlockTranslationCollection:
		return "app-cms-block-translation"
	case model.AppCmsBlock, AppCmsBlockCollection:
		return "app-cms-block"
	case model.AppPaymentMethod, AppPaymentMethodCollection:
		return "app-payment-method"
	case model.AppTemplate, AppTemplateCollection:
		return "app-template"
	case model.AppTranslation, AppTranslationCollection:
		return "app-translation"
	case model.App, AppCollection:
		return "app"
	case model.CategoryTag, CategoryTagCollection:
		return "category-tag"
	case model.CategoryTranslation, CategoryTranslationCollection:
		return "category-translation"
	case model.Category, CategoryCollection:
		return "category"
	case model.CmsBlock, CmsBlockCollection:
		return "cms-block"
	case model.CmsPageTranslation, CmsPageTranslationCollection:
		return "cms-page-translation"
	case model.CmsPage, CmsPageCollection:
		return "cms-page"
	case model.CmsSection, CmsSectionCollection:
		return "cms-section"
	case model.CmsSlotTranslation, CmsSlotTranslationCollection:
		return "cms-slot-translation"
	case model.CmsSlot, CmsSlotCollection:
		return "cms-slot"
	case model.CountryStateTranslation, CountryStateTranslationCollection:
		return "country-state-translation"
	case model.CountryState, CountryStateCollection:
		return "country-state"
	case model.CountryTranslation, CountryTranslationCollection:
		return "country-translation"
	case model.Country, CountryCollection:
		return "country"
	case model.CurrencyCountryRounding, CurrencyCountryRoundingCollection:
		return "currency-country-rounding"
	case model.CurrencyTranslation, CurrencyTranslationCollection:
		return "currency-translation"
	case model.Currency, CurrencyCollection:
		return "currency"
	case model.CustomFieldSetRelation, CustomFieldSetRelationCollection:
		return "custom-field-set-relation"
	case model.CustomFieldSet, CustomFieldSetCollection:
		return "custom-field-set"
	case model.CustomField, CustomFieldCollection:
		return "custom-field"
	case model.CustomerAddress, CustomerAddressCollection:
		return "customer-address"
	case model.CustomerGroupRegistrationSalesChannels, CustomerGroupRegistrationSalesChannelsCollection:
		return "customer-group-registration-sales-channels"
	case model.CustomerGroupTranslation, CustomerGroupTranslationCollection:
		return "customer-group-translation"
	case model.CustomerGroup, CustomerGroupCollection:
		return "customer-group"
	case model.CustomerRecovery, CustomerRecoveryCollection:
		return "customer-recovery"
	case model.CustomerTag, CustomerTagCollection:
		return "customer-tag"
	case model.CustomerWishlistProduct, CustomerWishlistProductCollection:
		return "customer-wishlist-product"
	case model.CustomerWishlist, CustomerWishlistCollection:
		return "customer-wishlist"
	case model.Customer, CustomerCollection:
		return "customer"
	case model.DeadMessage, DeadMessageCollection:
		return "dead-message"
	case model.DeliveryTimeTranslation, DeliveryTimeTranslationCollection:
		return "delivery-time-translation"
	case model.DeliveryTime, DeliveryTimeCollection:
		return "delivery-time"
	case model.DocumentBaseConfigSalesChannel, DocumentBaseConfigSalesChannelCollection:
		return "document-base-config-sales-channel"
	case model.DocumentBaseConfig, DocumentBaseConfigCollection:
		return "document-base-config"
	case model.DocumentTypeTranslation, DocumentTypeTranslationCollection:
		return "document-type-translation"
	case model.DocumentType, DocumentTypeCollection:
		return "document-type"
	case model.Document, DocumentCollection:
		return "document"
	case model.EventActionRule, EventActionRuleCollection:
		return "event-action-rule"
	case model.EventActionSalesChannel, EventActionSalesChannelCollection:
		return "event-action-sales-channel"
	case model.EventAction, EventActionCollection:
		return "event-action"
	case model.FlowSequence, FlowSequenceCollection:
		return "flow-sequence"
	case model.Flow, FlowCollection:
		return "flow"
	case model.ImportExportFile, ImportExportFileCollection:
		return "import-export-file"
	case model.ImportExportLog, ImportExportLogCollection:
		return "import-export-log"
	case model.ImportExportProfileTranslation, ImportExportProfileTranslationCollection:
		return "import-export-profile-translation"
	case model.ImportExportProfile, ImportExportProfileCollection:
		return "import-export-profile"
	case model.IntegrationRole, IntegrationRoleCollection:
		return "integration-role"
	case model.Integration, IntegrationCollection:
		return "integration"
	case model.LandingPageSalesChannel, LandingPageSalesChannelCollection:
		return "landing-page-sales-channel"
	case model.LandingPage, LandingPageCollection:
		return "landing-page"
	case model.Language, LanguageCollection:
		return "language"
	case model.LocaleTranslation, LocaleTranslationCollection:
		return "locale-translation"
	case model.Locale, LocaleCollection:
		return "locale"
	case model.LogEntry, LogEntryCollection:
		return "log-entry"
	case model.MailHeaderFooterTranslation, MailHeaderFooterTranslationCollection:
		return "mail-header-footer-translation"
	case model.MailHeaderFooter, MailHeaderFooterCollection:
		return "mail-header-footer"
	case model.MailTemplateMedia, MailTemplateMediaCollection:
		return "mail-template-media"
	case model.MailTemplateTranslation, MailTemplateTranslationCollection:
		return "mail-template-translation"
	case model.MailTemplateTypeTranslation, MailTemplateTypeTranslationCollection:
		return "mail-template-type-translation"
	case model.MailTemplateType, MailTemplateTypeCollection:
		return "mail-template-type"
	case model.MailTemplate, MailTemplateCollection:
		return "mail-template"
	case model.MainCategory, MainCategoryCollection:
		return "main-category"
	case model.MediaDefaultFolder, MediaDefaultFolderCollection:
		return "media-default-folder"
	case model.MediaFolderConfigurationMediaThumbnailSize, MediaFolderConfigurationMediaThumbnailSizeCollection:
		return "media-folder-configuration-media-thumbnail-size"
	case model.MediaFolderConfiguration, MediaFolderConfigurationCollection:
		return "media-folder-configuration"
	case model.MediaFolder, MediaFolderCollection:
		return "media-folder"
	case model.MediaTag, MediaTagCollection:
		return "media-tag"
	case model.MediaThumbnailSize, MediaThumbnailSizeCollection:
		return "media-thumbnail-size"
	case model.MediaThumbnail, MediaThumbnailCollection:
		return "media-thumbnail"
	case model.MediaTranslation, MediaTranslationCollection:
		return "media-translation"
	case model.Media, MediaCollection:
		return "media"
	case model.MessageQueueStats, MessageQueueStatsCollection:
		return "message-queue-stats"
	case model.NewsletterRecipientTag, NewsletterRecipientTagCollection:
		return "newsletter-recipient-tag"
	case model.NewsletterRecipient, NewsletterRecipientCollection:
		return "newsletter-recipient"
	case model.Notification, NotificationCollection:
		return "notification"
	case model.NumberRangeSalesChannel, NumberRangeSalesChannelCollection:
		return "number-range-sales-channel"
	case model.NumberRangeState, NumberRangeStateCollection:
		return "number-range-state"
	case model.NumberRangeTypeTranslation, NumberRangeTypeTranslationCollection:
		return "number-range-type-translation"
	case model.NumberRangeType, NumberRangeTypeCollection:
		return "number-range-type"
	case model.NumberRange, NumberRangeCollection:
		return "number-range"
	case model.OrderAddress, OrderAddressCollection:
		return "order-address"
	case model.OrderCustomer, OrderCustomerCollection:
		return "order-customer"
	case model.OrderDeliveryPosition, OrderDeliveryPositionCollection:
		return "order-delivery-position"
	case model.OrderDelivery, OrderDeliveryCollection:
		return "order-delivery"
	case model.OrderLineItem, OrderLineItemCollection:
		return "order-line-item"
	case model.OrderTag, OrderTagCollection:
		return "order-tag"
	case model.OrderTransaction, OrderTransactionCollection:
		return "order-transaction"
	case model.Order, OrderCollection:
		return "order"
	case model.PaymentMethodTranslation, PaymentMethodTranslationCollection:
		return "payment-method-translation"
	case model.PaymentMethod, PaymentMethodCollection:
		return "payment-method"
	case model.PluginTranslation, PluginTranslationCollection:
		return "plugin-translation"
	case model.Plugin, PluginCollection:
		return "plugin"
	case model.ProductCategoryTree, ProductCategoryTreeCollection:
		return "product-category-tree"
	case model.ProductCategory, ProductCategoryCollection:
		return "product-category"
	case model.ProductConfiguratorSetting, ProductConfiguratorSettingCollection:
		return "product-configurator-setting"
	case model.ProductCrossSellingAssignedProducts, ProductCrossSellingAssignedProductsCollection:
		return "product-cross-selling-assigned-products"
	case model.ProductCrossSellingTranslation, ProductCrossSellingTranslationCollection:
		return "product-cross-selling-translation"
	case model.ProductCrossSelling, ProductCrossSellingCollection:
		return "product-cross-selling"
	case model.ProductCustomFieldSet, ProductCustomFieldSetCollection:
		return "product-custom-field-set"
	case model.ProductExport, ProductExportCollection:
		return "product-export"
	case model.ProductFeatureSetTranslation, ProductFeatureSetTranslationCollection:
		return "product-feature-set-translation"
	case model.ProductFeatureSet, ProductFeatureSetCollection:
		return "product-feature-set"
	case model.ProductKeywordDictionary, ProductKeywordDictionaryCollection:
		return "product-keyword-dictionary"
	case model.ProductManufacturerTranslation, ProductManufacturerTranslationCollection:
		return "product-manufacturer-translation"
	case model.ProductManufacturer, ProductManufacturerCollection:
		return "product-manufacturer"
	case model.ProductMedia, ProductMediaCollection:
		return "product-media"
	case model.ProductOption, ProductOptionCollection:
		return "product-option"
	case model.ProductPrice, ProductPriceCollection:
		return "product-price"
	case model.ProductProperty, ProductPropertyCollection:
		return "product-property"
	case model.ProductReview, ProductReviewCollection:
		return "product-review"
	case model.ProductSearchConfigField, ProductSearchConfigFieldCollection:
		return "product-search-config-field"
	case model.ProductSearchConfig, ProductSearchConfigCollection:
		return "product-search-config"
	case model.ProductSearchKeyword, ProductSearchKeywordCollection:
		return "product-search-keyword"
	case model.ProductSortingTranslation, ProductSortingTranslationCollection:
		return "product-sorting-translation"
	case model.ProductSorting, ProductSortingCollection:
		return "product-sorting"
	case model.ProductStreamFilter, ProductStreamFilterCollection:
		return "product-stream-filter"
	case model.ProductStreamMapping, ProductStreamMappingCollection:
		return "product-stream-mapping"
	case model.ProductStreamTranslation, ProductStreamTranslationCollection:
		return "product-stream-translation"
	case model.ProductStream, ProductStreamCollection:
		return "product-stream"
	case model.ProductTag, ProductTagCollection:
		return "product-tag"
	case model.ProductTranslation, ProductTranslationCollection:
		return "product-translation"
	case model.ProductVisibility, ProductVisibilityCollection:
		return "product-visibility"
	case model.Product, ProductCollection:
		return "product"
	case model.PromotionCartRule, PromotionCartRuleCollection:
		return "promotion-cart-rule"
	case model.PromotionDiscountPrices, PromotionDiscountPricesCollection:
		return "promotion-discount-prices"
	case model.PromotionDiscountRule, PromotionDiscountRuleCollection:
		return "promotion-discount-rule"
	case model.PromotionDiscount, PromotionDiscountCollection:
		return "promotion-discount"
	case model.PromotionIndividualCode, PromotionIndividualCodeCollection:
		return "promotion-individual-code"
	case model.PromotionOrderRule, PromotionOrderRuleCollection:
		return "promotion-order-rule"
	case model.PromotionPersonaCustomer, PromotionPersonaCustomerCollection:
		return "promotion-persona-customer"
	case model.PromotionPersonaRule, PromotionPersonaRuleCollection:
		return "promotion-persona-rule"
	case model.PromotionSalesChannel, PromotionSalesChannelCollection:
		return "promotion-sales-channel"
	case model.PromotionSetgroupRule, PromotionSetgroupRuleCollection:
		return "promotion-setgroup-rule"
	case model.PromotionSetgroup, PromotionSetgroupCollection:
		return "promotion-setgroup"
	case model.PromotionTranslation, PromotionTranslationCollection:
		return "promotion-translation"
	case model.Promotion, PromotionCollection:
		return "promotion"
	case model.PropertyGroupOptionTranslation, PropertyGroupOptionTranslationCollection:
		return "property-group-option-translation"
	case model.PropertyGroupOption, PropertyGroupOptionCollection:
		return "property-group-option"
	case model.PropertyGroupTranslation, PropertyGroupTranslationCollection:
		return "property-group-translation"
	case model.PropertyGroup, PropertyGroupCollection:
		return "property-group"
	case model.RuleCondition, RuleConditionCollection:
		return "rule-condition"
	case model.Rule, RuleCollection:
		return "rule"
	case model.SalesChannelAnalytics, SalesChannelAnalyticsCollection:
		return "sales-channel-analytics"
	case model.SalesChannelCountry, SalesChannelCountryCollection:
		return "sales-channel-country"
	case model.SalesChannelCurrency, SalesChannelCurrencyCollection:
		return "sales-channel-currency"
	case model.SalesChannelDomain, SalesChannelDomainCollection:
		return "sales-channel-domain"
	case model.SalesChannelLanguage, SalesChannelLanguageCollection:
		return "sales-channel-language"
	case model.SalesChannelPaymentMethod, SalesChannelPaymentMethodCollection:
		return "sales-channel-payment-method"
	case model.SalesChannelShippingMethod, SalesChannelShippingMethodCollection:
		return "sales-channel-shipping-method"
	case model.SalesChannelTypeTranslation, SalesChannelTypeTranslationCollection:
		return "sales-channel-type-translation"
	case model.SalesChannelType, SalesChannelTypeCollection:
		return "sales-channel-type"
	case model.SalesChannelTranslation, SalesChannelTranslationCollection:
		return "sales-channel-translation"
	case model.SalesChannel, SalesChannelCollection:
		return "sales-channel"
	case model.SalutationTranslation, SalutationTranslationCollection:
		return "salutation-translation"
	case model.Salutation, SalutationCollection:
		return "salutation"
	case model.ScheduledTask, ScheduledTaskCollection:
		return "scheduled-task"
	case model.Script, ScriptCollection:
		return "script"
	case model.SeoUrlTemplate, SeoUrlTemplateCollection:
		return "seo-url-template"
	case model.SeoUrl, SeoUrlCollection:
		return "seo-url"
	case model.ShippingMethodPrice, ShippingMethodPriceCollection:
		return "shipping-method-price"
	case model.ShippingMethodTag, ShippingMethodTagCollection:
		return "shipping-method-tag"
	case model.ShippingMethodTranslation, ShippingMethodTranslationCollection:
		return "shipping-method-translation"
	case model.ShippingMethod, ShippingMethodCollection:
		return "shipping-method"
	case model.SnippetSet, SnippetSetCollection:
		return "snippet-set"
	case model.Snippet, SnippetCollection:
		return "snippet"
	case model.StateMachineHistory, StateMachineHistoryCollection:
		return "state-machine-history"
	case model.StateMachineStateTranslation, StateMachineStateTranslationCollection:
		return "state-machine-state-translation"
	case model.StateMachineState, StateMachineStateCollection:
		return "state-machine-state"
	case model.StateMachineTransition, StateMachineTransitionCollection:
		return "state-machine-transition"
	case model.StateMachineTranslation, StateMachineTranslationCollection:
		return "state-machine-translation"
	case model.StateMachine, StateMachineCollection:
		return "state-machine"
	case model.SystemConfig, SystemConfigCollection:
		return "system-config"
	case model.Tag, TagCollection:
		return "tag"
	case model.TaxRuleTypeTranslation, TaxRuleTypeTranslationCollection:
		return "tax-rule-type-translation"
	case model.TaxRuleType, TaxRuleTypeCollection:
		return "tax-rule-type"
	case model.TaxRule, TaxRuleCollection:
		return "tax-rule"
	case model.Tax, TaxCollection:
		return "tax"
	case model.ThemeMedia, ThemeMediaCollection:
		return "theme-media"
	case model.ThemeSalesChannel, ThemeSalesChannelCollection:
		return "theme-sales-channel"
	case model.ThemeTranslation, ThemeTranslationCollection:
		return "theme-translation"
	case model.Theme, ThemeCollection:
		return "theme"
	case model.UnitTranslation, UnitTranslationCollection:
		return "unit-translation"
	case model.Unit, UnitCollection:
		return "unit"
	case model.UserAccessKey, UserAccessKeyCollection:
		return "user-access-key"
	case model.UserConfig, UserConfigCollection:
		return "user-config"
	case model.UserRecovery, UserRecoveryCollection:
		return "user-recovery"
	case model.User, UserCollection:
		return "user"
	case model.VersionCommitData, VersionCommitDataCollection:
		return "version-commit-data"
	case model.Version, VersionCollection:
		return "version"
	case model.WebhookEventLog, WebhookEventLogCollection:
		return "webhook-event-log"
	case model.Webhook, WebhookCollection:
		return "webhook"
	}

	return "unknown"
}
