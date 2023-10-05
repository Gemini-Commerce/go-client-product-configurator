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

// ProductconfiguratoroptionBulkUpdateRequestUpdateEntity struct for ProductconfiguratoroptionBulkUpdateRequestUpdateEntity
type ProductconfiguratoroptionBulkUpdateRequestUpdateEntity struct {
	OptionId *string `json:"optionId,omitempty"`
	Payload *ProductconfiguratoroptionUpdatePayload `json:"payload,omitempty"`
	PayloadMask []string `json:"payloadMask,omitempty"`
}

// NewProductconfiguratoroptionBulkUpdateRequestUpdateEntity instantiates a new ProductconfiguratoroptionBulkUpdateRequestUpdateEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductconfiguratoroptionBulkUpdateRequestUpdateEntity() *ProductconfiguratoroptionBulkUpdateRequestUpdateEntity {
	this := ProductconfiguratoroptionBulkUpdateRequestUpdateEntity{}
	return &this
}

// NewProductconfiguratoroptionBulkUpdateRequestUpdateEntityWithDefaults instantiates a new ProductconfiguratoroptionBulkUpdateRequestUpdateEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductconfiguratoroptionBulkUpdateRequestUpdateEntityWithDefaults() *ProductconfiguratoroptionBulkUpdateRequestUpdateEntity {
	this := ProductconfiguratoroptionBulkUpdateRequestUpdateEntity{}
	return &this
}

// GetOptionId returns the OptionId field value if set, zero value otherwise.
func (o *ProductconfiguratoroptionBulkUpdateRequestUpdateEntity) GetOptionId() string {
	if o == nil || isNil(o.OptionId) {
		var ret string
		return ret
	}
	return *o.OptionId
}

// GetOptionIdOk returns a tuple with the OptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratoroptionBulkUpdateRequestUpdateEntity) GetOptionIdOk() (*string, bool) {
	if o == nil || isNil(o.OptionId) {
    return nil, false
	}
	return o.OptionId, true
}

// HasOptionId returns a boolean if a field has been set.
func (o *ProductconfiguratoroptionBulkUpdateRequestUpdateEntity) HasOptionId() bool {
	if o != nil && !isNil(o.OptionId) {
		return true
	}

	return false
}

// SetOptionId gets a reference to the given string and assigns it to the OptionId field.
func (o *ProductconfiguratoroptionBulkUpdateRequestUpdateEntity) SetOptionId(v string) {
	o.OptionId = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *ProductconfiguratoroptionBulkUpdateRequestUpdateEntity) GetPayload() ProductconfiguratoroptionUpdatePayload {
	if o == nil || isNil(o.Payload) {
		var ret ProductconfiguratoroptionUpdatePayload
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratoroptionBulkUpdateRequestUpdateEntity) GetPayloadOk() (*ProductconfiguratoroptionUpdatePayload, bool) {
	if o == nil || isNil(o.Payload) {
    return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *ProductconfiguratoroptionBulkUpdateRequestUpdateEntity) HasPayload() bool {
	if o != nil && !isNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given ProductconfiguratoroptionUpdatePayload and assigns it to the Payload field.
func (o *ProductconfiguratoroptionBulkUpdateRequestUpdateEntity) SetPayload(v ProductconfiguratoroptionUpdatePayload) {
	o.Payload = &v
}

// GetPayloadMask returns the PayloadMask field value if set, zero value otherwise.
func (o *ProductconfiguratoroptionBulkUpdateRequestUpdateEntity) GetPayloadMask() []string {
	if o == nil || isNil(o.PayloadMask) {
		var ret []string
		return ret
	}
	return o.PayloadMask
}

// GetPayloadMaskOk returns a tuple with the PayloadMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratoroptionBulkUpdateRequestUpdateEntity) GetPayloadMaskOk() ([]string, bool) {
	if o == nil || isNil(o.PayloadMask) {
    return nil, false
	}
	return o.PayloadMask, true
}

// HasPayloadMask returns a boolean if a field has been set.
func (o *ProductconfiguratoroptionBulkUpdateRequestUpdateEntity) HasPayloadMask() bool {
	if o != nil && !isNil(o.PayloadMask) {
		return true
	}

	return false
}

// SetPayloadMask gets a reference to the given []string and assigns it to the PayloadMask field.
func (o *ProductconfiguratoroptionBulkUpdateRequestUpdateEntity) SetPayloadMask(v []string) {
	o.PayloadMask = v
}

func (o ProductconfiguratoroptionBulkUpdateRequestUpdateEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.OptionId) {
		toSerialize["optionId"] = o.OptionId
	}
	if !isNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !isNil(o.PayloadMask) {
		toSerialize["payloadMask"] = o.PayloadMask
	}
	return json.Marshal(toSerialize)
}

type NullableProductconfiguratoroptionBulkUpdateRequestUpdateEntity struct {
	value *ProductconfiguratoroptionBulkUpdateRequestUpdateEntity
	isSet bool
}

func (v NullableProductconfiguratoroptionBulkUpdateRequestUpdateEntity) Get() *ProductconfiguratoroptionBulkUpdateRequestUpdateEntity {
	return v.value
}

func (v *NullableProductconfiguratoroptionBulkUpdateRequestUpdateEntity) Set(val *ProductconfiguratoroptionBulkUpdateRequestUpdateEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableProductconfiguratoroptionBulkUpdateRequestUpdateEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableProductconfiguratoroptionBulkUpdateRequestUpdateEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductconfiguratoroptionBulkUpdateRequestUpdateEntity(val *ProductconfiguratoroptionBulkUpdateRequestUpdateEntity) *NullableProductconfiguratoroptionBulkUpdateRequestUpdateEntity {
	return &NullableProductconfiguratoroptionBulkUpdateRequestUpdateEntity{value: val, isSet: true}
}

func (v NullableProductconfiguratoroptionBulkUpdateRequestUpdateEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductconfiguratoroptionBulkUpdateRequestUpdateEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


