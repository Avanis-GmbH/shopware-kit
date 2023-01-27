# Shopware Universal Client Kit

[![License](https://img.shields.io/badge/License-BSD_3--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)
[![golangci-lint](https://github.com/Avanis-GmbH/GoSUCK/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/Avanis-GmbH/GoSUCK/actions/workflows/golangci-lint.yml)
[![CodeQL](https://github.com/Avanis-GmbH/GoSUCK/actions/workflows/codeql.yml/badge.svg)](https://github.com/Avanis-GmbH/GoSUCK/actions/workflows/codeql.yml)

GoSuck is a universal API client for Shopware 6. It is based on the [Shopware 6 API](https://developer.shopware.com/docs/guides/api-guide/).
It can be used as a standalone client or directly with the included models, which are derived from the Shopware 6 API documentation.

## ⚙️ How it works

The package is divided into two parts. The first part is the client, which is responsible for the communication with the Shopware 6 API. The second part is the models, which are used to represent the data from the Shopware 6 API.

### 📡 Communication

The package defines a central client, which is responsible for the authorization and communication with the Shopware 6 API. And brings methods to execute a request to the Shopware 6 API.

### 📝 Models

The models are used to represent the data from the Shopware 6 API.
Each model is saved in a separate file and contains the type definition of one data entity.

For example, the `Order` model is saved in the `order.go` file and contains the type definition of the data entity.

```go
type Order struct {
	Addresses               []OrderAddress     `json:"addresses,omitempty"`
	AffiliateCode           string             `json:"affiliateCode,omitempty"`
	AmountNet               float64            `json:"amountNet,omitempty"`
	AmountTotal             float64            `json:"amountTotal,omitempty"`
	AutoIncrement           float64            `json:"autoIncrement,omitempty"`
	BillingAddress          *OrderAddress      `json:"billingAddress,omitempty"`
	BillingAddressId        string             `json:"billingAddressId"` // required
	BillingAddressVersionId string             `json:"billingAddressVersionId,omitempty"`
	CampaignCode            string             `json:"campaignCode,omitempty"`
	CreatedAt               time.Time          `json:"createdAt,omitempty"`
	CreatedBy               *User              `json:"createdBy,omitempty"`
    ...
}

```
> All attributes are marked with the json tag `omitempty`, this is done to reduce the size of the request body. If an attribute is required or a numeric type is used, the `omitempty` tag is removed.

### 💥 Collections

This package also contains a collection for each model. The collection is used to group a list of entities and represents the structure of an API response. Each collection embeds the `EntityCollection` struct and thus implements the `Collection` interface.

For example, the `OrderCollection` is used to group a list of `Order` entities.

```go
type OrderCollection struct {
    EntityCollection

    Data []Order `json:"data"`
}
```

## 📦 Examples

### Client

How to initialize a client:

```go
func main() {
	creds := com.NewIntegrationCredentials("<CLIENT-ID>", "<API-KEY>", []string{"write"})
	client, err := com.NewClient(context.Background(), "<URL>", creds, nil)
	if err != nil {
		panic(err)
	}
}
```

> It is possible to use the client without the models, but it is recommended to use the models, because they are directly derived from the Shopware 6 API documentation.

How to execute a basic request:

```go
	// Create a new request with the given context, method, url and criteria
	ctx := com.NewApiContext(context.Background())
	req, err := client.NewRequest(ctx, "GET", "/api/system-config", com.Criteria{})
	if err != nil {
		panic(err)
	}

	// Execute the request and decode the response into the given struct (v)
	v := &com.SystemConfigCollection{}
	_, err = client.Do(ctx.Context, req, v)
	if err != nil {
		panic(err)
	}
```

#### Search

How to execute a search request:

```go
    categoryCollection := &com.CategoryCollection{}
    err := client.Search(com.NewApiContext(context.Background()), com.Criteria{}, categoryCollection)
    if err != nil {
        panic(err)
    }

    fmt.Println(categoryCollection)
```

#### Update and Insert

Update and Insert actions are handled by the same method `Upsert`. The shopware API provides the `/api/_action/sync` endpoint, which can be used to update and insert data.

```go
	category := model.Category{
		Id:        Md5(fmt.Sprint("Category Name")), // Md5 is a helper function to generate a unique id as md5 hash
		Name:      "Category Name",
		ParentId:  "<PARENT-ID>",
		CmsPageId: "<CMS-PAGE-ID>",
	}

    // Upsert expects the given data to be a slice of the given model
	_, err = client.Upsert(com.NewApiContext(ctx), []model.Category{category})
	if err != nil {
		panic(err)
	}
```

#### Delete

How to execute a delete request:

```go
    // The second parameter is the model, which should be deleted and the third parameter is a slice of ids of the given model
	_, err = client.Delete(com.NewApiContext(ctx), model.Category{}, []string{Md5("Category Name")})
	if err != nil {
		panic(err)
	}
```
