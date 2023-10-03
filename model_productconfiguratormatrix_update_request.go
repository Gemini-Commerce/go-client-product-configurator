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

// ProductconfiguratormatrixUpdateRequest struct for ProductconfiguratormatrixUpdateRequest
type ProductconfiguratormatrixUpdateRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	MatrixId *string `json:"matrixId,omitempty"`
	Payload *ProductconfiguratormatrixUpdatePayload `json:"payload,omitempty"`
	PayloadMask []string `json:"payloadMask,omitempty"`
}

// NewProductconfiguratormatrixUpdateRequest instantiates a new ProductconfiguratormatrixUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductconfiguratormatrixUpdateRequest() *ProductconfiguratormatrixUpdateRequest {
	this := ProductconfiguratormatrixUpdateRequest{}
	return &this
}

// NewProductconfiguratormatrixUpdateRequestWithDefaults instantiates a new ProductconfiguratormatrixUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductconfiguratormatrixUpdateRequestWithDefaults() *ProductconfiguratormatrixUpdateRequest {
	this := ProductconfiguratormatrixUpdateRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *ProductconfiguratormatrixUpdateRequest) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratormatrixUpdateRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
    return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *ProductconfiguratormatrixUpdateRequest) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *ProductconfiguratormatrixUpdateRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetMatrixId returns the MatrixId field value if set, zero value otherwise.
func (o *ProductconfiguratormatrixUpdateRequest) GetMatrixId() string {
	if o == nil || isNil(o.MatrixId) {
		var ret string
		return ret
	}
	return *o.MatrixId
}

// GetMatrixIdOk returns a tuple with the MatrixId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratormatrixUpdateRequest) GetMatrixIdOk() (*string, bool) {
	if o == nil || isNil(o.MatrixId) {
    return nil, false
	}
	return o.MatrixId, true
}

// HasMatrixId returns a boolean if a field has been set.
func (o *ProductconfiguratormatrixUpdateRequest) HasMatrixId() bool {
	if o != nil && !isNil(o.MatrixId) {
		return true
	}

	return false
}

// SetMatrixId gets a reference to the given string and assigns it to the MatrixId field.
func (o *ProductconfiguratormatrixUpdateRequest) SetMatrixId(v string) {
	o.MatrixId = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *ProductconfiguratormatrixUpdateRequest) GetPayload() ProductconfiguratormatrixUpdatePayload {
	if o == nil || isNil(o.Payload) {
		var ret ProductconfiguratormatrixUpdatePayload
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratormatrixUpdateRequest) GetPayloadOk() (*ProductconfiguratormatrixUpdatePayload, bool) {
	if o == nil || isNil(o.Payload) {
    return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *ProductconfiguratormatrixUpdateRequest) HasPayload() bool {
	if o != nil && !isNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given ProductconfiguratormatrixUpdatePayload and assigns it to the Payload field.
func (o *ProductconfiguratormatrixUpdateRequest) SetPayload(v ProductconfiguratormatrixUpdatePayload) {
	o.Payload = &v
}

// GetPayloadMask returns the PayloadMask field value if set, zero value otherwise.
func (o *ProductconfiguratormatrixUpdateRequest) GetPayloadMask() []string {
	if o == nil || isNil(o.PayloadMask) {
		var ret []string
		return ret
	}
	return o.PayloadMask
}

// GetPayloadMaskOk returns a tuple with the PayloadMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratormatrixUpdateRequest) GetPayloadMaskOk() ([]string, bool) {
	if o == nil || isNil(o.PayloadMask) {
    return nil, false
	}
	return o.PayloadMask, true
}

// HasPayloadMask returns a boolean if a field has been set.
func (o *ProductconfiguratormatrixUpdateRequest) HasPayloadMask() bool {
	if o != nil && !isNil(o.PayloadMask) {
		return true
	}

	return false
}

// SetPayloadMask gets a reference to the given []string and assigns it to the PayloadMask field.
func (o *ProductconfiguratormatrixUpdateRequest) SetPayloadMask(v []string) {
	o.PayloadMask = v
}

func (o ProductconfiguratormatrixUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !isNil(o.MatrixId) {
		toSerialize["matrixId"] = o.MatrixId
	}
	if !isNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !isNil(o.PayloadMask) {
		toSerialize["payloadMask"] = o.PayloadMask
	}
	return json.Marshal(toSerialize)
}

type NullableProductconfiguratormatrixUpdateRequest struct {
	value *ProductconfiguratormatrixUpdateRequest
	isSet bool
}

func (v NullableProductconfiguratormatrixUpdateRequest) Get() *ProductconfiguratormatrixUpdateRequest {
	return v.value
}

func (v *NullableProductconfiguratormatrixUpdateRequest) Set(val *ProductconfiguratormatrixUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductconfiguratormatrixUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductconfiguratormatrixUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductconfiguratormatrixUpdateRequest(val *ProductconfiguratormatrixUpdateRequest) *NullableProductconfiguratormatrixUpdateRequest {
	return &NullableProductconfiguratormatrixUpdateRequest{value: val, isSet: true}
}

func (v NullableProductconfiguratormatrixUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductconfiguratormatrixUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


