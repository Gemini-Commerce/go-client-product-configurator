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

// checks if the ProductConfiguratorGetAvailableConfiguration2Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductConfiguratorGetAvailableConfiguration2Request{}

// ProductConfiguratorGetAvailableConfiguration2Request struct for ProductConfiguratorGetAvailableConfiguration2Request
type ProductConfiguratorGetAvailableConfiguration2Request struct {
	Selections []ProductconfiguratorconfigurationSelection `json:"selections,omitempty"`
	PropertyFilters []ConfigurationPropertyFilter `json:"propertyFilters,omitempty"`
	ConfiguratorId *string `json:"configuratorId,omitempty"`
}

// NewProductConfiguratorGetAvailableConfiguration2Request instantiates a new ProductConfiguratorGetAvailableConfiguration2Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductConfiguratorGetAvailableConfiguration2Request() *ProductConfiguratorGetAvailableConfiguration2Request {
	this := ProductConfiguratorGetAvailableConfiguration2Request{}
	return &this
}

// NewProductConfiguratorGetAvailableConfiguration2RequestWithDefaults instantiates a new ProductConfiguratorGetAvailableConfiguration2Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductConfiguratorGetAvailableConfiguration2RequestWithDefaults() *ProductConfiguratorGetAvailableConfiguration2Request {
	this := ProductConfiguratorGetAvailableConfiguration2Request{}
	return &this
}

// GetSelections returns the Selections field value if set, zero value otherwise.
func (o *ProductConfiguratorGetAvailableConfiguration2Request) GetSelections() []ProductconfiguratorconfigurationSelection {
	if o == nil || IsNil(o.Selections) {
		var ret []ProductconfiguratorconfigurationSelection
		return ret
	}
	return o.Selections
}

// GetSelectionsOk returns a tuple with the Selections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductConfiguratorGetAvailableConfiguration2Request) GetSelectionsOk() ([]ProductconfiguratorconfigurationSelection, bool) {
	if o == nil || IsNil(o.Selections) {
		return nil, false
	}
	return o.Selections, true
}

// HasSelections returns a boolean if a field has been set.
func (o *ProductConfiguratorGetAvailableConfiguration2Request) HasSelections() bool {
	if o != nil && !IsNil(o.Selections) {
		return true
	}

	return false
}

// SetSelections gets a reference to the given []ProductconfiguratorconfigurationSelection and assigns it to the Selections field.
func (o *ProductConfiguratorGetAvailableConfiguration2Request) SetSelections(v []ProductconfiguratorconfigurationSelection) {
	o.Selections = v
}

// GetPropertyFilters returns the PropertyFilters field value if set, zero value otherwise.
func (o *ProductConfiguratorGetAvailableConfiguration2Request) GetPropertyFilters() []ConfigurationPropertyFilter {
	if o == nil || IsNil(o.PropertyFilters) {
		var ret []ConfigurationPropertyFilter
		return ret
	}
	return o.PropertyFilters
}

// GetPropertyFiltersOk returns a tuple with the PropertyFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductConfiguratorGetAvailableConfiguration2Request) GetPropertyFiltersOk() ([]ConfigurationPropertyFilter, bool) {
	if o == nil || IsNil(o.PropertyFilters) {
		return nil, false
	}
	return o.PropertyFilters, true
}

// HasPropertyFilters returns a boolean if a field has been set.
func (o *ProductConfiguratorGetAvailableConfiguration2Request) HasPropertyFilters() bool {
	if o != nil && !IsNil(o.PropertyFilters) {
		return true
	}

	return false
}

// SetPropertyFilters gets a reference to the given []ConfigurationPropertyFilter and assigns it to the PropertyFilters field.
func (o *ProductConfiguratorGetAvailableConfiguration2Request) SetPropertyFilters(v []ConfigurationPropertyFilter) {
	o.PropertyFilters = v
}

// GetConfiguratorId returns the ConfiguratorId field value if set, zero value otherwise.
func (o *ProductConfiguratorGetAvailableConfiguration2Request) GetConfiguratorId() string {
	if o == nil || IsNil(o.ConfiguratorId) {
		var ret string
		return ret
	}
	return *o.ConfiguratorId
}

// GetConfiguratorIdOk returns a tuple with the ConfiguratorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductConfiguratorGetAvailableConfiguration2Request) GetConfiguratorIdOk() (*string, bool) {
	if o == nil || IsNil(o.ConfiguratorId) {
		return nil, false
	}
	return o.ConfiguratorId, true
}

// HasConfiguratorId returns a boolean if a field has been set.
func (o *ProductConfiguratorGetAvailableConfiguration2Request) HasConfiguratorId() bool {
	if o != nil && !IsNil(o.ConfiguratorId) {
		return true
	}

	return false
}

// SetConfiguratorId gets a reference to the given string and assigns it to the ConfiguratorId field.
func (o *ProductConfiguratorGetAvailableConfiguration2Request) SetConfiguratorId(v string) {
	o.ConfiguratorId = &v
}

func (o ProductConfiguratorGetAvailableConfiguration2Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductConfiguratorGetAvailableConfiguration2Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Selections) {
		toSerialize["selections"] = o.Selections
	}
	if !IsNil(o.PropertyFilters) {
		toSerialize["propertyFilters"] = o.PropertyFilters
	}
	if !IsNil(o.ConfiguratorId) {
		toSerialize["configuratorId"] = o.ConfiguratorId
	}
	return toSerialize, nil
}

type NullableProductConfiguratorGetAvailableConfiguration2Request struct {
	value *ProductConfiguratorGetAvailableConfiguration2Request
	isSet bool
}

func (v NullableProductConfiguratorGetAvailableConfiguration2Request) Get() *ProductConfiguratorGetAvailableConfiguration2Request {
	return v.value
}

func (v *NullableProductConfiguratorGetAvailableConfiguration2Request) Set(val *ProductConfiguratorGetAvailableConfiguration2Request) {
	v.value = val
	v.isSet = true
}

func (v NullableProductConfiguratorGetAvailableConfiguration2Request) IsSet() bool {
	return v.isSet
}

func (v *NullableProductConfiguratorGetAvailableConfiguration2Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductConfiguratorGetAvailableConfiguration2Request(val *ProductConfiguratorGetAvailableConfiguration2Request) *NullableProductConfiguratorGetAvailableConfiguration2Request {
	return &NullableProductConfiguratorGetAvailableConfiguration2Request{value: val, isSet: true}
}

func (v NullableProductConfiguratorGetAvailableConfiguration2Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductConfiguratorGetAvailableConfiguration2Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


