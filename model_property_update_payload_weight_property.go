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

// checks if the PropertyUpdatePayloadWeightProperty type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PropertyUpdatePayloadWeightProperty{}

// PropertyUpdatePayloadWeightProperty struct for PropertyUpdatePayloadWeightProperty
type PropertyUpdatePayloadWeightProperty struct {
	Weight               *float64 `json:"weight,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PropertyUpdatePayloadWeightProperty PropertyUpdatePayloadWeightProperty

// NewPropertyUpdatePayloadWeightProperty instantiates a new PropertyUpdatePayloadWeightProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPropertyUpdatePayloadWeightProperty() *PropertyUpdatePayloadWeightProperty {
	this := PropertyUpdatePayloadWeightProperty{}
	return &this
}

// NewPropertyUpdatePayloadWeightPropertyWithDefaults instantiates a new PropertyUpdatePayloadWeightProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertyUpdatePayloadWeightPropertyWithDefaults() *PropertyUpdatePayloadWeightProperty {
	this := PropertyUpdatePayloadWeightProperty{}
	return &this
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *PropertyUpdatePayloadWeightProperty) GetWeight() float64 {
	if o == nil || IsNil(o.Weight) {
		var ret float64
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyUpdatePayloadWeightProperty) GetWeightOk() (*float64, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *PropertyUpdatePayloadWeightProperty) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given float64 and assigns it to the Weight field.
func (o *PropertyUpdatePayloadWeightProperty) SetWeight(v float64) {
	o.Weight = &v
}

func (o PropertyUpdatePayloadWeightProperty) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PropertyUpdatePayloadWeightProperty) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PropertyUpdatePayloadWeightProperty) UnmarshalJSON(data []byte) (err error) {
	varPropertyUpdatePayloadWeightProperty := _PropertyUpdatePayloadWeightProperty{}

	err = json.Unmarshal(data, &varPropertyUpdatePayloadWeightProperty)

	if err != nil {
		return err
	}

	*o = PropertyUpdatePayloadWeightProperty(varPropertyUpdatePayloadWeightProperty)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "weight")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PropertyUpdatePayloadWeightProperty) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *PropertyUpdatePayloadWeightProperty) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullablePropertyUpdatePayloadWeightProperty struct {
	value *PropertyUpdatePayloadWeightProperty
	isSet bool
}

func (v NullablePropertyUpdatePayloadWeightProperty) Get() *PropertyUpdatePayloadWeightProperty {
	return v.value
}

func (v *NullablePropertyUpdatePayloadWeightProperty) Set(val *PropertyUpdatePayloadWeightProperty) {
	v.value = val
	v.isSet = true
}

func (v NullablePropertyUpdatePayloadWeightProperty) IsSet() bool {
	return v.isSet
}

func (v *NullablePropertyUpdatePayloadWeightProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropertyUpdatePayloadWeightProperty(val *PropertyUpdatePayloadWeightProperty) *NullablePropertyUpdatePayloadWeightProperty {
	return &NullablePropertyUpdatePayloadWeightProperty{value: val, isSet: true}
}

func (v NullablePropertyUpdatePayloadWeightProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropertyUpdatePayloadWeightProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
