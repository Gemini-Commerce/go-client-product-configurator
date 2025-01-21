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

// checks if the ConfigurationPropertyFilterGenericProperty type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigurationPropertyFilterGenericProperty{}

// ConfigurationPropertyFilterGenericProperty struct for ConfigurationPropertyFilterGenericProperty
type ConfigurationPropertyFilterGenericProperty struct {
	PropertyKey          *string `json:"propertyKey,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConfigurationPropertyFilterGenericProperty ConfigurationPropertyFilterGenericProperty

// NewConfigurationPropertyFilterGenericProperty instantiates a new ConfigurationPropertyFilterGenericProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigurationPropertyFilterGenericProperty() *ConfigurationPropertyFilterGenericProperty {
	this := ConfigurationPropertyFilterGenericProperty{}
	return &this
}

// NewConfigurationPropertyFilterGenericPropertyWithDefaults instantiates a new ConfigurationPropertyFilterGenericProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurationPropertyFilterGenericPropertyWithDefaults() *ConfigurationPropertyFilterGenericProperty {
	this := ConfigurationPropertyFilterGenericProperty{}
	return &this
}

// GetPropertyKey returns the PropertyKey field value if set, zero value otherwise.
func (o *ConfigurationPropertyFilterGenericProperty) GetPropertyKey() string {
	if o == nil || IsNil(o.PropertyKey) {
		var ret string
		return ret
	}
	return *o.PropertyKey
}

// GetPropertyKeyOk returns a tuple with the PropertyKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationPropertyFilterGenericProperty) GetPropertyKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PropertyKey) {
		return nil, false
	}
	return o.PropertyKey, true
}

// HasPropertyKey returns a boolean if a field has been set.
func (o *ConfigurationPropertyFilterGenericProperty) HasPropertyKey() bool {
	if o != nil && !IsNil(o.PropertyKey) {
		return true
	}

	return false
}

// SetPropertyKey gets a reference to the given string and assigns it to the PropertyKey field.
func (o *ConfigurationPropertyFilterGenericProperty) SetPropertyKey(v string) {
	o.PropertyKey = &v
}

func (o ConfigurationPropertyFilterGenericProperty) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigurationPropertyFilterGenericProperty) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PropertyKey) {
		toSerialize["propertyKey"] = o.PropertyKey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConfigurationPropertyFilterGenericProperty) UnmarshalJSON(data []byte) (err error) {
	varConfigurationPropertyFilterGenericProperty := _ConfigurationPropertyFilterGenericProperty{}

	err = json.Unmarshal(data, &varConfigurationPropertyFilterGenericProperty)

	if err != nil {
		return err
	}

	*o = ConfigurationPropertyFilterGenericProperty(varConfigurationPropertyFilterGenericProperty)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "propertyKey")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ConfigurationPropertyFilterGenericProperty) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *ConfigurationPropertyFilterGenericProperty) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableConfigurationPropertyFilterGenericProperty struct {
	value *ConfigurationPropertyFilterGenericProperty
	isSet bool
}

func (v NullableConfigurationPropertyFilterGenericProperty) Get() *ConfigurationPropertyFilterGenericProperty {
	return v.value
}

func (v *NullableConfigurationPropertyFilterGenericProperty) Set(val *ConfigurationPropertyFilterGenericProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurationPropertyFilterGenericProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurationPropertyFilterGenericProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurationPropertyFilterGenericProperty(val *ConfigurationPropertyFilterGenericProperty) *NullableConfigurationPropertyFilterGenericProperty {
	return &NullableConfigurationPropertyFilterGenericProperty{value: val, isSet: true}
}

func (v NullableConfigurationPropertyFilterGenericProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurationPropertyFilterGenericProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
