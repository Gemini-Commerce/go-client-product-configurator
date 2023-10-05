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

// OptionListOptionsRequest struct for OptionListOptionsRequest
type OptionListOptionsRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	StepId *string `json:"stepId,omitempty"`
	PageSize *int64 `json:"pageSize,omitempty"`
	PageToken *string `json:"pageToken,omitempty"`
}

// NewOptionListOptionsRequest instantiates a new OptionListOptionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptionListOptionsRequest() *OptionListOptionsRequest {
	this := OptionListOptionsRequest{}
	return &this
}

// NewOptionListOptionsRequestWithDefaults instantiates a new OptionListOptionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptionListOptionsRequestWithDefaults() *OptionListOptionsRequest {
	this := OptionListOptionsRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *OptionListOptionsRequest) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionListOptionsRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
    return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *OptionListOptionsRequest) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *OptionListOptionsRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetStepId returns the StepId field value if set, zero value otherwise.
func (o *OptionListOptionsRequest) GetStepId() string {
	if o == nil || isNil(o.StepId) {
		var ret string
		return ret
	}
	return *o.StepId
}

// GetStepIdOk returns a tuple with the StepId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionListOptionsRequest) GetStepIdOk() (*string, bool) {
	if o == nil || isNil(o.StepId) {
    return nil, false
	}
	return o.StepId, true
}

// HasStepId returns a boolean if a field has been set.
func (o *OptionListOptionsRequest) HasStepId() bool {
	if o != nil && !isNil(o.StepId) {
		return true
	}

	return false
}

// SetStepId gets a reference to the given string and assigns it to the StepId field.
func (o *OptionListOptionsRequest) SetStepId(v string) {
	o.StepId = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *OptionListOptionsRequest) GetPageSize() int64 {
	if o == nil || isNil(o.PageSize) {
		var ret int64
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionListOptionsRequest) GetPageSizeOk() (*int64, bool) {
	if o == nil || isNil(o.PageSize) {
    return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *OptionListOptionsRequest) HasPageSize() bool {
	if o != nil && !isNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int64 and assigns it to the PageSize field.
func (o *OptionListOptionsRequest) SetPageSize(v int64) {
	o.PageSize = &v
}

// GetPageToken returns the PageToken field value if set, zero value otherwise.
func (o *OptionListOptionsRequest) GetPageToken() string {
	if o == nil || isNil(o.PageToken) {
		var ret string
		return ret
	}
	return *o.PageToken
}

// GetPageTokenOk returns a tuple with the PageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionListOptionsRequest) GetPageTokenOk() (*string, bool) {
	if o == nil || isNil(o.PageToken) {
    return nil, false
	}
	return o.PageToken, true
}

// HasPageToken returns a boolean if a field has been set.
func (o *OptionListOptionsRequest) HasPageToken() bool {
	if o != nil && !isNil(o.PageToken) {
		return true
	}

	return false
}

// SetPageToken gets a reference to the given string and assigns it to the PageToken field.
func (o *OptionListOptionsRequest) SetPageToken(v string) {
	o.PageToken = &v
}

func (o OptionListOptionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !isNil(o.StepId) {
		toSerialize["stepId"] = o.StepId
	}
	if !isNil(o.PageSize) {
		toSerialize["pageSize"] = o.PageSize
	}
	if !isNil(o.PageToken) {
		toSerialize["pageToken"] = o.PageToken
	}
	return json.Marshal(toSerialize)
}

type NullableOptionListOptionsRequest struct {
	value *OptionListOptionsRequest
	isSet bool
}

func (v NullableOptionListOptionsRequest) Get() *OptionListOptionsRequest {
	return v.value
}

func (v *NullableOptionListOptionsRequest) Set(val *OptionListOptionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOptionListOptionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOptionListOptionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptionListOptionsRequest(val *OptionListOptionsRequest) *NullableOptionListOptionsRequest {
	return &NullableOptionListOptionsRequest{value: val, isSet: true}
}

func (v NullableOptionListOptionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptionListOptionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


