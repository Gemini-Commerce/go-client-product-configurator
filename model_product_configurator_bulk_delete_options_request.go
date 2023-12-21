/*
Product Configurator Service

## Introduction  This comprehensive guide will equip you with the knowledge to integrate and leverage our Product Configurator Service in your applications.  ## Quick Start  Get up and running in no time! Follow these steps to kickstart your integration:  1. **Authentication:** Obtain your integration JWT to authenticate your requests. 2. **Client Libraries:** Explore our GitHub repositories to grab client libraries in your preferred programming language. 3. **API Overview:** Familiarize yourself with our RESTful API using the OpenAPI specification.  ## Integration  ### API Overview  Our RESTful API is the gateway to unlocking the full potential of Product Configurator. Check out the detailed [API Reference](/docs/category/configurator) for a granular understanding of each endpoint and request/response format.  ### Client Libraries  To expedite your integration process, we provide client libraries for various programming languages. Find the one that suits your stack in our [GitHub repositories](https://github.com/Gemini-Commerce).  ### Authentication  Security is paramount. Learn how to authenticate your requests using JWT. This ensures a secure and reliable connection between your application and Product Configurator.  ## Configuration Management  ### Configurator Lifecycle  Understand the lifecycle of configurators, from draft to active and deleted. This flexibility allows you to manage configurations at your own pace.  ### Steps and Options  Configure product steps with ease and define options effortlessly. Explore the power of dependencies to create dynamic and intuitive configurations.  ### Matrices  Delve into matrices—your secret weapon. Explore price and weight matrices, and learn how configured steps influence properties and pricing.  ### Price Management  Unleash dynamic pricing with our versatile price matrices. From fixed prices to incremental structures, adapt to diverse pricing models effortlessly.  ## Security  Your data is in safe hands. Discover how Product Configurator ensures security through JWT authentication, safeguarding your sensitive information.  ## Backward Compatibility  Stay ahead of the curve. Learn about our versioning strategy, providing backward compatibility while allowing our service to evolve seamlessly.  ## Developer Support  Have questions? Need assistance? Write to us at [info@gemini-commerce.com](mailto:info@gemini-commerce.com) and we will get back to you.

API version: v1
Contact: info@gemini-commerce.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package product_configurator

import (
	"encoding/json"
)

// checks if the ProductConfiguratorBulkDeleteOptionsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductConfiguratorBulkDeleteOptionsRequest{}

// ProductConfiguratorBulkDeleteOptionsRequest struct for ProductConfiguratorBulkDeleteOptionsRequest
type ProductConfiguratorBulkDeleteOptionsRequest struct {
	OptionIds []string `json:"optionIds,omitempty"`
}

// NewProductConfiguratorBulkDeleteOptionsRequest instantiates a new ProductConfiguratorBulkDeleteOptionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductConfiguratorBulkDeleteOptionsRequest() *ProductConfiguratorBulkDeleteOptionsRequest {
	this := ProductConfiguratorBulkDeleteOptionsRequest{}
	return &this
}

// NewProductConfiguratorBulkDeleteOptionsRequestWithDefaults instantiates a new ProductConfiguratorBulkDeleteOptionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductConfiguratorBulkDeleteOptionsRequestWithDefaults() *ProductConfiguratorBulkDeleteOptionsRequest {
	this := ProductConfiguratorBulkDeleteOptionsRequest{}
	return &this
}

// GetOptionIds returns the OptionIds field value if set, zero value otherwise.
func (o *ProductConfiguratorBulkDeleteOptionsRequest) GetOptionIds() []string {
	if o == nil || IsNil(o.OptionIds) {
		var ret []string
		return ret
	}
	return o.OptionIds
}

// GetOptionIdsOk returns a tuple with the OptionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductConfiguratorBulkDeleteOptionsRequest) GetOptionIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.OptionIds) {
		return nil, false
	}
	return o.OptionIds, true
}

// HasOptionIds returns a boolean if a field has been set.
func (o *ProductConfiguratorBulkDeleteOptionsRequest) HasOptionIds() bool {
	if o != nil && !IsNil(o.OptionIds) {
		return true
	}

	return false
}

// SetOptionIds gets a reference to the given []string and assigns it to the OptionIds field.
func (o *ProductConfiguratorBulkDeleteOptionsRequest) SetOptionIds(v []string) {
	o.OptionIds = v
}

func (o ProductConfiguratorBulkDeleteOptionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductConfiguratorBulkDeleteOptionsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OptionIds) {
		toSerialize["optionIds"] = o.OptionIds
	}
	return toSerialize, nil
}

type NullableProductConfiguratorBulkDeleteOptionsRequest struct {
	value *ProductConfiguratorBulkDeleteOptionsRequest
	isSet bool
}

func (v NullableProductConfiguratorBulkDeleteOptionsRequest) Get() *ProductConfiguratorBulkDeleteOptionsRequest {
	return v.value
}

func (v *NullableProductConfiguratorBulkDeleteOptionsRequest) Set(val *ProductConfiguratorBulkDeleteOptionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductConfiguratorBulkDeleteOptionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductConfiguratorBulkDeleteOptionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductConfiguratorBulkDeleteOptionsRequest(val *ProductConfiguratorBulkDeleteOptionsRequest) *NullableProductConfiguratorBulkDeleteOptionsRequest {
	return &NullableProductConfiguratorBulkDeleteOptionsRequest{value: val, isSet: true}
}

func (v NullableProductConfiguratorBulkDeleteOptionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductConfiguratorBulkDeleteOptionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


