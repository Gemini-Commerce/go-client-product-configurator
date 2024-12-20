/*
Product Configurator Service

## Introduction  This comprehensive guide will equip you with the knowledge to integrate and leverage our Product Configurator Service in your applications.  ## Quick Start  Get up and running in no time! Follow these steps to kickstart your integration:  1. **Authentication:** Obtain your integration JWT to authenticate your requests. 2. **Client Libraries:** Explore our GitHub repositories to grab client libraries in your preferred programming language. 3. **API Overview:** Familiarize yourself with our RESTful API using the OpenAPI specification.  ## Integration  ### API Overview  Our RESTful API is the gateway to unlocking the full potential of Product Configurator. Check out the detailed [API Reference](/docs/category/configurator) for a granular understanding of each endpoint and request/response format.  ### Client Libraries  To expedite your integration process, we provide client libraries for various programming languages. Find the one that suits your stack in our [GitHub repositories](https://github.com/Gemini-Commerce).  ### Authentication  Security is paramount. Learn how to authenticate your requests using JWT. This ensures a secure and reliable connection between your application and Product Configurator.  ## Configuration Management  ### Configurator Lifecycle  Understand the lifecycle of configurators, from draft to active and deleted. This flexibility allows you to manage configurations at your own pace.  ### Steps and Options  Configure product steps with ease and define options effortlessly. Explore the power of dependencies to create dynamic and intuitive configurations.  ### Matrices  Delve into matrices—your secret weapon. Explore price and weight matrices, and learn how configured steps influence properties and pricing.  ### Price Management  Unleash dynamic pricing with our versatile price matrices. From fixed prices to incremental structures, adapt to diverse pricing models effortlessly.  ## Security  Your data is in safe hands. Discover how Product Configurator ensures security through JWT authentication, safeguarding your sensitive information.  ## Backward Compatibility  Stay ahead of the curve. Learn about our versioning strategy, providing backward compatibility while allowing our service to evolve seamlessly.  ## Developer Support  Have questions? Need assistance? Write to us at [info@gemini-commerce.com](mailto:info@gemini-commerce.com) and we will get back to you.

API version: v1
Contact: info@gemini-commerce.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package productconfigurator

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type ProductConfiguratorAPI interface {

	/*
		ProductConfiguratorGetProperty Method for ProductConfiguratorGetProperty

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param tenantId
		@param propertyId
		@return ApiProductConfiguratorGetPropertyRequest
	*/
	ProductConfiguratorGetProperty(ctx context.Context, tenantId string, propertyId string) ApiProductConfiguratorGetPropertyRequest

	// ProductConfiguratorGetPropertyExecute executes the request
	//  @return ProductconfiguratorpropertyEntity
	ProductConfiguratorGetPropertyExecute(r ApiProductConfiguratorGetPropertyRequest) (*ProductconfiguratorpropertyEntity, *http.Response, error)

	/*
		ProductConfiguratorListPropertiesByConfiguration Method for ProductConfiguratorListPropertiesByConfiguration

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param tenantId
		@param configuratorId
		@param pageSize
		@return ApiProductConfiguratorListPropertiesByConfigurationRequest
	*/
	ProductConfiguratorListPropertiesByConfiguration(ctx context.Context, tenantId string, configuratorId string, pageSize string) ApiProductConfiguratorListPropertiesByConfigurationRequest

	// ProductConfiguratorListPropertiesByConfigurationExecute executes the request
	//  @return PropertyListPropertiesByConfigurationResponse
	ProductConfiguratorListPropertiesByConfigurationExecute(r ApiProductConfiguratorListPropertiesByConfigurationRequest) (*PropertyListPropertiesByConfigurationResponse, *http.Response, error)
}

// ProductConfiguratorAPIService ProductConfiguratorAPI service
type ProductConfiguratorAPIService service

type ApiProductConfiguratorGetPropertyRequest struct {
	ctx        context.Context
	ApiService ProductConfiguratorAPI
	tenantId   string
	propertyId string
}

