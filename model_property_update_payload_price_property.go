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

// PropertyUpdatePayloadPriceProperty struct for PropertyUpdatePayloadPriceProperty
type PropertyUpdatePayloadPriceProperty struct {
	Price *ProductconfiguratorMoney `json:"price,omitempty"`
	PricelistGrn *string `json:"pricelistGrn,omitempty"`
}

// NewPropertyUpdatePayloadPriceProperty instantiates a new PropertyUpdatePayloadPriceProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPropertyUpdatePayloadPriceProperty() *PropertyUpdatePayloadPriceProperty {
	this := PropertyUpdatePayloadPriceProperty{}
	return &this
}

// NewPropertyUpdatePayloadPricePropertyWithDefaults instantiates a new PropertyUpdatePayloadPriceProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertyUpdatePayloadPricePropertyWithDefaults() *PropertyUpdatePayloadPriceProperty {
	this := PropertyUpdatePayloadPriceProperty{}
	return &this
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *PropertyUpdatePayloadPriceProperty) GetPrice() ProductconfiguratorMoney {
	if o == nil || isNil(o.Price) {
		var ret ProductconfiguratorMoney
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyUpdatePayloadPriceProperty) GetPriceOk() (*ProductconfiguratorMoney, bool) {
	if o == nil || isNil(o.Price) {
    return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *PropertyUpdatePayloadPriceProperty) HasPrice() bool {
	if o != nil && !isNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given ProductconfiguratorMoney and assigns it to the Price field.
func (o *PropertyUpdatePayloadPriceProperty) SetPrice(v ProductconfiguratorMoney) {
	o.Price = &v
}

// GetPricelistGrn returns the PricelistGrn field value if set, zero value otherwise.
func (o *PropertyUpdatePayloadPriceProperty) GetPricelistGrn() string {
	if o == nil || isNil(o.PricelistGrn) {
		var ret string
		return ret
	}
	return *o.PricelistGrn
}

// GetPricelistGrnOk returns a tuple with the PricelistGrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyUpdatePayloadPriceProperty) GetPricelistGrnOk() (*string, bool) {
	if o == nil || isNil(o.PricelistGrn) {
    return nil, false
	}
	return o.PricelistGrn, true
}

// HasPricelistGrn returns a boolean if a field has been set.
func (o *PropertyUpdatePayloadPriceProperty) HasPricelistGrn() bool {
	if o != nil && !isNil(o.PricelistGrn) {
		return true
	}

	return false
}

// SetPricelistGrn gets a reference to the given string and assigns it to the PricelistGrn field.
func (o *PropertyUpdatePayloadPriceProperty) SetPricelistGrn(v string) {
	o.PricelistGrn = &v
}

func (o PropertyUpdatePayloadPriceProperty) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !isNil(o.PricelistGrn) {
		toSerialize["pricelistGrn"] = o.PricelistGrn
	}
	return json.Marshal(toSerialize)
}

type NullablePropertyUpdatePayloadPriceProperty struct {
	value *PropertyUpdatePayloadPriceProperty
	isSet bool
}

func (v NullablePropertyUpdatePayloadPriceProperty) Get() *PropertyUpdatePayloadPriceProperty {
	return v.value
}

func (v *NullablePropertyUpdatePayloadPriceProperty) Set(val *PropertyUpdatePayloadPriceProperty) {
	v.value = val
	v.isSet = true
}

func (v NullablePropertyUpdatePayloadPriceProperty) IsSet() bool {
	return v.isSet
}

func (v *NullablePropertyUpdatePayloadPriceProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropertyUpdatePayloadPriceProperty(val *PropertyUpdatePayloadPriceProperty) *NullablePropertyUpdatePayloadPriceProperty {
	return &NullablePropertyUpdatePayloadPriceProperty{value: val, isSet: true}
}

func (v NullablePropertyUpdatePayloadPriceProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropertyUpdatePayloadPriceProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

