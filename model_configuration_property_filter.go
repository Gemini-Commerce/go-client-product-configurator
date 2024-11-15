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

// checks if the ConfigurationPropertyFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigurationPropertyFilter{}

// ConfigurationPropertyFilter struct for ConfigurationPropertyFilter
type ConfigurationPropertyFilter struct {
	GenericProperty *ConfigurationPropertyFilterGenericProperty `json:"genericProperty,omitempty"`
	PriceProperty map[string]interface{} `json:"priceProperty,omitempty"`
	WeightProperty map[string]interface{} `json:"weightProperty,omitempty"`
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
	toSerialize,err := o.ToMap()
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
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *ConfigurationPropertyFilter) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
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