func (r ApiProductConfiguratorGetPropertyRequest) Execute() (*ProductconfiguratorpropertyEntity, *http.Response, error) {
	return r.ApiService.ProductConfiguratorGetPropertyExecute(r)
}

/*
ProductConfiguratorGetProperty Method for ProductConfiguratorGetProperty

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param tenantId
	@param propertyId
	@return ApiProductConfiguratorGetPropertyRequest
*/
func (a *ProductConfiguratorAPIService) ProductConfiguratorGetProperty(ctx context.Context, tenantId string, propertyId string) ApiProductConfiguratorGetPropertyRequest {
	return ApiProductConfiguratorGetPropertyRequest{
		ApiService: a,
		ctx:        ctx,
		tenantId:   tenantId,
		propertyId: propertyId,
	}
}

// Execute executes the request
//
//	@return ProductconfiguratorpropertyEntity
func (a *ProductConfiguratorAPIService) ProductConfiguratorGetPropertyExecute(r ApiProductConfiguratorGetPropertyRequest) (*ProductconfiguratorpropertyEntity, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ProductconfiguratorpropertyEntity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductConfiguratorAPIService.ProductConfiguratorGetProperty")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/{tenantId}/property/{propertyId}"
	localVarPath = strings.Replace(localVarPath, "{"+"tenantId"+"}", url.PathEscape(parameterValueToString(r.tenantId, "tenantId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"propertyId"+"}", url.PathEscape(parameterValueToString(r.propertyId, "propertyId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v GooglerpcStatus
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiProductConfiguratorListPropertiesByConfigurationRequest struct {
	ctx            context.Context
	ApiService     ProductConfiguratorAPI
	tenantId       string
	configuratorId string
	pageSize       string
	body           *ProductConfiguratorListPropertiesByConfigurationRequest
}

func (r ApiProductConfiguratorListPropertiesByConfigurationRequest) Body(body ProductConfiguratorListPropertiesByConfigurationRequest) ApiProductConfiguratorListPropertiesByConfigurationRequest {
	r.body = &body
	return r
}

func (r ApiProductConfiguratorListPropertiesByConfigurationRequest) Execute() (*PropertyListPropertiesByConfigurationResponse, *http.Response, error) {
	return r.ApiService.ProductConfiguratorListPropertiesByConfigurationExecute(r)
}

/*
ProductConfiguratorListPropertiesByConfiguration Method for ProductConfiguratorListPropertiesByConfiguration

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param tenantId
	@param configuratorId
	@param pageSize
	@return ApiProductConfiguratorListPropertiesByConfigurationRequest
*/
func (a *ProductConfiguratorAPIService) ProductConfiguratorListPropertiesByConfiguration(ctx context.Context, tenantId string, configuratorId string, pageSize string) ApiProductConfiguratorListPropertiesByConfigurationRequest {
	return ApiProductConfiguratorListPropertiesByConfigurationRequest{
		ApiService:     a,
		ctx:            ctx,
		tenantId:       tenantId,
		configuratorId: configuratorId,
		pageSize:       pageSize,
	}
}

// Execute executes the request
//
//	@return PropertyListPropertiesByConfigurationResponse
func (a *ProductConfiguratorAPIService) ProductConfiguratorListPropertiesByConfigurationExecute(r ApiProductConfiguratorListPropertiesByConfigurationRequest) (*PropertyListPropertiesByConfigurationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PropertyListPropertiesByConfigurationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductConfiguratorAPIService.ProductConfiguratorListPropertiesByConfiguration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/{tenantId}/configurator/{configuratorId}/page-size/{pageSize}/properties"
	localVarPath = strings.Replace(localVarPath, "{"+"tenantId"+"}", url.PathEscape(parameterValueToString(r.tenantId, "tenantId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"configuratorId"+"}", url.PathEscape(parameterValueToString(r.configuratorId, "configuratorId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pageSize"+"}", url.PathEscape(parameterValueToString(r.pageSize, "pageSize")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v GooglerpcStatus
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
