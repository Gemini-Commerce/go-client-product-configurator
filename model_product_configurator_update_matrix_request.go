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

// checks if the ProductConfiguratorUpdateMatrixRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductConfiguratorUpdateMatrixRequest{}

// ProductConfiguratorUpdateMatrixRequest struct for ProductConfiguratorUpdateMatrixRequest
type ProductConfiguratorUpdateMatrixRequest struct {
	Payload              *ProductconfiguratormatrixUpdatePayload `json:"payload,omitempty"`
	PayloadMask          *string                                 `json:"payloadMask,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductConfiguratorUpdateMatrixRequest ProductConfiguratorUpdateMatrixRequest

// NewProductConfiguratorUpdateMatrixRequest instantiates a new ProductConfiguratorUpdateMatrixRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductConfiguratorUpdateMatrixRequest() *ProductConfiguratorUpdateMatrixRequest {
	this := ProductConfiguratorUpdateMatrixRequest{}
	return &this
}

// NewProductConfiguratorUpdateMatrixRequestWithDefaults instantiates a new ProductConfiguratorUpdateMatrixRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductConfiguratorUpdateMatrixRequestWithDefaults() *ProductConfiguratorUpdateMatrixRequest {
	this := ProductConfiguratorUpdateMatrixRequest{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *ProductConfiguratorUpdateMatrixRequest) GetPayload() ProductconfiguratormatrixUpdatePayload {
	if o == nil || IsNil(o.Payload) {
		var ret ProductconfiguratormatrixUpdatePayload
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductConfiguratorUpdateMatrixRequest) GetPayloadOk() (*ProductconfiguratormatrixUpdatePayload, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *ProductConfiguratorUpdateMatrixRequest) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given ProductconfiguratormatrixUpdatePayload and assigns it to the Payload field.
func (o *ProductConfiguratorUpdateMatrixRequest) SetPayload(v ProductconfiguratormatrixUpdatePayload) {
	o.Payload = &v
}

// GetPayloadMask returns the PayloadMask field value if set, zero value otherwise.
func (o *ProductConfiguratorUpdateMatrixRequest) GetPayloadMask() string {
	if o == nil || IsNil(o.PayloadMask) {
		var ret string
		return ret
	}
	return *o.PayloadMask
}

// GetPayloadMaskOk returns a tuple with the PayloadMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductConfiguratorUpdateMatrixRequest) GetPayloadMaskOk() (*string, bool) {
	if o == nil || IsNil(o.PayloadMask) {
		return nil, false
	}
	return o.PayloadMask, true
}

// HasPayloadMask returns a boolean if a field has been set.
func (o *ProductConfiguratorUpdateMatrixRequest) HasPayloadMask() bool {
	if o != nil && !IsNil(o.PayloadMask) {
		return true
	}

	return false
}

// SetPayloadMask gets a reference to the given string and assigns it to the PayloadMask field.
func (o *ProductConfiguratorUpdateMatrixRequest) SetPayloadMask(v string) {
	o.PayloadMask = &v
}

func (o ProductConfiguratorUpdateMatrixRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductConfiguratorUpdateMatrixRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.PayloadMask) {
		toSerialize["payloadMask"] = o.PayloadMask
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductConfiguratorUpdateMatrixRequest) UnmarshalJSON(data []byte) (err error) {
	varProductConfiguratorUpdateMatrixRequest := _ProductConfiguratorUpdateMatrixRequest{}

	err = json.Unmarshal(data, &varProductConfiguratorUpdateMatrixRequest)

	if err != nil {
		return err
	}

	*o = ProductConfiguratorUpdateMatrixRequest(varProductConfiguratorUpdateMatrixRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "payload")
		delete(additionalProperties, "payloadMask")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ProductConfiguratorUpdateMatrixRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *ProductConfiguratorUpdateMatrixRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableProductConfiguratorUpdateMatrixRequest struct {
	value *ProductConfiguratorUpdateMatrixRequest
	isSet bool
}

func (v NullableProductConfiguratorUpdateMatrixRequest) Get() *ProductConfiguratorUpdateMatrixRequest {
	return v.value
}

func (v *NullableProductConfiguratorUpdateMatrixRequest) Set(val *ProductConfiguratorUpdateMatrixRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductConfiguratorUpdateMatrixRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductConfiguratorUpdateMatrixRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductConfiguratorUpdateMatrixRequest(val *ProductConfiguratorUpdateMatrixRequest) *NullableProductConfiguratorUpdateMatrixRequest {
	return &NullableProductConfiguratorUpdateMatrixRequest{value: val, isSet: true}
}

func (v NullableProductConfiguratorUpdateMatrixRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductConfiguratorUpdateMatrixRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
