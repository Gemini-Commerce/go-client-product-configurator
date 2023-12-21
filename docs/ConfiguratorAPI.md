# \ConfiguratorAPI

All URIs are relative to *https://product-configurator.api.gogemini.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProductConfiguratorCopyConfigurator**](ConfiguratorAPI.md#ProductConfiguratorCopyConfigurator) | **Post** /v1/{tenantId}/product/{sourceConfiguratorId}/copy | Copy Configurator
[**ProductConfiguratorCreateConfigurator**](ConfiguratorAPI.md#ProductConfiguratorCreateConfigurator) | **Post** /v1/{tenantId}/product/{productId}/create | Create Configurator
[**ProductConfiguratorDeleteConfigurator**](ConfiguratorAPI.md#ProductConfiguratorDeleteConfigurator) | **Delete** /v1/{tenantId}/configurator/{configuratorId} | Delete Configurator
[**ProductConfiguratorGetConfiguratorByProductId**](ConfiguratorAPI.md#ProductConfiguratorGetConfiguratorByProductId) | **Get** /v1/{tenantId}/product/{productId} | Get Product Configurator by Product ID
[**ProductConfiguratorGetConfiguratorByProductId2**](ConfiguratorAPI.md#ProductConfiguratorGetConfiguratorByProductId2) | **Get** /v1/{tenantId}/product/{productId}/status/{status} | Get Product Configurator by Product ID
[**ProductConfiguratorListConfigurators**](ConfiguratorAPI.md#ProductConfiguratorListConfigurators) | **Post** /v1/{tenantId}/product/{productId}/page-size/{pageSize}/configurators | List Product Configurators
[**ProductConfiguratorUpdateConfigurator**](ConfiguratorAPI.md#ProductConfiguratorUpdateConfigurator) | **Put** /v1/{tenantId}/configurator/{configuratorId} | Update Configurator



## ProductConfiguratorCopyConfigurator

> ProductconfiguratorconfiguratorEntity ProductConfiguratorCopyConfigurator(ctx, tenantId, sourceConfiguratorId).Body(body).Execute()

Copy Configurator



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/gemini-commerce/go-client-product_configurator"
)

func main() {
    tenantId := "tenantId_example" // string | 
    sourceConfiguratorId := "sourceConfiguratorId_example" // string | 
    body := *openapiclient.NewProductConfiguratorCopyConfiguratorRequest() // ProductConfiguratorCopyConfiguratorRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfiguratorAPI.ProductConfiguratorCopyConfigurator(context.Background(), tenantId, sourceConfiguratorId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfiguratorAPI.ProductConfiguratorCopyConfigurator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorCopyConfigurator`: ProductconfiguratorconfiguratorEntity
    fmt.Fprintf(os.Stdout, "Response from `ConfiguratorAPI.ProductConfiguratorCopyConfigurator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**sourceConfiguratorId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductConfiguratorCopyConfiguratorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**ProductConfiguratorCopyConfiguratorRequest**](ProductConfiguratorCopyConfiguratorRequest.md) |  | 

### Return type

[**ProductconfiguratorconfiguratorEntity**](ProductconfiguratorconfiguratorEntity.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorCreateConfigurator

> ProductconfiguratorconfiguratorCreateRequest ProductConfiguratorCreateConfigurator(ctx, tenantId, productId).Body(body).Execute()

Create Configurator



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/gemini-commerce/go-client-product_configurator"
)

func main() {
    tenantId := "tenantId_example" // string | 
    productId := "productId_example" // string | 
    body := *openapiclient.NewProductConfiguratorCreateConfiguratorRequest() // ProductConfiguratorCreateConfiguratorRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfiguratorAPI.ProductConfiguratorCreateConfigurator(context.Background(), tenantId, productId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfiguratorAPI.ProductConfiguratorCreateConfigurator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorCreateConfigurator`: ProductconfiguratorconfiguratorCreateRequest
    fmt.Fprintf(os.Stdout, "Response from `ConfiguratorAPI.ProductConfiguratorCreateConfigurator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductConfiguratorCreateConfiguratorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**ProductConfiguratorCreateConfiguratorRequest**](ProductConfiguratorCreateConfiguratorRequest.md) |  | 

### Return type

[**ProductconfiguratorconfiguratorCreateRequest**](ProductconfiguratorconfiguratorCreateRequest.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorDeleteConfigurator

> map[string]interface{} ProductConfiguratorDeleteConfigurator(ctx, tenantId, configuratorId).Execute()

Delete Configurator



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/gemini-commerce/go-client-product_configurator"
)

func main() {
    tenantId := "tenantId_example" // string | 
    configuratorId := "configuratorId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfiguratorAPI.ProductConfiguratorDeleteConfigurator(context.Background(), tenantId, configuratorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfiguratorAPI.ProductConfiguratorDeleteConfigurator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorDeleteConfigurator`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfiguratorAPI.ProductConfiguratorDeleteConfigurator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**configuratorId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductConfiguratorDeleteConfiguratorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorGetConfiguratorByProductId

> ProductconfiguratorconfiguratorEntity ProductConfiguratorGetConfiguratorByProductId(ctx, tenantId, productId).Status(status).Execute()

Get Product Configurator by Product ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/gemini-commerce/go-client-product_configurator"
)

func main() {
    tenantId := "tenantId_example" // string | 
    productId := "productId_example" // string | 
    status := "status_example" // string |  (optional) (default to "UNKNOWN")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfiguratorAPI.ProductConfiguratorGetConfiguratorByProductId(context.Background(), tenantId, productId).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfiguratorAPI.ProductConfiguratorGetConfiguratorByProductId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorGetConfiguratorByProductId`: ProductconfiguratorconfiguratorEntity
    fmt.Fprintf(os.Stdout, "Response from `ConfiguratorAPI.ProductConfiguratorGetConfiguratorByProductId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductConfiguratorGetConfiguratorByProductIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **status** | **string** |  | [default to &quot;UNKNOWN&quot;]

### Return type

[**ProductconfiguratorconfiguratorEntity**](ProductconfiguratorconfiguratorEntity.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorGetConfiguratorByProductId2

> ProductconfiguratorconfiguratorEntity ProductConfiguratorGetConfiguratorByProductId2(ctx, tenantId, productId, status).Execute()

Get Product Configurator by Product ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/gemini-commerce/go-client-product_configurator"
)

func main() {
    tenantId := "tenantId_example" // string | 
    productId := "productId_example" // string | 
    status := "status_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfiguratorAPI.ProductConfiguratorGetConfiguratorByProductId2(context.Background(), tenantId, productId, status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfiguratorAPI.ProductConfiguratorGetConfiguratorByProductId2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorGetConfiguratorByProductId2`: ProductconfiguratorconfiguratorEntity
    fmt.Fprintf(os.Stdout, "Response from `ConfiguratorAPI.ProductConfiguratorGetConfiguratorByProductId2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**productId** | **string** |  | 
**status** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductConfiguratorGetConfiguratorByProductId2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ProductconfiguratorconfiguratorEntity**](ProductconfiguratorconfiguratorEntity.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorListConfigurators

> ConfiguratorListResponse ProductConfiguratorListConfigurators(ctx, tenantId, productId, pageSize).Body(body).Execute()

List Product Configurators



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/gemini-commerce/go-client-product_configurator"
)

func main() {
    tenantId := "tenantId_example" // string | 
    productId := "productId_example" // string | 
    pageSize := int64(789) // int64 | 
    body := *openapiclient.NewProductConfiguratorListPropertiesRequest() // ProductConfiguratorListPropertiesRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfiguratorAPI.ProductConfiguratorListConfigurators(context.Background(), tenantId, productId, pageSize).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfiguratorAPI.ProductConfiguratorListConfigurators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorListConfigurators`: ConfiguratorListResponse
    fmt.Fprintf(os.Stdout, "Response from `ConfiguratorAPI.ProductConfiguratorListConfigurators`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**productId** | **string** |  | 
**pageSize** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductConfiguratorListConfiguratorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**ProductConfiguratorListPropertiesRequest**](ProductConfiguratorListPropertiesRequest.md) |  | 

### Return type

[**ConfiguratorListResponse**](ConfiguratorListResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorUpdateConfigurator

> ProductconfiguratorconfiguratorEntity ProductConfiguratorUpdateConfigurator(ctx, tenantId, configuratorId).Body(body).Execute()

Update Configurator



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/gemini-commerce/go-client-product_configurator"
)

func main() {
    tenantId := "tenantId_example" // string | 
    configuratorId := "configuratorId_example" // string | 
    body := *openapiclient.NewProductConfiguratorUpdateConfiguratorRequest() // ProductConfiguratorUpdateConfiguratorRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfiguratorAPI.ProductConfiguratorUpdateConfigurator(context.Background(), tenantId, configuratorId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfiguratorAPI.ProductConfiguratorUpdateConfigurator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorUpdateConfigurator`: ProductconfiguratorconfiguratorEntity
    fmt.Fprintf(os.Stdout, "Response from `ConfiguratorAPI.ProductConfiguratorUpdateConfigurator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**configuratorId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductConfiguratorUpdateConfiguratorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**ProductConfiguratorUpdateConfiguratorRequest**](ProductConfiguratorUpdateConfiguratorRequest.md) |  | 

### Return type

[**ProductconfiguratorconfiguratorEntity**](ProductconfiguratorconfiguratorEntity.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

