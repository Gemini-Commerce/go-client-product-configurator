# GeminiCommerce\ProductConfigurator\ConfigurationAPI

All URIs are relative to *https://product-configurator.api.gogemini.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProductConfiguratorGetAvailableConfiguration**](ConfigurationAPI.md#ProductConfiguratorGetAvailableConfiguration) | **Get** /v1/{tenantId}/product/{productId}/configuration | Get Available Configuration
[**ProductConfiguratorGetAvailableConfiguration2**](ConfigurationAPI.md#ProductConfiguratorGetAvailableConfiguration2) | **Post** /v1/{tenantId}/product/{productId}/configuration | Get Available Configuration
[**ProductConfiguratorGetConfigurationFromSelections**](ConfigurationAPI.md#ProductConfiguratorGetConfigurationFromSelections) | **Post** /v1/{tenantId}/product/{productId}/configuration-from-selections | Get Configuration from Selections



## ProductConfiguratorGetAvailableConfiguration

> ConfigurationGetAvailableConfigurationResponse ProductConfiguratorGetAvailableConfiguration(ctx, tenantId, productId).ConfiguratorId(configuratorId).Execute()

Get Available Configuration



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
	productId := "productId_example" // string | 
	configuratorId := "configuratorId_example" // string | If not set, the service returns the active configurator (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.ProductConfiguratorGetAvailableConfiguration(context.Background(), tenantId, productId).ConfiguratorId(configuratorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.ProductConfiguratorGetAvailableConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductConfiguratorGetAvailableConfiguration`: ConfigurationGetAvailableConfigurationResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.ProductConfiguratorGetAvailableConfiguration`: %v\n", resp)
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


 **configuratorId** | **string** | If not set, the service returns the active configurator | 

### Return type

[**ConfigurationGetAvailableConfigurationResponse**](ConfigurationGetAvailableConfigurationResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization), [APIAuthorization](../README.md#APIAuthorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorGetAvailableConfiguration2

> ConfigurationGetAvailableConfigurationResponse ProductConfiguratorGetAvailableConfiguration2(ctx, tenantId, productId).Body(body).Execute()

Get Available Configuration



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
	productId := "productId_example" // string | 
	body := *openapiclient.NewProductConfiguratorGetAvailableConfiguration2Request() // ProductConfiguratorGetAvailableConfiguration2Request | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.ProductConfiguratorGetAvailableConfiguration2(context.Background(), tenantId, productId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.ProductConfiguratorGetAvailableConfiguration2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductConfiguratorGetAvailableConfiguration2`: ConfigurationGetAvailableConfigurationResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.ProductConfiguratorGetAvailableConfiguration2`: %v\n", resp)
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


 **body** | [**ProductConfiguratorGetAvailableConfiguration2Request**](ProductConfiguratorGetAvailableConfiguration2Request.md) |  | 

### Return type

[**ConfigurationGetAvailableConfigurationResponse**](ConfigurationGetAvailableConfigurationResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization), [APIAuthorization](../README.md#APIAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorGetConfigurationFromSelections

> ConfigurationGetConfigurationFromSelectionsResponse ProductConfiguratorGetConfigurationFromSelections(ctx, tenantId, productId).Body(body).Execute()

Get Configuration from Selections



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
	productId := "productId_example" // string | 
	body := *openapiclient.NewProductConfiguratorGetConfigurationFromSelectionsRequest() // ProductConfiguratorGetConfigurationFromSelectionsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.ProductConfiguratorGetConfigurationFromSelections(context.Background(), tenantId, productId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.ProductConfiguratorGetConfigurationFromSelections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductConfiguratorGetConfigurationFromSelections`: ConfigurationGetConfigurationFromSelectionsResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.ProductConfiguratorGetConfigurationFromSelections`: %v\n", resp)
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


 **body** | [**ProductConfiguratorGetConfigurationFromSelectionsRequest**](ProductConfiguratorGetConfigurationFromSelectionsRequest.md) |  | 

### Return type

[**ConfigurationGetConfigurationFromSelectionsResponse**](ConfigurationGetConfigurationFromSelectionsResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization), [APIAuthorization](../README.md#APIAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

