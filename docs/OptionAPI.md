# \OptionAPI

All URIs are relative to *https://product-configurator.api.gogemini.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProductConfiguratorBulkCreateOptions**](OptionAPI.md#ProductConfiguratorBulkCreateOptions) | **Post** /v1/{tenantId}/step/{stepId}/option/create/bulk | Bulk Create Options
[**ProductConfiguratorBulkDeleteOptions**](OptionAPI.md#ProductConfiguratorBulkDeleteOptions) | **Post** /v1/{tenantId}/option/delete/bulk | Bulk Delete Options
[**ProductConfiguratorBulkUpdateOptions**](OptionAPI.md#ProductConfiguratorBulkUpdateOptions) | **Put** /v1/{tenantId}/option/bulk | Bulk Update Options
[**ProductConfiguratorCopyOption**](OptionAPI.md#ProductConfiguratorCopyOption) | **Post** /v1/{tenantId}/option/{sourceOptionId}/copy | Copy Option
[**ProductConfiguratorCreateOption**](OptionAPI.md#ProductConfiguratorCreateOption) | **Post** /v1/{tenantId}/step/{stepId}/option/create | Create Option
[**ProductConfiguratorDeleteOption**](OptionAPI.md#ProductConfiguratorDeleteOption) | **Delete** /v1/{tenantId}/option/{optionId} | Delete Option
[**ProductConfiguratorListOptions**](OptionAPI.md#ProductConfiguratorListOptions) | **Post** /v1/{tenantId}/step/{stepId}/page-size/{pageSize}/options | List Options
[**ProductConfiguratorUpdateOption**](OptionAPI.md#ProductConfiguratorUpdateOption) | **Put** /v1/{tenantId}/option/{optionId} | Update Option



## ProductConfiguratorBulkCreateOptions

> ProductconfiguratoroptionBulkCreateResponse ProductConfiguratorBulkCreateOptions(ctx, tenantId, stepId).Body(body).Execute()

Bulk Create Options



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
    stepId := "stepId_example" // string | 
    body := *openapiclient.NewProductConfiguratorBulkCreateOptionsRequest() // ProductConfiguratorBulkCreateOptionsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OptionAPI.ProductConfiguratorBulkCreateOptions(context.Background(), tenantId, stepId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OptionAPI.ProductConfiguratorBulkCreateOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorBulkCreateOptions`: ProductconfiguratoroptionBulkCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `OptionAPI.ProductConfiguratorBulkCreateOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**stepId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductConfiguratorBulkCreateOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**ProductConfiguratorBulkCreateOptionsRequest**](ProductConfiguratorBulkCreateOptionsRequest.md) |  | 

### Return type

[**ProductconfiguratoroptionBulkCreateResponse**](ProductconfiguratoroptionBulkCreateResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization), [APIAuthorization](../README.md#APIAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorBulkDeleteOptions

> map[string]interface{} ProductConfiguratorBulkDeleteOptions(ctx, tenantId).Body(body).Execute()

Bulk Delete Options



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
    body := *openapiclient.NewProductConfiguratorBulkDeleteOptionsRequest() // ProductConfiguratorBulkDeleteOptionsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OptionAPI.ProductConfiguratorBulkDeleteOptions(context.Background(), tenantId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OptionAPI.ProductConfiguratorBulkDeleteOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorBulkDeleteOptions`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OptionAPI.ProductConfiguratorBulkDeleteOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductConfiguratorBulkDeleteOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ProductConfiguratorBulkDeleteOptionsRequest**](ProductConfiguratorBulkDeleteOptionsRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[standardAuthorization](../README.md#standardAuthorization), [APIAuthorization](../README.md#APIAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorBulkUpdateOptions

> ProductconfiguratoroptionBulkUpdateResponse ProductConfiguratorBulkUpdateOptions(ctx, tenantId).Body(body).Execute()

Bulk Update Options



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
    body := *openapiclient.NewProductConfiguratorBulkUpdateOptionsRequest() // ProductConfiguratorBulkUpdateOptionsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OptionAPI.ProductConfiguratorBulkUpdateOptions(context.Background(), tenantId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OptionAPI.ProductConfiguratorBulkUpdateOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorBulkUpdateOptions`: ProductconfiguratoroptionBulkUpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `OptionAPI.ProductConfiguratorBulkUpdateOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductConfiguratorBulkUpdateOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ProductConfiguratorBulkUpdateOptionsRequest**](ProductConfiguratorBulkUpdateOptionsRequest.md) |  | 

### Return type

[**ProductconfiguratoroptionBulkUpdateResponse**](ProductconfiguratoroptionBulkUpdateResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization), [APIAuthorization](../README.md#APIAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorCopyOption

> ProductconfiguratoroptionEntity ProductConfiguratorCopyOption(ctx, tenantId, sourceOptionId).Body(body).Execute()

Copy Option



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
    sourceOptionId := "sourceOptionId_example" // string | 
    body := *openapiclient.NewProductConfiguratorCopyOptionRequest() // ProductConfiguratorCopyOptionRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OptionAPI.ProductConfiguratorCopyOption(context.Background(), tenantId, sourceOptionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OptionAPI.ProductConfiguratorCopyOption``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorCopyOption`: ProductconfiguratoroptionEntity
    fmt.Fprintf(os.Stdout, "Response from `OptionAPI.ProductConfiguratorCopyOption`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**sourceOptionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductConfiguratorCopyOptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**ProductConfiguratorCopyOptionRequest**](ProductConfiguratorCopyOptionRequest.md) |  | 

### Return type

[**ProductconfiguratoroptionEntity**](ProductconfiguratoroptionEntity.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization), [APIAuthorization](../README.md#APIAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorCreateOption

> ProductconfiguratoroptionEntity ProductConfiguratorCreateOption(ctx, tenantId, stepId).Body(body).Execute()

Create Option



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
    stepId := "stepId_example" // string | 
    body := *openapiclient.NewProductConfiguratorCreateOptionRequest() // ProductConfiguratorCreateOptionRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OptionAPI.ProductConfiguratorCreateOption(context.Background(), tenantId, stepId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OptionAPI.ProductConfiguratorCreateOption``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorCreateOption`: ProductconfiguratoroptionEntity
    fmt.Fprintf(os.Stdout, "Response from `OptionAPI.ProductConfiguratorCreateOption`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**stepId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductConfiguratorCreateOptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**ProductConfiguratorCreateOptionRequest**](ProductConfiguratorCreateOptionRequest.md) |  | 

### Return type

[**ProductconfiguratoroptionEntity**](ProductconfiguratoroptionEntity.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization), [APIAuthorization](../README.md#APIAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorDeleteOption

> map[string]interface{} ProductConfiguratorDeleteOption(ctx, tenantId, optionId).Execute()

Delete Option



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
    optionId := "optionId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OptionAPI.ProductConfiguratorDeleteOption(context.Background(), tenantId, optionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OptionAPI.ProductConfiguratorDeleteOption``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorDeleteOption`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OptionAPI.ProductConfiguratorDeleteOption`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**optionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductConfiguratorDeleteOptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[standardAuthorization](../README.md#standardAuthorization), [APIAuthorization](../README.md#APIAuthorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorListOptions

> OptionListOptionsResponse ProductConfiguratorListOptions(ctx, tenantId, stepId, pageSize).Body(body).Execute()

List Options



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
    stepId := "stepId_example" // string | 
    pageSize := int64(789) // int64 | 
    body := *openapiclient.NewProductConfiguratorListPropertiesRequest() // ProductConfiguratorListPropertiesRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OptionAPI.ProductConfiguratorListOptions(context.Background(), tenantId, stepId, pageSize).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OptionAPI.ProductConfiguratorListOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorListOptions`: OptionListOptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `OptionAPI.ProductConfiguratorListOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**stepId** | **string** |  | 
**pageSize** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductConfiguratorListOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**ProductConfiguratorListPropertiesRequest**](ProductConfiguratorListPropertiesRequest.md) |  | 

### Return type

[**OptionListOptionsResponse**](OptionListOptionsResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization), [APIAuthorization](../README.md#APIAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorUpdateOption

> ProductconfiguratoroptionEntity ProductConfiguratorUpdateOption(ctx, tenantId, optionId).Body(body).Execute()

Update Option



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
    optionId := "optionId_example" // string | 
    body := *openapiclient.NewProductConfiguratorUpdateOptionRequest() // ProductConfiguratorUpdateOptionRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OptionAPI.ProductConfiguratorUpdateOption(context.Background(), tenantId, optionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OptionAPI.ProductConfiguratorUpdateOption``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorUpdateOption`: ProductconfiguratoroptionEntity
    fmt.Fprintf(os.Stdout, "Response from `OptionAPI.ProductConfiguratorUpdateOption`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**optionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductConfiguratorUpdateOptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**ProductConfiguratorUpdateOptionRequest**](ProductConfiguratorUpdateOptionRequest.md) |  | 

### Return type

[**ProductconfiguratoroptionEntity**](ProductconfiguratoroptionEntity.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization), [APIAuthorization](../README.md#APIAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

