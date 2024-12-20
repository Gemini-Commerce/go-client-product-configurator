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

// checks if the MatrixListMatricesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MatrixListMatricesResponse{}

// MatrixListMatricesResponse struct for MatrixListMatricesResponse
type MatrixListMatricesResponse struct {
	Matrices             []ProductconfiguratormatrixEntity `json:"matrices,omitempty"`
	NextPageToken        *string                           `json:"nextPageToken,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MatrixListMatricesResponse MatrixListMatricesResponse

// NewMatrixListMatricesResponse instantiates a new MatrixListMatricesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMatrixListMatricesResponse() *MatrixListMatricesResponse {
	this := MatrixListMatricesResponse{}
	return &this
}

// NewMatrixListMatricesResponseWithDefaults instantiates a new MatrixListMatricesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMatrixListMatricesResponseWithDefaults() *MatrixListMatricesResponse {
	this := MatrixListMatricesResponse{}
	return &this
}

// GetMatrices returns the Matrices field value if set, zero value otherwise.
func (o *MatrixListMatricesResponse) GetMatrices() []ProductconfiguratormatrixEntity {
	if o == nil || IsNil(o.Matrices) {
		var ret []ProductconfiguratormatrixEntity
		return ret
	}
	return o.Matrices
}

// GetMatricesOk returns a tuple with the Matrices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatrixListMatricesResponse) GetMatricesOk() ([]ProductconfiguratormatrixEntity, bool) {
	if o == nil || IsNil(o.Matrices) {
		return nil, false
	}
	return o.Matrices, true
}

// HasMatrices returns a boolean if a field has been set.
func (o *MatrixListMatricesResponse) HasMatrices() bool {
	if o != nil && !IsNil(o.Matrices) {
		return true
	}

	return false
}

// SetMatrices gets a reference to the given []ProductconfiguratormatrixEntity and assigns it to the Matrices field.
func (o *MatrixListMatricesResponse) SetMatrices(v []ProductconfiguratormatrixEntity) {
	o.Matrices = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *MatrixListMatricesResponse) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatrixListMatricesResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *MatrixListMatricesResponse) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *MatrixListMatricesResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o MatrixListMatricesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MatrixListMatricesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Matrices) {
		toSerialize["matrices"] = o.Matrices
	}
	if !IsNil(o.NextPageToken) {
		toSerialize["nextPageToken"] = o.NextPageToken
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MatrixListMatricesResponse) UnmarshalJSON(data []byte) (err error) {
	varMatrixListMatricesResponse := _MatrixListMatricesResponse{}

	err = json.Unmarshal(data, &varMatrixListMatricesResponse)

	if err != nil {
		return err
	}

	*o = MatrixListMatricesResponse(varMatrixListMatricesResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "matrices")
		delete(additionalProperties, "nextPageToken")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *MatrixListMatricesResponse) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *MatrixListMatricesResponse) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableMatrixListMatricesResponse struct {
	value *MatrixListMatricesResponse
	isSet bool
}

func (v NullableMatrixListMatricesResponse) Get() *MatrixListMatricesResponse {
	return v.value
}

func (v *NullableMatrixListMatricesResponse) Set(val *MatrixListMatricesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMatrixListMatricesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMatrixListMatricesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMatrixListMatricesResponse(val *MatrixListMatricesResponse) *NullableMatrixListMatricesResponse {
	return &NullableMatrixListMatricesResponse{value: val, isSet: true}
}

func (v NullableMatrixListMatricesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMatrixListMatricesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
