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

// MatrixListMatricesRequest struct for MatrixListMatricesRequest
type MatrixListMatricesRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	ConfiguratorId *string `json:"configuratorId,omitempty"`
	FilterMask []string `json:"filterMask,omitempty"`
	Filter *ListMatricesRequestFilter `json:"filter,omitempty"`
	PageSize *string `json:"pageSize,omitempty"`
	PageToken *string `json:"pageToken,omitempty"`
}

// NewMatrixListMatricesRequest instantiates a new MatrixListMatricesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMatrixListMatricesRequest() *MatrixListMatricesRequest {
	this := MatrixListMatricesRequest{}
	return &this
}

// NewMatrixListMatricesRequestWithDefaults instantiates a new MatrixListMatricesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMatrixListMatricesRequestWithDefaults() *MatrixListMatricesRequest {
	this := MatrixListMatricesRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *MatrixListMatricesRequest) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatrixListMatricesRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
    return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *MatrixListMatricesRequest) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *MatrixListMatricesRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetConfiguratorId returns the ConfiguratorId field value if set, zero value otherwise.
func (o *MatrixListMatricesRequest) GetConfiguratorId() string {
	if o == nil || isNil(o.ConfiguratorId) {
		var ret string
		return ret
	}
	return *o.ConfiguratorId
}

// GetConfiguratorIdOk returns a tuple with the ConfiguratorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatrixListMatricesRequest) GetConfiguratorIdOk() (*string, bool) {
	if o == nil || isNil(o.ConfiguratorId) {
    return nil, false
	}
	return o.ConfiguratorId, true
}

// HasConfiguratorId returns a boolean if a field has been set.
func (o *MatrixListMatricesRequest) HasConfiguratorId() bool {
	if o != nil && !isNil(o.ConfiguratorId) {
		return true
	}

	return false
}

// SetConfiguratorId gets a reference to the given string and assigns it to the ConfiguratorId field.
func (o *MatrixListMatricesRequest) SetConfiguratorId(v string) {
	o.ConfiguratorId = &v
}

// GetFilterMask returns the FilterMask field value if set, zero value otherwise.
func (o *MatrixListMatricesRequest) GetFilterMask() []string {
	if o == nil || isNil(o.FilterMask) {
		var ret []string
		return ret
	}
	return o.FilterMask
}

// GetFilterMaskOk returns a tuple with the FilterMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatrixListMatricesRequest) GetFilterMaskOk() ([]string, bool) {
	if o == nil || isNil(o.FilterMask) {
    return nil, false
	}
	return o.FilterMask, true
}

// HasFilterMask returns a boolean if a field has been set.
func (o *MatrixListMatricesRequest) HasFilterMask() bool {
	if o != nil && !isNil(o.FilterMask) {
		return true
	}

	return false
}

// SetFilterMask gets a reference to the given []string and assigns it to the FilterMask field.
func (o *MatrixListMatricesRequest) SetFilterMask(v []string) {
	o.FilterMask = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *MatrixListMatricesRequest) GetFilter() ListMatricesRequestFilter {
	if o == nil || isNil(o.Filter) {
		var ret ListMatricesRequestFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatrixListMatricesRequest) GetFilterOk() (*ListMatricesRequestFilter, bool) {
	if o == nil || isNil(o.Filter) {
    return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *MatrixListMatricesRequest) HasFilter() bool {
	if o != nil && !isNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given ListMatricesRequestFilter and assigns it to the Filter field.
func (o *MatrixListMatricesRequest) SetFilter(v ListMatricesRequestFilter) {
	o.Filter = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *MatrixListMatricesRequest) GetPageSize() string {
	if o == nil || isNil(o.PageSize) {
		var ret string
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatrixListMatricesRequest) GetPageSizeOk() (*string, bool) {
	if o == nil || isNil(o.PageSize) {
    return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *MatrixListMatricesRequest) HasPageSize() bool {
	if o != nil && !isNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given string and assigns it to the PageSize field.
func (o *MatrixListMatricesRequest) SetPageSize(v string) {
	o.PageSize = &v
}

// GetPageToken returns the PageToken field value if set, zero value otherwise.
func (o *MatrixListMatricesRequest) GetPageToken() string {
	if o == nil || isNil(o.PageToken) {
		var ret string
		return ret
	}
	return *o.PageToken
}

// GetPageTokenOk returns a tuple with the PageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatrixListMatricesRequest) GetPageTokenOk() (*string, bool) {
	if o == nil || isNil(o.PageToken) {
    return nil, false
	}
	return o.PageToken, true
}

// HasPageToken returns a boolean if a field has been set.
func (o *MatrixListMatricesRequest) HasPageToken() bool {
	if o != nil && !isNil(o.PageToken) {
		return true
	}

	return false
}

// SetPageToken gets a reference to the given string and assigns it to the PageToken field.
func (o *MatrixListMatricesRequest) SetPageToken(v string) {
	o.PageToken = &v
}

func (o MatrixListMatricesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !isNil(o.ConfiguratorId) {
		toSerialize["configuratorId"] = o.ConfiguratorId
	}
	if !isNil(o.FilterMask) {
		toSerialize["filterMask"] = o.FilterMask
	}
	if !isNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}
	if !isNil(o.PageSize) {
		toSerialize["pageSize"] = o.PageSize
	}
	if !isNil(o.PageToken) {
		toSerialize["pageToken"] = o.PageToken
	}
	return json.Marshal(toSerialize)
}

type NullableMatrixListMatricesRequest struct {
	value *MatrixListMatricesRequest
	isSet bool
}

func (v NullableMatrixListMatricesRequest) Get() *MatrixListMatricesRequest {
	return v.value
}

func (v *NullableMatrixListMatricesRequest) Set(val *MatrixListMatricesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMatrixListMatricesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMatrixListMatricesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMatrixListMatricesRequest(val *MatrixListMatricesRequest) *NullableMatrixListMatricesRequest {
	return &NullableMatrixListMatricesRequest{value: val, isSet: true}
}

func (v NullableMatrixListMatricesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMatrixListMatricesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

