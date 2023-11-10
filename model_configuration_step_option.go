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

// ConfigurationStepOption struct for ConfigurationStepOption
type ConfigurationStepOption struct {
	Id *string `json:"id,omitempty"`
	Grn *string `json:"grn,omitempty"`
	Label *LocalisationLocalizedText `json:"label,omitempty"`
	Swatch *OptionSwatch `json:"swatch,omitempty"`
	OptionProperties []ConfigurationOptionProperty `json:"optionProperties,omitempty"`
}

// NewConfigurationStepOption instantiates a new ConfigurationStepOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigurationStepOption() *ConfigurationStepOption {
	this := ConfigurationStepOption{}
	return &this
}

// NewConfigurationStepOptionWithDefaults instantiates a new ConfigurationStepOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurationStepOptionWithDefaults() *ConfigurationStepOption {
	this := ConfigurationStepOption{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ConfigurationStepOption) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationStepOption) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ConfigurationStepOption) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ConfigurationStepOption) SetId(v string) {
	o.Id = &v
}

// GetGrn returns the Grn field value if set, zero value otherwise.
func (o *ConfigurationStepOption) GetGrn() string {
	if o == nil || isNil(o.Grn) {
		var ret string
		return ret
	}
	return *o.Grn
}

// GetGrnOk returns a tuple with the Grn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationStepOption) GetGrnOk() (*string, bool) {
	if o == nil || isNil(o.Grn) {
    return nil, false
	}
	return o.Grn, true
}

// HasGrn returns a boolean if a field has been set.
func (o *ConfigurationStepOption) HasGrn() bool {
	if o != nil && !isNil(o.Grn) {
		return true
	}

	return false
}

// SetGrn gets a reference to the given string and assigns it to the Grn field.
func (o *ConfigurationStepOption) SetGrn(v string) {
	o.Grn = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ConfigurationStepOption) GetLabel() LocalisationLocalizedText {
	if o == nil || isNil(o.Label) {
		var ret LocalisationLocalizedText
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationStepOption) GetLabelOk() (*LocalisationLocalizedText, bool) {
	if o == nil || isNil(o.Label) {
    return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ConfigurationStepOption) HasLabel() bool {
	if o != nil && !isNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given LocalisationLocalizedText and assigns it to the Label field.
func (o *ConfigurationStepOption) SetLabel(v LocalisationLocalizedText) {
	o.Label = &v
}

// GetSwatch returns the Swatch field value if set, zero value otherwise.
func (o *ConfigurationStepOption) GetSwatch() OptionSwatch {
	if o == nil || isNil(o.Swatch) {
		var ret OptionSwatch
		return ret
	}
	return *o.Swatch
}

// GetSwatchOk returns a tuple with the Swatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationStepOption) GetSwatchOk() (*OptionSwatch, bool) {
	if o == nil || isNil(o.Swatch) {
    return nil, false
	}
	return o.Swatch, true
}

// HasSwatch returns a boolean if a field has been set.
func (o *ConfigurationStepOption) HasSwatch() bool {
	if o != nil && !isNil(o.Swatch) {
		return true
	}

	return false
}

// SetSwatch gets a reference to the given OptionSwatch and assigns it to the Swatch field.
func (o *ConfigurationStepOption) SetSwatch(v OptionSwatch) {
	o.Swatch = &v
}

// GetOptionProperties returns the OptionProperties field value if set, zero value otherwise.
func (o *ConfigurationStepOption) GetOptionProperties() []ConfigurationOptionProperty {
	if o == nil || isNil(o.OptionProperties) {
		var ret []ConfigurationOptionProperty
		return ret
	}
	return o.OptionProperties
}

// GetOptionPropertiesOk returns a tuple with the OptionProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationStepOption) GetOptionPropertiesOk() ([]ConfigurationOptionProperty, bool) {
	if o == nil || isNil(o.OptionProperties) {
    return nil, false
	}
	return o.OptionProperties, true
}

// HasOptionProperties returns a boolean if a field has been set.
func (o *ConfigurationStepOption) HasOptionProperties() bool {
	if o != nil && !isNil(o.OptionProperties) {
		return true
	}

	return false
}

// SetOptionProperties gets a reference to the given []ConfigurationOptionProperty and assigns it to the OptionProperties field.
func (o *ConfigurationStepOption) SetOptionProperties(v []ConfigurationOptionProperty) {
	o.OptionProperties = v
}

func (o ConfigurationStepOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Grn) {
		toSerialize["grn"] = o.Grn
	}
	if !isNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !isNil(o.Swatch) {
		toSerialize["swatch"] = o.Swatch
	}
	if !isNil(o.OptionProperties) {
		toSerialize["optionProperties"] = o.OptionProperties
	}
	return json.Marshal(toSerialize)
}

type NullableConfigurationStepOption struct {
	value *ConfigurationStepOption
	isSet bool
}

func (v NullableConfigurationStepOption) Get() *ConfigurationStepOption {
	return v.value
}

func (v *NullableConfigurationStepOption) Set(val *ConfigurationStepOption) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurationStepOption) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurationStepOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurationStepOption(val *ConfigurationStepOption) *NullableConfigurationStepOption {
	return &NullableConfigurationStepOption{value: val, isSet: true}
}

func (v NullableConfigurationStepOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurationStepOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


