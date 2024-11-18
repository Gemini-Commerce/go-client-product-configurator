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

// checks if the ProductConfiguratorBulkCreateOptionsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductConfiguratorBulkCreateOptionsRequest{}

// ProductConfiguratorBulkCreateOptionsRequest struct for ProductConfiguratorBulkCreateOptionsRequest
type ProductConfiguratorBulkCreateOptionsRequest struct {
	Options              []ProductconfiguratoroptionBulkCreateRequestCreateEntity `json:"options,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductConfiguratorBulkCreateOptionsRequest ProductConfiguratorBulkCreateOptionsRequest

// NewProductConfiguratorBulkCreateOptionsRequest instantiates a new ProductConfiguratorBulkCreateOptionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductConfiguratorBulkCreateOptionsRequest() *ProductConfiguratorBulkCreateOptionsRequest {
	this := ProductConfiguratorBulkCreateOptionsRequest{}
	return &this
}

// NewProductConfiguratorBulkCreateOptionsRequestWithDefaults instantiates a new ProductConfiguratorBulkCreateOptionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductConfiguratorBulkCreateOptionsRequestWithDefaults() *ProductConfiguratorBulkCreateOptionsRequest {
	this := ProductConfiguratorBulkCreateOptionsRequest{}
	return &this
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *ProductConfiguratorBulkCreateOptionsRequest) GetOptions() []ProductconfiguratoroptionBulkCreateRequestCreateEntity {
	if o == nil || IsNil(o.Options) {
		var ret []ProductconfiguratoroptionBulkCreateRequestCreateEntity
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductConfiguratorBulkCreateOptionsRequest) GetOptionsOk() ([]ProductconfiguratoroptionBulkCreateRequestCreateEntity, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ProductConfiguratorBulkCreateOptionsRequest) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []ProductconfiguratoroptionBulkCreateRequestCreateEntity and assigns it to the Options field.
func (o *ProductConfiguratorBulkCreateOptionsRequest) SetOptions(v []ProductconfiguratoroptionBulkCreateRequestCreateEntity) {
	o.Options = v
}

func (o ProductConfiguratorBulkCreateOptionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductConfiguratorBulkCreateOptionsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductConfiguratorBulkCreateOptionsRequest) UnmarshalJSON(data []byte) (err error) {
	varProductConfiguratorBulkCreateOptionsRequest := _ProductConfiguratorBulkCreateOptionsRequest{}

	err = json.Unmarshal(data, &varProductConfiguratorBulkCreateOptionsRequest)

	if err != nil {
		return err
	}

	*o = ProductConfiguratorBulkCreateOptionsRequest(varProductConfiguratorBulkCreateOptionsRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "options")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of wellknown types
func (o *ProductConfiguratorBulkCreateOptionsRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of wellknown types
func (o *ProductConfiguratorBulkCreateOptionsRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableProductConfiguratorBulkCreateOptionsRequest struct {
	value *ProductConfiguratorBulkCreateOptionsRequest
	isSet bool
}

func (v NullableProductConfiguratorBulkCreateOptionsRequest) Get() *ProductConfiguratorBulkCreateOptionsRequest {
	return v.value
}

func (v *NullableProductConfiguratorBulkCreateOptionsRequest) Set(val *ProductConfiguratorBulkCreateOptionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductConfiguratorBulkCreateOptionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductConfiguratorBulkCreateOptionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductConfiguratorBulkCreateOptionsRequest(val *ProductConfiguratorBulkCreateOptionsRequest) *NullableProductConfiguratorBulkCreateOptionsRequest {
	return &NullableProductConfiguratorBulkCreateOptionsRequest{value: val, isSet: true}
}

func (v NullableProductConfiguratorBulkCreateOptionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductConfiguratorBulkCreateOptionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
