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

// checks if the ProductconfiguratoroptionBulkCreateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductconfiguratoroptionBulkCreateResponse{}

// ProductconfiguratoroptionBulkCreateResponse struct for ProductconfiguratoroptionBulkCreateResponse
type ProductconfiguratoroptionBulkCreateResponse struct {
	Options              []ProductconfiguratoroptionEntity `json:"options,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductconfiguratoroptionBulkCreateResponse ProductconfiguratoroptionBulkCreateResponse

// NewProductconfiguratoroptionBulkCreateResponse instantiates a new ProductconfiguratoroptionBulkCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductconfiguratoroptionBulkCreateResponse() *ProductconfiguratoroptionBulkCreateResponse {
	this := ProductconfiguratoroptionBulkCreateResponse{}
	return &this
}

// NewProductconfiguratoroptionBulkCreateResponseWithDefaults instantiates a new ProductconfiguratoroptionBulkCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductconfiguratoroptionBulkCreateResponseWithDefaults() *ProductconfiguratoroptionBulkCreateResponse {
	this := ProductconfiguratoroptionBulkCreateResponse{}
	return &this
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *ProductconfiguratoroptionBulkCreateResponse) GetOptions() []ProductconfiguratoroptionEntity {
	if o == nil || IsNil(o.Options) {
		var ret []ProductconfiguratoroptionEntity
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratoroptionBulkCreateResponse) GetOptionsOk() ([]ProductconfiguratoroptionEntity, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ProductconfiguratoroptionBulkCreateResponse) IsSetOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []ProductconfiguratoroptionEntity and assigns it to the Options field.
func (o *ProductconfiguratoroptionBulkCreateResponse) SetOptions(v []ProductconfiguratoroptionEntity) {
	o.Options = v
}

func (o ProductconfiguratoroptionBulkCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductconfiguratoroptionBulkCreateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductconfiguratoroptionBulkCreateResponse) UnmarshalJSON(data []byte) (err error) {
	varProductconfiguratoroptionBulkCreateResponse := _ProductconfiguratoroptionBulkCreateResponse{}

	err = json.Unmarshal(data, &varProductconfiguratoroptionBulkCreateResponse)

	if err != nil {
		return err
	}

	*o = ProductconfiguratoroptionBulkCreateResponse(varProductconfiguratoroptionBulkCreateResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "options")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ProductconfiguratoroptionBulkCreateResponse) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *ProductconfiguratoroptionBulkCreateResponse) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableProductconfiguratoroptionBulkCreateResponse struct {
	value *ProductconfiguratoroptionBulkCreateResponse
	isSet bool
}

func (v NullableProductconfiguratoroptionBulkCreateResponse) Get() *ProductconfiguratoroptionBulkCreateResponse {
	return v.value
}

func (v *NullableProductconfiguratoroptionBulkCreateResponse) Set(val *ProductconfiguratoroptionBulkCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProductconfiguratoroptionBulkCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProductconfiguratoroptionBulkCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductconfiguratoroptionBulkCreateResponse(val *ProductconfiguratoroptionBulkCreateResponse) *NullableProductconfiguratoroptionBulkCreateResponse {
	return &NullableProductconfiguratoroptionBulkCreateResponse{value: val, isSet: true}
}

func (v NullableProductconfiguratoroptionBulkCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductconfiguratoroptionBulkCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
