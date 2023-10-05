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

// ProductconfiguratoroptionBulkDeleteRequest struct for ProductconfiguratoroptionBulkDeleteRequest
type ProductconfiguratoroptionBulkDeleteRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	OptionIds []string `json:"optionIds,omitempty"`
}

// NewProductconfiguratoroptionBulkDeleteRequest instantiates a new ProductconfiguratoroptionBulkDeleteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductconfiguratoroptionBulkDeleteRequest() *ProductconfiguratoroptionBulkDeleteRequest {
	this := ProductconfiguratoroptionBulkDeleteRequest{}
	return &this
}

// NewProductconfiguratoroptionBulkDeleteRequestWithDefaults instantiates a new ProductconfiguratoroptionBulkDeleteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductconfiguratoroptionBulkDeleteRequestWithDefaults() *ProductconfiguratoroptionBulkDeleteRequest {
	this := ProductconfiguratoroptionBulkDeleteRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *ProductconfiguratoroptionBulkDeleteRequest) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratoroptionBulkDeleteRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
    return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *ProductconfiguratoroptionBulkDeleteRequest) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *ProductconfiguratoroptionBulkDeleteRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetOptionIds returns the OptionIds field value if set, zero value otherwise.
func (o *ProductconfiguratoroptionBulkDeleteRequest) GetOptionIds() []string {
	if o == nil || isNil(o.OptionIds) {
		var ret []string
		return ret
	}
	return o.OptionIds
}

// GetOptionIdsOk returns a tuple with the OptionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratoroptionBulkDeleteRequest) GetOptionIdsOk() ([]string, bool) {
	if o == nil || isNil(o.OptionIds) {
    return nil, false
	}
	return o.OptionIds, true
}

// HasOptionIds returns a boolean if a field has been set.
func (o *ProductconfiguratoroptionBulkDeleteRequest) HasOptionIds() bool {
	if o != nil && !isNil(o.OptionIds) {
		return true
	}

	return false
}

// SetOptionIds gets a reference to the given []string and assigns it to the OptionIds field.
func (o *ProductconfiguratoroptionBulkDeleteRequest) SetOptionIds(v []string) {
	o.OptionIds = v
}

func (o ProductconfiguratoroptionBulkDeleteRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !isNil(o.OptionIds) {
		toSerialize["optionIds"] = o.OptionIds
	}
	return json.Marshal(toSerialize)
}

type NullableProductconfiguratoroptionBulkDeleteRequest struct {
	value *ProductconfiguratoroptionBulkDeleteRequest
	isSet bool
}

func (v NullableProductconfiguratoroptionBulkDeleteRequest) Get() *ProductconfiguratoroptionBulkDeleteRequest {
	return v.value
}

func (v *NullableProductconfiguratoroptionBulkDeleteRequest) Set(val *ProductconfiguratoroptionBulkDeleteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductconfiguratoroptionBulkDeleteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductconfiguratoroptionBulkDeleteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductconfiguratoroptionBulkDeleteRequest(val *ProductconfiguratoroptionBulkDeleteRequest) *NullableProductconfiguratoroptionBulkDeleteRequest {
	return &NullableProductconfiguratoroptionBulkDeleteRequest{value: val, isSet: true}
}

func (v NullableProductconfiguratoroptionBulkDeleteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductconfiguratoroptionBulkDeleteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


