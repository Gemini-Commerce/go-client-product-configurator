# \ProductConfiguratorAPI

All URIs are relative to *https://product-configurator.api.gogemini.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProductConfiguratorGetProperty**](ProductConfiguratorAPI.md#ProductConfiguratorGetProperty) | **Get** /v1/{tenantId}/property/{propertyId} | 
[**ProductConfiguratorListPropertiesByConfiguration**](ProductConfiguratorAPI.md#ProductConfiguratorListPropertiesByConfiguration) | **Post** /v1/{tenantId}/configurator/{configuratorId}/page-size/{pageSize}/properties | 



## ProductConfiguratorGetProperty

> ProductconfiguratorpropertyEntity ProductConfiguratorGetProperty(ctx, tenantId, propertyId).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductConfiguratorAPI.ProductConfiguratorGetProperty(context.Background(), tenantId, propertyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductConfiguratorAPI.ProductConfiguratorGetProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductConfiguratorGetProperty`: ProductconfiguratorpropertyEntity
	fmt.Fprintf(os.Stdout, "Response from `ProductConfiguratorAPI.ProductConfiguratorGetProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** |  | 
**propertyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductConfiguratorGetPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ProductconfiguratorpropertyEntity**](ProductconfiguratorpropertyEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorListPropertiesByConfiguration

> PropertyListPropertiesByConfigurationResponse ProductConfiguratorListPropertiesByConfiguration(ctx, tenantId, configuratorId, pageSize).Body(body).Execute()



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
	configuratorId := "configuratorId_example" // string | 
	pageSize := "pageSize_example" // string | 
	body := *openapiclient.NewProductConfiguratorListPropertiesByConfigurationRequest() // ProductConfiguratorListPropertiesByConfigurationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductConfiguratorAPI.ProductConfiguratorListPropertiesByConfiguration(context.Background(), tenantId, configuratorId, pageSize).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductConfiguratorAPI.ProductConfiguratorListPropertiesByConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductConfiguratorListPropertiesByConfiguration`: PropertyListPropertiesByConfigurationResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductConfiguratorAPI.ProductConfiguratorListPropertiesByConfiguration`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiProductConfiguratorListPropertiesByConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**ProductConfiguratorListPropertiesByConfigurationRequest**](ProductConfiguratorListPropertiesByConfigurationRequest.md) |  | 

### Return type

[**PropertyListPropertiesByConfigurationResponse**](PropertyListPropertiesByConfigurationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

