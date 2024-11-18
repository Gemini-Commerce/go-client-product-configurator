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

// checks if the ConfigurationGetAvailableConfigurationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigurationGetAvailableConfigurationResponse{}

// ConfigurationGetAvailableConfigurationResponse struct for ConfigurationGetAvailableConfigurationResponse
type ConfigurationGetAvailableConfigurationResponse struct {
	Configurator         *ConfigurationConfigurator                  `json:"configurator,omitempty"`
	MatchedProperties    []ConfigurationProperty                     `json:"matchedProperties,omitempty"`
	Selections           []ProductconfiguratorconfigurationSelection `json:"selections,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConfigurationGetAvailableConfigurationResponse ConfigurationGetAvailableConfigurationResponse

// NewConfigurationGetAvailableConfigurationResponse instantiates a new ConfigurationGetAvailableConfigurationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigurationGetAvailableConfigurationResponse() *ConfigurationGetAvailableConfigurationResponse {
	this := ConfigurationGetAvailableConfigurationResponse{}
	return &this
}

// NewConfigurationGetAvailableConfigurationResponseWithDefaults instantiates a new ConfigurationGetAvailableConfigurationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurationGetAvailableConfigurationResponseWithDefaults() *ConfigurationGetAvailableConfigurationResponse {
	this := ConfigurationGetAvailableConfigurationResponse{}
	return &this
}

// GetConfigurator returns the Configurator field value if set, zero value otherwise.
func (o *ConfigurationGetAvailableConfigurationResponse) GetConfigurator() ConfigurationConfigurator {
	if o == nil || IsNil(o.Configurator) {
		var ret ConfigurationConfigurator
		return ret
	}
	return *o.Configurator
}

// GetConfiguratorOk returns a tuple with the Configurator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationGetAvailableConfigurationResponse) GetConfiguratorOk() (*ConfigurationConfigurator, bool) {
	if o == nil || IsNil(o.Configurator) {
		return nil, false
	}
	return o.Configurator, true
}

// HasConfigurator returns a boolean if a field has been set.
func (o *ConfigurationGetAvailableConfigurationResponse) HasConfigurator() bool {
	if o != nil && !IsNil(o.Configurator) {
		return true
	}

	return false
}

// SetConfigurator gets a reference to the given ConfigurationConfigurator and assigns it to the Configurator field.
func (o *ConfigurationGetAvailableConfigurationResponse) SetConfigurator(v ConfigurationConfigurator) {
	o.Configurator = &v
}

// GetMatchedProperties returns the MatchedProperties field value if set, zero value otherwise.
func (o *ConfigurationGetAvailableConfigurationResponse) GetMatchedProperties() []ConfigurationProperty {
	if o == nil || IsNil(o.MatchedProperties) {
		var ret []ConfigurationProperty
		return ret
	}
	return o.MatchedProperties
}

// GetMatchedPropertiesOk returns a tuple with the MatchedProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationGetAvailableConfigurationResponse) GetMatchedPropertiesOk() ([]ConfigurationProperty, bool) {
	if o == nil || IsNil(o.MatchedProperties) {
		return nil, false
	}
	return o.MatchedProperties, true
}

// HasMatchedProperties returns a boolean if a field has been set.
func (o *ConfigurationGetAvailableConfigurationResponse) HasMatchedProperties() bool {
	if o != nil && !IsNil(o.MatchedProperties) {
		return true
	}

	return false
}

// SetMatchedProperties gets a reference to the given []ConfigurationProperty and assigns it to the MatchedProperties field.
func (o *ConfigurationGetAvailableConfigurationResponse) SetMatchedProperties(v []ConfigurationProperty) {
	o.MatchedProperties = v
}

// GetSelections returns the Selections field value if set, zero value otherwise.
func (o *ConfigurationGetAvailableConfigurationResponse) GetSelections() []ProductconfiguratorconfigurationSelection {
	if o == nil || IsNil(o.Selections) {
		var ret []ProductconfiguratorconfigurationSelection
		return ret
	}
	return o.Selections
}

// GetSelectionsOk returns a tuple with the Selections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationGetAvailableConfigurationResponse) GetSelectionsOk() ([]ProductconfiguratorconfigurationSelection, bool) {
	if o == nil || IsNil(o.Selections) {
		return nil, false
	}
	return o.Selections, true
}

// HasSelections returns a boolean if a field has been set.
func (o *ConfigurationGetAvailableConfigurationResponse) HasSelections() bool {
	if o != nil && !IsNil(o.Selections) {
		return true
	}

	return false
}

// SetSelections gets a reference to the given []ProductconfiguratorconfigurationSelection and assigns it to the Selections field.
func (o *ConfigurationGetAvailableConfigurationResponse) SetSelections(v []ProductconfiguratorconfigurationSelection) {
	o.Selections = v
}

func (o ConfigurationGetAvailableConfigurationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigurationGetAvailableConfigurationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Configurator) {
		toSerialize["configurator"] = o.Configurator
	}
	if !IsNil(o.MatchedProperties) {
		toSerialize["matchedProperties"] = o.MatchedProperties
	}
	if !IsNil(o.Selections) {
		toSerialize["selections"] = o.Selections
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConfigurationGetAvailableConfigurationResponse) UnmarshalJSON(data []byte) (err error) {
	varConfigurationGetAvailableConfigurationResponse := _ConfigurationGetAvailableConfigurationResponse{}

	err = json.Unmarshal(data, &varConfigurationGetAvailableConfigurationResponse)

	if err != nil {
		return err
	}

	*o = ConfigurationGetAvailableConfigurationResponse(varConfigurationGetAvailableConfigurationResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "configurator")
		delete(additionalProperties, "matchedProperties")
		delete(additionalProperties, "selections")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ConfigurationGetAvailableConfigurationResponse) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *ConfigurationGetAvailableConfigurationResponse) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableConfigurationGetAvailableConfigurationResponse struct {
	value *ConfigurationGetAvailableConfigurationResponse
	isSet bool
}

func (v NullableConfigurationGetAvailableConfigurationResponse) Get() *ConfigurationGetAvailableConfigurationResponse {
	return v.value
}

func (v *NullableConfigurationGetAvailableConfigurationResponse) Set(val *ConfigurationGetAvailableConfigurationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurationGetAvailableConfigurationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurationGetAvailableConfigurationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurationGetAvailableConfigurationResponse(val *ConfigurationGetAvailableConfigurationResponse) *NullableConfigurationGetAvailableConfigurationResponse {
	return &NullableConfigurationGetAvailableConfigurationResponse{value: val, isSet: true}
}

func (v NullableConfigurationGetAvailableConfigurationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurationGetAvailableConfigurationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
