# \DependencyAPI

All URIs are relative to *https://product-configurator.api.gogemini.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProductConfiguratorCreateDependency**](DependencyAPI.md#ProductConfiguratorCreateDependency) | **Post** /v1/{tenantId}/step/{stepId}/dependency/create | Create Dependency
[**ProductConfiguratorDeleteDependency**](DependencyAPI.md#ProductConfiguratorDeleteDependency) | **Delete** /v1/{tenantId}/dependency/{dependencyId} | Delete Dependency
[**ProductConfiguratorListDependencies**](DependencyAPI.md#ProductConfiguratorListDependencies) | **Post** /v1/{tenantId}/page-size/{pageSize}/dependencies | List Dependencies
[**ProductConfiguratorUpdateDependency**](DependencyAPI.md#ProductConfiguratorUpdateDependency) | **Put** /v1/{tenantId}/dependency/{dependencyId} | Update Dependency



## ProductConfiguratorCreateDependency

> ProductconfiguratordependencyEntity ProductConfiguratorCreateDependency(ctx, tenantId, stepId).Body(body).Execute()

Create Dependency



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
	stepId := "stepId_example" // string | 
	body := *openapiclient.NewProductConfiguratorCreateDependencyRequest() // ProductConfiguratorCreateDependencyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DependencyAPI.ProductConfiguratorCreateDependency(context.Background(), tenantId, stepId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DependencyAPI.ProductConfiguratorCreateDependency``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductConfiguratorCreateDependency`: ProductconfiguratordependencyEntity
	fmt.Fprintf(os.Stdout, "Response from `DependencyAPI.ProductConfiguratorCreateDependency`: %v\n", resp)
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


 **body** | [**ProductConfiguratorCreateDependencyRequest**](ProductConfiguratorCreateDependencyRequest.md) |  | 

### Return type

[**ProductconfiguratordependencyEntity**](ProductconfiguratordependencyEntity.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization), [APIAuthorization](../README.md#APIAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorDeleteDependency

> map[string]interface{} ProductConfiguratorDeleteDependency(ctx, tenantId, dependencyId).Execute()

Delete Dependency



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
	dependencyId := "dependencyId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DependencyAPI.ProductConfiguratorDeleteDependency(context.Background(), tenantId, dependencyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DependencyAPI.ProductConfiguratorDeleteDependency``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductConfiguratorDeleteDependency`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DependencyAPI.ProductConfiguratorDeleteDependency`: %v\n", resp)
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

[standardAuthorization](../README.md#standardAuthorization), [APIAuthorization](../README.md#APIAuthorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorListDependencies

> DependencyListDependenciesResponse ProductConfiguratorListDependencies(ctx, tenantId, pageSize).Body(body).Execute()

List Dependencies



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
	pageSize := int64(789) // int64 | 
	body := *openapiclient.NewProductConfiguratorListDependenciesRequest() // ProductConfiguratorListDependenciesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DependencyAPI.ProductConfiguratorListDependencies(context.Background(), tenantId, pageSize).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DependencyAPI.ProductConfiguratorListDependencies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductConfiguratorListDependencies`: DependencyListDependenciesResponse
	fmt.Fprintf(os.Stdout, "Response from `DependencyAPI.ProductConfiguratorListDependencies`: %v\n", resp)
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


 **body** | [**ProductConfiguratorListDependenciesRequest**](ProductConfiguratorListDependenciesRequest.md) |  | 

### Return type

[**DependencyListDependenciesResponse**](DependencyListDependenciesResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization), [APIAuthorization](../README.md#APIAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorUpdateDependency

> ProductconfiguratordependencyEntity ProductConfiguratorUpdateDependency(ctx, tenantId, dependencyId).Body(body).Execute()

Update Dependency



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
	dependencyId := "dependencyId_example" // string | 
	body := *openapiclient.NewProductConfiguratorUpdateDependencyRequest() // ProductConfiguratorUpdateDependencyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DependencyAPI.ProductConfiguratorUpdateDependency(context.Background(), tenantId, dependencyId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DependencyAPI.ProductConfiguratorUpdateDependency``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductConfiguratorUpdateDependency`: ProductconfiguratordependencyEntity
	fmt.Fprintf(os.Stdout, "Response from `DependencyAPI.ProductConfiguratorUpdateDependency`: %v\n", resp)
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


 **body** | [**ProductConfiguratorUpdateDependencyRequest**](ProductConfiguratorUpdateDependencyRequest.md) |  | 

### Return type

[**ProductconfiguratordependencyEntity**](ProductconfiguratordependencyEntity.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization), [APIAuthorization](../README.md#APIAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

