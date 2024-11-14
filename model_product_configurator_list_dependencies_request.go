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

// checks if the ProductConfiguratorListDependenciesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductConfiguratorListDependenciesRequest{}

// ProductConfiguratorListDependenciesRequest struct for ProductConfiguratorListDependenciesRequest
type ProductConfiguratorListDependenciesRequest struct {
	StepIds []string `json:"stepIds,omitempty"`
	PageToken *string `json:"pageToken,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductConfiguratorListDependenciesRequest ProductConfiguratorListDependenciesRequest

// NewProductConfiguratorListDependenciesRequest instantiates a new ProductConfiguratorListDependenciesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductConfiguratorListDependenciesRequest() *ProductConfiguratorListDependenciesRequest {
	this := ProductConfiguratorListDependenciesRequest{}
	return &this
}

// NewProductConfiguratorListDependenciesRequestWithDefaults instantiates a new ProductConfiguratorListDependenciesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductConfiguratorListDependenciesRequestWithDefaults() *ProductConfiguratorListDependenciesRequest {
	this := ProductConfiguratorListDependenciesRequest{}
	return &this
}

// GetStepIds returns the StepIds field value if set, zero value otherwise.
func (o *ProductConfiguratorListDependenciesRequest) GetStepIds() []string {
	if o == nil || IsNil(o.StepIds) {
		var ret []string
		return ret
	}
	return o.StepIds
}

// GetStepIdsOk returns a tuple with the StepIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductConfiguratorListDependenciesRequest) GetStepIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.StepIds) {
		return nil, false
	}
	return o.StepIds, true
}

// &#39;Has&#39;StepIds returns a boolean if a field has been set.
func (o *ProductConfiguratorListDependenciesRequest) &#39;Has&#39;StepIds() bool {
	if o != nil && !IsNil(o.StepIds) {
		return true
	}

	return false
}

// SetStepIds gets a reference to the given []string and assigns it to the StepIds field.
func (o *ProductConfiguratorListDependenciesRequest) SetStepIds(v []string) {
	o.StepIds = v
}

// GetPageToken returns the PageToken field value if set, zero value otherwise.
func (o *ProductConfiguratorListDependenciesRequest) GetPageToken() string {
	if o == nil || IsNil(o.PageToken) {
		var ret string
		return ret
	}
	return *o.PageToken
}

// GetPageTokenOk returns a tuple with the PageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductConfiguratorListDependenciesRequest) GetPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.PageToken) {
		return nil, false
	}
	return o.PageToken, true
}

// &#39;Has&#39;PageToken returns a boolean if a field has been set.
func (o *ProductConfiguratorListDependenciesRequest) &#39;Has&#39;PageToken() bool {
	if o != nil && !IsNil(o.PageToken) {
		return true
	}

	return false
}

// SetPageToken gets a reference to the given string and assigns it to the PageToken field.
func (o *ProductConfiguratorListDependenciesRequest) SetPageToken(v string) {
	o.PageToken = &v
}

func (o ProductConfiguratorListDependenciesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductConfiguratorListDependenciesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StepIds) {
		toSerialize["stepIds"] = o.StepIds
	}
	if !IsNil(o.PageToken) {
		toSerialize["pageToken"] = o.PageToken
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductConfiguratorListDependenciesRequest) UnmarshalJSON(data []byte) (err error) {
	varProductConfiguratorListDependenciesRequest := _ProductConfiguratorListDependenciesRequest{}

	err = json.Unmarshal(data, &varProductConfiguratorListDependenciesRequest)

	if err != nil {
		return err
	}

	*o = ProductConfiguratorListDependenciesRequest(varProductConfiguratorListDependenciesRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "stepIds")
		delete(additionalProperties, "pageToken")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ProductConfiguratorListDependenciesRequest) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *ProductConfiguratorListDependenciesRequest) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableProductConfiguratorListDependenciesRequest struct {
	value *ProductConfiguratorListDependenciesRequest
	isSet bool
}

func (v NullableProductConfiguratorListDependenciesRequest) Get() *ProductConfiguratorListDependenciesRequest {
	return v.value
}

func (v *NullableProductConfiguratorListDependenciesRequest) Set(val *ProductConfiguratorListDependenciesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductConfiguratorListDependenciesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductConfiguratorListDependenciesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductConfiguratorListDependenciesRequest(val *ProductConfiguratorListDependenciesRequest) *NullableProductConfiguratorListDependenciesRequest {
	return &NullableProductConfiguratorListDependenciesRequest{value: val, isSet: true}
}

func (v NullableProductConfiguratorListDependenciesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductConfiguratorListDependenciesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


