# \ProductConfiguratorApi

All URIs are relative to *https://product-configurator.api.gogemini.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProductConfiguratorBulkCreateOptions**](ProductConfiguratorApi.md#ProductConfiguratorBulkCreateOptions) | **Post** /v1/{tenantId}/step/{stepId}/option/create/bulk | 
[**ProductConfiguratorBulkCreateProperties**](ProductConfiguratorApi.md#ProductConfiguratorBulkCreateProperties) | **Post** /v1/{tenantId}/property/create/bulk | 
[**ProductConfiguratorBulkCreateSteps**](ProductConfiguratorApi.md#ProductConfiguratorBulkCreateSteps) | **Post** /v1/{tenantId}/configurator/{configuratorId}/step/create/bulk | 
[**ProductConfiguratorBulkDeleteOptions**](ProductConfiguratorApi.md#ProductConfiguratorBulkDeleteOptions) | **Post** /v1/{tenantId}/option/delete/bulk | 
[**ProductConfiguratorBulkDeleteSteps**](ProductConfiguratorApi.md#ProductConfiguratorBulkDeleteSteps) | **Post** /v1/{tenantId}/step/delete/bulk | 
[**ProductConfiguratorBulkUpdateOptions**](ProductConfiguratorApi.md#ProductConfiguratorBulkUpdateOptions) | **Put** /v1/{tenantId}/option/bulk | 
[**ProductConfiguratorBulkUpdateProperties**](ProductConfiguratorApi.md#ProductConfiguratorBulkUpdateProperties) | **Put** /v1/{tenantId}/properties/bulk | 
[**ProductConfiguratorCopyConfigurator**](ProductConfiguratorApi.md#ProductConfiguratorCopyConfigurator) | **Post** /v1/{tenantId}/product/{sourceConfiguratorId}/copy | 
[**ProductConfiguratorCopyOption**](ProductConfiguratorApi.md#ProductConfiguratorCopyOption) | **Post** /v1/{tenantId}/option/{sourceOptionId}/copy | 
[**ProductConfiguratorCopyStep**](ProductConfiguratorApi.md#ProductConfiguratorCopyStep) | **Post** /v1/{tenantId}/step/{sourceStepId}/copy | 
[**ProductConfiguratorCreateConfigurator**](ProductConfiguratorApi.md#ProductConfiguratorCreateConfigurator) | **Post** /v1/{tenantId}/product/{productId}/create | 
[**ProductConfiguratorCreateDependency**](ProductConfiguratorApi.md#ProductConfiguratorCreateDependency) | **Post** /v1/{tenantId}/step/{stepId}/dependency/create | 
[**ProductConfiguratorCreateMatrix**](ProductConfiguratorApi.md#ProductConfiguratorCreateMatrix) | **Post** /v1/{tenantId}/matrix/create | 
[**ProductConfiguratorCreateOption**](ProductConfiguratorApi.md#ProductConfiguratorCreateOption) | **Post** /v1/{tenantId}/step/{stepId}/option/create | 
[**ProductConfiguratorCreateProperty**](ProductConfiguratorApi.md#ProductConfiguratorCreateProperty) | **Post** /v1/{tenantId}/property/create | 
[**ProductConfiguratorCreateStep**](ProductConfiguratorApi.md#ProductConfiguratorCreateStep) | **Post** /v1/{tenantId}/configurator/{configuratorId}/step/create | 
[**ProductConfiguratorDeleteConfigurator**](ProductConfiguratorApi.md#ProductConfiguratorDeleteConfigurator) | **Delete** /v1/{tenantId}/configurator/{configuratorId} | 
[**ProductConfiguratorDeleteDependency**](ProductConfiguratorApi.md#ProductConfiguratorDeleteDependency) | **Delete** /v1/{tenantId}/dependency/{dependencyId} | 
[**ProductConfiguratorDeleteMatrix**](ProductConfiguratorApi.md#ProductConfiguratorDeleteMatrix) | **Delete** /v1/{tenantId}/matrix/{matrixId} | 
[**ProductConfiguratorDeleteOption**](ProductConfiguratorApi.md#ProductConfiguratorDeleteOption) | **Delete** /v1/{tenantId}/option/{optionId} | 
[**ProductConfiguratorDeleteStep**](ProductConfiguratorApi.md#ProductConfiguratorDeleteStep) | **Delete** /v1/{tenantId}/step/{stepId} | 
[**ProductConfiguratorGetAvailableConfiguration**](ProductConfiguratorApi.md#ProductConfiguratorGetAvailableConfiguration) | **Get** /v1/{tenantId}/product/{productId}/configuration | 
[**ProductConfiguratorGetAvailableConfiguration2**](ProductConfiguratorApi.md#ProductConfiguratorGetAvailableConfiguration2) | **Post** /v1/{tenantId}/product/{productId}/configuration | 
[**ProductConfiguratorGetConfigurationFromSelections**](ProductConfiguratorApi.md#ProductConfiguratorGetConfigurationFromSelections) | **Post** /v1/{tenantId}/product/{productId}/configuration-from-selections | 
[**ProductConfiguratorGetConfiguratorByProductId**](ProductConfiguratorApi.md#ProductConfiguratorGetConfiguratorByProductId) | **Get** /v1/{tenantId}/product/{productId} | 
[**ProductConfiguratorGetConfiguratorByProductId2**](ProductConfiguratorApi.md#ProductConfiguratorGetConfiguratorByProductId2) | **Get** /v1/{tenantId}/product/{productId}/status/{status} | 
[**ProductConfiguratorGetMatrix**](ProductConfiguratorApi.md#ProductConfiguratorGetMatrix) | **Get** /v1/{tenantId}/matrix/{matrixId} | 
[**ProductConfiguratorListConfigurators**](ProductConfiguratorApi.md#ProductConfiguratorListConfigurators) | **Post** /v1/{tenantId}/product/{productId}/page-size/{pageSize}/configurators | 
[**ProductConfiguratorListDependencies**](ProductConfiguratorApi.md#ProductConfiguratorListDependencies) | **Post** /v1/{tenantId}/page-size/{pageSize}/dependencies | 
[**ProductConfiguratorListMatrices**](ProductConfiguratorApi.md#ProductConfiguratorListMatrices) | **Post** /v1/{tenantId}/configurator/{configuratorId}/page-size/{pageSize}/matrices | 
[**ProductConfiguratorListOptions**](ProductConfiguratorApi.md#ProductConfiguratorListOptions) | **Post** /v1/{tenantId}/step/{stepId}/page-size/{pageSize}/options | 
[**ProductConfiguratorListProperties**](ProductConfiguratorApi.md#ProductConfiguratorListProperties) | **Post** /v1/{tenantId}/matrix/{matrixId}/page-size/{pageSize}/properties | 
[**ProductConfiguratorRemovePricelistFromMatrix**](ProductConfiguratorApi.md#ProductConfiguratorRemovePricelistFromMatrix) | **Delete** /v1/{tenantId}/matrix/{matrixId}/pricelist/{pricelistGrn} | 
[**ProductConfiguratorUpdateConfigurator**](ProductConfiguratorApi.md#ProductConfiguratorUpdateConfigurator) | **Put** /v1/{tenantId}/configurator/{configuratorId} | 
[**ProductConfiguratorUpdateDependency**](ProductConfiguratorApi.md#ProductConfiguratorUpdateDependency) | **Put** /v1/{tenantId}/dependency/{dependencyId} | 
[**ProductConfiguratorUpdateMatrix**](ProductConfiguratorApi.md#ProductConfiguratorUpdateMatrix) | **Put** /v1/{tenantId}/matrix/{matrixId} | 
[**ProductConfiguratorUpdateOption**](ProductConfiguratorApi.md#ProductConfiguratorUpdateOption) | **Put** /v1/{tenantId}/option/{optionId} | 
[**ProductConfiguratorUpdateProperty**](ProductConfiguratorApi.md#ProductConfiguratorUpdateProperty) | **Put** /v1/{tenantId}/property/{propertyId} | 
[**ProductConfiguratorUpdateStep**](ProductConfiguratorApi.md#ProductConfiguratorUpdateStep) | **Put** /v1/{tenantId}/step/{stepId} | 



## ProductConfiguratorBulkCreateOptions

> ProductconfiguratoroptionBulkCreateResponse ProductConfiguratorBulkCreateOptions(ctx, tenantId, stepId).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenantId := "tenantId_example" // string | 
    stepId := "stepId_example" // string | 
    body := *openapiclient.NewProductconfiguratoroptionBulkCreateRequest() // ProductconfiguratoroptionBulkCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductConfiguratorApi.ProductConfiguratorBulkCreateOptions(context.Background(), tenantId, stepId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductConfiguratorApi.ProductConfiguratorBulkCreateOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorBulkCreateOptions`: ProductconfiguratoroptionBulkCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductConfiguratorApi.ProductConfiguratorBulkCreateOptions`: %v\n", resp)
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


 **body** | [**ProductconfiguratoroptionBulkCreateRequest**](ProductconfiguratoroptionBulkCreateRequest.md) |  | 

### Return type

[**ProductconfiguratoroptionBulkCreateResponse**](ProductconfiguratoroptionBulkCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorBulkCreateProperties

> ProductconfiguratorpropertyBulkCreateResponse ProductConfiguratorBulkCreateProperties(ctx, tenantId).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenantId := "tenantId_example" // string | 
    body := *openapiclient.NewProductconfiguratorpropertyBulkCreateRequest() // ProductconfiguratorpropertyBulkCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductConfiguratorApi.ProductConfiguratorBulkCreateProperties(context.Background(), tenantId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductConfiguratorApi.ProductConfiguratorBulkCreateProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorBulkCreateProperties`: ProductconfiguratorpropertyBulkCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductConfiguratorApi.ProductConfiguratorBulkCreateProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductConfiguratorBulkCreatePropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ProductconfiguratorpropertyBulkCreateRequest**](ProductconfiguratorpropertyBulkCreateRequest.md) |  | 

### Return type

[**ProductconfiguratorpropertyBulkCreateResponse**](ProductconfiguratorpropertyBulkCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorBulkCreateSteps

> ProductconfiguratorstepBulkCreateResponse ProductConfiguratorBulkCreateSteps(ctx, tenantId, configuratorId).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenantId := "tenantId_example" // string | 
    configuratorId := "configuratorId_example" // string | 
    body := *openapiclient.NewProductconfiguratorstepBulkCreateRequest() // ProductconfiguratorstepBulkCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductConfiguratorApi.ProductConfiguratorBulkCreateSteps(context.Background(), tenantId, configuratorId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductConfiguratorApi.ProductConfiguratorBulkCreateSteps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorBulkCreateSteps`: ProductconfiguratorstepBulkCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductConfiguratorApi.ProductConfiguratorBulkCreateSteps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**configuratorId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductConfiguratorBulkCreateStepsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**ProductconfiguratorstepBulkCreateRequest**](ProductconfiguratorstepBulkCreateRequest.md) |  | 

### Return type

[**ProductconfiguratorstepBulkCreateResponse**](ProductconfiguratorstepBulkCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorBulkDeleteOptions

> map[string]interface{} ProductConfiguratorBulkDeleteOptions(ctx, tenantId).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenantId := "tenantId_example" // string | 
    body := *openapiclient.NewProductconfiguratoroptionBulkDeleteRequest() // ProductconfiguratoroptionBulkDeleteRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductConfiguratorApi.ProductConfiguratorBulkDeleteOptions(context.Background(), tenantId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductConfiguratorApi.ProductConfiguratorBulkDeleteOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorBulkDeleteOptions`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProductConfiguratorApi.ProductConfiguratorBulkDeleteOptions`: %v\n", resp)
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

 **body** | [**ProductconfiguratoroptionBulkDeleteRequest**](ProductconfiguratoroptionBulkDeleteRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorBulkDeleteSteps

> map[string]interface{} ProductConfiguratorBulkDeleteSteps(ctx, tenantId).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenantId := "tenantId_example" // string | 
    body := *openapiclient.NewProductconfiguratorstepBulkDeleteRequest() // ProductconfiguratorstepBulkDeleteRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductConfiguratorApi.ProductConfiguratorBulkDeleteSteps(context.Background(), tenantId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductConfiguratorApi.ProductConfiguratorBulkDeleteSteps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorBulkDeleteSteps`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProductConfiguratorApi.ProductConfiguratorBulkDeleteSteps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductConfiguratorBulkDeleteStepsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ProductconfiguratorstepBulkDeleteRequest**](ProductconfiguratorstepBulkDeleteRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorBulkUpdateOptions

> ProductconfiguratoroptionBulkUpdateResponse ProductConfiguratorBulkUpdateOptions(ctx, tenantId).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenantId := "tenantId_example" // string | 
    body := *openapiclient.NewProductconfiguratoroptionBulkUpdateRequest() // ProductconfiguratoroptionBulkUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductConfiguratorApi.ProductConfiguratorBulkUpdateOptions(context.Background(), tenantId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductConfiguratorApi.ProductConfiguratorBulkUpdateOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorBulkUpdateOptions`: ProductconfiguratoroptionBulkUpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductConfiguratorApi.ProductConfiguratorBulkUpdateOptions`: %v\n", resp)
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

 **body** | [**ProductconfiguratoroptionBulkUpdateRequest**](ProductconfiguratoroptionBulkUpdateRequest.md) |  | 

### Return type

[**ProductconfiguratoroptionBulkUpdateResponse**](ProductconfiguratoroptionBulkUpdateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorBulkUpdateProperties

> ProductconfiguratorpropertyBulkUpdateResponse ProductConfiguratorBulkUpdateProperties(ctx, tenantId).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenantId := "tenantId_example" // string | 
    body := *openapiclient.NewProductconfiguratorpropertyBulkUpdateRequest() // ProductconfiguratorpropertyBulkUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductConfiguratorApi.ProductConfiguratorBulkUpdateProperties(context.Background(), tenantId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductConfiguratorApi.ProductConfiguratorBulkUpdateProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorBulkUpdateProperties`: ProductconfiguratorpropertyBulkUpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductConfiguratorApi.ProductConfiguratorBulkUpdateProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductConfiguratorBulkUpdatePropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ProductconfiguratorpropertyBulkUpdateRequest**](ProductconfiguratorpropertyBulkUpdateRequest.md) |  | 

### Return type

[**ProductconfiguratorpropertyBulkUpdateResponse**](ProductconfiguratorpropertyBulkUpdateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorCopyConfigurator

> ProductconfiguratorconfiguratorEntity ProductConfiguratorCopyConfigurator(ctx, tenantId, sourceConfiguratorId).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenantId := "tenantId_example" // string | 
    sourceConfiguratorId := "sourceConfiguratorId_example" // string | 
    body := *openapiclient.NewProductconfiguratorconfiguratorCopyRequest() // ProductconfiguratorconfiguratorCopyRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductConfiguratorApi.ProductConfiguratorCopyConfigurator(context.Background(), tenantId, sourceConfiguratorId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductConfiguratorApi.ProductConfiguratorCopyConfigurator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorCopyConfigurator`: ProductconfiguratorconfiguratorEntity
    fmt.Fprintf(os.Stdout, "Response from `ProductConfiguratorApi.ProductConfiguratorCopyConfigurator`: %v\n", resp)
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


 **body** | [**ProductconfiguratorconfiguratorCopyRequest**](ProductconfiguratorconfiguratorCopyRequest.md) |  | 

### Return type

[**ProductconfiguratorconfiguratorEntity**](ProductconfiguratorconfiguratorEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorCopyOption

> ProductconfiguratoroptionEntity ProductConfiguratorCopyOption(ctx, tenantId, sourceOptionId).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenantId := "tenantId_example" // string | 
    sourceOptionId := "sourceOptionId_example" // string | 
    body := *openapiclient.NewProductconfiguratoroptionCopyRequest() // ProductconfiguratoroptionCopyRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductConfiguratorApi.ProductConfiguratorCopyOption(context.Background(), tenantId, sourceOptionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductConfiguratorApi.ProductConfiguratorCopyOption``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorCopyOption`: ProductconfiguratoroptionEntity
    fmt.Fprintf(os.Stdout, "Response from `ProductConfiguratorApi.ProductConfiguratorCopyOption`: %v\n", resp)
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


 **body** | [**ProductconfiguratoroptionCopyRequest**](ProductconfiguratoroptionCopyRequest.md) |  | 

### Return type

[**ProductconfiguratoroptionEntity**](ProductconfiguratoroptionEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorCopyStep

> ProductconfiguratorstepEntity ProductConfiguratorCopyStep(ctx, tenantId, sourceStepId).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenantId := "tenantId_example" // string | 
    sourceStepId := "sourceStepId_example" // string | 
    body := *openapiclient.NewProductconfiguratorstepCopyRequest() // ProductconfiguratorstepCopyRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductConfiguratorApi.ProductConfiguratorCopyStep(context.Background(), tenantId, sourceStepId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductConfiguratorApi.ProductConfiguratorCopyStep``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorCopyStep`: ProductconfiguratorstepEntity
    fmt.Fprintf(os.Stdout, "Response from `ProductConfiguratorApi.ProductConfiguratorCopyStep`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**sourceStepId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductConfiguratorCopyStepRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**ProductconfiguratorstepCopyRequest**](ProductconfiguratorstepCopyRequest.md) |  | 

### Return type

[**ProductconfiguratorstepEntity**](ProductconfiguratorstepEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorCreateConfigurator

> ProductconfiguratorconfiguratorEntity ProductConfiguratorCreateConfigurator(ctx, tenantId, productId).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenantId := "tenantId_example" // string | 
    productId := "productId_example" // string | 
    body := *openapiclient.NewProductconfiguratorconfiguratorCreateRequest() // ProductconfiguratorconfiguratorCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductConfiguratorApi.ProductConfiguratorCreateConfigurator(context.Background(), tenantId, productId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductConfiguratorApi.ProductConfiguratorCreateConfigurator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorCreateConfigurator`: ProductconfiguratorconfiguratorEntity
    fmt.Fprintf(os.Stdout, "Response from `ProductConfiguratorApi.ProductConfiguratorCreateConfigurator`: %v\n", resp)
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


 **body** | [**ProductconfiguratorconfiguratorCreateRequest**](ProductconfiguratorconfiguratorCreateRequest.md) |  | 

### Return type

[**ProductconfiguratorconfiguratorEntity**](ProductconfiguratorconfiguratorEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorCreateDependency

> ProductconfiguratordependencyEntity ProductConfiguratorCreateDependency(ctx, tenantId, stepId).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenantId := "tenantId_example" // string | 
    stepId := "stepId_example" // string | 
    body := *openapiclient.NewProductconfiguratordependencyCreateRequest() // ProductconfiguratordependencyCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductConfiguratorApi.ProductConfiguratorCreateDependency(context.Background(), tenantId, stepId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductConfiguratorApi.ProductConfiguratorCreateDependency``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorCreateDependency`: ProductconfiguratordependencyEntity
    fmt.Fprintf(os.Stdout, "Response from `ProductConfiguratorApi.ProductConfiguratorCreateDependency`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**stepId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductConfiguratorCreateDependencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**ProductconfiguratordependencyCreateRequest**](ProductconfiguratordependencyCreateRequest.md) |  | 

### Return type

[**ProductconfiguratordependencyEntity**](ProductconfiguratordependencyEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorCreateMatrix

> ProductconfiguratormatrixEntity ProductConfiguratorCreateMatrix(ctx, tenantId).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenantId := "tenantId_example" // string | 
    body := *openapiclient.NewProductconfiguratormatrixCreateRequest() // ProductconfiguratormatrixCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductConfiguratorApi.ProductConfiguratorCreateMatrix(context.Background(), tenantId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductConfiguratorApi.ProductConfiguratorCreateMatrix``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorCreateMatrix`: ProductconfiguratormatrixEntity
    fmt.Fprintf(os.Stdout, "Response from `ProductConfiguratorApi.ProductConfiguratorCreateMatrix`: %v\n", resp)
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

 **body** | [**ProductconfiguratormatrixCreateRequest**](ProductconfiguratormatrixCreateRequest.md) |  | 

### Return type

[**ProductconfiguratormatrixEntity**](ProductconfiguratormatrixEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorCreateOption

> ProductconfiguratoroptionEntity ProductConfiguratorCreateOption(ctx, tenantId, stepId).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenantId := "tenantId_example" // string | 
    stepId := "stepId_example" // string | 
    body := *openapiclient.NewProductconfiguratoroptionCreateRequest() // ProductconfiguratoroptionCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductConfiguratorApi.ProductConfiguratorCreateOption(context.Background(), tenantId, stepId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductConfiguratorApi.ProductConfiguratorCreateOption``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorCreateOption`: ProductconfiguratoroptionEntity
    fmt.Fprintf(os.Stdout, "Response from `ProductConfiguratorApi.ProductConfiguratorCreateOption`: %v\n", resp)
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


 **body** | [**ProductconfiguratoroptionCreateRequest**](ProductconfiguratoroptionCreateRequest.md) |  | 

### Return type

[**ProductconfiguratoroptionEntity**](ProductconfiguratoroptionEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorCreateProperty

> ProductconfiguratorpropertyEntity ProductConfiguratorCreateProperty(ctx, tenantId).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenantId := "tenantId_example" // string | 
    body := *openapiclient.NewProductconfiguratorpropertyCreateRequest() // ProductconfiguratorpropertyCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductConfiguratorApi.ProductConfiguratorCreateProperty(context.Background(), tenantId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductConfiguratorApi.ProductConfiguratorCreateProperty``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorCreateProperty`: ProductconfiguratorpropertyEntity
    fmt.Fprintf(os.Stdout, "Response from `ProductConfiguratorApi.ProductConfiguratorCreateProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductConfiguratorCreatePropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ProductconfiguratorpropertyCreateRequest**](ProductconfiguratorpropertyCreateRequest.md) |  | 

### Return type

[**ProductconfiguratorpropertyEntity**](ProductconfiguratorpropertyEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorCreateStep

> ProductconfiguratorstepEntity ProductConfiguratorCreateStep(ctx, tenantId, configuratorId).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenantId := "tenantId_example" // string | 
    configuratorId := "configuratorId_example" // string | 
    body := *openapiclient.NewProductconfiguratorstepCreateRequest() // ProductconfiguratorstepCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductConfiguratorApi.ProductConfiguratorCreateStep(context.Background(), tenantId, configuratorId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductConfiguratorApi.ProductConfiguratorCreateStep``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorCreateStep`: ProductconfiguratorstepEntity
    fmt.Fprintf(os.Stdout, "Response from `ProductConfiguratorApi.ProductConfiguratorCreateStep`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**configuratorId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductConfiguratorCreateStepRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**ProductconfiguratorstepCreateRequest**](ProductconfiguratorstepCreateRequest.md) |  | 

### Return type

[**ProductconfiguratorstepEntity**](ProductconfiguratorstepEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorDeleteConfigurator

> map[string]interface{} ProductConfiguratorDeleteConfigurator(ctx, tenantId, configuratorId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenantId := "tenantId_example" // string | 
    configuratorId := "configuratorId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductConfiguratorApi.ProductConfiguratorDeleteConfigurator(context.Background(), tenantId, configuratorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductConfiguratorApi.ProductConfiguratorDeleteConfigurator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorDeleteConfigurator`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProductConfiguratorApi.ProductConfiguratorDeleteConfigurator`: %v\n", resp)
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

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorDeleteDependency

> map[string]interface{} ProductConfiguratorDeleteDependency(ctx, tenantId, dependencyId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenantId := "tenantId_example" // string | 
    dependencyId := "dependencyId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductConfiguratorApi.ProductConfiguratorDeleteDependency(context.Background(), tenantId, dependencyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductConfiguratorApi.ProductConfiguratorDeleteDependency``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorDeleteDependency`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProductConfiguratorApi.ProductConfiguratorDeleteDependency`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**dependencyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductConfiguratorDeleteDependencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorDeleteMatrix

> map[string]interface{} ProductConfiguratorDeleteMatrix(ctx, tenantId, matrixId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenantId := "tenantId_example" // string | 
    matrixId := "matrixId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductConfiguratorApi.ProductConfiguratorDeleteMatrix(context.Background(), tenantId, matrixId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductConfiguratorApi.ProductConfiguratorDeleteMatrix``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorDeleteMatrix`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProductConfiguratorApi.ProductConfiguratorDeleteMatrix`: %v\n", resp)
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

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorDeleteOption

> map[string]interface{} ProductConfiguratorDeleteOption(ctx, tenantId, optionId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenantId := "tenantId_example" // string | 
    optionId := "optionId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductConfiguratorApi.ProductConfiguratorDeleteOption(context.Background(), tenantId, optionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductConfiguratorApi.ProductConfiguratorDeleteOption``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorDeleteOption`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProductConfiguratorApi.ProductConfiguratorDeleteOption`: %v\n", resp)
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

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorDeleteStep

> map[string]interface{} ProductConfiguratorDeleteStep(ctx, tenantId, stepId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenantId := "tenantId_example" // string | 
    stepId := "stepId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductConfiguratorApi.ProductConfiguratorDeleteStep(context.Background(), tenantId, stepId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductConfiguratorApi.ProductConfiguratorDeleteStep``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorDeleteStep`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProductConfiguratorApi.ProductConfiguratorDeleteStep`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**stepId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductConfiguratorDeleteStepRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorGetAvailableConfiguration

> ConfigurationGetAvailableConfigurationResponse ProductConfiguratorGetAvailableConfiguration(ctx, tenantId, productId).ConfiguratorId(configuratorId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenantId := "tenantId_example" // string | 
    productId := "productId_example" // string | 
    configuratorId := "configuratorId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductConfiguratorApi.ProductConfiguratorGetAvailableConfiguration(context.Background(), tenantId, productId).ConfiguratorId(configuratorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductConfiguratorApi.ProductConfiguratorGetAvailableConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorGetAvailableConfiguration`: ConfigurationGetAvailableConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductConfiguratorApi.ProductConfiguratorGetAvailableConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductConfiguratorGetAvailableConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **configuratorId** | **string** |  | 

### Return type

[**ConfigurationGetAvailableConfigurationResponse**](ConfigurationGetAvailableConfigurationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorGetAvailableConfiguration2

> ConfigurationGetAvailableConfigurationResponse ProductConfiguratorGetAvailableConfiguration2(ctx, tenantId, productId).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenantId := "tenantId_example" // string | 
    productId := "productId_example" // string | 
    body := *openapiclient.NewConfigurationGetAvailableConfigurationRequest() // ConfigurationGetAvailableConfigurationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductConfiguratorApi.ProductConfiguratorGetAvailableConfiguration2(context.Background(), tenantId, productId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductConfiguratorApi.ProductConfiguratorGetAvailableConfiguration2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorGetAvailableConfiguration2`: ConfigurationGetAvailableConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductConfiguratorApi.ProductConfiguratorGetAvailableConfiguration2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductConfiguratorGetAvailableConfiguration2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**ConfigurationGetAvailableConfigurationRequest**](ConfigurationGetAvailableConfigurationRequest.md) |  | 

### Return type

[**ConfigurationGetAvailableConfigurationResponse**](ConfigurationGetAvailableConfigurationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorGetConfigurationFromSelections

> ConfigurationGetConfigurationFromSelectionsResponse ProductConfiguratorGetConfigurationFromSelections(ctx, tenantId, productId).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenantId := "tenantId_example" // string | 
    productId := "productId_example" // string | 
    body := *openapiclient.NewConfigurationGetConfigurationFromSelectionsRequest() // ConfigurationGetConfigurationFromSelectionsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductConfiguratorApi.ProductConfiguratorGetConfigurationFromSelections(context.Background(), tenantId, productId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductConfiguratorApi.ProductConfiguratorGetConfigurationFromSelections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorGetConfigurationFromSelections`: ConfigurationGetConfigurationFromSelectionsResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductConfiguratorApi.ProductConfiguratorGetConfigurationFromSelections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductConfiguratorGetConfigurationFromSelectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**ConfigurationGetConfigurationFromSelectionsRequest**](ConfigurationGetConfigurationFromSelectionsRequest.md) |  | 

### Return type

[**ConfigurationGetConfigurationFromSelectionsResponse**](ConfigurationGetConfigurationFromSelectionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorGetConfiguratorByProductId

> ProductconfiguratorconfiguratorEntity ProductConfiguratorGetConfiguratorByProductId(ctx, tenantId, productId).Status(status).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenantId := "tenantId_example" // string | 
    productId := "productId_example" // string | 
    status := "status_example" // string |  (optional) (default to "UNKNOWN")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductConfiguratorApi.ProductConfiguratorGetConfiguratorByProductId(context.Background(), tenantId, productId).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductConfiguratorApi.ProductConfiguratorGetConfiguratorByProductId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorGetConfiguratorByProductId`: ProductconfiguratorconfiguratorEntity
    fmt.Fprintf(os.Stdout, "Response from `ProductConfiguratorApi.ProductConfiguratorGetConfiguratorByProductId`: %v\n", resp)
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

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorGetConfiguratorByProductId2

> ProductconfiguratorconfiguratorEntity ProductConfiguratorGetConfiguratorByProductId2(ctx, tenantId, productId, status).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenantId := "tenantId_example" // string | 
    productId := "productId_example" // string | 
    status := "status_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductConfiguratorApi.ProductConfiguratorGetConfiguratorByProductId2(context.Background(), tenantId, productId, status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductConfiguratorApi.ProductConfiguratorGetConfiguratorByProductId2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorGetConfiguratorByProductId2`: ProductconfiguratorconfiguratorEntity
    fmt.Fprintf(os.Stdout, "Response from `ProductConfiguratorApi.ProductConfiguratorGetConfiguratorByProductId2`: %v\n", resp)
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

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorGetMatrix

> ProductconfiguratormatrixEntity ProductConfiguratorGetMatrix(ctx, tenantId, matrixId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenantId := "tenantId_example" // string | 
    matrixId := "matrixId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductConfiguratorApi.ProductConfiguratorGetMatrix(context.Background(), tenantId, matrixId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductConfiguratorApi.ProductConfiguratorGetMatrix``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorGetMatrix`: ProductconfiguratormatrixEntity
    fmt.Fprintf(os.Stdout, "Response from `ProductConfiguratorApi.ProductConfiguratorGetMatrix`: %v\n", resp)
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

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorListConfigurators

> ConfiguratorListResponse ProductConfiguratorListConfigurators(ctx, tenantId, productId, pageSize).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenantId := "tenantId_example" // string | 
    productId := "productId_example" // string | 
    pageSize := int64(789) // int64 | 
    body := *openapiclient.NewConfiguratorListRequest() // ConfiguratorListRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductConfiguratorApi.ProductConfiguratorListConfigurators(context.Background(), tenantId, productId, pageSize).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductConfiguratorApi.ProductConfiguratorListConfigurators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorListConfigurators`: ConfiguratorListResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductConfiguratorApi.ProductConfiguratorListConfigurators`: %v\n", resp)
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



 **body** | [**ConfiguratorListRequest**](ConfiguratorListRequest.md) |  | 

### Return type

[**ConfiguratorListResponse**](ConfiguratorListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorListDependencies

> DependencyListDependenciesResponse ProductConfiguratorListDependencies(ctx, tenantId, pageSize).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenantId := "tenantId_example" // string | 
    pageSize := int64(789) // int64 | 
    body := *openapiclient.NewDependencyListDependenciesRequest() // DependencyListDependenciesRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductConfiguratorApi.ProductConfiguratorListDependencies(context.Background(), tenantId, pageSize).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductConfiguratorApi.ProductConfiguratorListDependencies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorListDependencies`: DependencyListDependenciesResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductConfiguratorApi.ProductConfiguratorListDependencies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**pageSize** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductConfiguratorListDependenciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**DependencyListDependenciesRequest**](DependencyListDependenciesRequest.md) |  | 

### Return type

[**DependencyListDependenciesResponse**](DependencyListDependenciesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorListMatrices

> MatrixListMatricesResponse ProductConfiguratorListMatrices(ctx, tenantId, configuratorId, pageSize).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenantId := "tenantId_example" // string | 
    configuratorId := "configuratorId_example" // string | 
    pageSize := "pageSize_example" // string | 
    body := *openapiclient.NewMatrixListMatricesRequest() // MatrixListMatricesRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductConfiguratorApi.ProductConfiguratorListMatrices(context.Background(), tenantId, configuratorId, pageSize).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductConfiguratorApi.ProductConfiguratorListMatrices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorListMatrices`: MatrixListMatricesResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductConfiguratorApi.ProductConfiguratorListMatrices`: %v\n", resp)
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



 **body** | [**MatrixListMatricesRequest**](MatrixListMatricesRequest.md) |  | 

### Return type

[**MatrixListMatricesResponse**](MatrixListMatricesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorListOptions

> OptionListOptionsResponse ProductConfiguratorListOptions(ctx, tenantId, stepId, pageSize).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenantId := "tenantId_example" // string | 
    stepId := "stepId_example" // string | 
    pageSize := int64(789) // int64 | 
    body := *openapiclient.NewOptionListOptionsRequest() // OptionListOptionsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductConfiguratorApi.ProductConfiguratorListOptions(context.Background(), tenantId, stepId, pageSize).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductConfiguratorApi.ProductConfiguratorListOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorListOptions`: OptionListOptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductConfiguratorApi.ProductConfiguratorListOptions`: %v\n", resp)
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



 **body** | [**OptionListOptionsRequest**](OptionListOptionsRequest.md) |  | 

### Return type

[**OptionListOptionsResponse**](OptionListOptionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorListProperties

> PropertyListPropertiesResponse ProductConfiguratorListProperties(ctx, tenantId, matrixId, pageSize).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenantId := "tenantId_example" // string | 
    matrixId := "matrixId_example" // string | 
    pageSize := "pageSize_example" // string | 
    body := *openapiclient.NewPropertyListPropertiesRequest() // PropertyListPropertiesRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductConfiguratorApi.ProductConfiguratorListProperties(context.Background(), tenantId, matrixId, pageSize).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductConfiguratorApi.ProductConfiguratorListProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorListProperties`: PropertyListPropertiesResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductConfiguratorApi.ProductConfiguratorListProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**matrixId** | **string** |  | 
**pageSize** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductConfiguratorListPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**PropertyListPropertiesRequest**](PropertyListPropertiesRequest.md) |  | 

### Return type

[**PropertyListPropertiesResponse**](PropertyListPropertiesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorRemovePricelistFromMatrix

> ProductconfiguratormatrixEntity ProductConfiguratorRemovePricelistFromMatrix(ctx, tenantId, matrixId, pricelistGrn).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenantId := "tenantId_example" // string | 
    matrixId := "matrixId_example" // string | 
    pricelistGrn := "pricelistGrn_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductConfiguratorApi.ProductConfiguratorRemovePricelistFromMatrix(context.Background(), tenantId, matrixId, pricelistGrn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductConfiguratorApi.ProductConfiguratorRemovePricelistFromMatrix``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorRemovePricelistFromMatrix`: ProductconfiguratormatrixEntity
    fmt.Fprintf(os.Stdout, "Response from `ProductConfiguratorApi.ProductConfiguratorRemovePricelistFromMatrix`: %v\n", resp)
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

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorUpdateConfigurator

> ProductconfiguratorconfiguratorEntity ProductConfiguratorUpdateConfigurator(ctx, tenantId, configuratorId).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenantId := "tenantId_example" // string | 
    configuratorId := "configuratorId_example" // string | 
    body := *openapiclient.NewProductconfiguratorconfiguratorUpdateRequest() // ProductconfiguratorconfiguratorUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductConfiguratorApi.ProductConfiguratorUpdateConfigurator(context.Background(), tenantId, configuratorId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductConfiguratorApi.ProductConfiguratorUpdateConfigurator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorUpdateConfigurator`: ProductconfiguratorconfiguratorEntity
    fmt.Fprintf(os.Stdout, "Response from `ProductConfiguratorApi.ProductConfiguratorUpdateConfigurator`: %v\n", resp)
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


 **body** | [**ProductconfiguratorconfiguratorUpdateRequest**](ProductconfiguratorconfiguratorUpdateRequest.md) |  | 

### Return type

[**ProductconfiguratorconfiguratorEntity**](ProductconfiguratorconfiguratorEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorUpdateDependency

> ProductconfiguratordependencyEntity ProductConfiguratorUpdateDependency(ctx, tenantId, dependencyId).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenantId := "tenantId_example" // string | 
    dependencyId := "dependencyId_example" // string | 
    body := *openapiclient.NewProductconfiguratordependencyUpdateRequest() // ProductconfiguratordependencyUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductConfiguratorApi.ProductConfiguratorUpdateDependency(context.Background(), tenantId, dependencyId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductConfiguratorApi.ProductConfiguratorUpdateDependency``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorUpdateDependency`: ProductconfiguratordependencyEntity
    fmt.Fprintf(os.Stdout, "Response from `ProductConfiguratorApi.ProductConfiguratorUpdateDependency`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**dependencyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductConfiguratorUpdateDependencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**ProductconfiguratordependencyUpdateRequest**](ProductconfiguratordependencyUpdateRequest.md) |  | 

### Return type

[**ProductconfiguratordependencyEntity**](ProductconfiguratordependencyEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorUpdateMatrix

> ProductconfiguratormatrixEntity ProductConfiguratorUpdateMatrix(ctx, tenantId, matrixId).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenantId := "tenantId_example" // string | 
    matrixId := "matrixId_example" // string | 
    body := *openapiclient.NewProductconfiguratormatrixUpdateRequest() // ProductconfiguratormatrixUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductConfiguratorApi.ProductConfiguratorUpdateMatrix(context.Background(), tenantId, matrixId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductConfiguratorApi.ProductConfiguratorUpdateMatrix``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorUpdateMatrix`: ProductconfiguratormatrixEntity
    fmt.Fprintf(os.Stdout, "Response from `ProductConfiguratorApi.ProductConfiguratorUpdateMatrix`: %v\n", resp)
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


 **body** | [**ProductconfiguratormatrixUpdateRequest**](ProductconfiguratormatrixUpdateRequest.md) |  | 

### Return type

[**ProductconfiguratormatrixEntity**](ProductconfiguratormatrixEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorUpdateOption

> ProductconfiguratoroptionEntity ProductConfiguratorUpdateOption(ctx, tenantId, optionId).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenantId := "tenantId_example" // string | 
    optionId := "optionId_example" // string | 
    body := *openapiclient.NewProductconfiguratoroptionUpdateRequest() // ProductconfiguratoroptionUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductConfiguratorApi.ProductConfiguratorUpdateOption(context.Background(), tenantId, optionId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductConfiguratorApi.ProductConfiguratorUpdateOption``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorUpdateOption`: ProductconfiguratoroptionEntity
    fmt.Fprintf(os.Stdout, "Response from `ProductConfiguratorApi.ProductConfiguratorUpdateOption`: %v\n", resp)
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


 **body** | [**ProductconfiguratoroptionUpdateRequest**](ProductconfiguratoroptionUpdateRequest.md) |  | 

### Return type

[**ProductconfiguratoroptionEntity**](ProductconfiguratoroptionEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorUpdateProperty

> ProductconfiguratorpropertyEntity ProductConfiguratorUpdateProperty(ctx, tenantId, propertyId).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenantId := "tenantId_example" // string | 
    propertyId := "propertyId_example" // string | 
    body := *openapiclient.NewProductconfiguratorpropertyUpdateRequest() // ProductconfiguratorpropertyUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductConfiguratorApi.ProductConfiguratorUpdateProperty(context.Background(), tenantId, propertyId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductConfiguratorApi.ProductConfiguratorUpdateProperty``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorUpdateProperty`: ProductconfiguratorpropertyEntity
    fmt.Fprintf(os.Stdout, "Response from `ProductConfiguratorApi.ProductConfiguratorUpdateProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**propertyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductConfiguratorUpdatePropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**ProductconfiguratorpropertyUpdateRequest**](ProductconfiguratorpropertyUpdateRequest.md) |  | 

### Return type

[**ProductconfiguratorpropertyEntity**](ProductconfiguratorpropertyEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorUpdateStep

> ProductconfiguratorstepEntity ProductConfiguratorUpdateStep(ctx, tenantId, stepId).Body(body).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenantId := "tenantId_example" // string | 
    stepId := "stepId_example" // string | 
    body := *openapiclient.NewProductconfiguratorstepUpdateRequest() // ProductconfiguratorstepUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductConfiguratorApi.ProductConfiguratorUpdateStep(context.Background(), tenantId, stepId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductConfiguratorApi.ProductConfiguratorUpdateStep``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductConfiguratorUpdateStep`: ProductconfiguratorstepEntity
    fmt.Fprintf(os.Stdout, "Response from `ProductConfiguratorApi.ProductConfiguratorUpdateStep`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**stepId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductConfiguratorUpdateStepRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**ProductconfiguratorstepUpdateRequest**](ProductconfiguratorstepUpdateRequest.md) |  | 

### Return type

[**ProductconfiguratorstepEntity**](ProductconfiguratorstepEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

