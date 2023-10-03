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

// ProductconfiguratormatrixUpdatePayload struct for ProductconfiguratormatrixUpdatePayload
type ProductconfiguratormatrixUpdatePayload struct {
	Label *string `json:"label,omitempty"`
	// default_property_id is the id of the property that will be used as the starting point to calculate the differences between the properties.
	DefaultPropertyId *string `json:"defaultPropertyId,omitempty"`
}

// NewProductconfiguratormatrixUpdatePayload instantiates a new ProductconfiguratormatrixUpdatePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductconfiguratormatrixUpdatePayload() *ProductconfiguratormatrixUpdatePayload {
	this := ProductconfiguratormatrixUpdatePayload{}
	return &this
}

// NewProductconfiguratormatrixUpdatePayloadWithDefaults instantiates a new ProductconfiguratormatrixUpdatePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductconfiguratormatrixUpdatePayloadWithDefaults() *ProductconfiguratormatrixUpdatePayload {
	this := ProductconfiguratormatrixUpdatePayload{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ProductconfiguratormatrixUpdatePayload) GetLabel() string {
	if o == nil || isNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratormatrixUpdatePayload) GetLabelOk() (*string, bool) {
	if o == nil || isNil(o.Label) {
    return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ProductconfiguratormatrixUpdatePayload) HasLabel() bool {
	if o != nil && !isNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ProductconfiguratormatrixUpdatePayload) SetLabel(v string) {
	o.Label = &v
}

// GetDefaultPropertyId returns the DefaultPropertyId field value if set, zero value otherwise.
func (o *ProductconfiguratormatrixUpdatePayload) GetDefaultPropertyId() string {
	if o == nil || isNil(o.DefaultPropertyId) {
		var ret string
		return ret
	}
	return *o.DefaultPropertyId
}

// GetDefaultPropertyIdOk returns a tuple with the DefaultPropertyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratormatrixUpdatePayload) GetDefaultPropertyIdOk() (*string, bool) {
	if o == nil || isNil(o.DefaultPropertyId) {
    return nil, false
	}
	return o.DefaultPropertyId, true
}

// HasDefaultPropertyId returns a boolean if a field has been set.
func (o *ProductconfiguratormatrixUpdatePayload) HasDefaultPropertyId() bool {
	if o != nil && !isNil(o.DefaultPropertyId) {
		return true
	}

	return false
}

// SetDefaultPropertyId gets a reference to the given string and assigns it to the DefaultPropertyId field.
func (o *ProductconfiguratormatrixUpdatePayload) SetDefaultPropertyId(v string) {
	o.DefaultPropertyId = &v
}

func (o ProductconfiguratormatrixUpdatePayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !isNil(o.DefaultPropertyId) {
		toSerialize["defaultPropertyId"] = o.DefaultPropertyId
	}
	return json.Marshal(toSerialize)
}

type NullableProductconfiguratormatrixUpdatePayload struct {
	value *ProductconfiguratormatrixUpdatePayload
	isSet bool
}

func (v NullableProductconfiguratormatrixUpdatePayload) Get() *ProductconfiguratormatrixUpdatePayload {
	return v.value
}

func (v *NullableProductconfiguratormatrixUpdatePayload) Set(val *ProductconfiguratormatrixUpdatePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableProductconfiguratormatrixUpdatePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableProductconfiguratormatrixUpdatePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductconfiguratormatrixUpdatePayload(val *ProductconfiguratormatrixUpdatePayload) *NullableProductconfiguratormatrixUpdatePayload {
	return &NullableProductconfiguratormatrixUpdatePayload{value: val, isSet: true}
}

func (v NullableProductconfiguratormatrixUpdatePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductconfiguratormatrixUpdatePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

