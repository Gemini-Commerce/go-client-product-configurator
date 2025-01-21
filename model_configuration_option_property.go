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

// checks if the ConfigurationOptionProperty type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigurationOptionProperty{}

// ConfigurationOptionProperty struct for ConfigurationOptionProperty
type ConfigurationOptionProperty struct {
	PropertyId    *string                          `json:"propertyId,omitempty"`
	PropertyValue *string                          `json:"propertyValue,omitempty"`
	PropertyType  *ProductconfiguratorPropertyType `json:"propertyType,omitempty"`
	// subtract_to_get_variation is a list of values and is used to calculate the variation from the property value.
	SubtractToGetVariation []string `json:"subtractToGetVariation,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _ConfigurationOptionProperty ConfigurationOptionProperty

// NewConfigurationOptionProperty instantiates a new ConfigurationOptionProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigurationOptionProperty() *ConfigurationOptionProperty {
	this := ConfigurationOptionProperty{}
	var propertyType ProductconfiguratorPropertyType = PRODUCTCONFIGURATORPROPERTYTYPE_UNKNOWN
	this.PropertyType = &propertyType
	return &this
}

// NewConfigurationOptionPropertyWithDefaults instantiates a new ConfigurationOptionProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurationOptionPropertyWithDefaults() *ConfigurationOptionProperty {
	this := ConfigurationOptionProperty{}
	var propertyType ProductconfiguratorPropertyType = PRODUCTCONFIGURATORPROPERTYTYPE_UNKNOWN
	this.PropertyType = &propertyType
	return &this
}

// GetPropertyId returns the PropertyId field value if set, zero value otherwise.
func (o *ConfigurationOptionProperty) GetPropertyId() string {
	if o == nil || IsNil(o.PropertyId) {
		var ret string
		return ret
	}
	return *o.PropertyId
}

// GetPropertyIdOk returns a tuple with the PropertyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationOptionProperty) GetPropertyIdOk() (*string, bool) {
	if o == nil || IsNil(o.PropertyId) {
		return nil, false
	}
	return o.PropertyId, true
}

// HasPropertyId returns a boolean if a field has been set.
func (o *ConfigurationOptionProperty) IsSetPropertyId() bool {
	if o != nil && !IsNil(o.PropertyId) {
		return true
	}

	return false
}

// SetPropertyId gets a reference to the given string and assigns it to the PropertyId field.
func (o *ConfigurationOptionProperty) SetPropertyId(v string) {
	o.PropertyId = &v
}

// GetPropertyValue returns the PropertyValue field value if set, zero value otherwise.
func (o *ConfigurationOptionProperty) GetPropertyValue() string {
	if o == nil || IsNil(o.PropertyValue) {
		var ret string
		return ret
	}
	return *o.PropertyValue
}

// GetPropertyValueOk returns a tuple with the PropertyValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationOptionProperty) GetPropertyValueOk() (*string, bool) {
	if o == nil || IsNil(o.PropertyValue) {
		return nil, false
	}
	return o.PropertyValue, true
}

// HasPropertyValue returns a boolean if a field has been set.
func (o *ConfigurationOptionProperty) IsSetPropertyValue() bool {
	if o != nil && !IsNil(o.PropertyValue) {
		return true
	}

	return false
}

// SetPropertyValue gets a reference to the given string and assigns it to the PropertyValue field.
func (o *ConfigurationOptionProperty) SetPropertyValue(v string) {
	o.PropertyValue = &v
}

// GetPropertyType returns the PropertyType field value if set, zero value otherwise.
func (o *ConfigurationOptionProperty) GetPropertyType() ProductconfiguratorPropertyType {
	if o == nil || IsNil(o.PropertyType) {
		var ret ProductconfiguratorPropertyType
		return ret
	}
	return *o.PropertyType
}

// GetPropertyTypeOk returns a tuple with the PropertyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationOptionProperty) GetPropertyTypeOk() (*ProductconfiguratorPropertyType, bool) {
	if o == nil || IsNil(o.PropertyType) {
		return nil, false
	}
	return o.PropertyType, true
}

// HasPropertyType returns a boolean if a field has been set.
func (o *ConfigurationOptionProperty) IsSetPropertyType() bool {
	if o != nil && !IsNil(o.PropertyType) {
		return true
	}

	return false
}

// SetPropertyType gets a reference to the given ProductconfiguratorPropertyType and assigns it to the PropertyType field.
func (o *ConfigurationOptionProperty) SetPropertyType(v ProductconfiguratorPropertyType) {
	o.PropertyType = &v
}

// GetSubtractToGetVariation returns the SubtractToGetVariation field value if set, zero value otherwise.
func (o *ConfigurationOptionProperty) GetSubtractToGetVariation() []string {
	if o == nil || IsNil(o.SubtractToGetVariation) {
		var ret []string
		return ret
	}
	return o.SubtractToGetVariation
}

// GetSubtractToGetVariationOk returns a tuple with the SubtractToGetVariation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationOptionProperty) GetSubtractToGetVariationOk() ([]string, bool) {
	if o == nil || IsNil(o.SubtractToGetVariation) {
		return nil, false
	}
	return o.SubtractToGetVariation, true
}

// HasSubtractToGetVariation returns a boolean if a field has been set.
func (o *ConfigurationOptionProperty) IsSetSubtractToGetVariation() bool {
	if o != nil && !IsNil(o.SubtractToGetVariation) {
		return true
	}

	return false
}

// SetSubtractToGetVariation gets a reference to the given []string and assigns it to the SubtractToGetVariation field.
func (o *ConfigurationOptionProperty) SetSubtractToGetVariation(v []string) {
	o.SubtractToGetVariation = v
}

func (o ConfigurationOptionProperty) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigurationOptionProperty) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PropertyId) {
		toSerialize["propertyId"] = o.PropertyId
	}
	if !IsNil(o.PropertyValue) {
		toSerialize["propertyValue"] = o.PropertyValue
	}
	if !IsNil(o.PropertyType) {
		toSerialize["propertyType"] = o.PropertyType
	}
	if !IsNil(o.SubtractToGetVariation) {
		toSerialize["subtractToGetVariation"] = o.SubtractToGetVariation
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConfigurationOptionProperty) UnmarshalJSON(data []byte) (err error) {
	varConfigurationOptionProperty := _ConfigurationOptionProperty{}

	err = json.Unmarshal(data, &varConfigurationOptionProperty)

	if err != nil {
		return err
	}

	*o = ConfigurationOptionProperty(varConfigurationOptionProperty)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "propertyId")
		delete(additionalProperties, "propertyValue")
		delete(additionalProperties, "propertyType")
		delete(additionalProperties, "subtractToGetVariation")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ConfigurationOptionProperty) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *ConfigurationOptionProperty) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableConfigurationOptionProperty struct {
	value *ConfigurationOptionProperty
	isSet bool
}

func (v NullableConfigurationOptionProperty) Get() *ConfigurationOptionProperty {
	return v.value
}

func (v *NullableConfigurationOptionProperty) Set(val *ConfigurationOptionProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurationOptionProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurationOptionProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurationOptionProperty(val *ConfigurationOptionProperty) *NullableConfigurationOptionProperty {
	return &NullableConfigurationOptionProperty{value: val, isSet: true}
}

func (v NullableConfigurationOptionProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurationOptionProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
