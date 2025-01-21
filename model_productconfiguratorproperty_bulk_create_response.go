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

// checks if the ProductconfiguratorpropertyBulkCreateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductconfiguratorpropertyBulkCreateResponse{}

// ProductconfiguratorpropertyBulkCreateResponse struct for ProductconfiguratorpropertyBulkCreateResponse
type ProductconfiguratorpropertyBulkCreateResponse struct {
	Properties           []ProductconfiguratorpropertyEntity `json:"properties,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductconfiguratorpropertyBulkCreateResponse ProductconfiguratorpropertyBulkCreateResponse

// NewProductconfiguratorpropertyBulkCreateResponse instantiates a new ProductconfiguratorpropertyBulkCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductconfiguratorpropertyBulkCreateResponse() *ProductconfiguratorpropertyBulkCreateResponse {
	this := ProductconfiguratorpropertyBulkCreateResponse{}
	return &this
}

// NewProductconfiguratorpropertyBulkCreateResponseWithDefaults instantiates a new ProductconfiguratorpropertyBulkCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductconfiguratorpropertyBulkCreateResponseWithDefaults() *ProductconfiguratorpropertyBulkCreateResponse {
	this := ProductconfiguratorpropertyBulkCreateResponse{}
	return &this
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *ProductconfiguratorpropertyBulkCreateResponse) GetProperties() []ProductconfiguratorpropertyEntity {
	if o == nil || IsNil(o.Properties) {
		var ret []ProductconfiguratorpropertyEntity
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorpropertyBulkCreateResponse) GetPropertiesOk() ([]ProductconfiguratorpropertyEntity, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *ProductconfiguratorpropertyBulkCreateResponse) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []ProductconfiguratorpropertyEntity and assigns it to the Properties field.
func (o *ProductconfiguratorpropertyBulkCreateResponse) SetProperties(v []ProductconfiguratorpropertyEntity) {
	o.Properties = v
}

func (o ProductconfiguratorpropertyBulkCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductconfiguratorpropertyBulkCreateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductconfiguratorpropertyBulkCreateResponse) UnmarshalJSON(data []byte) (err error) {
	varProductconfiguratorpropertyBulkCreateResponse := _ProductconfiguratorpropertyBulkCreateResponse{}

	err = json.Unmarshal(data, &varProductconfiguratorpropertyBulkCreateResponse)

	if err != nil {
		return err
	}

	*o = ProductconfiguratorpropertyBulkCreateResponse(varProductconfiguratorpropertyBulkCreateResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "properties")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ProductconfiguratorpropertyBulkCreateResponse) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *ProductconfiguratorpropertyBulkCreateResponse) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableProductconfiguratorpropertyBulkCreateResponse struct {
	value *ProductconfiguratorpropertyBulkCreateResponse
	isSet bool
}

func (v NullableProductconfiguratorpropertyBulkCreateResponse) Get() *ProductconfiguratorpropertyBulkCreateResponse {
	return v.value
}

func (v *NullableProductconfiguratorpropertyBulkCreateResponse) Set(val *ProductconfiguratorpropertyBulkCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProductconfiguratorpropertyBulkCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProductconfiguratorpropertyBulkCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductconfiguratorpropertyBulkCreateResponse(val *ProductconfiguratorpropertyBulkCreateResponse) *NullableProductconfiguratorpropertyBulkCreateResponse {
	return &NullableProductconfiguratorpropertyBulkCreateResponse{value: val, isSet: true}
}

func (v NullableProductconfiguratorpropertyBulkCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductconfiguratorpropertyBulkCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
