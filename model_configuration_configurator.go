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

// ConfigurationConfigurator struct for ConfigurationConfigurator
type ConfigurationConfigurator struct {
	Id *string `json:"id,omitempty"`
	Grn *string `json:"grn,omitempty"`
	Steps []ConfigurationStep `json:"steps,omitempty"`
}

// NewConfigurationConfigurator instantiates a new ConfigurationConfigurator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigurationConfigurator() *ConfigurationConfigurator {
	this := ConfigurationConfigurator{}
	return &this
}

// NewConfigurationConfiguratorWithDefaults instantiates a new ConfigurationConfigurator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurationConfiguratorWithDefaults() *ConfigurationConfigurator {
	this := ConfigurationConfigurator{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ConfigurationConfigurator) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationConfigurator) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ConfigurationConfigurator) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ConfigurationConfigurator) SetId(v string) {
	o.Id = &v
}

// GetGrn returns the Grn field value if set, zero value otherwise.
func (o *ConfigurationConfigurator) GetGrn() string {
	if o == nil || isNil(o.Grn) {
		var ret string
		return ret
	}
	return *o.Grn
}

// GetGrnOk returns a tuple with the Grn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationConfigurator) GetGrnOk() (*string, bool) {
	if o == nil || isNil(o.Grn) {
    return nil, false
	}
	return o.Grn, true
}

// HasGrn returns a boolean if a field has been set.
func (o *ConfigurationConfigurator) HasGrn() bool {
	if o != nil && !isNil(o.Grn) {
		return true
	}

	return false
}

// SetGrn gets a reference to the given string and assigns it to the Grn field.
func (o *ConfigurationConfigurator) SetGrn(v string) {
	o.Grn = &v
}

// GetSteps returns the Steps field value if set, zero value otherwise.
func (o *ConfigurationConfigurator) GetSteps() []ConfigurationStep {
	if o == nil || isNil(o.Steps) {
		var ret []ConfigurationStep
		return ret
	}
	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationConfigurator) GetStepsOk() ([]ConfigurationStep, bool) {
	if o == nil || isNil(o.Steps) {
    return nil, false
	}
	return o.Steps, true
}

// HasSteps returns a boolean if a field has been set.
func (o *ConfigurationConfigurator) HasSteps() bool {
	if o != nil && !isNil(o.Steps) {
		return true
	}

	return false
}

// SetSteps gets a reference to the given []ConfigurationStep and assigns it to the Steps field.
func (o *ConfigurationConfigurator) SetSteps(v []ConfigurationStep) {
	o.Steps = v
}

func (o ConfigurationConfigurator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Grn) {
		toSerialize["grn"] = o.Grn
	}
	if !isNil(o.Steps) {
		toSerialize["steps"] = o.Steps
	}
	return json.Marshal(toSerialize)
}

type NullableConfigurationConfigurator struct {
	value *ConfigurationConfigurator
	isSet bool
}

func (v NullableConfigurationConfigurator) Get() *ConfigurationConfigurator {
	return v.value
}

func (v *NullableConfigurationConfigurator) Set(val *ConfigurationConfigurator) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurationConfigurator) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurationConfigurator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurationConfigurator(val *ConfigurationConfigurator) *NullableConfigurationConfigurator {
	return &NullableConfigurationConfigurator{value: val, isSet: true}
}

func (v NullableConfigurationConfigurator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurationConfigurator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

