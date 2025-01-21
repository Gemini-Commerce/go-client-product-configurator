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

// checks if the ProductConfiguratorBulkDeleteOptionsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductConfiguratorBulkDeleteOptionsRequest{}

// ProductConfiguratorBulkDeleteOptionsRequest struct for ProductConfiguratorBulkDeleteOptionsRequest
type ProductConfiguratorBulkDeleteOptionsRequest struct {
	OptionIds            []string `json:"optionIds,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductConfiguratorBulkDeleteOptionsRequest ProductConfiguratorBulkDeleteOptionsRequest

// NewProductConfiguratorBulkDeleteOptionsRequest instantiates a new ProductConfiguratorBulkDeleteOptionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductConfiguratorBulkDeleteOptionsRequest() *ProductConfiguratorBulkDeleteOptionsRequest {
	this := ProductConfiguratorBulkDeleteOptionsRequest{}
	return &this
}

// NewProductConfiguratorBulkDeleteOptionsRequestWithDefaults instantiates a new ProductConfiguratorBulkDeleteOptionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductConfiguratorBulkDeleteOptionsRequestWithDefaults() *ProductConfiguratorBulkDeleteOptionsRequest {
	this := ProductConfiguratorBulkDeleteOptionsRequest{}
	return &this
}

// GetOptionIds returns the OptionIds field value if set, zero value otherwise.
func (o *ProductConfiguratorBulkDeleteOptionsRequest) GetOptionIds() []string {
	if o == nil || IsNil(o.OptionIds) {
		var ret []string
		return ret
	}
	return o.OptionIds
}

// GetOptionIdsOk returns a tuple with the OptionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductConfiguratorBulkDeleteOptionsRequest) GetOptionIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.OptionIds) {
		return nil, false
	}
	return o.OptionIds, true
}

// HasOptionIds returns a boolean if a field has been set.
func (o *ProductConfiguratorBulkDeleteOptionsRequest) HasOptionIds() bool {
	if o != nil && !IsNil(o.OptionIds) {
		return true
	}

	return false
}

// SetOptionIds gets a reference to the given []string and assigns it to the OptionIds field.
func (o *ProductConfiguratorBulkDeleteOptionsRequest) SetOptionIds(v []string) {
	o.OptionIds = v
}

func (o ProductConfiguratorBulkDeleteOptionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductConfiguratorBulkDeleteOptionsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OptionIds) {
		toSerialize["optionIds"] = o.OptionIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductConfiguratorBulkDeleteOptionsRequest) UnmarshalJSON(data []byte) (err error) {
	varProductConfiguratorBulkDeleteOptionsRequest := _ProductConfiguratorBulkDeleteOptionsRequest{}

	err = json.Unmarshal(data, &varProductConfiguratorBulkDeleteOptionsRequest)

	if err != nil {
		return err
	}

	*o = ProductConfiguratorBulkDeleteOptionsRequest(varProductConfiguratorBulkDeleteOptionsRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "optionIds")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ProductConfiguratorBulkDeleteOptionsRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *ProductConfiguratorBulkDeleteOptionsRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableProductConfiguratorBulkDeleteOptionsRequest struct {
	value *ProductConfiguratorBulkDeleteOptionsRequest
	isSet bool
}

func (v NullableProductConfiguratorBulkDeleteOptionsRequest) Get() *ProductConfiguratorBulkDeleteOptionsRequest {
	return v.value
}

func (v *NullableProductConfiguratorBulkDeleteOptionsRequest) Set(val *ProductConfiguratorBulkDeleteOptionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductConfiguratorBulkDeleteOptionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductConfiguratorBulkDeleteOptionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductConfiguratorBulkDeleteOptionsRequest(val *ProductConfiguratorBulkDeleteOptionsRequest) *NullableProductConfiguratorBulkDeleteOptionsRequest {
	return &NullableProductConfiguratorBulkDeleteOptionsRequest{value: val, isSet: true}
}

func (v NullableProductConfiguratorBulkDeleteOptionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductConfiguratorBulkDeleteOptionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
