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

// checks if the ProductconfiguratorconfigurationSelection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductconfiguratorconfigurationSelection{}

// ProductconfiguratorconfigurationSelection struct for ProductconfiguratorconfigurationSelection
type ProductconfiguratorconfigurationSelection struct {
	StepId               *string  `json:"stepId,omitempty"`
	OptionIds            []string `json:"optionIds,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductconfiguratorconfigurationSelection ProductconfiguratorconfigurationSelection

// NewProductconfiguratorconfigurationSelection instantiates a new ProductconfiguratorconfigurationSelection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductconfiguratorconfigurationSelection() *ProductconfiguratorconfigurationSelection {
	this := ProductconfiguratorconfigurationSelection{}
	return &this
}

// NewProductconfiguratorconfigurationSelectionWithDefaults instantiates a new ProductconfiguratorconfigurationSelection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductconfiguratorconfigurationSelectionWithDefaults() *ProductconfiguratorconfigurationSelection {
	this := ProductconfiguratorconfigurationSelection{}
	return &this
}

// GetStepId returns the StepId field value if set, zero value otherwise.
func (o *ProductconfiguratorconfigurationSelection) GetStepId() string {
	if o == nil || IsNil(o.StepId) {
		var ret string
		return ret
	}
	return *o.StepId
}

// GetStepIdOk returns a tuple with the StepId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorconfigurationSelection) GetStepIdOk() (*string, bool) {
	if o == nil || IsNil(o.StepId) {
		return nil, false
	}
	return o.StepId, true
}

// HasStepId returns a boolean if a field has been set.
func (o *ProductconfiguratorconfigurationSelection) IsSetStepId() bool {
	if o != nil && !IsNil(o.StepId) {
		return true
	}

	return false
}

// SetStepId gets a reference to the given string and assigns it to the StepId field.
func (o *ProductconfiguratorconfigurationSelection) SetStepId(v string) {
	o.StepId = &v
}

// GetOptionIds returns the OptionIds field value if set, zero value otherwise.
func (o *ProductconfiguratorconfigurationSelection) GetOptionIds() []string {
	if o == nil || IsNil(o.OptionIds) {
		var ret []string
		return ret
	}
	return o.OptionIds
}

// GetOptionIdsOk returns a tuple with the OptionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorconfigurationSelection) GetOptionIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.OptionIds) {
		return nil, false
	}
	return o.OptionIds, true
}

// HasOptionIds returns a boolean if a field has been set.
func (o *ProductconfiguratorconfigurationSelection) IsSetOptionIds() bool {
	if o != nil && !IsNil(o.OptionIds) {
		return true
	}

	return false
}

// SetOptionIds gets a reference to the given []string and assigns it to the OptionIds field.
func (o *ProductconfiguratorconfigurationSelection) SetOptionIds(v []string) {
	o.OptionIds = v
}

func (o ProductconfiguratorconfigurationSelection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductconfiguratorconfigurationSelection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StepId) {
		toSerialize["stepId"] = o.StepId
	}
	if !IsNil(o.OptionIds) {
		toSerialize["optionIds"] = o.OptionIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductconfiguratorconfigurationSelection) UnmarshalJSON(data []byte) (err error) {
	varProductconfiguratorconfigurationSelection := _ProductconfiguratorconfigurationSelection{}

	err = json.Unmarshal(data, &varProductconfiguratorconfigurationSelection)

	if err != nil {
		return err
	}

	*o = ProductconfiguratorconfigurationSelection(varProductconfiguratorconfigurationSelection)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "stepId")
		delete(additionalProperties, "optionIds")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ProductconfiguratorconfigurationSelection) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *ProductconfiguratorconfigurationSelection) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableProductconfiguratorconfigurationSelection struct {
	value *ProductconfiguratorconfigurationSelection
	isSet bool
}

func (v NullableProductconfiguratorconfigurationSelection) Get() *ProductconfiguratorconfigurationSelection {
	return v.value
}

func (v *NullableProductconfiguratorconfigurationSelection) Set(val *ProductconfiguratorconfigurationSelection) {
	v.value = val
	v.isSet = true
}

func (v NullableProductconfiguratorconfigurationSelection) IsSet() bool {
	return v.isSet
}

func (v *NullableProductconfiguratorconfigurationSelection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductconfiguratorconfigurationSelection(val *ProductconfiguratorconfigurationSelection) *NullableProductconfiguratorconfigurationSelection {
	return &NullableProductconfiguratorconfigurationSelection{value: val, isSet: true}
}

func (v NullableProductconfiguratorconfigurationSelection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductconfiguratorconfigurationSelection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
