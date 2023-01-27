package model

import "time"

type Product struct {
	Active                        bool                                  `json:"active,omitempty"`
	AutoIncrement                 float64                               `json:"autoIncrement"`
	Available                     bool                                  `json:"available,omitempty"`
	AvailableStock                float64                               `json:"availableStock"`
	CanonicalProduct              *Product                              `json:"canonicalProduct,omitempty"`
	CanonicalProductId            string                                `json:"canonicalProductId,omitempty"`
	Categories                    []Category                            `json:"categories,omitempty"`
	CategoriesRo                  []Category                            `json:"categoriesRo,omitempty"`
	CategoryIds                   interface{}                           `json:"categoryIds,omitempty"`
	CategoryTree                  interface{}                           `json:"categoryTree,omitempty"`
	CheapestPrice                 interface{}                           `json:"cheapestPrice,omitempty"`
	ChildCount                    float64                               `json:"childCount"`
	Children                      []Product                             `json:"children,omitempty"`
	CmsPage                       *CmsPage                              `json:"cmsPage,omitempty"`
	CmsPageId                     string                                `json:"cmsPageId,omitempty"`
	CmsPageVersionId              string                                `json:"cmsPageVersionId,omitempty"`
	ConfiguratorGroupConfig       interface{}                           `json:"configuratorGroupConfig,omitempty"` // deprecated
	ConfiguratorSettings          []ProductConfiguratorSetting          `json:"configuratorSettings,omitempty"`
	Cover                         *ProductMedia                         `json:"cover,omitempty"`
	CoverId                       string                                `json:"coverId,omitempty"`
	CreatedAt                     time.Time                             `json:"createdAt,omitempty"`
	CrossSellingAssignedProducts  []ProductCrossSellingAssignedProducts `json:"crossSellingAssignedProducts,omitempty"`
	CrossSellings                 []ProductCrossSelling                 `json:"crossSellings,omitempty"`
	CustomFields                  interface{}                           `json:"customFields,omitempty"`
	CustomFieldSets               []CustomFieldSet                      `json:"customFieldSets,omitempty"`
	CustomFieldSetSelectionActive bool                                  `json:"customFieldSetSelectionActive,omitempty"`
	CustomSearchKeywords          interface{}                           `json:"customSearchKeywords,omitempty"`
	DeliveryTime                  *DeliveryTime                         `json:"deliveryTime,omitempty"`
	DeliveryTimeId                string                                `json:"deliveryTimeId,omitempty"`
	Description                   string                                `json:"description,omitempty"`
	DisplayGroup                  string                                `json:"displayGroup,omitempty"`
	Ean                           string                                `json:"ean,omitempty"`
	FeatureSet                    *ProductFeatureSet                    `json:"featureSet,omitempty"`
	FeatureSetId                  string                                `json:"featureSetId,omitempty"`
	Height                        float64                               `json:"height"`
	Id                            string                                `json:"id,omitempty"`
	IsCloseout                    bool                                  `json:"isCloseout,omitempty"`
	Keywords                      string                                `json:"keywords,omitempty"`
	Length                        float64                               `json:"length"`
	MainCategories                []MainCategory                        `json:"mainCategories,omitempty"`
	MainVariantId                 string                                `json:"mainVariantId,omitempty"` // deprecated
	Manufacturer                  *ProductManufacturer                  `json:"manufacturer,omitempty"`
	ManufacturerId                string                                `json:"manufacturerId,omitempty"`
	ManufacturerNumber            string                                `json:"manufacturerNumber,omitempty"`
	MarkAsTopseller               bool                                  `json:"markAsTopseller,omitempty"`
	MaxPurchase                   float64                               `json:"maxPurchase"`
	Media                         []ProductMedia                        `json:"media,omitempty"`
	MetaDescription               string                                `json:"metaDescription,omitempty"`
	MetaTitle                     string                                `json:"metaTitle,omitempty"`
	MinPurchase                   float64                               `json:"minPurchase"`
	Name                          string                                `json:"name"` // required
	OptionIds                     interface{}                           `json:"optionIds,omitempty"`
	Options                       []PropertyGroupOption                 `json:"options,omitempty"`
	OrderLineItems                []OrderLineItem                       `json:"orderLineItems,omitempty"`
	PackUnit                      string                                `json:"packUnit,omitempty"`
	PackUnitPlural                string                                `json:"packUnitPlural,omitempty"`
	Parent                        *Product                              `json:"parent,omitempty"`
	ParentId                      string                                `json:"parentId,omitempty"`
	ParentVersionId               string                                `json:"parentVersionId,omitempty"`
	Price                         interface{}                           `json:"price"` // required
	Prices                        []ProductPrice                        `json:"prices,omitempty"`
	ProductManufacturerVersionId  string                                `json:"productManufacturerVersionId,omitempty"`
	ProductMediaVersionId         string                                `json:"productMediaVersionId,omitempty"`
	ProductNumber                 string                                `json:"productNumber"` // required
	ProductReviews                []ProductReview                       `json:"productReviews,omitempty"`
	Properties                    []PropertyGroupOption                 `json:"properties,omitempty"`
	PropertyIds                   interface{}                           `json:"propertyIds,omitempty"`
	PurchasePrices                interface{}                           `json:"purchasePrices,omitempty"`
	PurchaseSteps                 float64                               `json:"purchaseSteps"`
	PurchaseUnit                  float64                               `json:"purchaseUnit"`
	RatingAverage                 float64                               `json:"ratingAverage"`
	ReferenceUnit                 float64                               `json:"referenceUnit"`
	ReleaseDate                   time.Time                             `json:"releaseDate,omitempty"`
	RestockTime                   float64                               `json:"restockTime"`
	Sales                         float64                               `json:"sales"`
	SearchKeywords                []ProductSearchKeyword                `json:"searchKeywords,omitempty"`
	SeoUrls                       []SeoUrl                              `json:"seoUrls,omitempty"`
	ShippingFree                  bool                                  `json:"shippingFree,omitempty"`
	SlotConfig                    interface{}                           `json:"slotConfig,omitempty"`
	Stock                         float64                               `json:"stock"` // required
	StreamIds                     interface{}                           `json:"streamIds,omitempty"`
	Streams                       []ProductStream                       `json:"streams,omitempty"`
	TagIds                        interface{}                           `json:"tagIds,omitempty"`
	Tags                          []Tag                                 `json:"tags,omitempty"`
	Tax                           *Tax                                  `json:"tax,omitempty"`
	TaxId                         string                                `json:"taxId"` // required
	Translated                    interface{}                           `json:"translated,omitempty"`
	Translations                  []ProductTranslation                  `json:"translations,omitempty"`
	Unit                          *Unit                                 `json:"unit,omitempty"`
	UnitId                        string                                `json:"unitId,omitempty"`
	UpdatedAt                     time.Time                             `json:"updatedAt,omitempty"`
	VariantRestrictions           interface{}                           `json:"variantRestrictions,omitempty"`
	Variation                     interface{}                           `json:"variation,omitempty"`
	VersionId                     string                                `json:"versionId,omitempty"`
	Visibilities                  []ProductVisibility                   `json:"visibilities,omitempty"`
	Weight                        float64                               `json:"weight"`
	Width                         float64                               `json:"width"`
	Wishlists                     []CustomerWishlistProduct             `json:"wishlists,omitempty"`
}
