/*
Product Configurator Service

## Introduction  This comprehensive guide will equip you with the knowledge to integrate and leverage our Product Configurator Service in your applications.  ## Quick Start  Get up and running in no time! Follow these steps to kickstart your integration:  1. **Authentication:** Obtain your integration JWT to authenticate your requests. 2. **Client Libraries:** Explore our GitHub repositories to grab client libraries in your preferred programming language. 3. **API Overview:** Familiarize yourself with our RESTful API using the OpenAPI specification.  ## Integration  ### API Overview  Our RESTful API is the gateway to unlocking the full potential of Product Configurator. Check out the detailed [API Reference](/docs/category/configurator) for a granular understanding of each endpoint and request/response format.  ### Client Libraries  To expedite your integration process, we provide client libraries for various programming languages. Find the one that suits your stack in our [GitHub repositories](https://github.com/Gemini-Commerce).  ### Authentication  Security is paramount. Learn how to authenticate your requests using JWT. This ensures a secure and reliable connection between your application and Product Configurator.  ## Configuration Management  ### Configurator Lifecycle  Understand the lifecycle of configurators, from draft to active and deleted. This flexibility allows you to manage configurations at your own pace.  ### Steps and Options  Configure product steps with ease and define options effortlessly. Explore the power of dependencies to create dynamic and intuitive configurations.  ### Matrices  Delve into matrices—your secret weapon. Explore price and weight matrices, and learn how configured steps influence properties and pricing.  ### Price Management  Unleash dynamic pricing with our versatile price matrices. From fixed prices to incremental structures, adapt to diverse pricing models effortlessly.  ## Security  Your data is in safe hands. Discover how Product Configurator ensures security through JWT authentication, safeguarding your sensitive information.  ## Backward Compatibility  Stay ahead of the curve. Learn about our versioning strategy, providing backward compatibility while allowing our service to evolve seamlessly.  ## Developer Support  Have questions? Need assistance? Write to us at [info@gemini-commerce.com](mailto:info@gemini-commerce.com) and we will get back to you.

API version: v1
Contact: info@gemini-commerce.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package product_configurator

import (
	"encoding/json"
)

// checks if the ProtobufAny type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProtobufAny{}

// ProtobufAny struct for ProtobufAny
type ProtobufAny struct {
	Type *string `json:"@type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProtobufAny ProtobufAny

// NewProtobufAny instantiates a new ProtobufAny object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProtobufAny() *ProtobufAny {
	this := ProtobufAny{}
	return &this
}

// NewProtobufAnyWithDefaults instantiates a new ProtobufAny object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProtobufAnyWithDefaults() *ProtobufAny {
	this := ProtobufAny{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ProtobufAny) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProtobufAny) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ProtobufAny) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ProtobufAny) SetType(v string) {
	o.Type = &v
}

func (o ProtobufAny) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProtobufAny) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["@type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProtobufAny) UnmarshalJSON(bytes []byte) (err error) {
	varProtobufAny := _ProtobufAny{}

	err = json.Unmarshal(bytes, &varProtobufAny)

	if err != nil {
		return err
	}

	*o = ProtobufAny(varProtobufAny)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "@type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProtobufAny struct {
	value *ProtobufAny
	isSet bool
}

func (v NullableProtobufAny) Get() *ProtobufAny {
	return v.value
}

func (v *NullableProtobufAny) Set(val *ProtobufAny) {
	v.value = val
	v.isSet = true
}

func (v NullableProtobufAny) IsSet() bool {
	return v.isSet
}

func (v *NullableProtobufAny) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProtobufAny(val *ProtobufAny) *NullableProtobufAny {
	return &NullableProtobufAny{value: val, isSet: true}
}

func (v NullableProtobufAny) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProtobufAny) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


