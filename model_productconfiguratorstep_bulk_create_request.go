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

// ProductconfiguratorstepBulkCreateRequest struct for ProductconfiguratorstepBulkCreateRequest
type ProductconfiguratorstepBulkCreateRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	ConfiguratorId *string `json:"configuratorId,omitempty"`
	Steps []ProductconfiguratorstepBulkCreateRequestCreateEntity `json:"steps,omitempty"`
}

// NewProductconfiguratorstepBulkCreateRequest instantiates a new ProductconfiguratorstepBulkCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductconfiguratorstepBulkCreateRequest() *ProductconfiguratorstepBulkCreateRequest {
	this := ProductconfiguratorstepBulkCreateRequest{}
	return &this
}

// NewProductconfiguratorstepBulkCreateRequestWithDefaults instantiates a new ProductconfiguratorstepBulkCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductconfiguratorstepBulkCreateRequestWithDefaults() *ProductconfiguratorstepBulkCreateRequest {
	this := ProductconfiguratorstepBulkCreateRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *ProductconfiguratorstepBulkCreateRequest) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorstepBulkCreateRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
    return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *ProductconfiguratorstepBulkCreateRequest) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *ProductconfiguratorstepBulkCreateRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetConfiguratorId returns the ConfiguratorId field value if set, zero value otherwise.
func (o *ProductconfiguratorstepBulkCreateRequest) GetConfiguratorId() string {
	if o == nil || isNil(o.ConfiguratorId) {
		var ret string
		return ret
	}
	return *o.ConfiguratorId
}

// GetConfiguratorIdOk returns a tuple with the ConfiguratorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorstepBulkCreateRequest) GetConfiguratorIdOk() (*string, bool) {
	if o == nil || isNil(o.ConfiguratorId) {
    return nil, false
	}
	return o.ConfiguratorId, true
}

// HasConfiguratorId returns a boolean if a field has been set.
func (o *ProductconfiguratorstepBulkCreateRequest) HasConfiguratorId() bool {
	if o != nil && !isNil(o.ConfiguratorId) {
		return true
	}

	return false
}

// SetConfiguratorId gets a reference to the given string and assigns it to the ConfiguratorId field.
func (o *ProductconfiguratorstepBulkCreateRequest) SetConfiguratorId(v string) {
	o.ConfiguratorId = &v
}

// GetSteps returns the Steps field value if set, zero value otherwise.
func (o *ProductconfiguratorstepBulkCreateRequest) GetSteps() []ProductconfiguratorstepBulkCreateRequestCreateEntity {
	if o == nil || isNil(o.Steps) {
		var ret []ProductconfiguratorstepBulkCreateRequestCreateEntity
		return ret
	}
	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorstepBulkCreateRequest) GetStepsOk() ([]ProductconfiguratorstepBulkCreateRequestCreateEntity, bool) {
	if o == nil || isNil(o.Steps) {
    return nil, false
	}
	return o.Steps, true
}

// HasSteps returns a boolean if a field has been set.
func (o *ProductconfiguratorstepBulkCreateRequest) HasSteps() bool {
	if o != nil && !isNil(o.Steps) {
		return true
	}

	return false
}

// SetSteps gets a reference to the given []ProductconfiguratorstepBulkCreateRequestCreateEntity and assigns it to the Steps field.
func (o *ProductconfiguratorstepBulkCreateRequest) SetSteps(v []ProductconfiguratorstepBulkCreateRequestCreateEntity) {
	o.Steps = v
}

func (o ProductconfiguratorstepBulkCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !isNil(o.ConfiguratorId) {
		toSerialize["configuratorId"] = o.ConfiguratorId
	}
	if !isNil(o.Steps) {
		toSerialize["steps"] = o.Steps
	}
	return json.Marshal(toSerialize)
}

type NullableProductconfiguratorstepBulkCreateRequest struct {
	value *ProductconfiguratorstepBulkCreateRequest
	isSet bool
}

func (v NullableProductconfiguratorstepBulkCreateRequest) Get() *ProductconfiguratorstepBulkCreateRequest {
	return v.value
}

func (v *NullableProductconfiguratorstepBulkCreateRequest) Set(val *ProductconfiguratorstepBulkCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductconfiguratorstepBulkCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductconfiguratorstepBulkCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductconfiguratorstepBulkCreateRequest(val *ProductconfiguratorstepBulkCreateRequest) *NullableProductconfiguratorstepBulkCreateRequest {
	return &NullableProductconfiguratorstepBulkCreateRequest{value: val, isSet: true}
}

func (v NullableProductconfiguratorstepBulkCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductconfiguratorstepBulkCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


