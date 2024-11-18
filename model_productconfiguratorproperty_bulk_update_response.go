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

// checks if the ProductconfiguratorpropertyBulkUpdateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductconfiguratorpropertyBulkUpdateResponse{}

// ProductconfiguratorpropertyBulkUpdateResponse struct for ProductconfiguratorpropertyBulkUpdateResponse
type ProductconfiguratorpropertyBulkUpdateResponse struct {
	Properties           []ProductconfiguratorpropertyEntity `json:"properties,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductconfiguratorpropertyBulkUpdateResponse ProductconfiguratorpropertyBulkUpdateResponse

// NewProductconfiguratorpropertyBulkUpdateResponse instantiates a new ProductconfiguratorpropertyBulkUpdateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductconfiguratorpropertyBulkUpdateResponse() *ProductconfiguratorpropertyBulkUpdateResponse {
	this := ProductconfiguratorpropertyBulkUpdateResponse{}
	return &this
}

// NewProductconfiguratorpropertyBulkUpdateResponseWithDefaults instantiates a new ProductconfiguratorpropertyBulkUpdateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductconfiguratorpropertyBulkUpdateResponseWithDefaults() *ProductconfiguratorpropertyBulkUpdateResponse {
	this := ProductconfiguratorpropertyBulkUpdateResponse{}
	return &this
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *ProductconfiguratorpropertyBulkUpdateResponse) GetProperties() []ProductconfiguratorpropertyEntity {
	if o == nil || IsNil(o.Properties) {
		var ret []ProductconfiguratorpropertyEntity
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorpropertyBulkUpdateResponse) GetPropertiesOk() ([]ProductconfiguratorpropertyEntity, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *ProductconfiguratorpropertyBulkUpdateResponse) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []ProductconfiguratorpropertyEntity and assigns it to the Properties field.
func (o *ProductconfiguratorpropertyBulkUpdateResponse) SetProperties(v []ProductconfiguratorpropertyEntity) {
	o.Properties = v
}

func (o ProductconfiguratorpropertyBulkUpdateResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductconfiguratorpropertyBulkUpdateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductconfiguratorpropertyBulkUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	varProductconfiguratorpropertyBulkUpdateResponse := _ProductconfiguratorpropertyBulkUpdateResponse{}

	err = json.Unmarshal(data, &varProductconfiguratorpropertyBulkUpdateResponse)

	if err != nil {
		return err
	}

	*o = ProductconfiguratorpropertyBulkUpdateResponse(varProductconfiguratorpropertyBulkUpdateResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "properties")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of wellknown types
func (o *ProductconfiguratorpropertyBulkUpdateResponse) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of wellknown types
func (o *ProductconfiguratorpropertyBulkUpdateResponse) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableProductconfiguratorpropertyBulkUpdateResponse struct {
	value *ProductconfiguratorpropertyBulkUpdateResponse
	isSet bool
}

func (v NullableProductconfiguratorpropertyBulkUpdateResponse) Get() *ProductconfiguratorpropertyBulkUpdateResponse {
	return v.value
}

func (v *NullableProductconfiguratorpropertyBulkUpdateResponse) Set(val *ProductconfiguratorpropertyBulkUpdateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProductconfiguratorpropertyBulkUpdateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProductconfiguratorpropertyBulkUpdateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductconfiguratorpropertyBulkUpdateResponse(val *ProductconfiguratorpropertyBulkUpdateResponse) *NullableProductconfiguratorpropertyBulkUpdateResponse {
	return &NullableProductconfiguratorpropertyBulkUpdateResponse{value: val, isSet: true}
}

func (v NullableProductconfiguratorpropertyBulkUpdateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductconfiguratorpropertyBulkUpdateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
