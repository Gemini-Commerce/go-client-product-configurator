/*
Product Configurator Service

## Introduction  This comprehensive guide will equip you with the knowledge to integrate and leverage our Product Configurator Service in your applications.  ## Quick Start  Get up and running in no time! Follow these steps to kickstart your integration:  1. **Authentication:** Obtain your integration JWT to authenticate your requests. 2. **Client Libraries:** Explore our GitHub repositories to grab client libraries in your preferred programming language. 3. **API Overview:** Familiarize yourself with our RESTful API using the OpenAPI specification.  ## Integration  ### API Overview  Our RESTful API is the gateway to unlocking the full potential of Product Configurator. Check out the detailed [API Reference](/docs/category/configurator) for a granular understanding of each endpoint and request/response format.  ### Client Libraries  To expedite your integration process, we provide client libraries for various programming languages. Find the one that suits your stack in our [GitHub repositories](https://github.com/Gemini-Commerce).  ### Authentication  Security is paramount. Learn how to authenticate your requests using JWT. This ensures a secure and reliable connection between your application and Product Configurator.  ## Configuration Management  ### Configurator Lifecycle  Understand the lifecycle of configurators, from draft to active and deleted. This flexibility allows you to manage configurations at your own pace.  ### Steps and Options  Configure product steps with ease and define options effortlessly. Explore the power of dependencies to create dynamic and intuitive configurations.  ### Matrices  Delve into matrices—your secret weapon. Explore price and weight matrices, and learn how configured steps influence properties and pricing.  ### Price Management  Unleash dynamic pricing with our versatile price matrices. From fixed prices to incremental structures, adapt to diverse pricing models effortlessly.  ## Security  Your data is in safe hands. Discover how Product Configurator ensures security through JWT authentication, safeguarding your sensitive information.  ## Backward Compatibility  Stay ahead of the curve. Learn about our versioning strategy, providing backward compatibility while allowing our service to evolve seamlessly.  ## Developer Support  Have questions? Need assistance? Write to us at [info@gemini-commerce.com](mailto:info@gemini-commerce.com) and we will get back to you.

API version: v1
Contact: info@gemini-commerce.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package product-configurator

import (
	"encoding/json"
)

// checks if the ProductConfiguratorListPropertiesByConfigurationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductConfiguratorListPropertiesByConfigurationRequest{}

// ProductConfiguratorListPropertiesByConfigurationRequest struct for ProductConfiguratorListPropertiesByConfigurationRequest
type ProductConfiguratorListPropertiesByConfigurationRequest struct {
	Selections []PropertyListPropertiesByConfigurationRequestSelection `json:"selections,omitempty"`
	PropertyType *ProductconfiguratorPropertyType `json:"propertyType,omitempty"`
	PageToken *string `json:"pageToken,omitempty"`
}

// NewProductConfiguratorListPropertiesByConfigurationRequest instantiates a new ProductConfiguratorListPropertiesByConfigurationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductConfiguratorListPropertiesByConfigurationRequest() *ProductConfiguratorListPropertiesByConfigurationRequest {
	this := ProductConfiguratorListPropertiesByConfigurationRequest{}
	var propertyType ProductconfiguratorPropertyType = UNKNOWN
	this.PropertyType = &propertyType
	return &this
}

// NewProductConfiguratorListPropertiesByConfigurationRequestWithDefaults instantiates a new ProductConfiguratorListPropertiesByConfigurationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductConfiguratorListPropertiesByConfigurationRequestWithDefaults() *ProductConfiguratorListPropertiesByConfigurationRequest {
	this := ProductConfiguratorListPropertiesByConfigurationRequest{}
	var propertyType ProductconfiguratorPropertyType = UNKNOWN
	this.PropertyType = &propertyType
	return &this
}

// GetSelections returns the Selections field value if set, zero value otherwise.
func (o *ProductConfiguratorListPropertiesByConfigurationRequest) GetSelections() []PropertyListPropertiesByConfigurationRequestSelection {
	if o == nil || IsNil(o.Selections) {
		var ret []PropertyListPropertiesByConfigurationRequestSelection
		return ret
	}
	return o.Selections
}

// GetSelectionsOk returns a tuple with the Selections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductConfiguratorListPropertiesByConfigurationRequest) GetSelectionsOk() ([]PropertyListPropertiesByConfigurationRequestSelection, bool) {
	if o == nil || IsNil(o.Selections) {
		return nil, false
	}
	return o.Selections, true
}

// HasSelections returns a boolean if a field has been set.
func (o *ProductConfiguratorListPropertiesByConfigurationRequest) HasSelections() bool {
	if o != nil && !IsNil(o.Selections) {
		return true
	}

	return false
}

// SetSelections gets a reference to the given []PropertyListPropertiesByConfigurationRequestSelection and assigns it to the Selections field.
func (o *ProductConfiguratorListPropertiesByConfigurationRequest) SetSelections(v []PropertyListPropertiesByConfigurationRequestSelection) {
	o.Selections = v
}

// GetPropertyType returns the PropertyType field value if set, zero value otherwise.
func (o *ProductConfiguratorListPropertiesByConfigurationRequest) GetPropertyType() ProductconfiguratorPropertyType {
	if o == nil || IsNil(o.PropertyType) {
		var ret ProductconfiguratorPropertyType
		return ret
	}
	return *o.PropertyType
}

// GetPropertyTypeOk returns a tuple with the PropertyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductConfiguratorListPropertiesByConfigurationRequest) GetPropertyTypeOk() (*ProductconfiguratorPropertyType, bool) {
	if o == nil || IsNil(o.PropertyType) {
		return nil, false
	}
	return o.PropertyType, true
}

// HasPropertyType returns a boolean if a field has been set.
func (o *ProductConfiguratorListPropertiesByConfigurationRequest) HasPropertyType() bool {
	if o != nil && !IsNil(o.PropertyType) {
		return true
	}

	return false
}

// SetPropertyType gets a reference to the given ProductconfiguratorPropertyType and assigns it to the PropertyType field.
func (o *ProductConfiguratorListPropertiesByConfigurationRequest) SetPropertyType(v ProductconfiguratorPropertyType) {
	o.PropertyType = &v
}

// GetPageToken returns the PageToken field value if set, zero value otherwise.
func (o *ProductConfiguratorListPropertiesByConfigurationRequest) GetPageToken() string {
	if o == nil || IsNil(o.PageToken) {
		var ret string
		return ret
	}
	return *o.PageToken
}

// GetPageTokenOk returns a tuple with the PageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductConfiguratorListPropertiesByConfigurationRequest) GetPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.PageToken) {
		return nil, false
	}
	return o.PageToken, true
}

// HasPageToken returns a boolean if a field has been set.
func (o *ProductConfiguratorListPropertiesByConfigurationRequest) HasPageToken() bool {
	if o != nil && !IsNil(o.PageToken) {
		return true
	}

	return false
}

// SetPageToken gets a reference to the given string and assigns it to the PageToken field.
func (o *ProductConfiguratorListPropertiesByConfigurationRequest) SetPageToken(v string) {
	o.PageToken = &v
}

func (o ProductConfiguratorListPropertiesByConfigurationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductConfiguratorListPropertiesByConfigurationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Selections) {
		toSerialize["selections"] = o.Selections
	}
	if !IsNil(o.PropertyType) {
		toSerialize["propertyType"] = o.PropertyType
	}
	if !IsNil(o.PageToken) {
		toSerialize["pageToken"] = o.PageToken
	}
	return toSerialize, nil
}

type NullableProductConfiguratorListPropertiesByConfigurationRequest struct {
	value *ProductConfiguratorListPropertiesByConfigurationRequest
	isSet bool
}

func (v NullableProductConfiguratorListPropertiesByConfigurationRequest) Get() *ProductConfiguratorListPropertiesByConfigurationRequest {
	return v.value
}

func (v *NullableProductConfiguratorListPropertiesByConfigurationRequest) Set(val *ProductConfiguratorListPropertiesByConfigurationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductConfiguratorListPropertiesByConfigurationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductConfiguratorListPropertiesByConfigurationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductConfiguratorListPropertiesByConfigurationRequest(val *ProductConfiguratorListPropertiesByConfigurationRequest) *NullableProductConfiguratorListPropertiesByConfigurationRequest {
	return &NullableProductConfiguratorListPropertiesByConfigurationRequest{value: val, isSet: true}
}

func (v NullableProductConfiguratorListPropertiesByConfigurationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductConfiguratorListPropertiesByConfigurationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

