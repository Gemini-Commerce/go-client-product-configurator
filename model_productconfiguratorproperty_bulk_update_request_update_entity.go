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

// checks if the ProductconfiguratorpropertyBulkUpdateRequestUpdateEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductconfiguratorpropertyBulkUpdateRequestUpdateEntity{}

// ProductconfiguratorpropertyBulkUpdateRequestUpdateEntity struct for ProductconfiguratorpropertyBulkUpdateRequestUpdateEntity
type ProductconfiguratorpropertyBulkUpdateRequestUpdateEntity struct {
	PropertyId           *string                                   `json:"propertyId,omitempty"`
	Payload              *ProductconfiguratorpropertyUpdatePayload `json:"payload,omitempty"`
	PayloadMask          *string                                   `json:"payloadMask,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductconfiguratorpropertyBulkUpdateRequestUpdateEntity ProductconfiguratorpropertyBulkUpdateRequestUpdateEntity

// NewProductconfiguratorpropertyBulkUpdateRequestUpdateEntity instantiates a new ProductconfiguratorpropertyBulkUpdateRequestUpdateEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductconfiguratorpropertyBulkUpdateRequestUpdateEntity() *ProductconfiguratorpropertyBulkUpdateRequestUpdateEntity {
	this := ProductconfiguratorpropertyBulkUpdateRequestUpdateEntity{}
	return &this
}

// NewProductconfiguratorpropertyBulkUpdateRequestUpdateEntityWithDefaults instantiates a new ProductconfiguratorpropertyBulkUpdateRequestUpdateEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductconfiguratorpropertyBulkUpdateRequestUpdateEntityWithDefaults() *ProductconfiguratorpropertyBulkUpdateRequestUpdateEntity {
	this := ProductconfiguratorpropertyBulkUpdateRequestUpdateEntity{}
	return &this
}

// GetPropertyId returns the PropertyId field value if set, zero value otherwise.
func (o *ProductconfiguratorpropertyBulkUpdateRequestUpdateEntity) GetPropertyId() string {
	if o == nil || IsNil(o.PropertyId) {
		var ret string
		return ret
	}
	return *o.PropertyId
}

// GetPropertyIdOk returns a tuple with the PropertyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorpropertyBulkUpdateRequestUpdateEntity) GetPropertyIdOk() (*string, bool) {
	if o == nil || IsNil(o.PropertyId) {
		return nil, false
	}
	return o.PropertyId, true
}

// HasPropertyId returns a boolean if a field has been set.
func (o *ProductconfiguratorpropertyBulkUpdateRequestUpdateEntity) HasPropertyId() bool {
	if o != nil && !IsNil(o.PropertyId) {
		return true
	}

	return false
}

// SetPropertyId gets a reference to the given string and assigns it to the PropertyId field.
func (o *ProductconfiguratorpropertyBulkUpdateRequestUpdateEntity) SetPropertyId(v string) {
	o.PropertyId = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *ProductconfiguratorpropertyBulkUpdateRequestUpdateEntity) GetPayload() ProductconfiguratorpropertyUpdatePayload {
	if o == nil || IsNil(o.Payload) {
		var ret ProductconfiguratorpropertyUpdatePayload
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorpropertyBulkUpdateRequestUpdateEntity) GetPayloadOk() (*ProductconfiguratorpropertyUpdatePayload, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *ProductconfiguratorpropertyBulkUpdateRequestUpdateEntity) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given ProductconfiguratorpropertyUpdatePayload and assigns it to the Payload field.
func (o *ProductconfiguratorpropertyBulkUpdateRequestUpdateEntity) SetPayload(v ProductconfiguratorpropertyUpdatePayload) {
	o.Payload = &v
}

// GetPayloadMask returns the PayloadMask field value if set, zero value otherwise.
func (o *ProductconfiguratorpropertyBulkUpdateRequestUpdateEntity) GetPayloadMask() string {
	if o == nil || IsNil(o.PayloadMask) {
		var ret string
		return ret
	}
	return *o.PayloadMask
}

// GetPayloadMaskOk returns a tuple with the PayloadMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorpropertyBulkUpdateRequestUpdateEntity) GetPayloadMaskOk() (*string, bool) {
	if o == nil || IsNil(o.PayloadMask) {
		return nil, false
	}
	return o.PayloadMask, true
}

// HasPayloadMask returns a boolean if a field has been set.
func (o *ProductconfiguratorpropertyBulkUpdateRequestUpdateEntity) HasPayloadMask() bool {
	if o != nil && !IsNil(o.PayloadMask) {
		return true
	}

	return false
}

// SetPayloadMask gets a reference to the given string and assigns it to the PayloadMask field.
func (o *ProductconfiguratorpropertyBulkUpdateRequestUpdateEntity) SetPayloadMask(v string) {
	o.PayloadMask = &v
}

func (o ProductconfiguratorpropertyBulkUpdateRequestUpdateEntity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductconfiguratorpropertyBulkUpdateRequestUpdateEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PropertyId) {
		toSerialize["propertyId"] = o.PropertyId
	}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.PayloadMask) {
		toSerialize["payloadMask"] = o.PayloadMask
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductconfiguratorpropertyBulkUpdateRequestUpdateEntity) UnmarshalJSON(data []byte) (err error) {
	varProductconfiguratorpropertyBulkUpdateRequestUpdateEntity := _ProductconfiguratorpropertyBulkUpdateRequestUpdateEntity{}

	err = json.Unmarshal(data, &varProductconfiguratorpropertyBulkUpdateRequestUpdateEntity)

	if err != nil {
		return err
	}

	*o = ProductconfiguratorpropertyBulkUpdateRequestUpdateEntity(varProductconfiguratorpropertyBulkUpdateRequestUpdateEntity)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "propertyId")
		delete(additionalProperties, "payload")
		delete(additionalProperties, "payloadMask")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well_known types
func (o *ProductconfiguratorpropertyBulkUpdateRequestUpdateEntity) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well_known types
func (o *ProductconfiguratorpropertyBulkUpdateRequestUpdateEntity) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableProductconfiguratorpropertyBulkUpdateRequestUpdateEntity struct {
	value *ProductconfiguratorpropertyBulkUpdateRequestUpdateEntity
	isSet bool
}

func (v NullableProductconfiguratorpropertyBulkUpdateRequestUpdateEntity) Get() *ProductconfiguratorpropertyBulkUpdateRequestUpdateEntity {
	return v.value
}

func (v *NullableProductconfiguratorpropertyBulkUpdateRequestUpdateEntity) Set(val *ProductconfiguratorpropertyBulkUpdateRequestUpdateEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableProductconfiguratorpropertyBulkUpdateRequestUpdateEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableProductconfiguratorpropertyBulkUpdateRequestUpdateEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductconfiguratorpropertyBulkUpdateRequestUpdateEntity(val *ProductconfiguratorpropertyBulkUpdateRequestUpdateEntity) *NullableProductconfiguratorpropertyBulkUpdateRequestUpdateEntity {
	return &NullableProductconfiguratorpropertyBulkUpdateRequestUpdateEntity{value: val, isSet: true}
}

func (v NullableProductconfiguratorpropertyBulkUpdateRequestUpdateEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductconfiguratorpropertyBulkUpdateRequestUpdateEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
