/*
Product Configurator Service

## Introduction  This comprehensive guide will equip you with the knowledge to integrate and leverage our Product Configurator Service in your applications.  ## Quick Start  Get up and running in no time! Follow these steps to kickstart your integration:  1. **Authentication:** Obtain your integration JWT to authenticate your requests. 2. **Client Libraries:** Explore our GitHub repositories to grab client libraries in your preferred programming language. 3. **API Overview:** Familiarize yourself with our RESTful API using the OpenAPI specification.  ## Integration  ### API Overview  Our RESTful API is the gateway to unlocking the full potential of Product Configurator. Check out the detailed [API Reference](/docs/category/configurator) for a granular understanding of each endpoint and request/response format.  ### Client Libraries  To expedite your integration process, we provide client libraries for various programming languages. Find the one that suits your stack in our [GitHub repositories](https://github.com/Gemini_Commerce).  ### Authentication  Security is paramount. Learn how to authenticate your requests using JWT. This ensures a secure and reliable connection between your application and Product Configurator.  ## Configuration Management  ### Configurator Lifecycle  Understand the lifecycle of configurators, from draft to active and deleted. This flexibility allows you to manage configurations at your own pace.  ### Steps and Options  Configure product steps with ease and define options effortlessly. Explore the power of dependencies to create dynamic and intuitive configurations.  ### Matrices  Delve into matrices—your secret weapon. Explore price and weight matrices, and learn how configured steps influence properties and pricing.  ### Price Management  Unleash dynamic pricing with our versatile price matrices. From fixed prices to incremental structures, adapt to diverse pricing models effortlessly.  ## Security  Your data is in safe hands. Discover how Product Configurator ensures security through JWT authentication, safeguarding your sensitive information.  ## Backward Compatibility  Stay ahead of the curve. Learn about our versioning strategy, providing backward compatibility while allowing our service to evolve seamlessly.  ## Developer Support  Have questions? Need assistance? Write to us at [info@gemini_commerce.com](mailto:info@gemini_commerce.com) and we will get back to you.

API version: v1
Contact: info@gemini_commerce.com
*/

// Code generated by OpenAPI Generator (https://openapi_generator.tech); DO NOT EDIT.

package productconfigurator

import (
	"encoding/json"
)

// checks if the PropertyUpdatePayloadGenericProperty type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PropertyUpdatePayloadGenericProperty{}

// PropertyUpdatePayloadGenericProperty struct for PropertyUpdatePayloadGenericProperty
type PropertyUpdatePayloadGenericProperty struct {
	PropertyValue        *string `json:"propertyValue,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PropertyUpdatePayloadGenericProperty PropertyUpdatePayloadGenericProperty

// NewPropertyUpdatePayloadGenericProperty instantiates a new PropertyUpdatePayloadGenericProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPropertyUpdatePayloadGenericProperty() *PropertyUpdatePayloadGenericProperty {
	this := PropertyUpdatePayloadGenericProperty{}
	return &this
}

// NewPropertyUpdatePayloadGenericPropertyWithDefaults instantiates a new PropertyUpdatePayloadGenericProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertyUpdatePayloadGenericPropertyWithDefaults() *PropertyUpdatePayloadGenericProperty {
	this := PropertyUpdatePayloadGenericProperty{}
	return &this
}

// GetPropertyValue returns the PropertyValue field value if set, zero value otherwise.
func (o *PropertyUpdatePayloadGenericProperty) GetPropertyValue() string {
	if o == nil || IsNil(o.PropertyValue) {
		var ret string
		return ret
	}
	return *o.PropertyValue
}

// GetPropertyValueOk returns a tuple with the PropertyValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyUpdatePayloadGenericProperty) GetPropertyValueOk() (*string, bool) {
	if o == nil || IsNil(o.PropertyValue) {
		return nil, false
	}
	return o.PropertyValue, true
}

// HasPropertyValue returns a boolean if a field has been set.
func (o *PropertyUpdatePayloadGenericProperty) HasPropertyValue() bool {
	if o != nil && !IsNil(o.PropertyValue) {
		return true
	}

	return false
}

// SetPropertyValue gets a reference to the given string and assigns it to the PropertyValue field.
func (o *PropertyUpdatePayloadGenericProperty) SetPropertyValue(v string) {
	o.PropertyValue = &v
}

func (o PropertyUpdatePayloadGenericProperty) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PropertyUpdatePayloadGenericProperty) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PropertyValue) {
		toSerialize["propertyValue"] = o.PropertyValue
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PropertyUpdatePayloadGenericProperty) UnmarshalJSON(data []byte) (err error) {
	varPropertyUpdatePayloadGenericProperty := _PropertyUpdatePayloadGenericProperty{}

	err = json.Unmarshal(data, &varPropertyUpdatePayloadGenericProperty)

	if err != nil {
		return err
	}

	*o = PropertyUpdatePayloadGenericProperty(varPropertyUpdatePayloadGenericProperty)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "propertyValue")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well_known types
func (o *PropertyUpdatePayloadGenericProperty) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well_known types
func (o *PropertyUpdatePayloadGenericProperty) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullablePropertyUpdatePayloadGenericProperty struct {
	value *PropertyUpdatePayloadGenericProperty
	isSet bool
}

func (v NullablePropertyUpdatePayloadGenericProperty) Get() *PropertyUpdatePayloadGenericProperty {
	return v.value
}

func (v *NullablePropertyUpdatePayloadGenericProperty) Set(val *PropertyUpdatePayloadGenericProperty) {
	v.value = val
	v.isSet = true
}

func (v NullablePropertyUpdatePayloadGenericProperty) IsSet() bool {
	return v.isSet
}

func (v *NullablePropertyUpdatePayloadGenericProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropertyUpdatePayloadGenericProperty(val *PropertyUpdatePayloadGenericProperty) *NullablePropertyUpdatePayloadGenericProperty {
	return &NullablePropertyUpdatePayloadGenericProperty{value: val, isSet: true}
}

func (v NullablePropertyUpdatePayloadGenericProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropertyUpdatePayloadGenericProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
