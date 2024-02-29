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

// checks if the MatrixGenericType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MatrixGenericType{}

// MatrixGenericType struct for MatrixGenericType
type MatrixGenericType struct {
	PropertyKey *string `json:"propertyKey,omitempty"`
}

// NewMatrixGenericType instantiates a new MatrixGenericType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMatrixGenericType() *MatrixGenericType {
	this := MatrixGenericType{}
	return &this
}

// NewMatrixGenericTypeWithDefaults instantiates a new MatrixGenericType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMatrixGenericTypeWithDefaults() *MatrixGenericType {
	this := MatrixGenericType{}
	return &this
}

// GetPropertyKey returns the PropertyKey field value if set, zero value otherwise.
func (o *MatrixGenericType) GetPropertyKey() string {
	if o == nil || IsNil(o.PropertyKey) {
		var ret string
		return ret
	}
	return *o.PropertyKey
}

// GetPropertyKeyOk returns a tuple with the PropertyKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatrixGenericType) GetPropertyKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PropertyKey) {
		return nil, false
	}
	return o.PropertyKey, true
}

// HasPropertyKey returns a boolean if a field has been set.
func (o *MatrixGenericType) HasPropertyKey() bool {
	if o != nil && !IsNil(o.PropertyKey) {
		return true
	}

	return false
}

// SetPropertyKey gets a reference to the given string and assigns it to the PropertyKey field.
func (o *MatrixGenericType) SetPropertyKey(v string) {
	o.PropertyKey = &v
}

func (o MatrixGenericType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MatrixGenericType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PropertyKey) {
		toSerialize["propertyKey"] = o.PropertyKey
	}
	return toSerialize, nil
}

type NullableMatrixGenericType struct {
	value *MatrixGenericType
	isSet bool
}

func (v NullableMatrixGenericType) Get() *MatrixGenericType {
	return v.value
}

func (v *NullableMatrixGenericType) Set(val *MatrixGenericType) {
	v.value = val
	v.isSet = true
}

func (v NullableMatrixGenericType) IsSet() bool {
	return v.isSet
}

func (v *NullableMatrixGenericType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMatrixGenericType(val *MatrixGenericType) *NullableMatrixGenericType {
	return &NullableMatrixGenericType{value: val, isSet: true}
}

func (v NullableMatrixGenericType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMatrixGenericType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


