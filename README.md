# Go API client for openapi

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://product-configurator.api.gogemini.io*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ProductConfiguratorApi* | [**ProductConfiguratorAddPricelistToMatrix**](docs/ProductConfiguratorApi.md#productconfiguratoraddpricelisttomatrix) | **Post** /v1/{tenantId}/matrix/{matrixId}/pricelist/{pricelistGrn} | 
*ProductConfiguratorApi* | [**ProductConfiguratorBulkCreateOptions**](docs/ProductConfiguratorApi.md#productconfiguratorbulkcreateoptions) | **Post** /v1/{tenantId}/step/{stepId}/option/create/bulk | 
*ProductConfiguratorApi* | [**ProductConfiguratorBulkCreateProperties**](docs/ProductConfiguratorApi.md#productconfiguratorbulkcreateproperties) | **Post** /v1/{tenantId}/property/create/bulk | 
*ProductConfiguratorApi* | [**ProductConfiguratorBulkCreateSteps**](docs/ProductConfiguratorApi.md#productconfiguratorbulkcreatesteps) | **Post** /v1/{tenantId}/configurator/{configuratorId}/step/create/bulk | 
*ProductConfiguratorApi* | [**ProductConfiguratorBulkDeleteOptions**](docs/ProductConfiguratorApi.md#productconfiguratorbulkdeleteoptions) | **Post** /v1/{tenantId}/option/delete/bulk | 
*ProductConfiguratorApi* | [**ProductConfiguratorBulkDeleteSteps**](docs/ProductConfiguratorApi.md#productconfiguratorbulkdeletesteps) | **Post** /v1/{tenantId}/step/delete/bulk | 
*ProductConfiguratorApi* | [**ProductConfiguratorBulkUpdateOptions**](docs/ProductConfiguratorApi.md#productconfiguratorbulkupdateoptions) | **Put** /v1/{tenantId}/option/bulk | 
*ProductConfiguratorApi* | [**ProductConfiguratorBulkUpdateProperties**](docs/ProductConfiguratorApi.md#productconfiguratorbulkupdateproperties) | **Put** /v1/{tenantId}/properties/bulk | 
*ProductConfiguratorApi* | [**ProductConfiguratorCopyConfigurator**](docs/ProductConfiguratorApi.md#productconfiguratorcopyconfigurator) | **Post** /v1/{tenantId}/product/{sourceConfiguratorId}/copy | 
*ProductConfiguratorApi* | [**ProductConfiguratorCopyOption**](docs/ProductConfiguratorApi.md#productconfiguratorcopyoption) | **Post** /v1/{tenantId}/option/{sourceOptionId}/copy | 
*ProductConfiguratorApi* | [**ProductConfiguratorCopyStep**](docs/ProductConfiguratorApi.md#productconfiguratorcopystep) | **Post** /v1/{tenantId}/step/{sourceStepId}/copy | 
*ProductConfiguratorApi* | [**ProductConfiguratorCreateConfigurator**](docs/ProductConfiguratorApi.md#productconfiguratorcreateconfigurator) | **Post** /v1/{tenantId}/product/{productId}/create | 
*ProductConfiguratorApi* | [**ProductConfiguratorCreateDependency**](docs/ProductConfiguratorApi.md#productconfiguratorcreatedependency) | **Post** /v1/{tenantId}/step/{stepId}/dependency/create | 
*ProductConfiguratorApi* | [**ProductConfiguratorCreateMatrix**](docs/ProductConfiguratorApi.md#productconfiguratorcreatematrix) | **Post** /v1/{tenantId}/matrix/create | 
*ProductConfiguratorApi* | [**ProductConfiguratorCreateOption**](docs/ProductConfiguratorApi.md#productconfiguratorcreateoption) | **Post** /v1/{tenantId}/step/{stepId}/option/create | 
*ProductConfiguratorApi* | [**ProductConfiguratorCreateProperty**](docs/ProductConfiguratorApi.md#productconfiguratorcreateproperty) | **Post** /v1/{tenantId}/property/create | 
*ProductConfiguratorApi* | [**ProductConfiguratorCreateStep**](docs/ProductConfiguratorApi.md#productconfiguratorcreatestep) | **Post** /v1/{tenantId}/configurator/{configuratorId}/step/create | 
*ProductConfiguratorApi* | [**ProductConfiguratorDeleteConfigurator**](docs/ProductConfiguratorApi.md#productconfiguratordeleteconfigurator) | **Delete** /v1/{tenantId}/configurator/{configuratorId} | 
*ProductConfiguratorApi* | [**ProductConfiguratorDeleteDependency**](docs/ProductConfiguratorApi.md#productconfiguratordeletedependency) | **Delete** /v1/{tenantId}/dependency/{dependencyId} | 
*ProductConfiguratorApi* | [**ProductConfiguratorDeleteMatrix**](docs/ProductConfiguratorApi.md#productconfiguratordeletematrix) | **Delete** /v1/{tenantId}/matrix/{matrixId} | 
*ProductConfiguratorApi* | [**ProductConfiguratorDeleteOption**](docs/ProductConfiguratorApi.md#productconfiguratordeleteoption) | **Delete** /v1/{tenantId}/option/{optionId} | 
*ProductConfiguratorApi* | [**ProductConfiguratorDeleteStep**](docs/ProductConfiguratorApi.md#productconfiguratordeletestep) | **Delete** /v1/{tenantId}/step/{stepId} | 
*ProductConfiguratorApi* | [**ProductConfiguratorGetAvailableConfiguration**](docs/ProductConfiguratorApi.md#productconfiguratorgetavailableconfiguration) | **Get** /v1/{tenantId}/product/{productId}/configuration | 
*ProductConfiguratorApi* | [**ProductConfiguratorGetAvailableConfiguration2**](docs/ProductConfiguratorApi.md#productconfiguratorgetavailableconfiguration2) | **Post** /v1/{tenantId}/product/{productId}/configuration | 
*ProductConfiguratorApi* | [**ProductConfiguratorGetConfigurationFromSelections**](docs/ProductConfiguratorApi.md#productconfiguratorgetconfigurationfromselections) | **Post** /v1/{tenantId}/product/{productId}/configuration-from-selections | 
*ProductConfiguratorApi* | [**ProductConfiguratorGetConfiguratorByProductId**](docs/ProductConfiguratorApi.md#productconfiguratorgetconfiguratorbyproductid) | **Get** /v1/{tenantId}/product/{productId} | 
*ProductConfiguratorApi* | [**ProductConfiguratorGetConfiguratorByProductId2**](docs/ProductConfiguratorApi.md#productconfiguratorgetconfiguratorbyproductid2) | **Get** /v1/{tenantId}/product/{productId}/status/{status} | 
*ProductConfiguratorApi* | [**ProductConfiguratorGetMatrix**](docs/ProductConfiguratorApi.md#productconfiguratorgetmatrix) | **Get** /v1/{tenantId}/matrix/{matrixId} | 
*ProductConfiguratorApi* | [**ProductConfiguratorListConfigurators**](docs/ProductConfiguratorApi.md#productconfiguratorlistconfigurators) | **Post** /v1/{tenantId}/product/{productId}/page-size/{pageSize}/configurators | 
*ProductConfiguratorApi* | [**ProductConfiguratorListDependencies**](docs/ProductConfiguratorApi.md#productconfiguratorlistdependencies) | **Post** /v1/{tenantId}/page-size/{pageSize}/dependencies | 
*ProductConfiguratorApi* | [**ProductConfiguratorListMatrices**](docs/ProductConfiguratorApi.md#productconfiguratorlistmatrices) | **Post** /v1/{tenantId}/configurator/{configuratorId}/page-size/{pageSize}/matrices | 
*ProductConfiguratorApi* | [**ProductConfiguratorListOptions**](docs/ProductConfiguratorApi.md#productconfiguratorlistoptions) | **Post** /v1/{tenantId}/step/{stepId}/page-size/{pageSize}/options | 
*ProductConfiguratorApi* | [**ProductConfiguratorListProperties**](docs/ProductConfiguratorApi.md#productconfiguratorlistproperties) | **Post** /v1/{tenantId}/matrix/{matrixId}/page-size/{pageSize}/properties | 
*ProductConfiguratorApi* | [**ProductConfiguratorRemovePricelistFromMatrix**](docs/ProductConfiguratorApi.md#productconfiguratorremovepricelistfrommatrix) | **Delete** /v1/{tenantId}/matrix/{matrixId}/pricelist/{pricelistGrn} | 
*ProductConfiguratorApi* | [**ProductConfiguratorUpdateConfigurator**](docs/ProductConfiguratorApi.md#productconfiguratorupdateconfigurator) | **Put** /v1/{tenantId}/configurator/{configuratorId} | 
*ProductConfiguratorApi* | [**ProductConfiguratorUpdateDependency**](docs/ProductConfiguratorApi.md#productconfiguratorupdatedependency) | **Put** /v1/{tenantId}/dependency/{dependencyId} | 
*ProductConfiguratorApi* | [**ProductConfiguratorUpdateMatrix**](docs/ProductConfiguratorApi.md#productconfiguratorupdatematrix) | **Put** /v1/{tenantId}/matrix/{matrixId} | 
*ProductConfiguratorApi* | [**ProductConfiguratorUpdateOption**](docs/ProductConfiguratorApi.md#productconfiguratorupdateoption) | **Put** /v1/{tenantId}/option/{optionId} | 
*ProductConfiguratorApi* | [**ProductConfiguratorUpdateProperty**](docs/ProductConfiguratorApi.md#productconfiguratorupdateproperty) | **Put** /v1/{tenantId}/property/{propertyId} | 
*ProductConfiguratorApi* | [**ProductConfiguratorUpdateStep**](docs/ProductConfiguratorApi.md#productconfiguratorupdatestep) | **Put** /v1/{tenantId}/step/{stepId} | 


## Documentation For Models

 - [ConfigurationConfigurationStep](docs/ConfigurationConfigurationStep.md)
 - [ConfigurationConfigurator](docs/ConfigurationConfigurator.md)
 - [ConfigurationGetAvailableConfigurationRequest](docs/ConfigurationGetAvailableConfigurationRequest.md)
 - [ConfigurationGetAvailableConfigurationResponse](docs/ConfigurationGetAvailableConfigurationResponse.md)
 - [ConfigurationGetConfigurationFromSelectionsRequest](docs/ConfigurationGetConfigurationFromSelectionsRequest.md)
 - [ConfigurationGetConfigurationFromSelectionsResponse](docs/ConfigurationGetConfigurationFromSelectionsResponse.md)
 - [ConfigurationOptionProperty](docs/ConfigurationOptionProperty.md)
 - [ConfigurationProperty](docs/ConfigurationProperty.md)
 - [ConfigurationPropertyFilter](docs/ConfigurationPropertyFilter.md)
 - [ConfigurationPropertyFilterGenericProperty](docs/ConfigurationPropertyFilterGenericProperty.md)
 - [ConfigurationSelection](docs/ConfigurationSelection.md)
 - [ConfigurationStepOption](docs/ConfigurationStepOption.md)
 - [ConfiguratorListRequest](docs/ConfiguratorListRequest.md)
 - [ConfiguratorListResponse](docs/ConfiguratorListResponse.md)
 - [DependencyCondition](docs/DependencyCondition.md)
 - [DependencyListDependenciesRequest](docs/DependencyListDependenciesRequest.md)
 - [DependencyListDependenciesResponse](docs/DependencyListDependenciesResponse.md)
 - [GooglerpcStatus](docs/GooglerpcStatus.md)
 - [ListMatricesRequestFilter](docs/ListMatricesRequestFilter.md)
 - [LocalisationLocalizedText](docs/LocalisationLocalizedText.md)
 - [MatrixGenericType](docs/MatrixGenericType.md)
 - [MatrixListMatricesRequest](docs/MatrixListMatricesRequest.md)
 - [MatrixListMatricesResponse](docs/MatrixListMatricesResponse.md)
 - [MatrixPriceType](docs/MatrixPriceType.md)
 - [MatrixWeightType](docs/MatrixWeightType.md)
 - [OptionListOptionsRequest](docs/OptionListOptionsRequest.md)
 - [OptionListOptionsResponse](docs/OptionListOptionsResponse.md)
 - [OptionSwatch](docs/OptionSwatch.md)
 - [ProductconfiguratorMoney](docs/ProductconfiguratorMoney.md)
 - [ProductconfiguratorPropertyType](docs/ProductconfiguratorPropertyType.md)
 - [ProductconfiguratorWeightUnit](docs/ProductconfiguratorWeightUnit.md)
 - [ProductconfiguratorconfigurationOption](docs/ProductconfiguratorconfigurationOption.md)
 - [ProductconfiguratorconfigurationStep](docs/ProductconfiguratorconfigurationStep.md)
 - [ProductconfiguratorconfiguratorCopyRequest](docs/ProductconfiguratorconfiguratorCopyRequest.md)
 - [ProductconfiguratorconfiguratorCreateRequest](docs/ProductconfiguratorconfiguratorCreateRequest.md)
 - [ProductconfiguratorconfiguratorEntity](docs/ProductconfiguratorconfiguratorEntity.md)
 - [ProductconfiguratorconfiguratorStatus](docs/ProductconfiguratorconfiguratorStatus.md)
 - [ProductconfiguratorconfiguratorUpdatePayload](docs/ProductconfiguratorconfiguratorUpdatePayload.md)
 - [ProductconfiguratorconfiguratorUpdateRequest](docs/ProductconfiguratorconfiguratorUpdateRequest.md)
 - [ProductconfiguratordependencyCreateRequest](docs/ProductconfiguratordependencyCreateRequest.md)
 - [ProductconfiguratordependencyEntity](docs/ProductconfiguratordependencyEntity.md)
 - [ProductconfiguratordependencyType](docs/ProductconfiguratordependencyType.md)
 - [ProductconfiguratordependencyUpdatePayload](docs/ProductconfiguratordependencyUpdatePayload.md)
 - [ProductconfiguratordependencyUpdateRequest](docs/ProductconfiguratordependencyUpdateRequest.md)
 - [ProductconfiguratormatrixCreateRequest](docs/ProductconfiguratormatrixCreateRequest.md)
 - [ProductconfiguratormatrixEntity](docs/ProductconfiguratormatrixEntity.md)
 - [ProductconfiguratormatrixStep](docs/ProductconfiguratormatrixStep.md)
 - [ProductconfiguratormatrixUpdatePayload](docs/ProductconfiguratormatrixUpdatePayload.md)
 - [ProductconfiguratormatrixUpdateRequest](docs/ProductconfiguratormatrixUpdateRequest.md)
 - [ProductconfiguratoroptionBulkCreateRequest](docs/ProductconfiguratoroptionBulkCreateRequest.md)
 - [ProductconfiguratoroptionBulkCreateRequestCreateEntity](docs/ProductconfiguratoroptionBulkCreateRequestCreateEntity.md)
 - [ProductconfiguratoroptionBulkCreateResponse](docs/ProductconfiguratoroptionBulkCreateResponse.md)
 - [ProductconfiguratoroptionBulkDeleteRequest](docs/ProductconfiguratoroptionBulkDeleteRequest.md)
 - [ProductconfiguratoroptionBulkUpdateRequest](docs/ProductconfiguratoroptionBulkUpdateRequest.md)
 - [ProductconfiguratoroptionBulkUpdateRequestUpdateEntity](docs/ProductconfiguratoroptionBulkUpdateRequestUpdateEntity.md)
 - [ProductconfiguratoroptionBulkUpdateResponse](docs/ProductconfiguratoroptionBulkUpdateResponse.md)
 - [ProductconfiguratoroptionCopyRequest](docs/ProductconfiguratoroptionCopyRequest.md)
 - [ProductconfiguratoroptionCreateRequest](docs/ProductconfiguratoroptionCreateRequest.md)
 - [ProductconfiguratoroptionEntity](docs/ProductconfiguratoroptionEntity.md)
 - [ProductconfiguratoroptionUpdatePayload](docs/ProductconfiguratoroptionUpdatePayload.md)
 - [ProductconfiguratoroptionUpdateRequest](docs/ProductconfiguratoroptionUpdateRequest.md)
 - [ProductconfiguratorpropertyBulkCreateRequest](docs/ProductconfiguratorpropertyBulkCreateRequest.md)
 - [ProductconfiguratorpropertyBulkCreateRequestCreateEntity](docs/ProductconfiguratorpropertyBulkCreateRequestCreateEntity.md)
 - [ProductconfiguratorpropertyBulkCreateResponse](docs/ProductconfiguratorpropertyBulkCreateResponse.md)
 - [ProductconfiguratorpropertyBulkUpdateRequest](docs/ProductconfiguratorpropertyBulkUpdateRequest.md)
 - [ProductconfiguratorpropertyBulkUpdateRequestUpdateEntity](docs/ProductconfiguratorpropertyBulkUpdateRequestUpdateEntity.md)
 - [ProductconfiguratorpropertyBulkUpdateResponse](docs/ProductconfiguratorpropertyBulkUpdateResponse.md)
 - [ProductconfiguratorpropertyCreateRequest](docs/ProductconfiguratorpropertyCreateRequest.md)
 - [ProductconfiguratorpropertyEntity](docs/ProductconfiguratorpropertyEntity.md)
 - [ProductconfiguratorpropertyGenericProperty](docs/ProductconfiguratorpropertyGenericProperty.md)
 - [ProductconfiguratorpropertyPriceProperty](docs/ProductconfiguratorpropertyPriceProperty.md)
 - [ProductconfiguratorpropertyUpdatePayload](docs/ProductconfiguratorpropertyUpdatePayload.md)
 - [ProductconfiguratorpropertyUpdateRequest](docs/ProductconfiguratorpropertyUpdateRequest.md)
 - [ProductconfiguratorpropertyWeightProperty](docs/ProductconfiguratorpropertyWeightProperty.md)
 - [ProductconfiguratorstepBulkCreateRequest](docs/ProductconfiguratorstepBulkCreateRequest.md)
 - [ProductconfiguratorstepBulkCreateRequestCreateEntity](docs/ProductconfiguratorstepBulkCreateRequestCreateEntity.md)
 - [ProductconfiguratorstepBulkCreateResponse](docs/ProductconfiguratorstepBulkCreateResponse.md)
 - [ProductconfiguratorstepBulkDeleteRequest](docs/ProductconfiguratorstepBulkDeleteRequest.md)
 - [ProductconfiguratorstepCopyRequest](docs/ProductconfiguratorstepCopyRequest.md)
 - [ProductconfiguratorstepCreateRequest](docs/ProductconfiguratorstepCreateRequest.md)
 - [ProductconfiguratorstepEntity](docs/ProductconfiguratorstepEntity.md)
 - [ProductconfiguratorstepUpdatePayload](docs/ProductconfiguratorstepUpdatePayload.md)
 - [ProductconfiguratorstepUpdateRequest](docs/ProductconfiguratorstepUpdateRequest.md)
 - [PropertyListPropertiesRequest](docs/PropertyListPropertiesRequest.md)
 - [PropertyListPropertiesResponse](docs/PropertyListPropertiesResponse.md)
 - [PropertyUpdatePayloadGenericProperty](docs/PropertyUpdatePayloadGenericProperty.md)
 - [PropertyUpdatePayloadPriceProperty](docs/PropertyUpdatePayloadPriceProperty.md)
 - [PropertyUpdatePayloadWeightProperty](docs/PropertyUpdatePayloadWeightProperty.md)
 - [ProtobufAny](docs/ProtobufAny.md)


## Documentation For Authorization



### standardAuthorization


- **Type**: OAuth
- **Flow**: implicit
- **Authorization URL**: 
- **Scopes**: N/A

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

info@gemini-commerce.com

