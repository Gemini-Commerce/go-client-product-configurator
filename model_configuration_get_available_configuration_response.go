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

// ConfigurationGetAvailableConfigurationResponse struct for ConfigurationGetAvailableConfigurationResponse
type ConfigurationGetAvailableConfigurationResponse struct {
	Configurator *ConfigurationConfigurator `json:"configurator,omitempty"`
	Properties []ConfigurationProperty `json:"properties,omitempty"`
	MatchedProperties []ConfigurationProperty `json:"matchedProperties,omitempty"`
	Selections []ConfigurationSelection `json:"selections,omitempty"`
}

// NewConfigurationGetAvailableConfigurationResponse instantiates a new ConfigurationGetAvailableConfigurationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigurationGetAvailableConfigurationResponse() *ConfigurationGetAvailableConfigurationResponse {
	this := ConfigurationGetAvailableConfigurationResponse{}
	return &this
}

// NewConfigurationGetAvailableConfigurationResponseWithDefaults instantiates a new ConfigurationGetAvailableConfigurationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurationGetAvailableConfigurationResponseWithDefaults() *ConfigurationGetAvailableConfigurationResponse {
	this := ConfigurationGetAvailableConfigurationResponse{}
	return &this
}

// GetConfigurator returns the Configurator field value if set, zero value otherwise.
func (o *ConfigurationGetAvailableConfigurationResponse) GetConfigurator() ConfigurationConfigurator {
	if o == nil || isNil(o.Configurator) {
		var ret ConfigurationConfigurator
		return ret
	}
	return *o.Configurator
}

// GetConfiguratorOk returns a tuple with the Configurator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationGetAvailableConfigurationResponse) GetConfiguratorOk() (*ConfigurationConfigurator, bool) {
	if o == nil || isNil(o.Configurator) {
    return nil, false
	}
	return o.Configurator, true
}

// HasConfigurator returns a boolean if a field has been set.
func (o *ConfigurationGetAvailableConfigurationResponse) HasConfigurator() bool {
	if o != nil && !isNil(o.Configurator) {
		return true
	}

	return false
}

// SetConfigurator gets a reference to the given ConfigurationConfigurator and assigns it to the Configurator field.
func (o *ConfigurationGetAvailableConfigurationResponse) SetConfigurator(v ConfigurationConfigurator) {
	o.Configurator = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *ConfigurationGetAvailableConfigurationResponse) GetProperties() []ConfigurationProperty {
	if o == nil || isNil(o.Properties) {
		var ret []ConfigurationProperty
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationGetAvailableConfigurationResponse) GetPropertiesOk() ([]ConfigurationProperty, bool) {
	if o == nil || isNil(o.Properties) {
    return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *ConfigurationGetAvailableConfigurationResponse) HasProperties() bool {
	if o != nil && !isNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []ConfigurationProperty and assigns it to the Properties field.
func (o *ConfigurationGetAvailableConfigurationResponse) SetProperties(v []ConfigurationProperty) {
	o.Properties = v
}

// GetMatchedProperties returns the MatchedProperties field value if set, zero value otherwise.
func (o *ConfigurationGetAvailableConfigurationResponse) GetMatchedProperties() []ConfigurationProperty {
	if o == nil || isNil(o.MatchedProperties) {
		var ret []ConfigurationProperty
		return ret
	}
	return o.MatchedProperties
}

// GetMatchedPropertiesOk returns a tuple with the MatchedProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationGetAvailableConfigurationResponse) GetMatchedPropertiesOk() ([]ConfigurationProperty, bool) {
	if o == nil || isNil(o.MatchedProperties) {
    return nil, false
	}
	return o.MatchedProperties, true
}

// HasMatchedProperties returns a boolean if a field has been set.
func (o *ConfigurationGetAvailableConfigurationResponse) HasMatchedProperties() bool {
	if o != nil && !isNil(o.MatchedProperties) {
		return true
	}

	return false
}

// SetMatchedProperties gets a reference to the given []ConfigurationProperty and assigns it to the MatchedProperties field.
func (o *ConfigurationGetAvailableConfigurationResponse) SetMatchedProperties(v []ConfigurationProperty) {
	o.MatchedProperties = v
}

// GetSelections returns the Selections field value if set, zero value otherwise.
func (o *ConfigurationGetAvailableConfigurationResponse) GetSelections() []ConfigurationSelection {
	if o == nil || isNil(o.Selections) {
		var ret []ConfigurationSelection
		return ret
	}
	return o.Selections
}

// GetSelectionsOk returns a tuple with the Selections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationGetAvailableConfigurationResponse) GetSelectionsOk() ([]ConfigurationSelection, bool) {
	if o == nil || isNil(o.Selections) {
    return nil, false
	}
	return o.Selections, true
}

// HasSelections returns a boolean if a field has been set.
func (o *ConfigurationGetAvailableConfigurationResponse) HasSelections() bool {
	if o != nil && !isNil(o.Selections) {
		return true
	}

	return false
}

// SetSelections gets a reference to the given []ConfigurationSelection and assigns it to the Selections field.
func (o *ConfigurationGetAvailableConfigurationResponse) SetSelections(v []ConfigurationSelection) {
	o.Selections = v
}

func (o ConfigurationGetAvailableConfigurationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Configurator) {
		toSerialize["configurator"] = o.Configurator
	}
	if !isNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	if !isNil(o.MatchedProperties) {
		toSerialize["matchedProperties"] = o.MatchedProperties
	}
	if !isNil(o.Selections) {
		toSerialize["selections"] = o.Selections
	}
	return json.Marshal(toSerialize)
}

type NullableConfigurationGetAvailableConfigurationResponse struct {
	value *ConfigurationGetAvailableConfigurationResponse
	isSet bool
}

func (v NullableConfigurationGetAvailableConfigurationResponse) Get() *ConfigurationGetAvailableConfigurationResponse {
	return v.value
}

func (v *NullableConfigurationGetAvailableConfigurationResponse) Set(val *ConfigurationGetAvailableConfigurationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurationGetAvailableConfigurationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurationGetAvailableConfigurationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurationGetAvailableConfigurationResponse(val *ConfigurationGetAvailableConfigurationResponse) *NullableConfigurationGetAvailableConfigurationResponse {
	return &NullableConfigurationGetAvailableConfigurationResponse{value: val, isSet: true}
}

func (v NullableConfigurationGetAvailableConfigurationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurationGetAvailableConfigurationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


