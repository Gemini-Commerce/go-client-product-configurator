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

// checks if the ProductconfiguratorconfiguratorUpdatePayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductconfiguratorconfiguratorUpdatePayload{}

// ProductconfiguratorconfiguratorUpdatePayload struct for ProductconfiguratorconfiguratorUpdatePayload
type ProductconfiguratorconfiguratorUpdatePayload struct {
	Label                *string                                `json:"label,omitempty"`
	Status               *ProductconfiguratorconfiguratorStatus `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductconfiguratorconfiguratorUpdatePayload ProductconfiguratorconfiguratorUpdatePayload

// NewProductconfiguratorconfiguratorUpdatePayload instantiates a new ProductconfiguratorconfiguratorUpdatePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductconfiguratorconfiguratorUpdatePayload() *ProductconfiguratorconfiguratorUpdatePayload {
	this := ProductconfiguratorconfiguratorUpdatePayload{}
	var status ProductconfiguratorconfiguratorStatus = PRODUCTCONFIGURATORCONFIGURATORSTATUS_UNKNOWN
	this.Status = &status
	return &this
}

// NewProductconfiguratorconfiguratorUpdatePayloadWithDefaults instantiates a new ProductconfiguratorconfiguratorUpdatePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductconfiguratorconfiguratorUpdatePayloadWithDefaults() *ProductconfiguratorconfiguratorUpdatePayload {
	this := ProductconfiguratorconfiguratorUpdatePayload{}
	var status ProductconfiguratorconfiguratorStatus = PRODUCTCONFIGURATORCONFIGURATORSTATUS_UNKNOWN
	this.Status = &status
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ProductconfiguratorconfiguratorUpdatePayload) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorconfiguratorUpdatePayload) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ProductconfiguratorconfiguratorUpdatePayload) IsSetLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ProductconfiguratorconfiguratorUpdatePayload) SetLabel(v string) {
	o.Label = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ProductconfiguratorconfiguratorUpdatePayload) GetStatus() ProductconfiguratorconfiguratorStatus {
	if o == nil || IsNil(o.Status) {
		var ret ProductconfiguratorconfiguratorStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorconfiguratorUpdatePayload) GetStatusOk() (*ProductconfiguratorconfiguratorStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ProductconfiguratorconfiguratorUpdatePayload) IsSetStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ProductconfiguratorconfiguratorStatus and assigns it to the Status field.
func (o *ProductconfiguratorconfiguratorUpdatePayload) SetStatus(v ProductconfiguratorconfiguratorStatus) {
	o.Status = &v
}

func (o ProductconfiguratorconfiguratorUpdatePayload) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductconfiguratorconfiguratorUpdatePayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductconfiguratorconfiguratorUpdatePayload) UnmarshalJSON(data []byte) (err error) {
	varProductconfiguratorconfiguratorUpdatePayload := _ProductconfiguratorconfiguratorUpdatePayload{}

	err = json.Unmarshal(data, &varProductconfiguratorconfiguratorUpdatePayload)

	if err != nil {
		return err
	}

	*o = ProductconfiguratorconfiguratorUpdatePayload(varProductconfiguratorconfiguratorUpdatePayload)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "label")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ProductconfiguratorconfiguratorUpdatePayload) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *ProductconfiguratorconfiguratorUpdatePayload) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableProductconfiguratorconfiguratorUpdatePayload struct {
	value *ProductconfiguratorconfiguratorUpdatePayload
	isSet bool
}

func (v NullableProductconfiguratorconfiguratorUpdatePayload) Get() *ProductconfiguratorconfiguratorUpdatePayload {
	return v.value
}

func (v *NullableProductconfiguratorconfiguratorUpdatePayload) Set(val *ProductconfiguratorconfiguratorUpdatePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableProductconfiguratorconfiguratorUpdatePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableProductconfiguratorconfiguratorUpdatePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductconfiguratorconfiguratorUpdatePayload(val *ProductconfiguratorconfiguratorUpdatePayload) *NullableProductconfiguratorconfiguratorUpdatePayload {
	return &NullableProductconfiguratorconfiguratorUpdatePayload{value: val, isSet: true}
}

func (v NullableProductconfiguratorconfiguratorUpdatePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductconfiguratorconfiguratorUpdatePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
