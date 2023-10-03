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

// LocalisationLocalizedText struct for LocalisationLocalizedText
type LocalisationLocalizedText struct {
	Value *map[string]string `json:"value,omitempty"`
}

// NewLocalisationLocalizedText instantiates a new LocalisationLocalizedText object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocalisationLocalizedText() *LocalisationLocalizedText {
	this := LocalisationLocalizedText{}
	return &this
}

// NewLocalisationLocalizedTextWithDefaults instantiates a new LocalisationLocalizedText object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocalisationLocalizedTextWithDefaults() *LocalisationLocalizedText {
	this := LocalisationLocalizedText{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *LocalisationLocalizedText) GetValue() map[string]string {
	if o == nil || isNil(o.Value) {
		var ret map[string]string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalisationLocalizedText) GetValueOk() (*map[string]string, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *LocalisationLocalizedText) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given map[string]string and assigns it to the Value field.
func (o *LocalisationLocalizedText) SetValue(v map[string]string) {
	o.Value = &v
}

func (o LocalisationLocalizedText) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableLocalisationLocalizedText struct {
	value *LocalisationLocalizedText
	isSet bool
}

func (v NullableLocalisationLocalizedText) Get() *LocalisationLocalizedText {
	return v.value
}

func (v *NullableLocalisationLocalizedText) Set(val *LocalisationLocalizedText) {
	v.value = val
	v.isSet = true
}

func (v NullableLocalisationLocalizedText) IsSet() bool {
	return v.isSet
}

func (v *NullableLocalisationLocalizedText) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocalisationLocalizedText(val *LocalisationLocalizedText) *NullableLocalisationLocalizedText {
	return &NullableLocalisationLocalizedText{value: val, isSet: true}
}

func (v NullableLocalisationLocalizedText) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocalisationLocalizedText) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


