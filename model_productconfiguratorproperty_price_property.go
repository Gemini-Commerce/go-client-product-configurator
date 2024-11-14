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

// checks if the ProductconfiguratorpropertyPriceProperty type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductconfiguratorpropertyPriceProperty{}

// ProductconfiguratorpropertyPriceProperty struct for ProductconfiguratorpropertyPriceProperty
type ProductconfiguratorpropertyPriceProperty struct {
	Price *ProductconfiguratorMoney `json:"price,omitempty"`
	PricelistGrn *string `json:"pricelistGrn,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductconfiguratorpropertyPriceProperty ProductconfiguratorpropertyPriceProperty

// NewProductconfiguratorpropertyPriceProperty instantiates a new ProductconfiguratorpropertyPriceProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductconfiguratorpropertyPriceProperty() *ProductconfiguratorpropertyPriceProperty {
	this := ProductconfiguratorpropertyPriceProperty{}
	return &this
}

// NewProductconfiguratorpropertyPricePropertyWithDefaults instantiates a new ProductconfiguratorpropertyPriceProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductconfiguratorpropertyPricePropertyWithDefaults() *ProductconfiguratorpropertyPriceProperty {
	this := ProductconfiguratorpropertyPriceProperty{}
	return &this
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *ProductconfiguratorpropertyPriceProperty) GetPrice() ProductconfiguratorMoney {
	if o == nil || IsNil(o.Price) {
		var ret ProductconfiguratorMoney
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorpropertyPriceProperty) GetPriceOk() (*ProductconfiguratorMoney, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// &#39;Has&#39;Price returns a boolean if a field has been set.
func (o *ProductconfiguratorpropertyPriceProperty) &#39;Has&#39;Price() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given ProductconfiguratorMoney and assigns it to the Price field.
func (o *ProductconfiguratorpropertyPriceProperty) SetPrice(v ProductconfiguratorMoney) {
	o.Price = &v
}

// GetPricelistGrn returns the PricelistGrn field value if set, zero value otherwise.
func (o *ProductconfiguratorpropertyPriceProperty) GetPricelistGrn() string {
	if o == nil || IsNil(o.PricelistGrn) {
		var ret string
		return ret
	}
	return *o.PricelistGrn
}

// GetPricelistGrnOk returns a tuple with the PricelistGrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorpropertyPriceProperty) GetPricelistGrnOk() (*string, bool) {
	if o == nil || IsNil(o.PricelistGrn) {
		return nil, false
	}
	return o.PricelistGrn, true
}

// &#39;Has&#39;PricelistGrn returns a boolean if a field has been set.
func (o *ProductconfiguratorpropertyPriceProperty) &#39;Has&#39;PricelistGrn() bool {
	if o != nil && !IsNil(o.PricelistGrn) {
		return true
	}

	return false
}

// SetPricelistGrn gets a reference to the given string and assigns it to the PricelistGrn field.
func (o *ProductconfiguratorpropertyPriceProperty) SetPricelistGrn(v string) {
	o.PricelistGrn = &v
}

func (o ProductconfiguratorpropertyPriceProperty) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductconfiguratorpropertyPriceProperty) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.PricelistGrn) {
		toSerialize["pricelistGrn"] = o.PricelistGrn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductconfiguratorpropertyPriceProperty) UnmarshalJSON(data []byte) (err error) {
	varProductconfiguratorpropertyPriceProperty := _ProductconfiguratorpropertyPriceProperty{}

	err = json.Unmarshal(data, &varProductconfiguratorpropertyPriceProperty)

	if err != nil {
		return err
	}

	*o = ProductconfiguratorpropertyPriceProperty(varProductconfiguratorpropertyPriceProperty)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "price")
		delete(additionalProperties, "pricelistGrn")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ProductconfiguratorpropertyPriceProperty) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *ProductconfiguratorpropertyPriceProperty) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableProductconfiguratorpropertyPriceProperty struct {
	value *ProductconfiguratorpropertyPriceProperty
	isSet bool
}

func (v NullableProductconfiguratorpropertyPriceProperty) Get() *ProductconfiguratorpropertyPriceProperty {
	return v.value
}

func (v *NullableProductconfiguratorpropertyPriceProperty) Set(val *ProductconfiguratorpropertyPriceProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableProductconfiguratorpropertyPriceProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableProductconfiguratorpropertyPriceProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductconfiguratorpropertyPriceProperty(val *ProductconfiguratorpropertyPriceProperty) *NullableProductconfiguratorpropertyPriceProperty {
	return &NullableProductconfiguratorpropertyPriceProperty{value: val, isSet: true}
}

func (v NullableProductconfiguratorpropertyPriceProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductconfiguratorpropertyPriceProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


