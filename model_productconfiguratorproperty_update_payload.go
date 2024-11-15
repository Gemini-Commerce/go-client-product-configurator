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

// checks if the ProductconfiguratorpropertyUpdatePayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductconfiguratorpropertyUpdatePayload{}

// ProductconfiguratorpropertyUpdatePayload struct for ProductconfiguratorpropertyUpdatePayload
type ProductconfiguratorpropertyUpdatePayload struct {
	GenericProperty *PropertyUpdatePayloadGenericProperty `json:"genericProperty,omitempty"`
	WeightProperty *PropertyUpdatePayloadWeightProperty `json:"weightProperty,omitempty"`
	PriceProperty *PropertyUpdatePayloadPriceProperty `json:"priceProperty,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductconfiguratorpropertyUpdatePayload ProductconfiguratorpropertyUpdatePayload

// NewProductconfiguratorpropertyUpdatePayload instantiates a new ProductconfiguratorpropertyUpdatePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductconfiguratorpropertyUpdatePayload() *ProductconfiguratorpropertyUpdatePayload {
	this := ProductconfiguratorpropertyUpdatePayload{}
	return &this
}

// NewProductconfiguratorpropertyUpdatePayloadWithDefaults instantiates a new ProductconfiguratorpropertyUpdatePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductconfiguratorpropertyUpdatePayloadWithDefaults() *ProductconfiguratorpropertyUpdatePayload {
	this := ProductconfiguratorpropertyUpdatePayload{}
	return &this
}

// GetGenericProperty returns the GenericProperty field value if set, zero value otherwise.
func (o *ProductconfiguratorpropertyUpdatePayload) GetGenericProperty() PropertyUpdatePayloadGenericProperty {
	if o == nil || IsNil(o.GenericProperty) {
		var ret PropertyUpdatePayloadGenericProperty
		return ret
	}
	return *o.GenericProperty
}

// GetGenericPropertyOk returns a tuple with the GenericProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorpropertyUpdatePayload) GetGenericPropertyOk() (*PropertyUpdatePayloadGenericProperty, bool) {
	if o == nil || IsNil(o.GenericProperty) {
		return nil, false
	}
	return o.GenericProperty, true
}

// HasGenericProperty returns a boolean if a field has been set.
func (o *ProductconfiguratorpropertyUpdatePayload) HasGenericProperty() bool {
	if o != nil && !IsNil(o.GenericProperty) {
		return true
	}

	return false
}

// SetGenericProperty gets a reference to the given PropertyUpdatePayloadGenericProperty and assigns it to the GenericProperty field.
func (o *ProductconfiguratorpropertyUpdatePayload) SetGenericProperty(v PropertyUpdatePayloadGenericProperty) {
	o.GenericProperty = &v
}

// GetWeightProperty returns the WeightProperty field value if set, zero value otherwise.
func (o *ProductconfiguratorpropertyUpdatePayload) GetWeightProperty() PropertyUpdatePayloadWeightProperty {
	if o == nil || IsNil(o.WeightProperty) {
		var ret PropertyUpdatePayloadWeightProperty
		return ret
	}
	return *o.WeightProperty
}

// GetWeightPropertyOk returns a tuple with the WeightProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorpropertyUpdatePayload) GetWeightPropertyOk() (*PropertyUpdatePayloadWeightProperty, bool) {
	if o == nil || IsNil(o.WeightProperty) {
		return nil, false
	}
	return o.WeightProperty, true
}

// HasWeightProperty returns a boolean if a field has been set.
func (o *ProductconfiguratorpropertyUpdatePayload) HasWeightProperty() bool {
	if o != nil && !IsNil(o.WeightProperty) {
		return true
	}

	return false
}

// SetWeightProperty gets a reference to the given PropertyUpdatePayloadWeightProperty and assigns it to the WeightProperty field.
func (o *ProductconfiguratorpropertyUpdatePayload) SetWeightProperty(v PropertyUpdatePayloadWeightProperty) {
	o.WeightProperty = &v
}

// GetPriceProperty returns the PriceProperty field value if set, zero value otherwise.
func (o *ProductconfiguratorpropertyUpdatePayload) GetPriceProperty() PropertyUpdatePayloadPriceProperty {
	if o == nil || IsNil(o.PriceProperty) {
		var ret PropertyUpdatePayloadPriceProperty
		return ret
	}
	return *o.PriceProperty
}

// GetPricePropertyOk returns a tuple with the PriceProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorpropertyUpdatePayload) GetPricePropertyOk() (*PropertyUpdatePayloadPriceProperty, bool) {
	if o == nil || IsNil(o.PriceProperty) {
		return nil, false
	}
	return o.PriceProperty, true
}

// HasPriceProperty returns a boolean if a field has been set.
func (o *ProductconfiguratorpropertyUpdatePayload) HasPriceProperty() bool {
	if o != nil && !IsNil(o.PriceProperty) {
		return true
	}

	return false
}

// SetPriceProperty gets a reference to the given PropertyUpdatePayloadPriceProperty and assigns it to the PriceProperty field.
func (o *ProductconfiguratorpropertyUpdatePayload) SetPriceProperty(v PropertyUpdatePayloadPriceProperty) {
	o.PriceProperty = &v
}

func (o ProductconfiguratorpropertyUpdatePayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductconfiguratorpropertyUpdatePayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GenericProperty) {
		toSerialize["genericProperty"] = o.GenericProperty
	}
	if !IsNil(o.WeightProperty) {
		toSerialize["weightProperty"] = o.WeightProperty
	}
	if !IsNil(o.PriceProperty) {
		toSerialize["priceProperty"] = o.PriceProperty
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductconfiguratorpropertyUpdatePayload) UnmarshalJSON(data []byte) (err error) {
	varProductconfiguratorpropertyUpdatePayload := _ProductconfiguratorpropertyUpdatePayload{}

	err = json.Unmarshal(data, &varProductconfiguratorpropertyUpdatePayload)

	if err != nil {
		return err
	}

	*o = ProductconfiguratorpropertyUpdatePayload(varProductconfiguratorpropertyUpdatePayload)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "genericProperty")
		delete(additionalProperties, "weightProperty")
		delete(additionalProperties, "priceProperty")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ProductconfiguratorpropertyUpdatePayload) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *ProductconfiguratorpropertyUpdatePayload) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableProductconfiguratorpropertyUpdatePayload struct {
	value *ProductconfiguratorpropertyUpdatePayload
	isSet bool
}

func (v NullableProductconfiguratorpropertyUpdatePayload) Get() *ProductconfiguratorpropertyUpdatePayload {
	return v.value
}

func (v *NullableProductconfiguratorpropertyUpdatePayload) Set(val *ProductconfiguratorpropertyUpdatePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableProductconfiguratorpropertyUpdatePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableProductconfiguratorpropertyUpdatePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductconfiguratorpropertyUpdatePayload(val *ProductconfiguratorpropertyUpdatePayload) *NullableProductconfiguratorpropertyUpdatePayload {
	return &NullableProductconfiguratorpropertyUpdatePayload{value: val, isSet: true}
}

func (v NullableProductconfiguratorpropertyUpdatePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductconfiguratorpropertyUpdatePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


