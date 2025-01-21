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

// checks if the ConfigurationPropertyFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigurationPropertyFilter{}

// ConfigurationPropertyFilter struct for ConfigurationPropertyFilter
type ConfigurationPropertyFilter struct {
	GenericProperty      *ConfigurationPropertyFilterGenericProperty `json:"genericProperty,omitempty"`
	PriceProperty        map[string]interface{}                      `json:"priceProperty,omitempty"`
	WeightProperty       map[string]interface{}                      `json:"weightProperty,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConfigurationPropertyFilter ConfigurationPropertyFilter

// NewConfigurationPropertyFilter instantiates a new ConfigurationPropertyFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigurationPropertyFilter() *ConfigurationPropertyFilter {
	this := ConfigurationPropertyFilter{}
	return &this
}

// NewConfigurationPropertyFilterWithDefaults instantiates a new ConfigurationPropertyFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurationPropertyFilterWithDefaults() *ConfigurationPropertyFilter {
	this := ConfigurationPropertyFilter{}
	return &this
}

// GetGenericProperty returns the GenericProperty field value if set, zero value otherwise.
func (o *ConfigurationPropertyFilter) GetGenericProperty() ConfigurationPropertyFilterGenericProperty {
	if o == nil || IsNil(o.GenericProperty) {
		var ret ConfigurationPropertyFilterGenericProperty
		return ret
	}
	return *o.GenericProperty
}

// GetGenericPropertyOk returns a tuple with the GenericProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationPropertyFilter) GetGenericPropertyOk() (*ConfigurationPropertyFilterGenericProperty, bool) {
	if o == nil || IsNil(o.GenericProperty) {
		return nil, false
	}
	return o.GenericProperty, true
}

// HasGenericProperty returns a boolean if a field has been set.
func (o *ConfigurationPropertyFilter) HasGenericProperty() bool {
	if o != nil && !IsNil(o.GenericProperty) {
		return true
	}

	return false
}

// SetGenericProperty gets a reference to the given ConfigurationPropertyFilterGenericProperty and assigns it to the GenericProperty field.
func (o *ConfigurationPropertyFilter) SetGenericProperty(v ConfigurationPropertyFilterGenericProperty) {
	o.GenericProperty = &v
}

// GetPriceProperty returns the PriceProperty field value if set, zero value otherwise.
func (o *ConfigurationPropertyFilter) GetPriceProperty() map[string]interface{} {
	if o == nil || IsNil(o.PriceProperty) {
		var ret map[string]interface{}
		return ret
	}
	return o.PriceProperty
}

// GetPricePropertyOk returns a tuple with the PriceProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationPropertyFilter) GetPricePropertyOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.PriceProperty) {
		return map[string]interface{}{}, false
	}
	return o.PriceProperty, true
}

// HasPriceProperty returns a boolean if a field has been set.
func (o *ConfigurationPropertyFilter) HasPriceProperty() bool {
	if o != nil && !IsNil(o.PriceProperty) {
		return true
	}

	return false
}

// SetPriceProperty gets a reference to the given map[string]interface{} and assigns it to the PriceProperty field.
func (o *ConfigurationPropertyFilter) SetPriceProperty(v map[string]interface{}) {
	o.PriceProperty = v
}

// GetWeightProperty returns the WeightProperty field value if set, zero value otherwise.
func (o *ConfigurationPropertyFilter) GetWeightProperty() map[string]interface{} {
	if o == nil || IsNil(o.WeightProperty) {
		var ret map[string]interface{}
		return ret
	}
	return o.WeightProperty
}

// GetWeightPropertyOk returns a tuple with the WeightProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationPropertyFilter) GetWeightPropertyOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.WeightProperty) {
		return map[string]interface{}{}, false
	}
	return o.WeightProperty, true
}

// HasWeightProperty returns a boolean if a field has been set.
func (o *ConfigurationPropertyFilter) HasWeightProperty() bool {
	if o != nil && !IsNil(o.WeightProperty) {
		return true
	}

	return false
}

// SetWeightProperty gets a reference to the given map[string]interface{} and assigns it to the WeightProperty field.
func (o *ConfigurationPropertyFilter) SetWeightProperty(v map[string]interface{}) {
	o.WeightProperty = v
}

func (o ConfigurationPropertyFilter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigurationPropertyFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GenericProperty) {
		toSerialize["genericProperty"] = o.GenericProperty
	}
	if !IsNil(o.PriceProperty) {
		toSerialize["priceProperty"] = o.PriceProperty
	}
	if !IsNil(o.WeightProperty) {
		toSerialize["weightProperty"] = o.WeightProperty
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConfigurationPropertyFilter) UnmarshalJSON(data []byte) (err error) {
	varConfigurationPropertyFilter := _ConfigurationPropertyFilter{}

	err = json.Unmarshal(data, &varConfigurationPropertyFilter)

	if err != nil {
		return err
	}

	*o = ConfigurationPropertyFilter(varConfigurationPropertyFilter)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "genericProperty")
		delete(additionalProperties, "priceProperty")
		delete(additionalProperties, "weightProperty")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ConfigurationPropertyFilter) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *ConfigurationPropertyFilter) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableConfigurationPropertyFilter struct {
	value *ConfigurationPropertyFilter
	isSet bool
}

func (v NullableConfigurationPropertyFilter) Get() *ConfigurationPropertyFilter {
	return v.value
}

func (v *NullableConfigurationPropertyFilter) Set(val *ConfigurationPropertyFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurationPropertyFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurationPropertyFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurationPropertyFilter(val *ConfigurationPropertyFilter) *NullableConfigurationPropertyFilter {
	return &NullableConfigurationPropertyFilter{value: val, isSet: true}
}

func (v NullableConfigurationPropertyFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurationPropertyFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
