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

// ProductconfiguratorconfiguratorUpdatePayload struct for ProductconfiguratorconfiguratorUpdatePayload
type ProductconfiguratorconfiguratorUpdatePayload struct {
	Label *string `json:"label,omitempty"`
	Status *ProductconfiguratorconfiguratorStatus `json:"status,omitempty"`
}

// NewProductconfiguratorconfiguratorUpdatePayload instantiates a new ProductconfiguratorconfiguratorUpdatePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductconfiguratorconfiguratorUpdatePayload() *ProductconfiguratorconfiguratorUpdatePayload {
	this := ProductconfiguratorconfiguratorUpdatePayload{}
	var status ProductconfiguratorconfiguratorStatus = UNKNOWN
	this.Status = &status
	return &this
}

// NewProductconfiguratorconfiguratorUpdatePayloadWithDefaults instantiates a new ProductconfiguratorconfiguratorUpdatePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductconfiguratorconfiguratorUpdatePayloadWithDefaults() *ProductconfiguratorconfiguratorUpdatePayload {
	this := ProductconfiguratorconfiguratorUpdatePayload{}
	var status ProductconfiguratorconfiguratorStatus = UNKNOWN
	this.Status = &status
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ProductconfiguratorconfiguratorUpdatePayload) GetLabel() string {
	if o == nil || isNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorconfiguratorUpdatePayload) GetLabelOk() (*string, bool) {
	if o == nil || isNil(o.Label) {
    return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ProductconfiguratorconfiguratorUpdatePayload) HasLabel() bool {
	if o != nil && !isNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ProductconfiguratorconfiguratorUpdatePayload) SetLabel(v string) {
	o.Label = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ProductconfiguratorconfiguratorUpdatePayload) GetStatus() ProductconfiguratorconfiguratorStatus {
	if o == nil || isNil(o.Status) {
		var ret ProductconfiguratorconfiguratorStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorconfiguratorUpdatePayload) GetStatusOk() (*ProductconfiguratorconfiguratorStatus, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ProductconfiguratorconfiguratorUpdatePayload) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ProductconfiguratorconfiguratorStatus and assigns it to the Status field.
func (o *ProductconfiguratorconfiguratorUpdatePayload) SetStatus(v ProductconfiguratorconfiguratorStatus) {
	o.Status = &v
}

func (o ProductconfiguratorconfiguratorUpdatePayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableProductconfiguratorconfiguratorUpdatePayload struct {
	value *ProductconfiguratorconfiguratorUpdatePayload
	isSet bool
}

func (v NullableProductconfiguratorconfiguratorUpdatePayload) Get() *ProductconfiguratorconfiguratorUpdatePayload {
	return v.value
}

func (v *NullableProductconfiguratorconfiguratorUpdatePayload) Set(val *ProductconfiguratorconfiguratorUpdatePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableProductconfiguratorconfiguratorUpdatePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableProductconfiguratorconfiguratorUpdatePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductconfiguratorconfiguratorUpdatePayload(val *ProductconfiguratorconfiguratorUpdatePayload) *NullableProductconfiguratorconfiguratorUpdatePayload {
	return &NullableProductconfiguratorconfiguratorUpdatePayload{value: val, isSet: true}
}

func (v NullableProductconfiguratorconfiguratorUpdatePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductconfiguratorconfiguratorUpdatePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


