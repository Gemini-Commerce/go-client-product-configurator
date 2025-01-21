# Go API client for productconfigurator


The Product Configurator Service is a versatile platform designed to manage dynamic product configurations. 
It enables the creation, updating, and management of product configurations through steps, options, and dependencies, 
ensuring granular control over customization.

## Core Components
1. **Configurators**
   - Create and manage configurators linked to products.
   - Support for configurator states (`ACTIVE`, `DRAFT`, etc.) and versioning.

2. **Steps**
   - Define logical sequences to guide users through the configuration process.
   - Include localized labels, descriptions, and selection rules.

3. **Options**
   - Add and manage options available for each step.
   - Support for visual content (`Swatch`) and configurable quantities.

4. **Dependencies**
   - Create rules between options and steps to control dynamic interactions.
   - Manage complex conditions across configurations.

5. **Matrices**
   - Use matrices to handle prices, weights, and other properties.
   - Support for dynamic customization based on user selections.

6. **Properties**
   - Add custom attributes and properties to configurators.

7. **Configuration Management**
   - Retrieve available or user-specific configurations.
   - Create optimized configurations to enhance the user experience.

## Key Features
- **Security**: Authenticate every request with JWT, ensuring safety and reliability.
- **Flexibility**: Bulk operations (create, update, delete) for steps, options, and dependencies.
- **Scalability**: Suitable for large volumes of configurations and complex personalization scenarios.
- **Backward Compatibility**: Version management to minimize the impact of changes on existing clients.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v1
- Package version: 1.0.0
- Generator version: 7.9.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import productconfigurator "github.com/Gemini-Commerce/go-client-product-configurator"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `productconfigurator.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), productconfigurator.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `productconfigurator.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), productconfigurator.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `productconfigurator.ContextOperationServerIndices` and `productconfigurator.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), productconfigurator.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), productconfigurator.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://product-configurator.api.gogemini.io*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ConfigurationAPI* | [**ProductConfiguratorGetAvailableConfiguration**](docs/ConfigurationAPI.md#productconfiguratorgetavailableconfiguration) | **Get** /v1/{tenantId}/product/{productId}/configuration | Get Available Configuration
*ConfigurationAPI* | [**ProductConfiguratorGetAvailableConfiguration2**](docs/ConfigurationAPI.md#productconfiguratorgetavailableconfiguration2) | **Post** /v1/{tenantId}/product/{productId}/configuration | Get Available Configuration
*ConfigurationAPI* | [**ProductConfiguratorGetConfigurationFromSelections**](docs/ConfigurationAPI.md#productconfiguratorgetconfigurationfromselections) | **Post** /v1/{tenantId}/product/{productId}/configuration-from-selections | Get Configuration from Selections
*ConfiguratorAPI* | [**ProductConfiguratorCopyConfigurator**](docs/ConfiguratorAPI.md#productconfiguratorcopyconfigurator) | **Post** /v1/{tenantId}/product/{sourceConfiguratorId}/copy | Copy Configurator
*ConfiguratorAPI* | [**ProductConfiguratorCreateConfigurator**](docs/ConfiguratorAPI.md#productconfiguratorcreateconfigurator) | **Post** /v1/{tenantId}/product/{productId}/create | Create Configurator
*ConfiguratorAPI* | [**ProductConfiguratorDeleteConfigurator**](docs/ConfiguratorAPI.md#productconfiguratordeleteconfigurator) | **Delete** /v1/{tenantId}/configurator/{configuratorId} | Delete Configurator
*ConfiguratorAPI* | [**ProductConfiguratorGetConfiguratorByProductId**](docs/ConfiguratorAPI.md#productconfiguratorgetconfiguratorbyproductid) | **Get** /v1/{tenantId}/product/{productId} | Get Configurator by Product ID
*ConfiguratorAPI* | [**ProductConfiguratorGetConfiguratorByProductId2**](docs/ConfiguratorAPI.md#productconfiguratorgetconfiguratorbyproductid2) | **Get** /v1/{tenantId}/product/{productId}/status/{status} | Get Configurator by Product ID
*ConfiguratorAPI* | [**ProductConfiguratorListConfigurators**](docs/ConfiguratorAPI.md#productconfiguratorlistconfigurators) | **Post** /v1/{tenantId}/product/{productId}/page-size/{pageSize}/configurators | List Product Configurators
*ConfiguratorAPI* | [**ProductConfiguratorUpdateConfigurator**](docs/ConfiguratorAPI.md#productconfiguratorupdateconfigurator) | **Put** /v1/{tenantId}/configurator/{configuratorId} | Update Configurator
*DependencyAPI* | [**ProductConfiguratorCreateDependency**](docs/DependencyAPI.md#productconfiguratorcreatedependency) | **Post** /v1/{tenantId}/step/{stepId}/dependency/create | Create Dependency
*DependencyAPI* | [**ProductConfiguratorDeleteDependency**](docs/DependencyAPI.md#productconfiguratordeletedependency) | **Delete** /v1/{tenantId}/dependency/{dependencyId} | Delete Dependency
*DependencyAPI* | [**ProductConfiguratorListDependencies**](docs/DependencyAPI.md#productconfiguratorlistdependencies) | **Post** /v1/{tenantId}/page-size/{pageSize}/dependencies | List Dependencies
*DependencyAPI* | [**ProductConfiguratorUpdateDependency**](docs/DependencyAPI.md#productconfiguratorupdatedependency) | **Put** /v1/{tenantId}/dependency/{dependencyId} | Update Dependency
*MatrixAPI* | [**ProductConfiguratorCreateMatrix**](docs/MatrixAPI.md#productconfiguratorcreatematrix) | **Post** /v1/{tenantId}/matrix/create | Create Matrix
*MatrixAPI* | [**ProductConfiguratorDeleteMatrix**](docs/MatrixAPI.md#productconfiguratordeletematrix) | **Delete** /v1/{tenantId}/matrix/{matrixId} | Delete Matrix
*MatrixAPI* | [**ProductConfiguratorGetMatrix**](docs/MatrixAPI.md#productconfiguratorgetmatrix) | **Get** /v1/{tenantId}/matrix/{matrixId} | Get Matrix
*MatrixAPI* | [**ProductConfiguratorListMatrices**](docs/MatrixAPI.md#productconfiguratorlistmatrices) | **Post** /v1/{tenantId}/configurator/{configuratorId}/page-size/{pageSize}/matrices | List Matrices
*MatrixAPI* | [**ProductConfiguratorRemovePricelistFromMatrix**](docs/MatrixAPI.md#productconfiguratorremovepricelistfrommatrix) | **Delete** /v1/{tenantId}/matrix/{matrixId}/pricelist/{pricelistGrn} | Remove Pricelist from Matrix
*MatrixAPI* | [**ProductConfiguratorUpdateMatrix**](docs/MatrixAPI.md#productconfiguratorupdatematrix) | **Put** /v1/{tenantId}/matrix/{matrixId} | Update Matrix
*OptionAPI* | [**ProductConfiguratorBulkCreateOptions**](docs/OptionAPI.md#productconfiguratorbulkcreateoptions) | **Post** /v1/{tenantId}/step/{stepId}/option/create/bulk | Bulk Create Options
*OptionAPI* | [**ProductConfiguratorBulkDeleteOptions**](docs/OptionAPI.md#productconfiguratorbulkdeleteoptions) | **Post** /v1/{tenantId}/option/delete/bulk | Bulk Delete Options
*OptionAPI* | [**ProductConfiguratorBulkUpdateOptions**](docs/OptionAPI.md#productconfiguratorbulkupdateoptions) | **Put** /v1/{tenantId}/option/bulk | Bulk Update Options
*OptionAPI* | [**ProductConfiguratorCopyOption**](docs/OptionAPI.md#productconfiguratorcopyoption) | **Post** /v1/{tenantId}/option/{sourceOptionId}/copy | Copy Option
*OptionAPI* | [**ProductConfiguratorCreateOption**](docs/OptionAPI.md#productconfiguratorcreateoption) | **Post** /v1/{tenantId}/step/{stepId}/option/create | Create Option
*OptionAPI* | [**ProductConfiguratorDeleteOption**](docs/OptionAPI.md#productconfiguratordeleteoption) | **Delete** /v1/{tenantId}/option/{optionId} | Delete Option
*OptionAPI* | [**ProductConfiguratorListOptions**](docs/OptionAPI.md#productconfiguratorlistoptions) | **Post** /v1/{tenantId}/step/{stepId}/page-size/{pageSize}/options | List Options
*OptionAPI* | [**ProductConfiguratorUpdateOption**](docs/OptionAPI.md#productconfiguratorupdateoption) | **Put** /v1/{tenantId}/option/{optionId} | Update Option
*ProductConfiguratorAPI* | [**ProductConfiguratorGetConfiguratorById**](docs/ProductConfiguratorAPI.md#productconfiguratorgetconfiguratorbyid) | **Get** /v1/{tenantId}/configurator/{configuratorId} | 
*ProductConfiguratorAPI* | [**ProductConfiguratorGetProperty**](docs/ProductConfiguratorAPI.md#productconfiguratorgetproperty) | **Get** /v1/{tenantId}/property/{propertyId} | 
*ProductConfiguratorAPI* | [**ProductConfiguratorListPropertiesByConfiguration**](docs/ProductConfiguratorAPI.md#productconfiguratorlistpropertiesbyconfiguration) | **Post** /v1/{tenantId}/configurator/{configuratorId}/page-size/{pageSize}/properties | 
*PropertyAPI* | [**ProductConfiguratorBulkCreateProperties**](docs/PropertyAPI.md#productconfiguratorbulkcreateproperties) | **Post** /v1/{tenantId}/property/create/bulk | Bulk Create Properties
*PropertyAPI* | [**ProductConfiguratorBulkUpdateProperties**](docs/PropertyAPI.md#productconfiguratorbulkupdateproperties) | **Put** /v1/{tenantId}/properties/bulk | Bulk Update Properties
*PropertyAPI* | [**ProductConfiguratorCreateProperty**](docs/PropertyAPI.md#productconfiguratorcreateproperty) | **Post** /v1/{tenantId}/property/create | Create Property
*PropertyAPI* | [**ProductConfiguratorListProperties**](docs/PropertyAPI.md#productconfiguratorlistproperties) | **Post** /v1/{tenantId}/matrix/{matrixId}/page-size/{pageSize}/properties | List Properties
*PropertyAPI* | [**ProductConfiguratorUpdateProperty**](docs/PropertyAPI.md#productconfiguratorupdateproperty) | **Put** /v1/{tenantId}/property/{propertyId} | Update Property
*StepAPI* | [**ProductConfiguratorBulkCreateSteps**](docs/StepAPI.md#productconfiguratorbulkcreatesteps) | **Post** /v1/{tenantId}/configurator/{configuratorId}/step/create/bulk | Bulk Create Steps
*StepAPI* | [**ProductConfiguratorBulkDeleteSteps**](docs/StepAPI.md#productconfiguratorbulkdeletesteps) | **Post** /v1/{tenantId}/step/delete/bulk | Bulk Delete Steps
*StepAPI* | [**ProductConfiguratorCopyStep**](docs/StepAPI.md#productconfiguratorcopystep) | **Post** /v1/{tenantId}/step/{sourceStepId}/copy | Copy Step
*StepAPI* | [**ProductConfiguratorCreateStep**](docs/StepAPI.md#productconfiguratorcreatestep) | **Post** /v1/{tenantId}/configurator/{configuratorId}/step/create | Create Step
*StepAPI* | [**ProductConfiguratorDeleteStep**](docs/StepAPI.md#productconfiguratordeletestep) | **Delete** /v1/{tenantId}/step/{stepId} | Delete Step
*StepAPI* | [**ProductConfiguratorUpdateStep**](docs/StepAPI.md#productconfiguratorupdatestep) | **Put** /v1/{tenantId}/step/{stepId} | Update Step


## Documentation For Models

 - [ConfigurationConfigurationStep](docs/ConfigurationConfigurationStep.md)
 - [ConfigurationConfigurator](docs/ConfigurationConfigurator.md)
 - [ConfigurationGetAvailableConfigurationResponse](docs/ConfigurationGetAvailableConfigurationResponse.md)
 - [ConfigurationGetConfigurationFromSelectionsResponse](docs/ConfigurationGetConfigurationFromSelectionsResponse.md)
 - [ConfigurationOptionProperty](docs/ConfigurationOptionProperty.md)
 - [ConfigurationProperty](docs/ConfigurationProperty.md)
 - [ConfigurationPropertyFilter](docs/ConfigurationPropertyFilter.md)
 - [ConfigurationPropertyFilterGenericProperty](docs/ConfigurationPropertyFilterGenericProperty.md)
 - [ConfigurationStepOption](docs/ConfigurationStepOption.md)
 - [ConfiguratorListResponse](docs/ConfiguratorListResponse.md)
 - [DependencyCondition](docs/DependencyCondition.md)
 - [DependencyListDependenciesResponse](docs/DependencyListDependenciesResponse.md)
 - [GooglerpcStatus](docs/GooglerpcStatus.md)
 - [ListMatricesRequestFilter](docs/ListMatricesRequestFilter.md)
 - [LocalisationLocalizedText](docs/LocalisationLocalizedText.md)
 - [MatrixGenericType](docs/MatrixGenericType.md)
 - [MatrixListMatricesResponse](docs/MatrixListMatricesResponse.md)
 - [MatrixPriceType](docs/MatrixPriceType.md)
 - [MatrixWeightType](docs/MatrixWeightType.md)
 - [OptionListOptionsResponse](docs/OptionListOptionsResponse.md)
 - [OptionSwatch](docs/OptionSwatch.md)
 - [ProductConfiguratorBulkCreateOptionsRequest](docs/ProductConfiguratorBulkCreateOptionsRequest.md)
 - [ProductConfiguratorBulkCreatePropertiesRequest](docs/ProductConfiguratorBulkCreatePropertiesRequest.md)
 - [ProductConfiguratorBulkCreateStepsRequest](docs/ProductConfiguratorBulkCreateStepsRequest.md)
 - [ProductConfiguratorBulkDeleteOptionsRequest](docs/ProductConfiguratorBulkDeleteOptionsRequest.md)
 - [ProductConfiguratorBulkDeleteStepsRequest](docs/ProductConfiguratorBulkDeleteStepsRequest.md)
 - [ProductConfiguratorBulkUpdateOptionsRequest](docs/ProductConfiguratorBulkUpdateOptionsRequest.md)
 - [ProductConfiguratorBulkUpdatePropertiesRequest](docs/ProductConfiguratorBulkUpdatePropertiesRequest.md)
 - [ProductConfiguratorCopyConfiguratorRequest](docs/ProductConfiguratorCopyConfiguratorRequest.md)
 - [ProductConfiguratorCopyOptionRequest](docs/ProductConfiguratorCopyOptionRequest.md)
 - [ProductConfiguratorCopyStepRequest](docs/ProductConfiguratorCopyStepRequest.md)
 - [ProductConfiguratorCreateConfiguratorRequest](docs/ProductConfiguratorCreateConfiguratorRequest.md)
 - [ProductConfiguratorCreateDependencyRequest](docs/ProductConfiguratorCreateDependencyRequest.md)
 - [ProductConfiguratorCreateMatrixRequest](docs/ProductConfiguratorCreateMatrixRequest.md)
 - [ProductConfiguratorCreateOptionRequest](docs/ProductConfiguratorCreateOptionRequest.md)
 - [ProductConfiguratorCreatePropertyRequest](docs/ProductConfiguratorCreatePropertyRequest.md)
 - [ProductConfiguratorCreateStepRequest](docs/ProductConfiguratorCreateStepRequest.md)
 - [ProductConfiguratorGetAvailableConfiguration2Request](docs/ProductConfiguratorGetAvailableConfiguration2Request.md)
 - [ProductConfiguratorGetConfigurationFromSelectionsRequest](docs/ProductConfiguratorGetConfigurationFromSelectionsRequest.md)
 - [ProductConfiguratorListDependenciesRequest](docs/ProductConfiguratorListDependenciesRequest.md)
 - [ProductConfiguratorListMatricesRequest](docs/ProductConfiguratorListMatricesRequest.md)
 - [ProductConfiguratorListPropertiesByConfigurationRequest](docs/ProductConfiguratorListPropertiesByConfigurationRequest.md)
 - [ProductConfiguratorListPropertiesRequest](docs/ProductConfiguratorListPropertiesRequest.md)
 - [ProductConfiguratorUpdateConfiguratorRequest](docs/ProductConfiguratorUpdateConfiguratorRequest.md)
 - [ProductConfiguratorUpdateDependencyRequest](docs/ProductConfiguratorUpdateDependencyRequest.md)
 - [ProductConfiguratorUpdateMatrixRequest](docs/ProductConfiguratorUpdateMatrixRequest.md)
 - [ProductConfiguratorUpdateOptionRequest](docs/ProductConfiguratorUpdateOptionRequest.md)
 - [ProductConfiguratorUpdatePropertyRequest](docs/ProductConfiguratorUpdatePropertyRequest.md)
 - [ProductConfiguratorUpdateStepRequest](docs/ProductConfiguratorUpdateStepRequest.md)
 - [ProductconfiguratorMoney](docs/ProductconfiguratorMoney.md)
 - [ProductconfiguratorPropertyMode](docs/ProductconfiguratorPropertyMode.md)
 - [ProductconfiguratorPropertyType](docs/ProductconfiguratorPropertyType.md)
 - [ProductconfiguratorWeightUnit](docs/ProductconfiguratorWeightUnit.md)
 - [ProductconfiguratorconfigurationOption](docs/ProductconfiguratorconfigurationOption.md)
 - [ProductconfiguratorconfigurationSelection](docs/ProductconfiguratorconfigurationSelection.md)
 - [ProductconfiguratorconfigurationStep](docs/ProductconfiguratorconfigurationStep.md)
 - [ProductconfiguratorconfiguratorEntity](docs/ProductconfiguratorconfiguratorEntity.md)
 - [ProductconfiguratorconfiguratorStatus](docs/ProductconfiguratorconfiguratorStatus.md)
 - [ProductconfiguratorconfiguratorUpdatePayload](docs/ProductconfiguratorconfiguratorUpdatePayload.md)
 - [ProductconfiguratordependencyEntity](docs/ProductconfiguratordependencyEntity.md)
 - [ProductconfiguratordependencyType](docs/ProductconfiguratordependencyType.md)
 - [ProductconfiguratordependencyUpdatePayload](docs/ProductconfiguratordependencyUpdatePayload.md)
 - [ProductconfiguratormatrixEntity](docs/ProductconfiguratormatrixEntity.md)
 - [ProductconfiguratormatrixStep](docs/ProductconfiguratormatrixStep.md)
 - [ProductconfiguratormatrixUpdatePayload](docs/ProductconfiguratormatrixUpdatePayload.md)
 - [ProductconfiguratoroptionBulkCreateRequestCreateEntity](docs/ProductconfiguratoroptionBulkCreateRequestCreateEntity.md)
 - [ProductconfiguratoroptionBulkCreateResponse](docs/ProductconfiguratoroptionBulkCreateResponse.md)
 - [ProductconfiguratoroptionBulkUpdateRequestUpdateEntity](docs/ProductconfiguratoroptionBulkUpdateRequestUpdateEntity.md)
 - [ProductconfiguratoroptionBulkUpdateResponse](docs/ProductconfiguratoroptionBulkUpdateResponse.md)
 - [ProductconfiguratoroptionEntity](docs/ProductconfiguratoroptionEntity.md)
 - [ProductconfiguratoroptionUpdatePayload](docs/ProductconfiguratoroptionUpdatePayload.md)
 - [ProductconfiguratorpropertyBulkCreateRequestCreateEntity](docs/ProductconfiguratorpropertyBulkCreateRequestCreateEntity.md)
 - [ProductconfiguratorpropertyBulkCreateResponse](docs/ProductconfiguratorpropertyBulkCreateResponse.md)
 - [ProductconfiguratorpropertyBulkUpdateRequestUpdateEntity](docs/ProductconfiguratorpropertyBulkUpdateRequestUpdateEntity.md)
 - [ProductconfiguratorpropertyBulkUpdateResponse](docs/ProductconfiguratorpropertyBulkUpdateResponse.md)
 - [ProductconfiguratorpropertyEntity](docs/ProductconfiguratorpropertyEntity.md)
 - [ProductconfiguratorpropertyGenericProperty](docs/ProductconfiguratorpropertyGenericProperty.md)
 - [ProductconfiguratorpropertyPriceProperty](docs/ProductconfiguratorpropertyPriceProperty.md)
 - [ProductconfiguratorpropertyUpdatePayload](docs/ProductconfiguratorpropertyUpdatePayload.md)
 - [ProductconfiguratorpropertyWeightProperty](docs/ProductconfiguratorpropertyWeightProperty.md)
 - [ProductconfiguratorstepBulkCreateRequestCreateEntity](docs/ProductconfiguratorstepBulkCreateRequestCreateEntity.md)
 - [ProductconfiguratorstepBulkCreateResponse](docs/ProductconfiguratorstepBulkCreateResponse.md)
 - [ProductconfiguratorstepEntity](docs/ProductconfiguratorstepEntity.md)
 - [ProductconfiguratorstepUpdatePayload](docs/ProductconfiguratorstepUpdatePayload.md)
 - [PropertyListPropertiesByConfigurationRequestSelection](docs/PropertyListPropertiesByConfigurationRequestSelection.md)
 - [PropertyListPropertiesByConfigurationResponse](docs/PropertyListPropertiesByConfigurationResponse.md)
 - [PropertyListPropertiesResponse](docs/PropertyListPropertiesResponse.md)
 - [PropertyUpdatePayloadGenericProperty](docs/PropertyUpdatePayloadGenericProperty.md)
 - [PropertyUpdatePayloadPriceProperty](docs/PropertyUpdatePayloadPriceProperty.md)
 - [PropertyUpdatePayloadWeightProperty](docs/PropertyUpdatePayloadWeightProperty.md)
 - [ProtobufAny](docs/ProtobufAny.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### APIAuthorization

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: APIAuthorization and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		productconfigurator.ContextAPIKeys,
		map[string]productconfigurator.APIKey{
			"APIAuthorization": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```

### standardAuthorization


- **Type**: OAuth
- **Flow**: implicit
- **Authorization URL**: https://iambackoffice.gogemini.io/iambackoffice.IamBackoffice/Login
- **Scopes**: N/A

Example

```go
auth := context.WithValue(context.Background(), productconfigurator.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```go
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, productconfigurator.ContextOAuth2, tokenSource)
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

