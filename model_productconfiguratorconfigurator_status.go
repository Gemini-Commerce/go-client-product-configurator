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
	"fmt"
)

// ProductconfiguratorconfiguratorStatus the model 'ProductconfiguratorconfiguratorStatus'
type ProductconfiguratorconfiguratorStatus string

// List of productconfiguratorconfiguratorStatus
const (
	UNKNOWN ProductconfiguratorconfiguratorStatus = "UNKNOWN"
	ACTIVE ProductconfiguratorconfiguratorStatus = "ACTIVE"
	DRAFT ProductconfiguratorconfiguratorStatus = "DRAFT"
	DISABLED ProductconfiguratorconfiguratorStatus = "DISABLED"
)

// All allowed values of ProductconfiguratorconfiguratorStatus enum
var AllowedProductconfiguratorconfiguratorStatusEnumValues = []ProductconfiguratorconfiguratorStatus{
	"UNKNOWN",
	"ACTIVE",
	"DRAFT",
	"DISABLED",
}

func (v *ProductconfiguratorconfiguratorStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProductconfiguratorconfiguratorStatus(value)
	for _, existing := range AllowedProductconfiguratorconfiguratorStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProductconfiguratorconfiguratorStatus", value)
}

// NewProductconfiguratorconfiguratorStatusFromValue returns a pointer to a valid ProductconfiguratorconfiguratorStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProductconfiguratorconfiguratorStatusFromValue(v string) (*ProductconfiguratorconfiguratorStatus, error) {
	ev := ProductconfiguratorconfiguratorStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProductconfiguratorconfiguratorStatus: valid values are %v", v, AllowedProductconfiguratorconfiguratorStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProductconfiguratorconfiguratorStatus) IsValid() bool {
	for _, existing := range AllowedProductconfiguratorconfiguratorStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to productconfiguratorconfiguratorStatus value
func (v ProductconfiguratorconfiguratorStatus) Ptr() *ProductconfiguratorconfiguratorStatus {
	return &v
}

type NullableProductconfiguratorconfiguratorStatus struct {
	value *ProductconfiguratorconfiguratorStatus
	isSet bool
}

func (v NullableProductconfiguratorconfiguratorStatus) Get() *ProductconfiguratorconfiguratorStatus {
	return v.value
}

func (v *NullableProductconfiguratorconfiguratorStatus) Set(val *ProductconfiguratorconfiguratorStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableProductconfiguratorconfiguratorStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableProductconfiguratorconfiguratorStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductconfiguratorconfiguratorStatus(val *ProductconfiguratorconfiguratorStatus) *NullableProductconfiguratorconfiguratorStatus {
	return &NullableProductconfiguratorconfiguratorStatus{value: val, isSet: true}
}

func (v NullableProductconfiguratorconfiguratorStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductconfiguratorconfiguratorStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

