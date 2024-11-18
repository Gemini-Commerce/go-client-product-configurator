/*
Product Configurator Service

## Introduction  This comprehensive guide will equip you with the knowledge to integrate and leverage our Product Configurator Service in your applications.  ## Quick Start  Get up and running in no time! Follow these steps to kickstart your integration:  1. **Authentication:** Obtain your integration JWT to authenticate your requests. 2. **Client Libraries:** Explore our GitHub repositories to grab client libraries in your preferred programming language. 3. **API Overview:** Familiarize yourself with our RESTful API using the OpenAPI specification.  ## Integration  ### API Overview  Our RESTful API is the gateway to unlocking the full potential of Product Configurator. Check out the detailed [API Reference](/docs/category/configurator) for a granular understanding of each endpoint and request/response format.  ### Client Libraries  To expedite your integration process, we provide client libraries for various programming languages. Find the one that suits your stack in our [GitHub repositories](https://github.com/Gemini-Commerce).  ### Authentication  Security is paramount. Learn how to authenticate your requests using JWT. This ensures a secure and reliable connection between your application and Product Configurator.  ## Configuration Management  ### Configurator Lifecycle  Understand the lifecycle of configurators, from draft to active and deleted. This flexibility allows you to manage configurations at your own pace.  ### Steps and Options  Configure product steps with ease and define options effortlessly. Explore the power of dependencies to create dynamic and intuitive configurations.  ### Matrices  Delve into matrices—your secret weapon. Explore price and weight matrices, and learn how configured steps influence properties and pricing.  ### Price Management  Unleash dynamic pricing with our versatile price matrices. From fixed prices to incremental structures, adapt to diverse pricing models effortlessly.  ## Security  Your data is in safe hands. Discover how Product Configurator ensures security through JWT authentication, safeguarding your sensitive information.  ## Backward Compatibility  Stay ahead of the curve. Learn about our versioning strategy, providing backward compatibility while allowing our service to evolve seamlessly.  ## Developer Support  Have questions? Need assistance? Write to us at [info@gemini-commerce.com](mailto:info@gemini-commerce.com) and we will get back to you.

API version: v1
Contact: info@gemini-commerce.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package productconfigurator

import (
	"encoding/json"
)

// checks if the ProductconfiguratorMoney type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductconfiguratorMoney{}

// ProductconfiguratorMoney Represents an amount of money.
type ProductconfiguratorMoney struct {
	// The whole units of the amount. For example if `currencyCode` is `\"USD\"`, then 1 unit is one US dollar.
	Units *string `json:"units,omitempty"`
	// Number of micro (10^-6) units of the amount. The value must be between -999,999 and +999,999 inclusive. If `units` is positive, `micros` must be positive or zero. If `units` is zero, `micros` can be positive, zero, or negative. If `units` is negative, `micros` must be negative or zero. For example $-1.75 is represented as `units`=-1 and `micros`=-750,000.
	Micros               *int32 `json:"micros,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductconfiguratorMoney ProductconfiguratorMoney

// NewProductconfiguratorMoney instantiates a new ProductconfiguratorMoney object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductconfiguratorMoney() *ProductconfiguratorMoney {
	this := ProductconfiguratorMoney{}
	return &this
}

// NewProductconfiguratorMoneyWithDefaults instantiates a new ProductconfiguratorMoney object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductconfiguratorMoneyWithDefaults() *ProductconfiguratorMoney {
	this := ProductconfiguratorMoney{}
	return &this
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *ProductconfiguratorMoney) GetUnits() string {
	if o == nil || IsNil(o.Units) {
		var ret string
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorMoney) GetUnitsOk() (*string, bool) {
	if o == nil || IsNil(o.Units) {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *ProductconfiguratorMoney) HasUnits() bool {
	if o != nil && !IsNil(o.Units) {
		return true
	}

	return false
}

// SetUnits gets a reference to the given string and assigns it to the Units field.
func (o *ProductconfiguratorMoney) SetUnits(v string) {
	o.Units = &v
}

// GetMicros returns the Micros field value if set, zero value otherwise.
func (o *ProductconfiguratorMoney) GetMicros() int32 {
	if o == nil || IsNil(o.Micros) {
		var ret int32
		return ret
	}
	return *o.Micros
}

// GetMicrosOk returns a tuple with the Micros field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorMoney) GetMicrosOk() (*int32, bool) {
	if o == nil || IsNil(o.Micros) {
		return nil, false
	}
	return o.Micros, true
}

// HasMicros returns a boolean if a field has been set.
func (o *ProductconfiguratorMoney) HasMicros() bool {
	if o != nil && !IsNil(o.Micros) {
		return true
	}

	return false
}

// SetMicros gets a reference to the given int32 and assigns it to the Micros field.
func (o *ProductconfiguratorMoney) SetMicros(v int32) {
	o.Micros = &v
}

func (o ProductconfiguratorMoney) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductconfiguratorMoney) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Units) {
		toSerialize["units"] = o.Units
	}
	if !IsNil(o.Micros) {
		toSerialize["micros"] = o.Micros
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductconfiguratorMoney) UnmarshalJSON(data []byte) (err error) {
	varProductconfiguratorMoney := _ProductconfiguratorMoney{}

	err = json.Unmarshal(data, &varProductconfiguratorMoney)

	if err != nil {
		return err
	}

	*o = ProductconfiguratorMoney(varProductconfiguratorMoney)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "units")
		delete(additionalProperties, "micros")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ProductconfiguratorMoney) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *ProductconfiguratorMoney) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableProductconfiguratorMoney struct {
	value *ProductconfiguratorMoney
	isSet bool
}

func (v NullableProductconfiguratorMoney) Get() *ProductconfiguratorMoney {
	return v.value
}

func (v *NullableProductconfiguratorMoney) Set(val *ProductconfiguratorMoney) {
	v.value = val
	v.isSet = true
}

func (v NullableProductconfiguratorMoney) IsSet() bool {
	return v.isSet
}

func (v *NullableProductconfiguratorMoney) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductconfiguratorMoney(val *ProductconfiguratorMoney) *NullableProductconfiguratorMoney {
	return &NullableProductconfiguratorMoney{value: val, isSet: true}
}

func (v NullableProductconfiguratorMoney) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductconfiguratorMoney) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
