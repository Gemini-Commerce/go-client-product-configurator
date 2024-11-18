/*
Product Configurator Service

## Introduction  This comprehensive guide will equip you with the knowledge to integrate and leverage our Product Configurator Service in your applications.  ## Quick Start  Get up and running in no time! Follow these steps to kickstart your integration:  1. **Authentication:** Obtain your integration JWT to authenticate your requests. 2. **Client Libraries:** Explore our GitHub repositories to grab client libraries in your preferred programming language. 3. **API Overview:** Familiarize yourself with our RESTful API using the OpenAPI specification.  ## Integration  ### API Overview  Our RESTful API is the gateway to unlocking the full potential of Product Configurator. Check out the detailed [API Reference](/docs/category/configurator) for a granular understanding of each endpoint and request/response format.  ### Client Libraries  To expedite your integration process, we provide client libraries for various programming languages. Find the one that suits your stack in our [GitHub repositories](https://github.com/GeminiCommerce).  ### Authentication  Security is paramount. Learn how to authenticate your requests using JWT. This ensures a secure and reliable connection between your application and Product Configurator.  ## Configuration Management  ### Configurator Lifecycle  Understand the lifecycle of configurators, from draft to active and deleted. This flexibility allows you to manage configurations at your own pace.  ### Steps and Options  Configure product steps with ease and define options effortlessly. Explore the power of dependencies to create dynamic and intuitive configurations.  ### Matrices  Delve into matrices—your secret weapon. Explore price and weight matrices, and learn how configured steps influence properties and pricing.  ### Price Management  Unleash dynamic pricing with our versatile price matrices. From fixed prices to incremental structures, adapt to diverse pricing models effortlessly.  ## Security  Your data is in safe hands. Discover how Product Configurator ensures security through JWT authentication, safeguarding your sensitive information.  ## Backward Compatibility  Stay ahead of the curve. Learn about our versioning strategy, providing backward compatibility while allowing our service to evolve seamlessly.  ## Developer Support  Have questions? Need assistance? Write to us at [info@geminicommerce.com](mailto:info@geminicommerce.com) and we will get back to you.

API version: v1
Contact: info@geminicommerce.com
*/

// Code generated by OpenAPI Generator (https://openapigenerator.tech); DO NOT EDIT.

package productconfigurator

import (
	"encoding/json"
)

// checks if the LocalisationLocalizedText type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocalisationLocalizedText{}

// LocalisationLocalizedText struct for LocalisationLocalizedText
type LocalisationLocalizedText struct {
	Value                *map[string]string `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LocalisationLocalizedText LocalisationLocalizedText

// NewLocalisationLocalizedText instantiates a new LocalisationLocalizedText object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocalisationLocalizedText() *LocalisationLocalizedText {
	this := LocalisationLocalizedText{}
	return &this
}

// NewLocalisationLocalizedTextWithDefaults instantiates a new LocalisationLocalizedText object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocalisationLocalizedTextWithDefaults() *LocalisationLocalizedText {
	this := LocalisationLocalizedText{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *LocalisationLocalizedText) GetValue() map[string]string {
	if o == nil || IsNil(o.Value) {
		var ret map[string]string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalisationLocalizedText) GetValueOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *LocalisationLocalizedText) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given map[string]string and assigns it to the Value field.
func (o *LocalisationLocalizedText) SetValue(v map[string]string) {
	o.Value = &v
}

func (o LocalisationLocalizedText) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocalisationLocalizedText) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LocalisationLocalizedText) UnmarshalJSON(data []byte) (err error) {
	varLocalisationLocalizedText := _LocalisationLocalizedText{}

	err = json.Unmarshal(data, &varLocalisationLocalizedText)

	if err != nil {
		return err
	}

	*o = LocalisationLocalizedText(varLocalisationLocalizedText)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLocalisationLocalizedText struct {
	value *LocalisationLocalizedText
	isSet bool
}

func (v NullableLocalisationLocalizedText) Get() *LocalisationLocalizedText {
	return v.value
}

func (v *NullableLocalisationLocalizedText) Set(val *LocalisationLocalizedText) {
	v.value = val
	v.isSet = true
}

func (v NullableLocalisationLocalizedText) IsSet() bool {
	return v.isSet
}

func (v *NullableLocalisationLocalizedText) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocalisationLocalizedText(val *LocalisationLocalizedText) *NullableLocalisationLocalizedText {
	return &NullableLocalisationLocalizedText{value: val, isSet: true}
}

func (v NullableLocalisationLocalizedText) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocalisationLocalizedText) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
