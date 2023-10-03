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

// ConfigurationSelection struct for ConfigurationSelection
type ConfigurationSelection struct {
	StepId *string `json:"stepId,omitempty"`
	OptionId *string `json:"optionId,omitempty"`
}

// NewConfigurationSelection instantiates a new ConfigurationSelection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigurationSelection() *ConfigurationSelection {
	this := ConfigurationSelection{}
	return &this
}

// NewConfigurationSelectionWithDefaults instantiates a new ConfigurationSelection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurationSelectionWithDefaults() *ConfigurationSelection {
	this := ConfigurationSelection{}
	return &this
}

// GetStepId returns the StepId field value if set, zero value otherwise.
func (o *ConfigurationSelection) GetStepId() string {
	if o == nil || isNil(o.StepId) {
		var ret string
		return ret
	}
	return *o.StepId
}

// GetStepIdOk returns a tuple with the StepId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationSelection) GetStepIdOk() (*string, bool) {
	if o == nil || isNil(o.StepId) {
    return nil, false
	}
	return o.StepId, true
}

// HasStepId returns a boolean if a field has been set.
func (o *ConfigurationSelection) HasStepId() bool {
	if o != nil && !isNil(o.StepId) {
		return true
	}

	return false
}

// SetStepId gets a reference to the given string and assigns it to the StepId field.
func (o *ConfigurationSelection) SetStepId(v string) {
	o.StepId = &v
}

// GetOptionId returns the OptionId field value if set, zero value otherwise.
func (o *ConfigurationSelection) GetOptionId() string {
	if o == nil || isNil(o.OptionId) {
		var ret string
		return ret
	}
	return *o.OptionId
}

// GetOptionIdOk returns a tuple with the OptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationSelection) GetOptionIdOk() (*string, bool) {
	if o == nil || isNil(o.OptionId) {
    return nil, false
	}
	return o.OptionId, true
}

// HasOptionId returns a boolean if a field has been set.
func (o *ConfigurationSelection) HasOptionId() bool {
	if o != nil && !isNil(o.OptionId) {
		return true
	}

	return false
}

// SetOptionId gets a reference to the given string and assigns it to the OptionId field.
func (o *ConfigurationSelection) SetOptionId(v string) {
	o.OptionId = &v
}

func (o ConfigurationSelection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.StepId) {
		toSerialize["stepId"] = o.StepId
	}
	if !isNil(o.OptionId) {
		toSerialize["optionId"] = o.OptionId
	}
	return json.Marshal(toSerialize)
}

type NullableConfigurationSelection struct {
	value *ConfigurationSelection
	isSet bool
}

func (v NullableConfigurationSelection) Get() *ConfigurationSelection {
	return v.value
}

func (v *NullableConfigurationSelection) Set(val *ConfigurationSelection) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurationSelection) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurationSelection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurationSelection(val *ConfigurationSelection) *NullableConfigurationSelection {
	return &NullableConfigurationSelection{value: val, isSet: true}
}

func (v NullableConfigurationSelection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurationSelection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


