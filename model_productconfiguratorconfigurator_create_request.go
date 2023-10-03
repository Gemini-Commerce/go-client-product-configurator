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

// ProductconfiguratorconfiguratorCreateRequest struct for ProductconfiguratorconfiguratorCreateRequest
type ProductconfiguratorconfiguratorCreateRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	ProductId *string `json:"productId,omitempty"`
	Label *string `json:"label,omitempty"`
}

// NewProductconfiguratorconfiguratorCreateRequest instantiates a new ProductconfiguratorconfiguratorCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductconfiguratorconfiguratorCreateRequest() *ProductconfiguratorconfiguratorCreateRequest {
	this := ProductconfiguratorconfiguratorCreateRequest{}
	return &this
}

// NewProductconfiguratorconfiguratorCreateRequestWithDefaults instantiates a new ProductconfiguratorconfiguratorCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductconfiguratorconfiguratorCreateRequestWithDefaults() *ProductconfiguratorconfiguratorCreateRequest {
	this := ProductconfiguratorconfiguratorCreateRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *ProductconfiguratorconfiguratorCreateRequest) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorconfiguratorCreateRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
    return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *ProductconfiguratorconfiguratorCreateRequest) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *ProductconfiguratorconfiguratorCreateRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *ProductconfiguratorconfiguratorCreateRequest) GetProductId() string {
	if o == nil || isNil(o.ProductId) {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorconfiguratorCreateRequest) GetProductIdOk() (*string, bool) {
	if o == nil || isNil(o.ProductId) {
    return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *ProductconfiguratorconfiguratorCreateRequest) HasProductId() bool {
	if o != nil && !isNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *ProductconfiguratorconfiguratorCreateRequest) SetProductId(v string) {
	o.ProductId = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ProductconfiguratorconfiguratorCreateRequest) GetLabel() string {
	if o == nil || isNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorconfiguratorCreateRequest) GetLabelOk() (*string, bool) {
	if o == nil || isNil(o.Label) {
    return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ProductconfiguratorconfiguratorCreateRequest) HasLabel() bool {
	if o != nil && !isNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ProductconfiguratorconfiguratorCreateRequest) SetLabel(v string) {
	o.Label = &v
}

func (o ProductconfiguratorconfiguratorCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !isNil(o.ProductId) {
		toSerialize["productId"] = o.ProductId
	}
	if !isNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	return json.Marshal(toSerialize)
}

type NullableProductconfiguratorconfiguratorCreateRequest struct {
	value *ProductconfiguratorconfiguratorCreateRequest
	isSet bool
}

func (v NullableProductconfiguratorconfiguratorCreateRequest) Get() *ProductconfiguratorconfiguratorCreateRequest {
	return v.value
}

func (v *NullableProductconfiguratorconfiguratorCreateRequest) Set(val *ProductconfiguratorconfiguratorCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductconfiguratorconfiguratorCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductconfiguratorconfiguratorCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductconfiguratorconfiguratorCreateRequest(val *ProductconfiguratorconfiguratorCreateRequest) *NullableProductconfiguratorconfiguratorCreateRequest {
	return &NullableProductconfiguratorconfiguratorCreateRequest{value: val, isSet: true}
}

func (v NullableProductconfiguratorconfiguratorCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductconfiguratorconfiguratorCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

