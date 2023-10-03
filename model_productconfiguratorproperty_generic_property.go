/*
Product Configurator Service

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
Contact: info@gemini-commerce.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package product_configurator

import (
	"encoding/json"
)

// ProductconfiguratorpropertyGenericProperty struct for ProductconfiguratorpropertyGenericProperty
type ProductconfiguratorpropertyGenericProperty struct {
	PropertyKey *string `json:"propertyKey,omitempty"`
	PropertyValue *string `json:"propertyValue,omitempty"`
}

// NewProductconfiguratorpropertyGenericProperty instantiates a new ProductconfiguratorpropertyGenericProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductconfiguratorpropertyGenericProperty() *ProductconfiguratorpropertyGenericProperty {
	this := ProductconfiguratorpropertyGenericProperty{}
	return &this
}

// NewProductconfiguratorpropertyGenericPropertyWithDefaults instantiates a new ProductconfiguratorpropertyGenericProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductconfiguratorpropertyGenericPropertyWithDefaults() *ProductconfiguratorpropertyGenericProperty {
	this := ProductconfiguratorpropertyGenericProperty{}
	return &this
}

// GetPropertyKey returns the PropertyKey field value if set, zero value otherwise.
func (o *ProductconfiguratorpropertyGenericProperty) GetPropertyKey() string {
	if o == nil || isNil(o.PropertyKey) {
		var ret string
		return ret
	}
	return *o.PropertyKey
}

// GetPropertyKeyOk returns a tuple with the PropertyKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorpropertyGenericProperty) GetPropertyKeyOk() (*string, bool) {
	if o == nil || isNil(o.PropertyKey) {
    return nil, false
	}
	return o.PropertyKey, true
}

// HasPropertyKey returns a boolean if a field has been set.
func (o *ProductconfiguratorpropertyGenericProperty) HasPropertyKey() bool {
	if o != nil && !isNil(o.PropertyKey) {
		return true
	}

	return false
}

// SetPropertyKey gets a reference to the given string and assigns it to the PropertyKey field.
func (o *ProductconfiguratorpropertyGenericProperty) SetPropertyKey(v string) {
	o.PropertyKey = &v
}

// GetPropertyValue returns the PropertyValue field value if set, zero value otherwise.
func (o *ProductconfiguratorpropertyGenericProperty) GetPropertyValue() string {
	if o == nil || isNil(o.PropertyValue) {
		var ret string
		return ret
	}
	return *o.PropertyValue
}

// GetPropertyValueOk returns a tuple with the PropertyValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorpropertyGenericProperty) GetPropertyValueOk() (*string, bool) {
	if o == nil || isNil(o.PropertyValue) {
    return nil, false
	}
	return o.PropertyValue, true
}

// HasPropertyValue returns a boolean if a field has been set.
func (o *ProductconfiguratorpropertyGenericProperty) HasPropertyValue() bool {
	if o != nil && !isNil(o.PropertyValue) {
		return true
	}

	return false
}

// SetPropertyValue gets a reference to the given string and assigns it to the PropertyValue field.
func (o *ProductconfiguratorpropertyGenericProperty) SetPropertyValue(v string) {
	o.PropertyValue = &v
}

func (o ProductconfiguratorpropertyGenericProperty) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PropertyKey) {
		toSerialize["propertyKey"] = o.PropertyKey
	}
	if !isNil(o.PropertyValue) {
		toSerialize["propertyValue"] = o.PropertyValue
	}
	return json.Marshal(toSerialize)
}

type NullableProductconfiguratorpropertyGenericProperty struct {
	value *ProductconfiguratorpropertyGenericProperty
	isSet bool
}

func (v NullableProductconfiguratorpropertyGenericProperty) Get() *ProductconfiguratorpropertyGenericProperty {
	return v.value
}

func (v *NullableProductconfiguratorpropertyGenericProperty) Set(val *ProductconfiguratorpropertyGenericProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableProductconfiguratorpropertyGenericProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableProductconfiguratorpropertyGenericProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductconfiguratorpropertyGenericProperty(val *ProductconfiguratorpropertyGenericProperty) *NullableProductconfiguratorpropertyGenericProperty {
	return &NullableProductconfiguratorpropertyGenericProperty{value: val, isSet: true}
}

func (v NullableProductconfiguratorpropertyGenericProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductconfiguratorpropertyGenericProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

