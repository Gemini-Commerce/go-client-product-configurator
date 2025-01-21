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

// checks if the ProductConfiguratorCopyOptionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductConfiguratorCopyOptionRequest{}

// ProductConfiguratorCopyOptionRequest struct for ProductConfiguratorCopyOptionRequest
type ProductConfiguratorCopyOptionRequest struct {
	TargetStepId         *string `json:"targetStepId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductConfiguratorCopyOptionRequest ProductConfiguratorCopyOptionRequest

// NewProductConfiguratorCopyOptionRequest instantiates a new ProductConfiguratorCopyOptionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductConfiguratorCopyOptionRequest() *ProductConfiguratorCopyOptionRequest {
	this := ProductConfiguratorCopyOptionRequest{}
	return &this
}

// NewProductConfiguratorCopyOptionRequestWithDefaults instantiates a new ProductConfiguratorCopyOptionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductConfiguratorCopyOptionRequestWithDefaults() *ProductConfiguratorCopyOptionRequest {
	this := ProductConfiguratorCopyOptionRequest{}
	return &this
}

// GetTargetStepId returns the TargetStepId field value if set, zero value otherwise.
func (o *ProductConfiguratorCopyOptionRequest) GetTargetStepId() string {
	if o == nil || IsNil(o.TargetStepId) {
		var ret string
		return ret
	}
	return *o.TargetStepId
}

// GetTargetStepIdOk returns a tuple with the TargetStepId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductConfiguratorCopyOptionRequest) GetTargetStepIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetStepId) {
		return nil, false
	}
	return o.TargetStepId, true
}

// HasTargetStepId returns a boolean if a field has been set.
func (o *ProductConfiguratorCopyOptionRequest) HasTargetStepId() bool {
	if o != nil && !IsNil(o.TargetStepId) {
		return true
	}

	return false
}

// SetTargetStepId gets a reference to the given string and assigns it to the TargetStepId field.
func (o *ProductConfiguratorCopyOptionRequest) SetTargetStepId(v string) {
	o.TargetStepId = &v
}

func (o ProductConfiguratorCopyOptionRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductConfiguratorCopyOptionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TargetStepId) {
		toSerialize["targetStepId"] = o.TargetStepId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductConfiguratorCopyOptionRequest) UnmarshalJSON(data []byte) (err error) {
	varProductConfiguratorCopyOptionRequest := _ProductConfiguratorCopyOptionRequest{}

	err = json.Unmarshal(data, &varProductConfiguratorCopyOptionRequest)

	if err != nil {
		return err
	}

	*o = ProductConfiguratorCopyOptionRequest(varProductConfiguratorCopyOptionRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "targetStepId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ProductConfiguratorCopyOptionRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *ProductConfiguratorCopyOptionRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableProductConfiguratorCopyOptionRequest struct {
	value *ProductConfiguratorCopyOptionRequest
	isSet bool
}

func (v NullableProductConfiguratorCopyOptionRequest) Get() *ProductConfiguratorCopyOptionRequest {
	return v.value
}

func (v *NullableProductConfiguratorCopyOptionRequest) Set(val *ProductConfiguratorCopyOptionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductConfiguratorCopyOptionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductConfiguratorCopyOptionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductConfiguratorCopyOptionRequest(val *ProductConfiguratorCopyOptionRequest) *NullableProductConfiguratorCopyOptionRequest {
	return &NullableProductConfiguratorCopyOptionRequest{value: val, isSet: true}
}

func (v NullableProductConfiguratorCopyOptionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductConfiguratorCopyOptionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
