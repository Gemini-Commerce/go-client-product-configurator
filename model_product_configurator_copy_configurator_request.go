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

// checks if the ProductConfiguratorCopyConfiguratorRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductConfiguratorCopyConfiguratorRequest{}

// ProductConfiguratorCopyConfiguratorRequest struct for ProductConfiguratorCopyConfiguratorRequest
type ProductConfiguratorCopyConfiguratorRequest struct {
	TargetProductId *string `json:"targetProductId,omitempty"`
	CopyDependencies *bool `json:"copyDependencies,omitempty"`
	CopyMatrices *bool `json:"copyMatrices,omitempty"`
}

// NewProductConfiguratorCopyConfiguratorRequest instantiates a new ProductConfiguratorCopyConfiguratorRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductConfiguratorCopyConfiguratorRequest() *ProductConfiguratorCopyConfiguratorRequest {
	this := ProductConfiguratorCopyConfiguratorRequest{}
	return &this
}

// NewProductConfiguratorCopyConfiguratorRequestWithDefaults instantiates a new ProductConfiguratorCopyConfiguratorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductConfiguratorCopyConfiguratorRequestWithDefaults() *ProductConfiguratorCopyConfiguratorRequest {
	this := ProductConfiguratorCopyConfiguratorRequest{}
	return &this
}

// GetTargetProductId returns the TargetProductId field value if set, zero value otherwise.
func (o *ProductConfiguratorCopyConfiguratorRequest) GetTargetProductId() string {
	if o == nil || IsNil(o.TargetProductId) {
		var ret string
		return ret
	}
	return *o.TargetProductId
}

// GetTargetProductIdOk returns a tuple with the TargetProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductConfiguratorCopyConfiguratorRequest) GetTargetProductIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetProductId) {
		return nil, false
	}
	return o.TargetProductId, true
}

// HasTargetProductId returns a boolean if a field has been set.
func (o *ProductConfiguratorCopyConfiguratorRequest) HasTargetProductId() bool {
	if o != nil && !IsNil(o.TargetProductId) {
		return true
	}

	return false
}

// SetTargetProductId gets a reference to the given string and assigns it to the TargetProductId field.
func (o *ProductConfiguratorCopyConfiguratorRequest) SetTargetProductId(v string) {
	o.TargetProductId = &v
}

// GetCopyDependencies returns the CopyDependencies field value if set, zero value otherwise.
func (o *ProductConfiguratorCopyConfiguratorRequest) GetCopyDependencies() bool {
	if o == nil || IsNil(o.CopyDependencies) {
		var ret bool
		return ret
	}
	return *o.CopyDependencies
}

// GetCopyDependenciesOk returns a tuple with the CopyDependencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductConfiguratorCopyConfiguratorRequest) GetCopyDependenciesOk() (*bool, bool) {
	if o == nil || IsNil(o.CopyDependencies) {
		return nil, false
	}
	return o.CopyDependencies, true
}

// HasCopyDependencies returns a boolean if a field has been set.
func (o *ProductConfiguratorCopyConfiguratorRequest) HasCopyDependencies() bool {
	if o != nil && !IsNil(o.CopyDependencies) {
		return true
	}

	return false
}

// SetCopyDependencies gets a reference to the given bool and assigns it to the CopyDependencies field.
func (o *ProductConfiguratorCopyConfiguratorRequest) SetCopyDependencies(v bool) {
	o.CopyDependencies = &v
}

// GetCopyMatrices returns the CopyMatrices field value if set, zero value otherwise.
func (o *ProductConfiguratorCopyConfiguratorRequest) GetCopyMatrices() bool {
	if o == nil || IsNil(o.CopyMatrices) {
		var ret bool
		return ret
	}
	return *o.CopyMatrices
}

// GetCopyMatricesOk returns a tuple with the CopyMatrices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductConfiguratorCopyConfiguratorRequest) GetCopyMatricesOk() (*bool, bool) {
	if o == nil || IsNil(o.CopyMatrices) {
		return nil, false
	}
	return o.CopyMatrices, true
}

// HasCopyMatrices returns a boolean if a field has been set.
func (o *ProductConfiguratorCopyConfiguratorRequest) HasCopyMatrices() bool {
	if o != nil && !IsNil(o.CopyMatrices) {
		return true
	}

	return false
}

// SetCopyMatrices gets a reference to the given bool and assigns it to the CopyMatrices field.
func (o *ProductConfiguratorCopyConfiguratorRequest) SetCopyMatrices(v bool) {
	o.CopyMatrices = &v
}

func (o ProductConfiguratorCopyConfiguratorRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductConfiguratorCopyConfiguratorRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TargetProductId) {
		toSerialize["targetProductId"] = o.TargetProductId
	}
	if !IsNil(o.CopyDependencies) {
		toSerialize["copyDependencies"] = o.CopyDependencies
	}
	if !IsNil(o.CopyMatrices) {
		toSerialize["copyMatrices"] = o.CopyMatrices
	}
	return toSerialize, nil
}

type NullableProductConfiguratorCopyConfiguratorRequest struct {
	value *ProductConfiguratorCopyConfiguratorRequest
	isSet bool
}

func (v NullableProductConfiguratorCopyConfiguratorRequest) Get() *ProductConfiguratorCopyConfiguratorRequest {
	return v.value
}

func (v *NullableProductConfiguratorCopyConfiguratorRequest) Set(val *ProductConfiguratorCopyConfiguratorRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductConfiguratorCopyConfiguratorRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductConfiguratorCopyConfiguratorRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductConfiguratorCopyConfiguratorRequest(val *ProductConfiguratorCopyConfiguratorRequest) *NullableProductConfiguratorCopyConfiguratorRequest {
	return &NullableProductConfiguratorCopyConfiguratorRequest{value: val, isSet: true}
}

func (v NullableProductConfiguratorCopyConfiguratorRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductConfiguratorCopyConfiguratorRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


