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

// ConfigurationOptionProperty struct for ConfigurationOptionProperty
type ConfigurationOptionProperty struct {
	PropertyId *string `json:"propertyId,omitempty"`
	PropertyValue *string `json:"propertyValue,omitempty"`
	PropertyType *ProductconfiguratorPropertyType `json:"propertyType,omitempty"`
	// subtract_to_get_variation is a list of values and is used to calculate the variation from the property value.
	SubtractToGetVariation []string `json:"subtractToGetVariation,omitempty"`
}

// NewConfigurationOptionProperty instantiates a new ConfigurationOptionProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigurationOptionProperty() *ConfigurationOptionProperty {
	this := ConfigurationOptionProperty{}
	var propertyType ProductconfiguratorPropertyType = UNKNOWN
	this.PropertyType = &propertyType
	return &this
}

// NewConfigurationOptionPropertyWithDefaults instantiates a new ConfigurationOptionProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurationOptionPropertyWithDefaults() *ConfigurationOptionProperty {
	this := ConfigurationOptionProperty{}
	var propertyType ProductconfiguratorPropertyType = UNKNOWN
	this.PropertyType = &propertyType
	return &this
}

// GetPropertyId returns the PropertyId field value if set, zero value otherwise.
func (o *ConfigurationOptionProperty) GetPropertyId() string {
	if o == nil || isNil(o.PropertyId) {
		var ret string
		return ret
	}
	return *o.PropertyId
}

// GetPropertyIdOk returns a tuple with the PropertyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationOptionProperty) GetPropertyIdOk() (*string, bool) {
	if o == nil || isNil(o.PropertyId) {
    return nil, false
	}
	return o.PropertyId, true
}

// HasPropertyId returns a boolean if a field has been set.
func (o *ConfigurationOptionProperty) HasPropertyId() bool {
	if o != nil && !isNil(o.PropertyId) {
		return true
	}

	return false
}

// SetPropertyId gets a reference to the given string and assigns it to the PropertyId field.
func (o *ConfigurationOptionProperty) SetPropertyId(v string) {
	o.PropertyId = &v
}

// GetPropertyValue returns the PropertyValue field value if set, zero value otherwise.
func (o *ConfigurationOptionProperty) GetPropertyValue() string {
	if o == nil || isNil(o.PropertyValue) {
		var ret string
		return ret
	}
	return *o.PropertyValue
}

// GetPropertyValueOk returns a tuple with the PropertyValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationOptionProperty) GetPropertyValueOk() (*string, bool) {
	if o == nil || isNil(o.PropertyValue) {
    return nil, false
	}
	return o.PropertyValue, true
}

// HasPropertyValue returns a boolean if a field has been set.
func (o *ConfigurationOptionProperty) HasPropertyValue() bool {
	if o != nil && !isNil(o.PropertyValue) {
		return true
	}

	return false
}

// SetPropertyValue gets a reference to the given string and assigns it to the PropertyValue field.
func (o *ConfigurationOptionProperty) SetPropertyValue(v string) {
	o.PropertyValue = &v
}

// GetPropertyType returns the PropertyType field value if set, zero value otherwise.
func (o *ConfigurationOptionProperty) GetPropertyType() ProductconfiguratorPropertyType {
	if o == nil || isNil(o.PropertyType) {
		var ret ProductconfiguratorPropertyType
		return ret
	}
	return *o.PropertyType
}

// GetPropertyTypeOk returns a tuple with the PropertyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationOptionProperty) GetPropertyTypeOk() (*ProductconfiguratorPropertyType, bool) {
	if o == nil || isNil(o.PropertyType) {
    return nil, false
	}
	return o.PropertyType, true
}

// HasPropertyType returns a boolean if a field has been set.
func (o *ConfigurationOptionProperty) HasPropertyType() bool {
	if o != nil && !isNil(o.PropertyType) {
		return true
	}

	return false
}

// SetPropertyType gets a reference to the given ProductconfiguratorPropertyType and assigns it to the PropertyType field.
func (o *ConfigurationOptionProperty) SetPropertyType(v ProductconfiguratorPropertyType) {
	o.PropertyType = &v
}

// GetSubtractToGetVariation returns the SubtractToGetVariation field value if set, zero value otherwise.
func (o *ConfigurationOptionProperty) GetSubtractToGetVariation() []string {
	if o == nil || isNil(o.SubtractToGetVariation) {
		var ret []string
		return ret
	}
	return o.SubtractToGetVariation
}

// GetSubtractToGetVariationOk returns a tuple with the SubtractToGetVariation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationOptionProperty) GetSubtractToGetVariationOk() ([]string, bool) {
	if o == nil || isNil(o.SubtractToGetVariation) {
    return nil, false
	}
	return o.SubtractToGetVariation, true
}

// HasSubtractToGetVariation returns a boolean if a field has been set.
func (o *ConfigurationOptionProperty) HasSubtractToGetVariation() bool {
	if o != nil && !isNil(o.SubtractToGetVariation) {
		return true
	}

	return false
}

// SetSubtractToGetVariation gets a reference to the given []string and assigns it to the SubtractToGetVariation field.
func (o *ConfigurationOptionProperty) SetSubtractToGetVariation(v []string) {
	o.SubtractToGetVariation = v
}

func (o ConfigurationOptionProperty) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PropertyId) {
		toSerialize["propertyId"] = o.PropertyId
	}
	if !isNil(o.PropertyValue) {
		toSerialize["propertyValue"] = o.PropertyValue
	}
	if !isNil(o.PropertyType) {
		toSerialize["propertyType"] = o.PropertyType
	}
	if !isNil(o.SubtractToGetVariation) {
		toSerialize["subtractToGetVariation"] = o.SubtractToGetVariation
	}
	return json.Marshal(toSerialize)
}

type NullableConfigurationOptionProperty struct {
	value *ConfigurationOptionProperty
	isSet bool
}

func (v NullableConfigurationOptionProperty) Get() *ConfigurationOptionProperty {
	return v.value
}

func (v *NullableConfigurationOptionProperty) Set(val *ConfigurationOptionProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurationOptionProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurationOptionProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurationOptionProperty(val *ConfigurationOptionProperty) *NullableConfigurationOptionProperty {
	return &NullableConfigurationOptionProperty{value: val, isSet: true}
}

func (v NullableConfigurationOptionProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurationOptionProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


