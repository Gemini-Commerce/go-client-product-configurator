/*
Product Configurator Service

## Introduction  This comprehensive guide will equip you with the knowledge to integrate and leverage our Product Configurator Service in your applications.  ## Quick Start  Get up and running in no time! Follow these steps to kickstart your integration:  1. **Authentication:** Obtain your integration JWT to authenticate your requests. 2. **Client Libraries:** Explore our GitHub repositories to grab client libraries in your preferred programming language. 3. **API Overview:** Familiarize yourself with our RESTful API using the OpenAPI specification.  ## Integration  ### API Overview  Our RESTful API is the gateway to unlocking the full potential of Product Configurator. Check out the detailed [API Reference](/docs/category/configurator) for a granular understanding of each endpoint and request/response format.  ### Client Libraries  To expedite your integration process, we provide client libraries for various programming languages. Find the one that suits your stack in our [GitHub repositories](https://github.com/Gemini-Commerce).  ### Authentication  Security is paramount. Learn how to authenticate your requests using JWT. This ensures a secure and reliable connection between your application and Product Configurator.  ## Configuration Management  ### Configurator Lifecycle  Understand the lifecycle of configurators, from draft to active and deleted. This flexibility allows you to manage configurations at your own pace.  ### Steps and Options  Configure product steps with ease and define options effortlessly. Explore the power of dependencies to create dynamic and intuitive configurations.  ### Matrices  Delve into matrices—your secret weapon. Explore price and weight matrices, and learn how configured steps influence properties and pricing.  ### Price Management  Unleash dynamic pricing with our versatile price matrices. From fixed prices to incremental structures, adapt to diverse pricing models effortlessly.  ## Security  Your data is in safe hands. Discover how Product Configurator ensures security through JWT authentication, safeguarding your sensitive information.  ## Backward Compatibility  Stay ahead of the curve. Learn about our versioning strategy, providing backward compatibility while allowing our service to evolve seamlessly.  ## Developer Support  Have questions? Need assistance? Write to us at [info@gemini-commerce.com](mailto:info@gemini-commerce.com) and we will get back to you.

API version: v1
Contact: info@gemini-commerce.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package productconfigurator

import (
	"encoding/json"
)

// checks if the ConfigurationProperty type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigurationProperty{}

// ConfigurationProperty struct for ConfigurationProperty
type ConfigurationProperty struct {
	Id                   *string                          `json:"id,omitempty"`
	Grn                  *string                          `json:"grn,omitempty"`
	OptionIds            []string                         `json:"optionIds,omitempty"`
	PropertyKey          *string                          `json:"propertyKey,omitempty"`
	PropertyValue        *string                          `json:"propertyValue,omitempty"`
	PropertyType         *ProductconfiguratorPropertyType `json:"propertyType,omitempty"`
	PropertyMode         *ProductconfiguratorPropertyMode `json:"propertyMode,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConfigurationProperty ConfigurationProperty

// NewConfigurationProperty instantiates a new ConfigurationProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigurationProperty() *ConfigurationProperty {
	this := ConfigurationProperty{}
	var propertyType ProductconfiguratorPropertyType = PRODUCTCONFIGURATORPROPERTYTYPE_UNKNOWN
	this.PropertyType = &propertyType
	var propertyMode ProductconfiguratorPropertyMode = PRODUCTCONFIGURATORPROPERTYMODE_UNKNOWN
	this.PropertyMode = &propertyMode
	return &this
}

// NewConfigurationPropertyWithDefaults instantiates a new ConfigurationProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurationPropertyWithDefaults() *ConfigurationProperty {
	this := ConfigurationProperty{}
	var propertyType ProductconfiguratorPropertyType = PRODUCTCONFIGURATORPROPERTYTYPE_UNKNOWN
	this.PropertyType = &propertyType
	var propertyMode ProductconfiguratorPropertyMode = PRODUCTCONFIGURATORPROPERTYMODE_UNKNOWN
	this.PropertyMode = &propertyMode
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ConfigurationProperty) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationProperty) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ConfigurationProperty) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ConfigurationProperty) SetId(v string) {
	o.Id = &v
}

// GetGrn returns the Grn field value if set, zero value otherwise.
func (o *ConfigurationProperty) GetGrn() string {
	if o == nil || IsNil(o.Grn) {
		var ret string
		return ret
	}
	return *o.Grn
}

// GetGrnOk returns a tuple with the Grn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationProperty) GetGrnOk() (*string, bool) {
	if o == nil || IsNil(o.Grn) {
		return nil, false
	}
	return o.Grn, true
}

// HasGrn returns a boolean if a field has been set.
func (o *ConfigurationProperty) HasGrn() bool {
	if o != nil && !IsNil(o.Grn) {
		return true
	}

	return false
}

// SetGrn gets a reference to the given string and assigns it to the Grn field.
func (o *ConfigurationProperty) SetGrn(v string) {
	o.Grn = &v
}

// GetOptionIds returns the OptionIds field value if set, zero value otherwise.
func (o *ConfigurationProperty) GetOptionIds() []string {
	if o == nil || IsNil(o.OptionIds) {
		var ret []string
		return ret
	}
	return o.OptionIds
}

// GetOptionIdsOk returns a tuple with the OptionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationProperty) GetOptionIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.OptionIds) {
		return nil, false
	}
	return o.OptionIds, true
}

// HasOptionIds returns a boolean if a field has been set.
func (o *ConfigurationProperty) HasOptionIds() bool {
	if o != nil && !IsNil(o.OptionIds) {
		return true
	}

	return false
}

// SetOptionIds gets a reference to the given []string and assigns it to the OptionIds field.
func (o *ConfigurationProperty) SetOptionIds(v []string) {
	o.OptionIds = v
}

// GetPropertyKey returns the PropertyKey field value if set, zero value otherwise.
func (o *ConfigurationProperty) GetPropertyKey() string {
	if o == nil || IsNil(o.PropertyKey) {
		var ret string
		return ret
	}
	return *o.PropertyKey
}

// GetPropertyKeyOk returns a tuple with the PropertyKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationProperty) GetPropertyKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PropertyKey) {
		return nil, false
	}
	return o.PropertyKey, true
}

// HasPropertyKey returns a boolean if a field has been set.
func (o *ConfigurationProperty) HasPropertyKey() bool {
	if o != nil && !IsNil(o.PropertyKey) {
		return true
	}

	return false
}

// SetPropertyKey gets a reference to the given string and assigns it to the PropertyKey field.
func (o *ConfigurationProperty) SetPropertyKey(v string) {
	o.PropertyKey = &v
}

// GetPropertyValue returns the PropertyValue field value if set, zero value otherwise.
func (o *ConfigurationProperty) GetPropertyValue() string {
	if o == nil || IsNil(o.PropertyValue) {
		var ret string
		return ret
	}
	return *o.PropertyValue
}

// GetPropertyValueOk returns a tuple with the PropertyValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationProperty) GetPropertyValueOk() (*string, bool) {
	if o == nil || IsNil(o.PropertyValue) {
		return nil, false
	}
	return o.PropertyValue, true
}

// HasPropertyValue returns a boolean if a field has been set.
func (o *ConfigurationProperty) HasPropertyValue() bool {
	if o != nil && !IsNil(o.PropertyValue) {
		return true
	}

	return false
}

// SetPropertyValue gets a reference to the given string and assigns it to the PropertyValue field.
func (o *ConfigurationProperty) SetPropertyValue(v string) {
	o.PropertyValue = &v
}

// GetPropertyType returns the PropertyType field value if set, zero value otherwise.
func (o *ConfigurationProperty) GetPropertyType() ProductconfiguratorPropertyType {
	if o == nil || IsNil(o.PropertyType) {
		var ret ProductconfiguratorPropertyType
		return ret
	}
	return *o.PropertyType
}

// GetPropertyTypeOk returns a tuple with the PropertyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationProperty) GetPropertyTypeOk() (*ProductconfiguratorPropertyType, bool) {
	if o == nil || IsNil(o.PropertyType) {
		return nil, false
	}
	return o.PropertyType, true
}

// HasPropertyType returns a boolean if a field has been set.
func (o *ConfigurationProperty) HasPropertyType() bool {
	if o != nil && !IsNil(o.PropertyType) {
		return true
	}

	return false
}

// SetPropertyType gets a reference to the given ProductconfiguratorPropertyType and assigns it to the PropertyType field.
func (o *ConfigurationProperty) SetPropertyType(v ProductconfiguratorPropertyType) {
	o.PropertyType = &v
}

// GetPropertyMode returns the PropertyMode field value if set, zero value otherwise.
func (o *ConfigurationProperty) GetPropertyMode() ProductconfiguratorPropertyMode {
	if o == nil || IsNil(o.PropertyMode) {
		var ret ProductconfiguratorPropertyMode
		return ret
	}
	return *o.PropertyMode
}

// GetPropertyModeOk returns a tuple with the PropertyMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationProperty) GetPropertyModeOk() (*ProductconfiguratorPropertyMode, bool) {
	if o == nil || IsNil(o.PropertyMode) {
		return nil, false
	}
	return o.PropertyMode, true
}

// HasPropertyMode returns a boolean if a field has been set.
func (o *ConfigurationProperty) HasPropertyMode() bool {
	if o != nil && !IsNil(o.PropertyMode) {
		return true
	}

	return false
}

// SetPropertyMode gets a reference to the given ProductconfiguratorPropertyMode and assigns it to the PropertyMode field.
func (o *ConfigurationProperty) SetPropertyMode(v ProductconfiguratorPropertyMode) {
	o.PropertyMode = &v
}

func (o ConfigurationProperty) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigurationProperty) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Grn) {
		toSerialize["grn"] = o.Grn
	}
	if !IsNil(o.OptionIds) {
		toSerialize["optionIds"] = o.OptionIds
	}
	if !IsNil(o.PropertyKey) {
		toSerialize["propertyKey"] = o.PropertyKey
	}
	if !IsNil(o.PropertyValue) {
		toSerialize["propertyValue"] = o.PropertyValue
	}
	if !IsNil(o.PropertyType) {
		toSerialize["propertyType"] = o.PropertyType
	}
	if !IsNil(o.PropertyMode) {
		toSerialize["propertyMode"] = o.PropertyMode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConfigurationProperty) UnmarshalJSON(data []byte) (err error) {
	varConfigurationProperty := _ConfigurationProperty{}

	err = json.Unmarshal(data, &varConfigurationProperty)

	if err != nil {
		return err
	}

	*o = ConfigurationProperty(varConfigurationProperty)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "grn")
		delete(additionalProperties, "optionIds")
		delete(additionalProperties, "propertyKey")
		delete(additionalProperties, "propertyValue")
		delete(additionalProperties, "propertyType")
		delete(additionalProperties, "propertyMode")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ConfigurationProperty) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *ConfigurationProperty) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableConfigurationProperty struct {
	value *ConfigurationProperty
	isSet bool
}

func (v NullableConfigurationProperty) Get() *ConfigurationProperty {
	return v.value
}

func (v *NullableConfigurationProperty) Set(val *ConfigurationProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurationProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurationProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurationProperty(val *ConfigurationProperty) *NullableConfigurationProperty {
	return &NullableConfigurationProperty{value: val, isSet: true}
}

func (v NullableConfigurationProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurationProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
