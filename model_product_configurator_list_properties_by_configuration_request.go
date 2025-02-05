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

// checks if the ProductConfiguratorListPropertiesByConfigurationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductConfiguratorListPropertiesByConfigurationRequest{}

// ProductConfiguratorListPropertiesByConfigurationRequest struct for ProductConfiguratorListPropertiesByConfigurationRequest
type ProductConfiguratorListPropertiesByConfigurationRequest struct {
	Selections           []PropertyListPropertiesByConfigurationRequestSelection `json:"selections,omitempty"`
	PropertyType         *ProductconfiguratorPropertyType                        `json:"propertyType,omitempty"`
	PageToken            *string                                                 `json:"pageToken,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductConfiguratorListPropertiesByConfigurationRequest ProductConfiguratorListPropertiesByConfigurationRequest

// NewProductConfiguratorListPropertiesByConfigurationRequest instantiates a new ProductConfiguratorListPropertiesByConfigurationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductConfiguratorListPropertiesByConfigurationRequest() *ProductConfiguratorListPropertiesByConfigurationRequest {
	this := ProductConfiguratorListPropertiesByConfigurationRequest{}
	var propertyType ProductconfiguratorPropertyType = PRODUCTCONFIGURATORPROPERTYTYPE_UNKNOWN
	this.PropertyType = &propertyType
	return &this
}

// NewProductConfiguratorListPropertiesByConfigurationRequestWithDefaults instantiates a new ProductConfiguratorListPropertiesByConfigurationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductConfiguratorListPropertiesByConfigurationRequestWithDefaults() *ProductConfiguratorListPropertiesByConfigurationRequest {
	this := ProductConfiguratorListPropertiesByConfigurationRequest{}
	var propertyType ProductconfiguratorPropertyType = PRODUCTCONFIGURATORPROPERTYTYPE_UNKNOWN
	this.PropertyType = &propertyType
	return &this
}

// GetSelections returns the Selections field value if set, zero value otherwise.
func (o *ProductConfiguratorListPropertiesByConfigurationRequest) GetSelections() []PropertyListPropertiesByConfigurationRequestSelection {
	if o == nil || IsNil(o.Selections) {
		var ret []PropertyListPropertiesByConfigurationRequestSelection
		return ret
	}
	return o.Selections
}

// GetSelectionsOk returns a tuple with the Selections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductConfiguratorListPropertiesByConfigurationRequest) GetSelectionsOk() ([]PropertyListPropertiesByConfigurationRequestSelection, bool) {
	if o == nil || IsNil(o.Selections) {
		return nil, false
	}
	return o.Selections, true
}

// HasSelections returns a boolean if a field has been set.
func (o *ProductConfiguratorListPropertiesByConfigurationRequest) IsSetSelections() bool {
	if o != nil && !IsNil(o.Selections) {
		return true
	}

	return false
}

// SetSelections gets a reference to the given []PropertyListPropertiesByConfigurationRequestSelection and assigns it to the Selections field.
func (o *ProductConfiguratorListPropertiesByConfigurationRequest) SetSelections(v []PropertyListPropertiesByConfigurationRequestSelection) {
	o.Selections = v
}

// GetPropertyType returns the PropertyType field value if set, zero value otherwise.
func (o *ProductConfiguratorListPropertiesByConfigurationRequest) GetPropertyType() ProductconfiguratorPropertyType {
	if o == nil || IsNil(o.PropertyType) {
		var ret ProductconfiguratorPropertyType
		return ret
	}
	return *o.PropertyType
}

// GetPropertyTypeOk returns a tuple with the PropertyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductConfiguratorListPropertiesByConfigurationRequest) GetPropertyTypeOk() (*ProductconfiguratorPropertyType, bool) {
	if o == nil || IsNil(o.PropertyType) {
		return nil, false
	}
	return o.PropertyType, true
}

// HasPropertyType returns a boolean if a field has been set.
func (o *ProductConfiguratorListPropertiesByConfigurationRequest) IsSetPropertyType() bool {
	if o != nil && !IsNil(o.PropertyType) {
		return true
	}

	return false
}

// SetPropertyType gets a reference to the given ProductconfiguratorPropertyType and assigns it to the PropertyType field.
func (o *ProductConfiguratorListPropertiesByConfigurationRequest) SetPropertyType(v ProductconfiguratorPropertyType) {
	o.PropertyType = &v
}

// GetPageToken returns the PageToken field value if set, zero value otherwise.
func (o *ProductConfiguratorListPropertiesByConfigurationRequest) GetPageToken() string {
	if o == nil || IsNil(o.PageToken) {
		var ret string
		return ret
	}
	return *o.PageToken
}

// GetPageTokenOk returns a tuple with the PageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductConfiguratorListPropertiesByConfigurationRequest) GetPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.PageToken) {
		return nil, false
	}
	return o.PageToken, true
}

// HasPageToken returns a boolean if a field has been set.
func (o *ProductConfiguratorListPropertiesByConfigurationRequest) IsSetPageToken() bool {
	if o != nil && !IsNil(o.PageToken) {
		return true
	}

	return false
}

// SetPageToken gets a reference to the given string and assigns it to the PageToken field.
func (o *ProductConfiguratorListPropertiesByConfigurationRequest) SetPageToken(v string) {
	o.PageToken = &v
}

func (o ProductConfiguratorListPropertiesByConfigurationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductConfiguratorListPropertiesByConfigurationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Selections) {
		toSerialize["selections"] = o.Selections
	}
	if !IsNil(o.PropertyType) {
		toSerialize["propertyType"] = o.PropertyType
	}
	if !IsNil(o.PageToken) {
		toSerialize["pageToken"] = o.PageToken
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductConfiguratorListPropertiesByConfigurationRequest) UnmarshalJSON(data []byte) (err error) {
	varProductConfiguratorListPropertiesByConfigurationRequest := _ProductConfiguratorListPropertiesByConfigurationRequest{}

	err = json.Unmarshal(data, &varProductConfiguratorListPropertiesByConfigurationRequest)

	if err != nil {
		return err
	}

	*o = ProductConfiguratorListPropertiesByConfigurationRequest(varProductConfiguratorListPropertiesByConfigurationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "selections")
		delete(additionalProperties, "propertyType")
		delete(additionalProperties, "pageToken")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ProductConfiguratorListPropertiesByConfigurationRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *ProductConfiguratorListPropertiesByConfigurationRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableProductConfiguratorListPropertiesByConfigurationRequest struct {
	value *ProductConfiguratorListPropertiesByConfigurationRequest
	isSet bool
}

func (v NullableProductConfiguratorListPropertiesByConfigurationRequest) Get() *ProductConfiguratorListPropertiesByConfigurationRequest {
	return v.value
}

func (v *NullableProductConfiguratorListPropertiesByConfigurationRequest) Set(val *ProductConfiguratorListPropertiesByConfigurationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductConfiguratorListPropertiesByConfigurationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductConfiguratorListPropertiesByConfigurationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductConfiguratorListPropertiesByConfigurationRequest(val *ProductConfiguratorListPropertiesByConfigurationRequest) *NullableProductConfiguratorListPropertiesByConfigurationRequest {
	return &NullableProductConfiguratorListPropertiesByConfigurationRequest{value: val, isSet: true}
}

func (v NullableProductConfiguratorListPropertiesByConfigurationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductConfiguratorListPropertiesByConfigurationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
