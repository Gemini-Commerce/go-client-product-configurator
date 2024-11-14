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

// checks if the ProductConfiguratorUpdateStepRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductConfiguratorUpdateStepRequest{}

// ProductConfiguratorUpdateStepRequest struct for ProductConfiguratorUpdateStepRequest
type ProductConfiguratorUpdateStepRequest struct {
	Payload *ProductconfiguratorstepUpdatePayload `json:"payload,omitempty"`
	PayloadMask *string `json:"payloadMask,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductConfiguratorUpdateStepRequest ProductConfiguratorUpdateStepRequest

// NewProductConfiguratorUpdateStepRequest instantiates a new ProductConfiguratorUpdateStepRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductConfiguratorUpdateStepRequest() *ProductConfiguratorUpdateStepRequest {
	this := ProductConfiguratorUpdateStepRequest{}
	return &this
}

// NewProductConfiguratorUpdateStepRequestWithDefaults instantiates a new ProductConfiguratorUpdateStepRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductConfiguratorUpdateStepRequestWithDefaults() *ProductConfiguratorUpdateStepRequest {
	this := ProductConfiguratorUpdateStepRequest{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *ProductConfiguratorUpdateStepRequest) GetPayload() ProductconfiguratorstepUpdatePayload {
	if o == nil || IsNil(o.Payload) {
		var ret ProductconfiguratorstepUpdatePayload
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductConfiguratorUpdateStepRequest) GetPayloadOk() (*ProductconfiguratorstepUpdatePayload, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// &#39;Has&#39;Payload returns a boolean if a field has been set.
func (o *ProductConfiguratorUpdateStepRequest) &#39;Has&#39;Payload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given ProductconfiguratorstepUpdatePayload and assigns it to the Payload field.
func (o *ProductConfiguratorUpdateStepRequest) SetPayload(v ProductconfiguratorstepUpdatePayload) {
	o.Payload = &v
}

// GetPayloadMask returns the PayloadMask field value if set, zero value otherwise.
func (o *ProductConfiguratorUpdateStepRequest) GetPayloadMask() string {
	if o == nil || IsNil(o.PayloadMask) {
		var ret string
		return ret
	}
	return *o.PayloadMask
}

// GetPayloadMaskOk returns a tuple with the PayloadMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductConfiguratorUpdateStepRequest) GetPayloadMaskOk() (*string, bool) {
	if o == nil || IsNil(o.PayloadMask) {
		return nil, false
	}
	return o.PayloadMask, true
}

// &#39;Has&#39;PayloadMask returns a boolean if a field has been set.
func (o *ProductConfiguratorUpdateStepRequest) &#39;Has&#39;PayloadMask() bool {
	if o != nil && !IsNil(o.PayloadMask) {
		return true
	}

	return false
}

// SetPayloadMask gets a reference to the given string and assigns it to the PayloadMask field.
func (o *ProductConfiguratorUpdateStepRequest) SetPayloadMask(v string) {
	o.PayloadMask = &v
}

func (o ProductConfiguratorUpdateStepRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductConfiguratorUpdateStepRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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

func (o *ProductConfiguratorUpdateStepRequest) UnmarshalJSON(data []byte) (err error) {
	varProductConfiguratorUpdateStepRequest := _ProductConfiguratorUpdateStepRequest{}

	err = json.Unmarshal(data, &varProductConfiguratorUpdateStepRequest)

	if err != nil {
		return err
	}

	*o = ProductConfiguratorUpdateStepRequest(varProductConfiguratorUpdateStepRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "payload")
		delete(additionalProperties, "payloadMask")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ProductConfiguratorUpdateStepRequest) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *ProductConfiguratorUpdateStepRequest) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableProductConfiguratorUpdateStepRequest struct {
	value *ProductConfiguratorUpdateStepRequest
	isSet bool
}

func (v NullableProductConfiguratorUpdateStepRequest) Get() *ProductConfiguratorUpdateStepRequest {
	return v.value
}

func (v *NullableProductConfiguratorUpdateStepRequest) Set(val *ProductConfiguratorUpdateStepRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductConfiguratorUpdateStepRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductConfiguratorUpdateStepRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductConfiguratorUpdateStepRequest(val *ProductConfiguratorUpdateStepRequest) *NullableProductConfiguratorUpdateStepRequest {
	return &NullableProductConfiguratorUpdateStepRequest{value: val, isSet: true}
}

func (v NullableProductConfiguratorUpdateStepRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductConfiguratorUpdateStepRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


