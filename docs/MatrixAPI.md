# \MatrixAPI

All URIs are relative to *https://product-configurator.api.gogemini.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProductConfiguratorCreateMatrix**](MatrixAPI.md#ProductConfiguratorCreateMatrix) | **Post** /v1/{tenantId}/matrix/create | Create Matrix
[**ProductConfiguratorDeleteMatrix**](MatrixAPI.md#ProductConfiguratorDeleteMatrix) | **Delete** /v1/{tenantId}/matrix/{matrixId} | Delete Matrix
[**ProductConfiguratorGetMatrix**](MatrixAPI.md#ProductConfiguratorGetMatrix) | **Get** /v1/{tenantId}/matrix/{matrixId} | Get Matrix
[**ProductConfiguratorListMatrices**](MatrixAPI.md#ProductConfiguratorListMatrices) | **Post** /v1/{tenantId}/configurator/{configuratorId}/page-size/{pageSize}/matrices | List Matrices
[**ProductConfiguratorRemovePricelistFromMatrix**](MatrixAPI.md#ProductConfiguratorRemovePricelistFromMatrix) | **Delete** /v1/{tenantId}/matrix/{matrixId}/pricelist/{pricelistGrn} | Remove Pricelist from Matrix
[**ProductConfiguratorUpdateMatrix**](MatrixAPI.md#ProductConfiguratorUpdateMatrix) | **Put** /v1/{tenantId}/matrix/{matrixId} | Update Matrix



## ProductConfiguratorCreateMatrix

> ProductconfiguratormatrixEntity ProductConfiguratorCreateMatrix(ctx, tenantId).Body(body).Execute()

Create Matrix



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
    body := *openapiclient.NewProductConfiguratorCreateMatrixRequest() // ProductConfiguratorCreateMatrixRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MatrixAPI.ProductConfiguratorCreateMatrix(context.Background(), tenantId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MatrixAPI.ProductConfiguratorCreateMatrix``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorCreateMatrix`: ProductconfiguratormatrixEntity
    fmt.Fprintf(os.Stdout, "Response from `MatrixAPI.ProductConfiguratorCreateMatrix`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductConfiguratorCreateMatrixRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ProductConfiguratorCreateMatrixRequest**](ProductConfiguratorCreateMatrixRequest.md) |  | 

### Return type

[**ProductconfiguratormatrixEntity**](ProductconfiguratormatrixEntity.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization), [APIAuthorization](../README.md#APIAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorDeleteMatrix

> map[string]interface{} ProductConfiguratorDeleteMatrix(ctx, tenantId, matrixId).Execute()

Delete Matrix



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
    matrixId := "matrixId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MatrixAPI.ProductConfiguratorDeleteMatrix(context.Background(), tenantId, matrixId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MatrixAPI.ProductConfiguratorDeleteMatrix``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorDeleteMatrix`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MatrixAPI.ProductConfiguratorDeleteMatrix`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**matrixId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductConfiguratorDeleteMatrixRequest struct via the builder pattern


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


## ProductConfiguratorGetMatrix

> ProductconfiguratormatrixEntity ProductConfiguratorGetMatrix(ctx, tenantId, matrixId).Execute()

Get Matrix



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
    matrixId := "matrixId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MatrixAPI.ProductConfiguratorGetMatrix(context.Background(), tenantId, matrixId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MatrixAPI.ProductConfiguratorGetMatrix``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorGetMatrix`: ProductconfiguratormatrixEntity
    fmt.Fprintf(os.Stdout, "Response from `MatrixAPI.ProductConfiguratorGetMatrix`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**matrixId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductConfiguratorGetMatrixRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ProductconfiguratormatrixEntity**](ProductconfiguratormatrixEntity.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization), [APIAuthorization](../README.md#APIAuthorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorListMatrices

> MatrixListMatricesResponse ProductConfiguratorListMatrices(ctx, tenantId, configuratorId, pageSize).Body(body).Execute()

List Matrices



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
    pageSize := "pageSize_example" // string | 
    body := *openapiclient.NewProductConfiguratorListMatricesRequest() // ProductConfiguratorListMatricesRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MatrixAPI.ProductConfiguratorListMatrices(context.Background(), tenantId, configuratorId, pageSize).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MatrixAPI.ProductConfiguratorListMatrices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorListMatrices`: MatrixListMatricesResponse
    fmt.Fprintf(os.Stdout, "Response from `MatrixAPI.ProductConfiguratorListMatrices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**configuratorId** | **string** |  | 
**pageSize** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductConfiguratorListMatricesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**ProductConfiguratorListMatricesRequest**](ProductConfiguratorListMatricesRequest.md) |  | 

### Return type

[**MatrixListMatricesResponse**](MatrixListMatricesResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization), [APIAuthorization](../README.md#APIAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorRemovePricelistFromMatrix

> ProductconfiguratormatrixEntity ProductConfiguratorRemovePricelistFromMatrix(ctx, tenantId, matrixId, pricelistGrn).Execute()

Remove Pricelist from Matrix



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
    matrixId := "matrixId_example" // string | 
    pricelistGrn := "pricelistGrn_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MatrixAPI.ProductConfiguratorRemovePricelistFromMatrix(context.Background(), tenantId, matrixId, pricelistGrn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MatrixAPI.ProductConfiguratorRemovePricelistFromMatrix``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorRemovePricelistFromMatrix`: ProductconfiguratormatrixEntity
    fmt.Fprintf(os.Stdout, "Response from `MatrixAPI.ProductConfiguratorRemovePricelistFromMatrix`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**matrixId** | **string** |  | 
**pricelistGrn** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductConfiguratorRemovePricelistFromMatrixRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ProductconfiguratormatrixEntity**](ProductconfiguratormatrixEntity.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization), [APIAuthorization](../README.md#APIAuthorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorUpdateMatrix

> ProductconfiguratormatrixEntity ProductConfiguratorUpdateMatrix(ctx, tenantId, matrixId).Body(body).Execute()

Update Matrix



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
    matrixId := "matrixId_example" // string | 
    body := *openapiclient.NewProductConfiguratorUpdateMatrixRequest() // ProductConfiguratorUpdateMatrixRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MatrixAPI.ProductConfiguratorUpdateMatrix(context.Background(), tenantId, matrixId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MatrixAPI.ProductConfiguratorUpdateMatrix``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorUpdateMatrix`: ProductconfiguratormatrixEntity
    fmt.Fprintf(os.Stdout, "Response from `MatrixAPI.ProductConfiguratorUpdateMatrix`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**matrixId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductConfiguratorUpdateMatrixRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**ProductConfiguratorUpdateMatrixRequest**](ProductConfiguratorUpdateMatrixRequest.md) |  | 

### Return type

[**ProductconfiguratormatrixEntity**](ProductconfiguratormatrixEntity.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization), [APIAuthorization](../README.md#APIAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

