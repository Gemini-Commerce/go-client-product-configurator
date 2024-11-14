# \StepAPI

All URIs are relative to *https://product-configurator.api.gogemini.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProductConfiguratorBulkCreateSteps**](StepAPI.md#ProductConfiguratorBulkCreateSteps) | **Post** /v1/{tenantId}/configurator/{configuratorId}/step/create/bulk | Bulk Create Steps
[**ProductConfiguratorBulkDeleteSteps**](StepAPI.md#ProductConfiguratorBulkDeleteSteps) | **Post** /v1/{tenantId}/step/delete/bulk | Bulk Delete Steps
[**ProductConfiguratorCopyStep**](StepAPI.md#ProductConfiguratorCopyStep) | **Post** /v1/{tenantId}/step/{sourceStepId}/copy | Copy Step
[**ProductConfiguratorCreateStep**](StepAPI.md#ProductConfiguratorCreateStep) | **Post** /v1/{tenantId}/configurator/{configuratorId}/step/create | Create Step
[**ProductConfiguratorDeleteStep**](StepAPI.md#ProductConfiguratorDeleteStep) | **Delete** /v1/{tenantId}/step/{stepId} | Delete Step
[**ProductConfiguratorUpdateStep**](StepAPI.md#ProductConfiguratorUpdateStep) | **Put** /v1/{tenantId}/step/{stepId} | Update Step



## ProductConfiguratorBulkCreateSteps

> ProductconfiguratorstepBulkCreateResponse ProductConfiguratorBulkCreateSteps(ctx, tenantId, configuratorId).Body(body).Execute()

Bulk Create Steps



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
	body := *openapiclient.NewProductConfiguratorBulkCreateStepsRequest() // ProductConfiguratorBulkCreateStepsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StepAPI.ProductConfiguratorBulkCreateSteps(context.Background(), tenantId, configuratorId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StepAPI.ProductConfiguratorBulkCreateSteps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductConfiguratorBulkCreateSteps`: ProductconfiguratorstepBulkCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `StepAPI.ProductConfiguratorBulkCreateSteps`: %v\n", resp)
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


 **body** | [**ProductConfiguratorBulkCreateStepsRequest**](ProductConfiguratorBulkCreateStepsRequest.md) |  | 

### Return type

[**ProductconfiguratorstepBulkCreateResponse**](ProductconfiguratorstepBulkCreateResponse.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization), [APIAuthorization](../README.md#APIAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorBulkDeleteSteps

> map[string]interface{} ProductConfiguratorBulkDeleteSteps(ctx, tenantId).Body(body).Execute()

Bulk Delete Steps



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
	body := *openapiclient.NewProductConfiguratorBulkDeleteStepsRequest() // ProductConfiguratorBulkDeleteStepsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StepAPI.ProductConfiguratorBulkDeleteSteps(context.Background(), tenantId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StepAPI.ProductConfiguratorBulkDeleteSteps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductConfiguratorBulkDeleteSteps`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `StepAPI.ProductConfiguratorBulkDeleteSteps`: %v\n", resp)
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

 **body** | [**ProductConfiguratorBulkDeleteStepsRequest**](ProductConfiguratorBulkDeleteStepsRequest.md) |  | 

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


## ProductConfiguratorCopyStep

> ProductconfiguratorstepEntity ProductConfiguratorCopyStep(ctx, tenantId, sourceStepId).Body(body).Execute()

Copy Step



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
	sourceStepId := "sourceStepId_example" // string | 
	body := *openapiclient.NewProductConfiguratorCopyStepRequest() // ProductConfiguratorCopyStepRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StepAPI.ProductConfiguratorCopyStep(context.Background(), tenantId, sourceStepId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StepAPI.ProductConfiguratorCopyStep``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductConfiguratorCopyStep`: ProductconfiguratorstepEntity
	fmt.Fprintf(os.Stdout, "Response from `StepAPI.ProductConfiguratorCopyStep`: %v\n", resp)
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


 **body** | [**ProductConfiguratorCopyStepRequest**](ProductConfiguratorCopyStepRequest.md) |  | 

### Return type

[**ProductconfiguratorstepEntity**](ProductconfiguratorstepEntity.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization), [APIAuthorization](../README.md#APIAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorCreateStep

> ProductconfiguratorstepEntity ProductConfiguratorCreateStep(ctx, tenantId, configuratorId).Body(body).Execute()

Create Step



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
	body := *openapiclient.NewProductConfiguratorCreateStepRequest() // ProductConfiguratorCreateStepRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StepAPI.ProductConfiguratorCreateStep(context.Background(), tenantId, configuratorId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StepAPI.ProductConfiguratorCreateStep``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductConfiguratorCreateStep`: ProductconfiguratorstepEntity
	fmt.Fprintf(os.Stdout, "Response from `StepAPI.ProductConfiguratorCreateStep`: %v\n", resp)
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


 **body** | [**ProductConfiguratorCreateStepRequest**](ProductConfiguratorCreateStepRequest.md) |  | 

### Return type

[**ProductconfiguratorstepEntity**](ProductconfiguratorstepEntity.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization), [APIAuthorization](../README.md#APIAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorDeleteStep

> GooglerpcStatus ProductConfiguratorDeleteStep(ctx, tenantId, stepId).Execute()

Delete Step



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StepAPI.ProductConfiguratorDeleteStep(context.Background(), tenantId, stepId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StepAPI.ProductConfiguratorDeleteStep``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductConfiguratorDeleteStep`: GooglerpcStatus
	fmt.Fprintf(os.Stdout, "Response from `StepAPI.ProductConfiguratorDeleteStep`: %v\n", resp)
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

[**GooglerpcStatus**](GooglerpcStatus.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization), [APIAuthorization](../README.md#APIAuthorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductConfiguratorUpdateStep

> ProductconfiguratorstepEntity ProductConfiguratorUpdateStep(ctx, tenantId, stepId).Body(body).Execute()

Update Step



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
	body := *openapiclient.NewProductConfiguratorUpdateStepRequest() // ProductConfiguratorUpdateStepRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StepAPI.ProductConfiguratorUpdateStep(context.Background(), tenantId, stepId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StepAPI.ProductConfiguratorUpdateStep``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProductConfiguratorUpdateStep`: ProductconfiguratorstepEntity
	fmt.Fprintf(os.Stdout, "Response from `StepAPI.ProductConfiguratorUpdateStep`: %v\n", resp)
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


 **body** | [**ProductConfiguratorUpdateStepRequest**](ProductConfiguratorUpdateStepRequest.md) |  | 

### Return type

[**ProductconfiguratorstepEntity**](ProductconfiguratorstepEntity.md)

### Authorization

[standardAuthorization](../README.md#standardAuthorization), [APIAuthorization](../README.md#APIAuthorization)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

