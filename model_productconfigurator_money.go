/*
Product Configurator Service

 The Product Configurator Service is a versatile platform designed to manage dynamic product configurations.  It enables the creation, updating, and management of product configurations through steps, options, and dependencies,  ensuring granular control over customization.  ## Core Components 1. **Configurators**    - Create and manage configurators linked to products.    - Support for configurator states (`ACTIVE`, `DRAFT`, etc.) and versioning.  2. **Steps**    - Define logical sequences to guide users through the configuration process.    - Include localized labels, descriptions, and selection rules.  3. **Options**    - Add and manage options available for each step.    - Support for visual content (`Swatch`) and configurable quantities.  4. **Dependencies**    - Create rules between options and steps to control dynamic interactions.    - Manage complex conditions across configurations.  5. **Matrices**    - Use matrices to handle prices, weights, and other properties.    - Support for dynamic customization based on user selections.  6. **Properties**    - Add custom attributes and properties to configurators.  7. **Configuration Management**    - Retrieve available or user-specific configurations.    - Create optimized configurations to enhance the user experience.  ## Key Features - **Security**: Authenticate every request with JWT, ensuring safety and reliability. - **Flexibility**: Bulk operations (create, update, delete) for steps, options, and dependencies. - **Scalability**: Suitable for large volumes of configurations and complex personalization scenarios. - **Backward Compatibility**: Version management to minimize the impact of changes on existing clients.

API version: v1
Contact: info@gemini-commerce.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package productconfigurator

import (
	"encoding/json"
)

// checks if the ProductconfiguratorMoney type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductconfiguratorMoney{}

// ProductconfiguratorMoney Represents an amount of money.
type ProductconfiguratorMoney struct {
	// The whole units of the amount. For example if `currencyCode` is `\"USD\"`, then 1 unit is one US dollar.
	Units *string `json:"units,omitempty"`
	// Number of micro (10^-6) units of the amount. The value must be between -999,999 and +999,999 inclusive. If `units` is positive, `micros` must be positive or zero. If `units` is zero, `micros` can be positive, zero, or negative. If `units` is negative, `micros` must be negative or zero. For example $-1.75 is represented as `units`=-1 and `micros`=-750,000.
	Micros               *int32 `json:"micros,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductconfiguratorMoney ProductconfiguratorMoney

// NewProductconfiguratorMoney instantiates a new ProductconfiguratorMoney object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductconfiguratorMoney() *ProductconfiguratorMoney {
	this := ProductconfiguratorMoney{}
	return &this
}

// NewProductconfiguratorMoneyWithDefaults instantiates a new ProductconfiguratorMoney object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductconfiguratorMoneyWithDefaults() *ProductconfiguratorMoney {
	this := ProductconfiguratorMoney{}
	return &this
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *ProductconfiguratorMoney) GetUnits() string {
	if o == nil || IsNil(o.Units) {
		var ret string
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorMoney) GetUnitsOk() (*string, bool) {
	if o == nil || IsNil(o.Units) {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *ProductconfiguratorMoney) HasUnits() bool {
	if o != nil && !IsNil(o.Units) {
		return true
	}

	return false
}

// SetUnits gets a reference to the given string and assigns it to the Units field.
func (o *ProductconfiguratorMoney) SetUnits(v string) {
	o.Units = &v
}

// GetMicros returns the Micros field value if set, zero value otherwise.
func (o *ProductconfiguratorMoney) GetMicros() int32 {
	if o == nil || IsNil(o.Micros) {
		var ret int32
		return ret
	}
	return *o.Micros
}

// GetMicrosOk returns a tuple with the Micros field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorMoney) GetMicrosOk() (*int32, bool) {
	if o == nil || IsNil(o.Micros) {
		return nil, false
	}
	return o.Micros, true
}

// HasMicros returns a boolean if a field has been set.
func (o *ProductconfiguratorMoney) HasMicros() bool {
	if o != nil && !IsNil(o.Micros) {
		return true
	}

	return false
}

// SetMicros gets a reference to the given int32 and assigns it to the Micros field.
func (o *ProductconfiguratorMoney) SetMicros(v int32) {
	o.Micros = &v
}

func (o ProductconfiguratorMoney) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductconfiguratorMoney) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Units) {
		toSerialize["units"] = o.Units
	}
	if !IsNil(o.Micros) {
		toSerialize["micros"] = o.Micros
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductconfiguratorMoney) UnmarshalJSON(data []byte) (err error) {
	varProductconfiguratorMoney := _ProductconfiguratorMoney{}

	err = json.Unmarshal(data, &varProductconfiguratorMoney)

	if err != nil {
		return err
	}

	*o = ProductconfiguratorMoney(varProductconfiguratorMoney)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "units")
		delete(additionalProperties, "micros")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ProductconfiguratorMoney) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *ProductconfiguratorMoney) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableProductconfiguratorMoney struct {
	value *ProductconfiguratorMoney
	isSet bool
}

func (v NullableProductconfiguratorMoney) Get() *ProductconfiguratorMoney {
	return v.value
}

func (v *NullableProductconfiguratorMoney) Set(val *ProductconfiguratorMoney) {
	v.value = val
	v.isSet = true
}

func (v NullableProductconfiguratorMoney) IsSet() bool {
	return v.isSet
}

func (v *NullableProductconfiguratorMoney) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductconfiguratorMoney(val *ProductconfiguratorMoney) *NullableProductconfiguratorMoney {
	return &NullableProductconfiguratorMoney{value: val, isSet: true}
}

func (v NullableProductconfiguratorMoney) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductconfiguratorMoney) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
