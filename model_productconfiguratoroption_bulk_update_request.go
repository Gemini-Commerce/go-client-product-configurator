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

// ProductconfiguratoroptionBulkUpdateRequest struct for ProductconfiguratoroptionBulkUpdateRequest
type ProductconfiguratoroptionBulkUpdateRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	Options []ProductconfiguratoroptionBulkUpdateRequestUpdateEntity `json:"options,omitempty"`
}

// NewProductconfiguratoroptionBulkUpdateRequest instantiates a new ProductconfiguratoroptionBulkUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductconfiguratoroptionBulkUpdateRequest() *ProductconfiguratoroptionBulkUpdateRequest {
	this := ProductconfiguratoroptionBulkUpdateRequest{}
	return &this
}

// NewProductconfiguratoroptionBulkUpdateRequestWithDefaults instantiates a new ProductconfiguratoroptionBulkUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductconfiguratoroptionBulkUpdateRequestWithDefaults() *ProductconfiguratoroptionBulkUpdateRequest {
	this := ProductconfiguratoroptionBulkUpdateRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *ProductconfiguratoroptionBulkUpdateRequest) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratoroptionBulkUpdateRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
    return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *ProductconfiguratoroptionBulkUpdateRequest) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *ProductconfiguratoroptionBulkUpdateRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *ProductconfiguratoroptionBulkUpdateRequest) GetOptions() []ProductconfiguratoroptionBulkUpdateRequestUpdateEntity {
	if o == nil || isNil(o.Options) {
		var ret []ProductconfiguratoroptionBulkUpdateRequestUpdateEntity
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratoroptionBulkUpdateRequest) GetOptionsOk() ([]ProductconfiguratoroptionBulkUpdateRequestUpdateEntity, bool) {
	if o == nil || isNil(o.Options) {
    return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ProductconfiguratoroptionBulkUpdateRequest) HasOptions() bool {
	if o != nil && !isNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []ProductconfiguratoroptionBulkUpdateRequestUpdateEntity and assigns it to the Options field.
func (o *ProductconfiguratoroptionBulkUpdateRequest) SetOptions(v []ProductconfiguratoroptionBulkUpdateRequestUpdateEntity) {
	o.Options = v
}

func (o ProductconfiguratoroptionBulkUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !isNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableProductconfiguratoroptionBulkUpdateRequest struct {
	value *ProductconfiguratoroptionBulkUpdateRequest
	isSet bool
}

func (v NullableProductconfiguratoroptionBulkUpdateRequest) Get() *ProductconfiguratoroptionBulkUpdateRequest {
	return v.value
}

func (v *NullableProductconfiguratoroptionBulkUpdateRequest) Set(val *ProductconfiguratoroptionBulkUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductconfiguratoroptionBulkUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductconfiguratoroptionBulkUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductconfiguratoroptionBulkUpdateRequest(val *ProductconfiguratoroptionBulkUpdateRequest) *NullableProductconfiguratoroptionBulkUpdateRequest {
	return &NullableProductconfiguratoroptionBulkUpdateRequest{value: val, isSet: true}
}

func (v NullableProductconfiguratoroptionBulkUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductconfiguratoroptionBulkUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


