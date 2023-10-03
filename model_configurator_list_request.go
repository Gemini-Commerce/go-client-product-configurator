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

// ConfiguratorListRequest struct for ConfiguratorListRequest
type ConfiguratorListRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	ProductId *string `json:"productId,omitempty"`
	PageSize *int64 `json:"pageSize,omitempty"`
	PageToken *string `json:"pageToken,omitempty"`
}

// NewConfiguratorListRequest instantiates a new ConfiguratorListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfiguratorListRequest() *ConfiguratorListRequest {
	this := ConfiguratorListRequest{}
	return &this
}

// NewConfiguratorListRequestWithDefaults instantiates a new ConfiguratorListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfiguratorListRequestWithDefaults() *ConfiguratorListRequest {
	this := ConfiguratorListRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *ConfiguratorListRequest) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfiguratorListRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
    return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *ConfiguratorListRequest) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *ConfiguratorListRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *ConfiguratorListRequest) GetProductId() string {
	if o == nil || isNil(o.ProductId) {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfiguratorListRequest) GetProductIdOk() (*string, bool) {
	if o == nil || isNil(o.ProductId) {
    return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *ConfiguratorListRequest) HasProductId() bool {
	if o != nil && !isNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *ConfiguratorListRequest) SetProductId(v string) {
	o.ProductId = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *ConfiguratorListRequest) GetPageSize() int64 {
	if o == nil || isNil(o.PageSize) {
		var ret int64
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfiguratorListRequest) GetPageSizeOk() (*int64, bool) {
	if o == nil || isNil(o.PageSize) {
    return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *ConfiguratorListRequest) HasPageSize() bool {
	if o != nil && !isNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int64 and assigns it to the PageSize field.
func (o *ConfiguratorListRequest) SetPageSize(v int64) {
	o.PageSize = &v
}

// GetPageToken returns the PageToken field value if set, zero value otherwise.
func (o *ConfiguratorListRequest) GetPageToken() string {
	if o == nil || isNil(o.PageToken) {
		var ret string
		return ret
	}
	return *o.PageToken
}

// GetPageTokenOk returns a tuple with the PageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfiguratorListRequest) GetPageTokenOk() (*string, bool) {
	if o == nil || isNil(o.PageToken) {
    return nil, false
	}
	return o.PageToken, true
}

// HasPageToken returns a boolean if a field has been set.
func (o *ConfiguratorListRequest) HasPageToken() bool {
	if o != nil && !isNil(o.PageToken) {
		return true
	}

	return false
}

// SetPageToken gets a reference to the given string and assigns it to the PageToken field.
func (o *ConfiguratorListRequest) SetPageToken(v string) {
	o.PageToken = &v
}

func (o ConfiguratorListRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !isNil(o.ProductId) {
		toSerialize["productId"] = o.ProductId
	}
	if !isNil(o.PageSize) {
		toSerialize["pageSize"] = o.PageSize
	}
	if !isNil(o.PageToken) {
		toSerialize["pageToken"] = o.PageToken
	}
	return json.Marshal(toSerialize)
}

type NullableConfiguratorListRequest struct {
	value *ConfiguratorListRequest
	isSet bool
}

func (v NullableConfiguratorListRequest) Get() *ConfiguratorListRequest {
	return v.value
}

func (v *NullableConfiguratorListRequest) Set(val *ConfiguratorListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableConfiguratorListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableConfiguratorListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfiguratorListRequest(val *ConfiguratorListRequest) *NullableConfiguratorListRequest {
	return &NullableConfiguratorListRequest{value: val, isSet: true}
}

func (v NullableConfiguratorListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfiguratorListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


