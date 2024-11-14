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

// checks if the ConfigurationGetConfigurationFromSelectionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigurationGetConfigurationFromSelectionsResponse{}

// ConfigurationGetConfigurationFromSelectionsResponse struct for ConfigurationGetConfigurationFromSelectionsResponse
type ConfigurationGetConfigurationFromSelectionsResponse struct {
	Steps []ConfigurationConfigurationStep `json:"steps,omitempty"`
	MatchedProperties []ConfigurationProperty `json:"matchedProperties,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConfigurationGetConfigurationFromSelectionsResponse ConfigurationGetConfigurationFromSelectionsResponse

// NewConfigurationGetConfigurationFromSelectionsResponse instantiates a new ConfigurationGetConfigurationFromSelectionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigurationGetConfigurationFromSelectionsResponse() *ConfigurationGetConfigurationFromSelectionsResponse {
	this := ConfigurationGetConfigurationFromSelectionsResponse{}
	return &this
}

// NewConfigurationGetConfigurationFromSelectionsResponseWithDefaults instantiates a new ConfigurationGetConfigurationFromSelectionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurationGetConfigurationFromSelectionsResponseWithDefaults() *ConfigurationGetConfigurationFromSelectionsResponse {
	this := ConfigurationGetConfigurationFromSelectionsResponse{}
	return &this
}

// GetSteps returns the Steps field value if set, zero value otherwise.
func (o *ConfigurationGetConfigurationFromSelectionsResponse) GetSteps() []ConfigurationConfigurationStep {
	if o == nil || IsNil(o.Steps) {
		var ret []ConfigurationConfigurationStep
		return ret
	}
	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationGetConfigurationFromSelectionsResponse) GetStepsOk() ([]ConfigurationConfigurationStep, bool) {
	if o == nil || IsNil(o.Steps) {
		return nil, false
	}
	return o.Steps, true
}

// &#39;Has&#39;Steps returns a boolean if a field has been set.
func (o *ConfigurationGetConfigurationFromSelectionsResponse) &#39;Has&#39;Steps() bool {
	if o != nil && !IsNil(o.Steps) {
		return true
	}

	return false
}

// SetSteps gets a reference to the given []ConfigurationConfigurationStep and assigns it to the Steps field.
func (o *ConfigurationGetConfigurationFromSelectionsResponse) SetSteps(v []ConfigurationConfigurationStep) {
	o.Steps = v
}

// GetMatchedProperties returns the MatchedProperties field value if set, zero value otherwise.
func (o *ConfigurationGetConfigurationFromSelectionsResponse) GetMatchedProperties() []ConfigurationProperty {
	if o == nil || IsNil(o.MatchedProperties) {
		var ret []ConfigurationProperty
		return ret
	}
	return o.MatchedProperties
}

// GetMatchedPropertiesOk returns a tuple with the MatchedProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationGetConfigurationFromSelectionsResponse) GetMatchedPropertiesOk() ([]ConfigurationProperty, bool) {
	if o == nil || IsNil(o.MatchedProperties) {
		return nil, false
	}
	return o.MatchedProperties, true
}

// &#39;Has&#39;MatchedProperties returns a boolean if a field has been set.
func (o *ConfigurationGetConfigurationFromSelectionsResponse) &#39;Has&#39;MatchedProperties() bool {
	if o != nil && !IsNil(o.MatchedProperties) {
		return true
	}

	return false
}

// SetMatchedProperties gets a reference to the given []ConfigurationProperty and assigns it to the MatchedProperties field.
func (o *ConfigurationGetConfigurationFromSelectionsResponse) SetMatchedProperties(v []ConfigurationProperty) {
	o.MatchedProperties = v
}

func (o ConfigurationGetConfigurationFromSelectionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigurationGetConfigurationFromSelectionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Steps) {
		toSerialize["steps"] = o.Steps
	}
	if !IsNil(o.MatchedProperties) {
		toSerialize["matchedProperties"] = o.MatchedProperties
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConfigurationGetConfigurationFromSelectionsResponse) UnmarshalJSON(data []byte) (err error) {
	varConfigurationGetConfigurationFromSelectionsResponse := _ConfigurationGetConfigurationFromSelectionsResponse{}

	err = json.Unmarshal(data, &varConfigurationGetConfigurationFromSelectionsResponse)

	if err != nil {
		return err
	}

	*o = ConfigurationGetConfigurationFromSelectionsResponse(varConfigurationGetConfigurationFromSelectionsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "steps")
		delete(additionalProperties, "matchedProperties")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ConfigurationGetConfigurationFromSelectionsResponse) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *ConfigurationGetConfigurationFromSelectionsResponse) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableConfigurationGetConfigurationFromSelectionsResponse struct {
	value *ConfigurationGetConfigurationFromSelectionsResponse
	isSet bool
}

func (v NullableConfigurationGetConfigurationFromSelectionsResponse) Get() *ConfigurationGetConfigurationFromSelectionsResponse {
	return v.value
}

func (v *NullableConfigurationGetConfigurationFromSelectionsResponse) Set(val *ConfigurationGetConfigurationFromSelectionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurationGetConfigurationFromSelectionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurationGetConfigurationFromSelectionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurationGetConfigurationFromSelectionsResponse(val *ConfigurationGetConfigurationFromSelectionsResponse) *NullableConfigurationGetConfigurationFromSelectionsResponse {
	return &NullableConfigurationGetConfigurationFromSelectionsResponse{value: val, isSet: true}
}

func (v NullableConfigurationGetConfigurationFromSelectionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurationGetConfigurationFromSelectionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


