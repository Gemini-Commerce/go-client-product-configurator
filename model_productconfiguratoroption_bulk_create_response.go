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

// ProductconfiguratoroptionBulkCreateResponse struct for ProductconfiguratoroptionBulkCreateResponse
type ProductconfiguratoroptionBulkCreateResponse struct {
	Options []ProductconfiguratoroptionEntity `json:"options,omitempty"`
}

// NewProductconfiguratoroptionBulkCreateResponse instantiates a new ProductconfiguratoroptionBulkCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductconfiguratoroptionBulkCreateResponse() *ProductconfiguratoroptionBulkCreateResponse {
	this := ProductconfiguratoroptionBulkCreateResponse{}
	return &this
}

// NewProductconfiguratoroptionBulkCreateResponseWithDefaults instantiates a new ProductconfiguratoroptionBulkCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductconfiguratoroptionBulkCreateResponseWithDefaults() *ProductconfiguratoroptionBulkCreateResponse {
	this := ProductconfiguratoroptionBulkCreateResponse{}
	return &this
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *ProductconfiguratoroptionBulkCreateResponse) GetOptions() []ProductconfiguratoroptionEntity {
	if o == nil || isNil(o.Options) {
		var ret []ProductconfiguratoroptionEntity
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratoroptionBulkCreateResponse) GetOptionsOk() ([]ProductconfiguratoroptionEntity, bool) {
	if o == nil || isNil(o.Options) {
    return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ProductconfiguratoroptionBulkCreateResponse) HasOptions() bool {
	if o != nil && !isNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []ProductconfiguratoroptionEntity and assigns it to the Options field.
func (o *ProductconfiguratoroptionBulkCreateResponse) SetOptions(v []ProductconfiguratoroptionEntity) {
	o.Options = v
}

func (o ProductconfiguratoroptionBulkCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableProductconfiguratoroptionBulkCreateResponse struct {
	value *ProductconfiguratoroptionBulkCreateResponse
	isSet bool
}

func (v NullableProductconfiguratoroptionBulkCreateResponse) Get() *ProductconfiguratoroptionBulkCreateResponse {
	return v.value
}

func (v *NullableProductconfiguratoroptionBulkCreateResponse) Set(val *ProductconfiguratoroptionBulkCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProductconfiguratoroptionBulkCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProductconfiguratoroptionBulkCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductconfiguratoroptionBulkCreateResponse(val *ProductconfiguratoroptionBulkCreateResponse) *NullableProductconfiguratoroptionBulkCreateResponse {
	return &NullableProductconfiguratoroptionBulkCreateResponse{value: val, isSet: true}
}

func (v NullableProductconfiguratoroptionBulkCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductconfiguratoroptionBulkCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


