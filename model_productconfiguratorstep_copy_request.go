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

// ProductconfiguratorstepCopyRequest struct for ProductconfiguratorstepCopyRequest
type ProductconfiguratorstepCopyRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	SourceStepId *string `json:"sourceStepId,omitempty"`
	TargetConfiguratorId *string `json:"targetConfiguratorId,omitempty"`
}

// NewProductconfiguratorstepCopyRequest instantiates a new ProductconfiguratorstepCopyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductconfiguratorstepCopyRequest() *ProductconfiguratorstepCopyRequest {
	this := ProductconfiguratorstepCopyRequest{}
	return &this
}

// NewProductconfiguratorstepCopyRequestWithDefaults instantiates a new ProductconfiguratorstepCopyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductconfiguratorstepCopyRequestWithDefaults() *ProductconfiguratorstepCopyRequest {
	this := ProductconfiguratorstepCopyRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *ProductconfiguratorstepCopyRequest) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorstepCopyRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
    return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *ProductconfiguratorstepCopyRequest) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *ProductconfiguratorstepCopyRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetSourceStepId returns the SourceStepId field value if set, zero value otherwise.
func (o *ProductconfiguratorstepCopyRequest) GetSourceStepId() string {
	if o == nil || isNil(o.SourceStepId) {
		var ret string
		return ret
	}
	return *o.SourceStepId
}

// GetSourceStepIdOk returns a tuple with the SourceStepId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorstepCopyRequest) GetSourceStepIdOk() (*string, bool) {
	if o == nil || isNil(o.SourceStepId) {
    return nil, false
	}
	return o.SourceStepId, true
}

// HasSourceStepId returns a boolean if a field has been set.
func (o *ProductconfiguratorstepCopyRequest) HasSourceStepId() bool {
	if o != nil && !isNil(o.SourceStepId) {
		return true
	}

	return false
}

// SetSourceStepId gets a reference to the given string and assigns it to the SourceStepId field.
func (o *ProductconfiguratorstepCopyRequest) SetSourceStepId(v string) {
	o.SourceStepId = &v
}

// GetTargetConfiguratorId returns the TargetConfiguratorId field value if set, zero value otherwise.
func (o *ProductconfiguratorstepCopyRequest) GetTargetConfiguratorId() string {
	if o == nil || isNil(o.TargetConfiguratorId) {
		var ret string
		return ret
	}
	return *o.TargetConfiguratorId
}

// GetTargetConfiguratorIdOk returns a tuple with the TargetConfiguratorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorstepCopyRequest) GetTargetConfiguratorIdOk() (*string, bool) {
	if o == nil || isNil(o.TargetConfiguratorId) {
    return nil, false
	}
	return o.TargetConfiguratorId, true
}

// HasTargetConfiguratorId returns a boolean if a field has been set.
func (o *ProductconfiguratorstepCopyRequest) HasTargetConfiguratorId() bool {
	if o != nil && !isNil(o.TargetConfiguratorId) {
		return true
	}

	return false
}

// SetTargetConfiguratorId gets a reference to the given string and assigns it to the TargetConfiguratorId field.
func (o *ProductconfiguratorstepCopyRequest) SetTargetConfiguratorId(v string) {
	o.TargetConfiguratorId = &v
}

func (o ProductconfiguratorstepCopyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !isNil(o.SourceStepId) {
		toSerialize["sourceStepId"] = o.SourceStepId
	}
	if !isNil(o.TargetConfiguratorId) {
		toSerialize["targetConfiguratorId"] = o.TargetConfiguratorId
	}
	return json.Marshal(toSerialize)
}

type NullableProductconfiguratorstepCopyRequest struct {
	value *ProductconfiguratorstepCopyRequest
	isSet bool
}

func (v NullableProductconfiguratorstepCopyRequest) Get() *ProductconfiguratorstepCopyRequest {
	return v.value
}

func (v *NullableProductconfiguratorstepCopyRequest) Set(val *ProductconfiguratorstepCopyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductconfiguratorstepCopyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductconfiguratorstepCopyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductconfiguratorstepCopyRequest(val *ProductconfiguratorstepCopyRequest) *NullableProductconfiguratorstepCopyRequest {
	return &NullableProductconfiguratorstepCopyRequest{value: val, isSet: true}
}

func (v NullableProductconfiguratorstepCopyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductconfiguratorstepCopyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

