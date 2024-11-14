# \PropertyAPI

All URIs are relative to *https://product-configurator.api.gogemini.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProductConfiguratorBulkCreateProperties**](PropertyAPI.md#ProductConfiguratorBulkCreateProperties) | **Post** /v1/{tenantId}/property/create/bulk | Bulk Create Properties
[**ProductConfiguratorBulkUpdateProperties**](PropertyAPI.md#ProductConfiguratorBulkUpdateProperties) | **Put** /v1/{tenantId}/properties/bulk | Bulk Update Properties
[**ProductConfiguratorCreateProperty**](PropertyAPI.md#ProductConfiguratorCreateProperty) | **Post** /v1/{tenantId}/property/create | Create Property
[**ProductConfiguratorListProperties**](PropertyAPI.md#ProductConfiguratorListProperties) | **Post** /v1/{tenantId}/matrix/{matrixId}/page-size/{pageSize}/properties | List Properties
[**ProductConfiguratorUpdateProperty**](PropertyAPI.md#ProductConfiguratorUpdateProperty) | **Put** /v1/{tenantId}/property/{propertyId} | Update Property



## ProductConfiguratorBulkCreateProperties

> ProductconfiguratorpropertyBulkCreateResponse ProductConfiguratorBulkCreateProperties(ctx, tenantId).Body(body).Execute()

Bulk Create Properties



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-product-configurator"
)

func main() {
	tenantId := "tenantId_example" // string | 
	body := *openapiclient.NewProductConfiguratorBulkCreatePropertiesRequest() // ProductConfiguratorBulkCreatePropertiesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PropertyAPI.ProductConfiguratorBulkCreateProperties(context.Background(), tenantId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PropertyAPI.ProductConfiguratorBulkCreateProperties``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductConfiguratorBulkCreateProperties`: ProductconfiguratorpropertyBulkCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `PropertyAPI.ProductConfiguratorBulkCreateProperties`: %v\n", resp)
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

 **body** | [**ProductConfiguratorBulkCreatePropertiesRequest**](ProductConfiguratorBulkCreatePropertiesRequest.md) |  | 

### Return type

[**ProductconfiguratorpropertyBulkCreateResponse**](ProductconfiguratorpropertyBulkCreateResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization), [APIAuthorization](../README.md#APIAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorBulkUpdateProperties

> ProductconfiguratorpropertyBulkUpdateResponse ProductConfiguratorBulkUpdateProperties(ctx, tenantId).Body(body).Execute()

Bulk Update Properties



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-product-configurator"
)

func main() {
	tenantId := "tenantId_example" // string | 
	body := *openapiclient.NewProductConfiguratorBulkUpdatePropertiesRequest() // ProductConfiguratorBulkUpdatePropertiesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PropertyAPI.ProductConfiguratorBulkUpdateProperties(context.Background(), tenantId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PropertyAPI.ProductConfiguratorBulkUpdateProperties``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductConfiguratorBulkUpdateProperties`: ProductconfiguratorpropertyBulkUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `PropertyAPI.ProductConfiguratorBulkUpdateProperties`: %v\n", resp)
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

 **body** | [**ProductConfiguratorBulkUpdatePropertiesRequest**](ProductConfiguratorBulkUpdatePropertiesRequest.md) |  | 

### Return type

[**ProductconfiguratorpropertyBulkUpdateResponse**](ProductconfiguratorpropertyBulkUpdateResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization), [APIAuthorization](../README.md#APIAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorCreateProperty

> ProductconfiguratorpropertyEntity ProductConfiguratorCreateProperty(ctx, tenantId).Body(body).Execute()

Create Property



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-product-configurator"
)

func main() {
	tenantId := "tenantId_example" // string | 
	body := *openapiclient.NewProductConfiguratorCreatePropertyRequest() // ProductConfiguratorCreatePropertyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PropertyAPI.ProductConfiguratorCreateProperty(context.Background(), tenantId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PropertyAPI.ProductConfiguratorCreateProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductConfiguratorCreateProperty`: ProductconfiguratorpropertyEntity
	fmt.Fprintf(os.Stdout, "Response from `PropertyAPI.ProductConfiguratorCreateProperty`: %v\n", resp)
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

 **body** | [**ProductConfiguratorCreatePropertyRequest**](ProductConfiguratorCreatePropertyRequest.md) |  | 

### Return type

[**ProductconfiguratorpropertyEntity**](ProductconfiguratorpropertyEntity.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization), [APIAuthorization](../README.md#APIAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorListProperties

> PropertyListPropertiesResponse ProductConfiguratorListProperties(ctx, tenantId, matrixId, pageSize).Body(body).Execute()

List Properties



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-product-configurator"
)

func main() {
	tenantId := "tenantId_example" // string | 
	matrixId := "matrixId_example" // string | 
	pageSize := "pageSize_example" // string | 
	body := *openapiclient.NewProductConfiguratorListPropertiesRequest() // ProductConfiguratorListPropertiesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PropertyAPI.ProductConfiguratorListProperties(context.Background(), tenantId, matrixId, pageSize).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PropertyAPI.ProductConfiguratorListProperties``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductConfiguratorListProperties`: PropertyListPropertiesResponse
	fmt.Fprintf(os.Stdout, "Response from `PropertyAPI.ProductConfiguratorListProperties`: %v\n", resp)
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



 **body** | [**ProductConfiguratorListPropertiesRequest**](ProductConfiguratorListPropertiesRequest.md) |  | 

### Return type

[**PropertyListPropertiesResponse**](PropertyListPropertiesResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization), [APIAuthorization](../README.md#APIAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorUpdateProperty

> ProductconfiguratorpropertyEntity ProductConfiguratorUpdateProperty(ctx, tenantId, propertyId).Body(body).Execute()

Update Property



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Gemini-Commerce/go-client-product-configurator"
)

func main() {
	tenantId := "tenantId_example" // string | 
	propertyId := "propertyId_example" // string | 
	body := *openapiclient.NewProductConfiguratorUpdatePropertyRequest() // ProductConfiguratorUpdatePropertyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PropertyAPI.ProductConfiguratorUpdateProperty(context.Background(), tenantId, propertyId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PropertyAPI.ProductConfiguratorUpdateProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductConfiguratorUpdateProperty`: ProductconfiguratorpropertyEntity
	fmt.Fprintf(os.Stdout, "Response from `PropertyAPI.ProductConfiguratorUpdateProperty`: %v\n", resp)
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


 **body** | [**ProductConfiguratorUpdatePropertyRequest**](ProductConfiguratorUpdatePropertyRequest.md) |  | 

### Return type

[**ProductconfiguratorpropertyEntity**](ProductconfiguratorpropertyEntity.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization), [APIAuthorization](../README.md#APIAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

