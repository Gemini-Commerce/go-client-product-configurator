/*
Product Configurator Service

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
Contact: info@gemini-commerce.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ProductconfiguratorMoney Represents an amount of money.
type ProductconfiguratorMoney struct {
	// The whole units of the amount. For example if `currencyCode` is `\"USD\"`, then 1 unit is one US dollar.
	Units *string `json:"units,omitempty"`
	// Number of micro (10^-6) units of the amount. The value must be between -999,999 and +999,999 inclusive. If `units` is positive, `micros` must be positive or zero. If `units` is zero, `micros` can be positive, zero, or negative. If `units` is negative, `micros` must be negative or zero. For example $-1.75 is represented as `units`=-1 and `micros`=-750,000.
	Micros *int32 `json:"micros,omitempty"`
}

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
	if o == nil || isNil(o.Units) {
		var ret string
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorMoney) GetUnitsOk() (*string, bool) {
	if o == nil || isNil(o.Units) {
    return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *ProductconfiguratorMoney) HasUnits() bool {
	if o != nil && !isNil(o.Units) {
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
	if o == nil || isNil(o.Micros) {
		var ret int32
		return ret
	}
	return *o.Micros
}

// GetMicrosOk returns a tuple with the Micros field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorMoney) GetMicrosOk() (*int32, bool) {
	if o == nil || isNil(o.Micros) {
    return nil, false
	}
	return o.Micros, true
}

// HasMicros returns a boolean if a field has been set.
func (o *ProductconfiguratorMoney) HasMicros() bool {
	if o != nil && !isNil(o.Micros) {
		return true
	}

	return false
}

// SetMicros gets a reference to the given int32 and assigns it to the Micros field.
func (o *ProductconfiguratorMoney) SetMicros(v int32) {
	o.Micros = &v
}

func (o ProductconfiguratorMoney) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Units) {
		toSerialize["units"] = o.Units
	}
	if !isNil(o.Micros) {
		toSerialize["micros"] = o.Micros
	}
	return json.Marshal(toSerialize)
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

