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

// ProductconfiguratorpropertyBulkCreateResponse struct for ProductconfiguratorpropertyBulkCreateResponse
type ProductconfiguratorpropertyBulkCreateResponse struct {
	Properties []ProductconfiguratorpropertyEntity `json:"properties,omitempty"`
}

// NewProductconfiguratorpropertyBulkCreateResponse instantiates a new ProductconfiguratorpropertyBulkCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductconfiguratorpropertyBulkCreateResponse() *ProductconfiguratorpropertyBulkCreateResponse {
	this := ProductconfiguratorpropertyBulkCreateResponse{}
	return &this
}

// NewProductconfiguratorpropertyBulkCreateResponseWithDefaults instantiates a new ProductconfiguratorpropertyBulkCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductconfiguratorpropertyBulkCreateResponseWithDefaults() *ProductconfiguratorpropertyBulkCreateResponse {
	this := ProductconfiguratorpropertyBulkCreateResponse{}
	return &this
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *ProductconfiguratorpropertyBulkCreateResponse) GetProperties() []ProductconfiguratorpropertyEntity {
	if o == nil || isNil(o.Properties) {
		var ret []ProductconfiguratorpropertyEntity
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorpropertyBulkCreateResponse) GetPropertiesOk() ([]ProductconfiguratorpropertyEntity, bool) {
	if o == nil || isNil(o.Properties) {
    return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *ProductconfiguratorpropertyBulkCreateResponse) HasProperties() bool {
	if o != nil && !isNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []ProductconfiguratorpropertyEntity and assigns it to the Properties field.
func (o *ProductconfiguratorpropertyBulkCreateResponse) SetProperties(v []ProductconfiguratorpropertyEntity) {
	o.Properties = v
}

func (o ProductconfiguratorpropertyBulkCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	return json.Marshal(toSerialize)
}

type NullableProductconfiguratorpropertyBulkCreateResponse struct {
	value *ProductconfiguratorpropertyBulkCreateResponse
	isSet bool
}

func (v NullableProductconfiguratorpropertyBulkCreateResponse) Get() *ProductconfiguratorpropertyBulkCreateResponse {
	return v.value
}

func (v *NullableProductconfiguratorpropertyBulkCreateResponse) Set(val *ProductconfiguratorpropertyBulkCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProductconfiguratorpropertyBulkCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProductconfiguratorpropertyBulkCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductconfiguratorpropertyBulkCreateResponse(val *ProductconfiguratorpropertyBulkCreateResponse) *NullableProductconfiguratorpropertyBulkCreateResponse {
	return &NullableProductconfiguratorpropertyBulkCreateResponse{value: val, isSet: true}
}

func (v NullableProductconfiguratorpropertyBulkCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductconfiguratorpropertyBulkCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


