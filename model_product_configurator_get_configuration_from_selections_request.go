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

// checks if the ProductConfiguratorGetConfigurationFromSelectionsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductConfiguratorGetConfigurationFromSelectionsRequest{}

// ProductConfiguratorGetConfigurationFromSelectionsRequest struct for ProductConfiguratorGetConfigurationFromSelectionsRequest
type ProductConfiguratorGetConfigurationFromSelectionsRequest struct {
	Selections           []ProductconfiguratorconfigurationSelection `json:"selections,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductConfiguratorGetConfigurationFromSelectionsRequest ProductConfiguratorGetConfigurationFromSelectionsRequest

// NewProductConfiguratorGetConfigurationFromSelectionsRequest instantiates a new ProductConfiguratorGetConfigurationFromSelectionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductConfiguratorGetConfigurationFromSelectionsRequest() *ProductConfiguratorGetConfigurationFromSelectionsRequest {
	this := ProductConfiguratorGetConfigurationFromSelectionsRequest{}
	return &this
}

// NewProductConfiguratorGetConfigurationFromSelectionsRequestWithDefaults instantiates a new ProductConfiguratorGetConfigurationFromSelectionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductConfiguratorGetConfigurationFromSelectionsRequestWithDefaults() *ProductConfiguratorGetConfigurationFromSelectionsRequest {
	this := ProductConfiguratorGetConfigurationFromSelectionsRequest{}
	return &this
}

// GetSelections returns the Selections field value if set, zero value otherwise.
func (o *ProductConfiguratorGetConfigurationFromSelectionsRequest) GetSelections() []ProductconfiguratorconfigurationSelection {
	if o == nil || IsNil(o.Selections) {
		var ret []ProductconfiguratorconfigurationSelection
		return ret
	}
	return o.Selections
}

// GetSelectionsOk returns a tuple with the Selections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductConfiguratorGetConfigurationFromSelectionsRequest) GetSelectionsOk() ([]ProductconfiguratorconfigurationSelection, bool) {
	if o == nil || IsNil(o.Selections) {
		return nil, false
	}
	return o.Selections, true
}

// HasSelections returns a boolean if a field has been set.
func (o *ProductConfiguratorGetConfigurationFromSelectionsRequest) HasSelections() bool {
	if o != nil && !IsNil(o.Selections) {
		return true
	}

	return false
}

// SetSelections gets a reference to the given []ProductconfiguratorconfigurationSelection and assigns it to the Selections field.
func (o *ProductConfiguratorGetConfigurationFromSelectionsRequest) SetSelections(v []ProductconfiguratorconfigurationSelection) {
	o.Selections = v
}

func (o ProductConfiguratorGetConfigurationFromSelectionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductConfiguratorGetConfigurationFromSelectionsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Selections) {
		toSerialize["selections"] = o.Selections
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductConfiguratorGetConfigurationFromSelectionsRequest) UnmarshalJSON(data []byte) (err error) {
	varProductConfiguratorGetConfigurationFromSelectionsRequest := _ProductConfiguratorGetConfigurationFromSelectionsRequest{}

	err = json.Unmarshal(data, &varProductConfiguratorGetConfigurationFromSelectionsRequest)

	if err != nil {
		return err
	}

	*o = ProductConfiguratorGetConfigurationFromSelectionsRequest(varProductConfiguratorGetConfigurationFromSelectionsRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "selections")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ProductConfiguratorGetConfigurationFromSelectionsRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *ProductConfiguratorGetConfigurationFromSelectionsRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableProductConfiguratorGetConfigurationFromSelectionsRequest struct {
	value *ProductConfiguratorGetConfigurationFromSelectionsRequest
	isSet bool
}

func (v NullableProductConfiguratorGetConfigurationFromSelectionsRequest) Get() *ProductConfiguratorGetConfigurationFromSelectionsRequest {
	return v.value
}

func (v *NullableProductConfiguratorGetConfigurationFromSelectionsRequest) Set(val *ProductConfiguratorGetConfigurationFromSelectionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductConfiguratorGetConfigurationFromSelectionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductConfiguratorGetConfigurationFromSelectionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductConfiguratorGetConfigurationFromSelectionsRequest(val *ProductConfiguratorGetConfigurationFromSelectionsRequest) *NullableProductConfiguratorGetConfigurationFromSelectionsRequest {
	return &NullableProductConfiguratorGetConfigurationFromSelectionsRequest{value: val, isSet: true}
}

func (v NullableProductConfiguratorGetConfigurationFromSelectionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductConfiguratorGetConfigurationFromSelectionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
