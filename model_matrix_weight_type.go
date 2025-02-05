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

// checks if the MatrixWeightType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MatrixWeightType{}

// MatrixWeightType struct for MatrixWeightType
type MatrixWeightType struct {
	WeightUnit           *ProductconfiguratorWeightUnit `json:"weightUnit,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MatrixWeightType MatrixWeightType

// NewMatrixWeightType instantiates a new MatrixWeightType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMatrixWeightType() *MatrixWeightType {
	this := MatrixWeightType{}
	var weightUnit ProductconfiguratorWeightUnit = PRODUCTCONFIGURATORWEIGHTUNIT_UNKNOWN
	this.WeightUnit = &weightUnit
	return &this
}

// NewMatrixWeightTypeWithDefaults instantiates a new MatrixWeightType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMatrixWeightTypeWithDefaults() *MatrixWeightType {
	this := MatrixWeightType{}
	var weightUnit ProductconfiguratorWeightUnit = PRODUCTCONFIGURATORWEIGHTUNIT_UNKNOWN
	this.WeightUnit = &weightUnit
	return &this
}

// GetWeightUnit returns the WeightUnit field value if set, zero value otherwise.
func (o *MatrixWeightType) GetWeightUnit() ProductconfiguratorWeightUnit {
	if o == nil || IsNil(o.WeightUnit) {
		var ret ProductconfiguratorWeightUnit
		return ret
	}
	return *o.WeightUnit
}

// GetWeightUnitOk returns a tuple with the WeightUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatrixWeightType) GetWeightUnitOk() (*ProductconfiguratorWeightUnit, bool) {
	if o == nil || IsNil(o.WeightUnit) {
		return nil, false
	}
	return o.WeightUnit, true
}

// HasWeightUnit returns a boolean if a field has been set.
func (o *MatrixWeightType) IsSetWeightUnit() bool {
	if o != nil && !IsNil(o.WeightUnit) {
		return true
	}

	return false
}

// SetWeightUnit gets a reference to the given ProductconfiguratorWeightUnit and assigns it to the WeightUnit field.
func (o *MatrixWeightType) SetWeightUnit(v ProductconfiguratorWeightUnit) {
	o.WeightUnit = &v
}

func (o MatrixWeightType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MatrixWeightType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.WeightUnit) {
		toSerialize["weightUnit"] = o.WeightUnit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MatrixWeightType) UnmarshalJSON(data []byte) (err error) {
	varMatrixWeightType := _MatrixWeightType{}

	err = json.Unmarshal(data, &varMatrixWeightType)

	if err != nil {
		return err
	}

	*o = MatrixWeightType(varMatrixWeightType)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "weightUnit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *MatrixWeightType) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *MatrixWeightType) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableMatrixWeightType struct {
	value *MatrixWeightType
	isSet bool
}

func (v NullableMatrixWeightType) Get() *MatrixWeightType {
	return v.value
}

func (v *NullableMatrixWeightType) Set(val *MatrixWeightType) {
	v.value = val
	v.isSet = true
}

func (v NullableMatrixWeightType) IsSet() bool {
	return v.isSet
}

func (v *NullableMatrixWeightType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMatrixWeightType(val *MatrixWeightType) *NullableMatrixWeightType {
	return &NullableMatrixWeightType{value: val, isSet: true}
}

func (v NullableMatrixWeightType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMatrixWeightType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
