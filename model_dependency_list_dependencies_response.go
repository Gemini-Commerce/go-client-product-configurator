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

// checks if the DependencyListDependenciesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DependencyListDependenciesResponse{}

// DependencyListDependenciesResponse struct for DependencyListDependenciesResponse
type DependencyListDependenciesResponse struct {
	Dependencies         []ProductconfiguratordependencyEntity `json:"dependencies,omitempty"`
	NextPageToken        *string                               `json:"nextPageToken,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DependencyListDependenciesResponse DependencyListDependenciesResponse

// NewDependencyListDependenciesResponse instantiates a new DependencyListDependenciesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDependencyListDependenciesResponse() *DependencyListDependenciesResponse {
	this := DependencyListDependenciesResponse{}
	return &this
}

// NewDependencyListDependenciesResponseWithDefaults instantiates a new DependencyListDependenciesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDependencyListDependenciesResponseWithDefaults() *DependencyListDependenciesResponse {
	this := DependencyListDependenciesResponse{}
	return &this
}

// GetDependencies returns the Dependencies field value if set, zero value otherwise.
func (o *DependencyListDependenciesResponse) GetDependencies() []ProductconfiguratordependencyEntity {
	if o == nil || IsNil(o.Dependencies) {
		var ret []ProductconfiguratordependencyEntity
		return ret
	}
	return o.Dependencies
}

// GetDependenciesOk returns a tuple with the Dependencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DependencyListDependenciesResponse) GetDependenciesOk() ([]ProductconfiguratordependencyEntity, bool) {
	if o == nil || IsNil(o.Dependencies) {
		return nil, false
	}
	return o.Dependencies, true
}

// HasDependencies returns a boolean if a field has been set.
func (o *DependencyListDependenciesResponse) HasDependencies() bool {
	if o != nil && !IsNil(o.Dependencies) {
		return true
	}

	return false
}

// SetDependencies gets a reference to the given []ProductconfiguratordependencyEntity and assigns it to the Dependencies field.
func (o *DependencyListDependenciesResponse) SetDependencies(v []ProductconfiguratordependencyEntity) {
	o.Dependencies = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *DependencyListDependenciesResponse) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DependencyListDependenciesResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *DependencyListDependenciesResponse) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *DependencyListDependenciesResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o DependencyListDependenciesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DependencyListDependenciesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Dependencies) {
		toSerialize["dependencies"] = o.Dependencies
	}
	if !IsNil(o.NextPageToken) {
		toSerialize["nextPageToken"] = o.NextPageToken
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DependencyListDependenciesResponse) UnmarshalJSON(data []byte) (err error) {
	varDependencyListDependenciesResponse := _DependencyListDependenciesResponse{}

	err = json.Unmarshal(data, &varDependencyListDependenciesResponse)

	if err != nil {
		return err
	}

	*o = DependencyListDependenciesResponse(varDependencyListDependenciesResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "dependencies")
		delete(additionalProperties, "nextPageToken")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *DependencyListDependenciesResponse) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *DependencyListDependenciesResponse) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableDependencyListDependenciesResponse struct {
	value *DependencyListDependenciesResponse
	isSet bool
}

func (v NullableDependencyListDependenciesResponse) Get() *DependencyListDependenciesResponse {
	return v.value
}

func (v *NullableDependencyListDependenciesResponse) Set(val *DependencyListDependenciesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDependencyListDependenciesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDependencyListDependenciesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDependencyListDependenciesResponse(val *DependencyListDependenciesResponse) *NullableDependencyListDependenciesResponse {
	return &NullableDependencyListDependenciesResponse{value: val, isSet: true}
}

func (v NullableDependencyListDependenciesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDependencyListDependenciesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
