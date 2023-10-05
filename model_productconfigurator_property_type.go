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
	"fmt"
)

// ProductconfiguratorPropertyType the model 'ProductconfiguratorPropertyType'
type ProductconfiguratorPropertyType string

// List of productconfiguratorPropertyType
const (
	PRODUCTCONFIGURATORPROPERTYTYPE_UNKNOWN ProductconfiguratorPropertyType = "PROPERTY_TYPE_UNKNOWN"
	PRODUCTCONFIGURATORPROPERTYTYPE_GENERIC ProductconfiguratorPropertyType = "PROPERTY_TYPE_GENERIC"
	PRODUCTCONFIGURATORPROPERTYTYPE_PRICE ProductconfiguratorPropertyType = "PROPERTY_TYPE_PRICE"
)

// All allowed values of ProductconfiguratorPropertyType enum
var AllowedProductconfiguratorPropertyTypeEnumValues = []ProductconfiguratorPropertyType{
	"PROPERTY_TYPE_UNKNOWN",
	"PROPERTY_TYPE_GENERIC",
	"PROPERTY_TYPE_PRICE",
}

func (v *ProductconfiguratorPropertyType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProductconfiguratorPropertyType(value)
	for _, existing := range AllowedProductconfiguratorPropertyTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProductconfiguratorPropertyType", value)
}

// NewProductconfiguratorPropertyTypeFromValue returns a pointer to a valid ProductconfiguratorPropertyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProductconfiguratorPropertyTypeFromValue(v string) (*ProductconfiguratorPropertyType, error) {
	ev := ProductconfiguratorPropertyType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProductconfiguratorPropertyType: valid values are %v", v, AllowedProductconfiguratorPropertyTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProductconfiguratorPropertyType) IsValid() bool {
	for _, existing := range AllowedProductconfiguratorPropertyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to productconfiguratorPropertyType value
func (v ProductconfiguratorPropertyType) Ptr() *ProductconfiguratorPropertyType {
	return &v
}

type NullableProductconfiguratorPropertyType struct {
	value *ProductconfiguratorPropertyType
	isSet bool
}

func (v NullableProductconfiguratorPropertyType) Get() *ProductconfiguratorPropertyType {
	return v.value
}

func (v *NullableProductconfiguratorPropertyType) Set(val *ProductconfiguratorPropertyType) {
	v.value = val
	v.isSet = true
}

func (v NullableProductconfiguratorPropertyType) IsSet() bool {
	return v.isSet
}

func (v *NullableProductconfiguratorPropertyType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductconfiguratorPropertyType(val *ProductconfiguratorPropertyType) *NullableProductconfiguratorPropertyType {
	return &NullableProductconfiguratorPropertyType{value: val, isSet: true}
}

func (v NullableProductconfiguratorPropertyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductconfiguratorPropertyType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

