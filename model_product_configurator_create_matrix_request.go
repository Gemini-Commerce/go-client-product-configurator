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

// checks if the ProductConfiguratorCreateMatrixRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductConfiguratorCreateMatrixRequest{}

// ProductConfiguratorCreateMatrixRequest struct for ProductConfiguratorCreateMatrixRequest
type ProductConfiguratorCreateMatrixRequest struct {
	ConfiguratorId *string `json:"configuratorId,omitempty"`
	Label *string `json:"label,omitempty"`
	GenericType *MatrixGenericType `json:"genericType,omitempty"`
	PriceType *MatrixPriceType `json:"priceType,omitempty"`
	WeightType *MatrixWeightType `json:"weightType,omitempty"`
	Steps []ProductconfiguratormatrixStep `json:"steps,omitempty"`
	PropertiesMode *ProductconfiguratorPropertyMode `json:"propertiesMode,omitempty"`
}

// NewProductConfiguratorCreateMatrixRequest instantiates a new ProductConfiguratorCreateMatrixRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductConfiguratorCreateMatrixRequest() *ProductConfiguratorCreateMatrixRequest {
	this := ProductConfiguratorCreateMatrixRequest{}
	var propertiesMode ProductconfiguratorPropertyMode = UNKNOWN
	this.PropertiesMode = &propertiesMode
	return &this
}

// NewProductConfiguratorCreateMatrixRequestWithDefaults instantiates a new ProductConfiguratorCreateMatrixRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductConfiguratorCreateMatrixRequestWithDefaults() *ProductConfiguratorCreateMatrixRequest {
	this := ProductConfiguratorCreateMatrixRequest{}
	var propertiesMode ProductconfiguratorPropertyMode = UNKNOWN
	this.PropertiesMode = &propertiesMode
	return &this
}

// GetConfiguratorId returns the ConfiguratorId field value if set, zero value otherwise.
func (o *ProductConfiguratorCreateMatrixRequest) GetConfiguratorId() string {
	if o == nil || IsNil(o.ConfiguratorId) {
		var ret string
		return ret
	}
	return *o.ConfiguratorId
}

// GetConfiguratorIdOk returns a tuple with the ConfiguratorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductConfiguratorCreateMatrixRequest) GetConfiguratorIdOk() (*string, bool) {
	if o == nil || IsNil(o.ConfiguratorId) {
		return nil, false
	}
	return o.ConfiguratorId, true
}

// HasConfiguratorId returns a boolean if a field has been set.
func (o *ProductConfiguratorCreateMatrixRequest) HasConfiguratorId() bool {
	if o != nil && !IsNil(o.ConfiguratorId) {
		return true
	}

	return false
}

// SetConfiguratorId gets a reference to the given string and assigns it to the ConfiguratorId field.
func (o *ProductConfiguratorCreateMatrixRequest) SetConfiguratorId(v string) {
	o.ConfiguratorId = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ProductConfiguratorCreateMatrixRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductConfiguratorCreateMatrixRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ProductConfiguratorCreateMatrixRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ProductConfiguratorCreateMatrixRequest) SetLabel(v string) {
	o.Label = &v
}

// GetGenericType returns the GenericType field value if set, zero value otherwise.
func (o *ProductConfiguratorCreateMatrixRequest) GetGenericType() MatrixGenericType {
	if o == nil || IsNil(o.GenericType) {
		var ret MatrixGenericType
		return ret
	}
	return *o.GenericType
}

// GetGenericTypeOk returns a tuple with the GenericType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductConfiguratorCreateMatrixRequest) GetGenericTypeOk() (*MatrixGenericType, bool) {
	if o == nil || IsNil(o.GenericType) {
		return nil, false
	}
	return o.GenericType, true
}

// HasGenericType returns a boolean if a field has been set.
func (o *ProductConfiguratorCreateMatrixRequest) HasGenericType() bool {
	if o != nil && !IsNil(o.GenericType) {
		return true
	}

	return false
}

// SetGenericType gets a reference to the given MatrixGenericType and assigns it to the GenericType field.
func (o *ProductConfiguratorCreateMatrixRequest) SetGenericType(v MatrixGenericType) {
	o.GenericType = &v
}

// GetPriceType returns the PriceType field value if set, zero value otherwise.
func (o *ProductConfiguratorCreateMatrixRequest) GetPriceType() MatrixPriceType {
	if o == nil || IsNil(o.PriceType) {
		var ret MatrixPriceType
		return ret
	}
	return *o.PriceType
}

// GetPriceTypeOk returns a tuple with the PriceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductConfiguratorCreateMatrixRequest) GetPriceTypeOk() (*MatrixPriceType, bool) {
	if o == nil || IsNil(o.PriceType) {
		return nil, false
	}
	return o.PriceType, true
}

// HasPriceType returns a boolean if a field has been set.
func (o *ProductConfiguratorCreateMatrixRequest) HasPriceType() bool {
	if o != nil && !IsNil(o.PriceType) {
		return true
	}

	return false
}

// SetPriceType gets a reference to the given MatrixPriceType and assigns it to the PriceType field.
func (o *ProductConfiguratorCreateMatrixRequest) SetPriceType(v MatrixPriceType) {
	o.PriceType = &v
}

// GetWeightType returns the WeightType field value if set, zero value otherwise.
func (o *ProductConfiguratorCreateMatrixRequest) GetWeightType() MatrixWeightType {
	if o == nil || IsNil(o.WeightType) {
		var ret MatrixWeightType
		return ret
	}
	return *o.WeightType
}

// GetWeightTypeOk returns a tuple with the WeightType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductConfiguratorCreateMatrixRequest) GetWeightTypeOk() (*MatrixWeightType, bool) {
	if o == nil || IsNil(o.WeightType) {
		return nil, false
	}
	return o.WeightType, true
}

// HasWeightType returns a boolean if a field has been set.
func (o *ProductConfiguratorCreateMatrixRequest) HasWeightType() bool {
	if o != nil && !IsNil(o.WeightType) {
		return true
	}

	return false
}

// SetWeightType gets a reference to the given MatrixWeightType and assigns it to the WeightType field.
func (o *ProductConfiguratorCreateMatrixRequest) SetWeightType(v MatrixWeightType) {
	o.WeightType = &v
}

// GetSteps returns the Steps field value if set, zero value otherwise.
func (o *ProductConfiguratorCreateMatrixRequest) GetSteps() []ProductconfiguratormatrixStep {
	if o == nil || IsNil(o.Steps) {
		var ret []ProductconfiguratormatrixStep
		return ret
	}
	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductConfiguratorCreateMatrixRequest) GetStepsOk() ([]ProductconfiguratormatrixStep, bool) {
	if o == nil || IsNil(o.Steps) {
		return nil, false
	}
	return o.Steps, true
}

// HasSteps returns a boolean if a field has been set.
func (o *ProductConfiguratorCreateMatrixRequest) HasSteps() bool {
	if o != nil && !IsNil(o.Steps) {
		return true
	}

	return false
}

// SetSteps gets a reference to the given []ProductconfiguratormatrixStep and assigns it to the Steps field.
func (o *ProductConfiguratorCreateMatrixRequest) SetSteps(v []ProductconfiguratormatrixStep) {
	o.Steps = v
}

// GetPropertiesMode returns the PropertiesMode field value if set, zero value otherwise.
func (o *ProductConfiguratorCreateMatrixRequest) GetPropertiesMode() ProductconfiguratorPropertyMode {
	if o == nil || IsNil(o.PropertiesMode) {
		var ret ProductconfiguratorPropertyMode
		return ret
	}
	return *o.PropertiesMode
}

// GetPropertiesModeOk returns a tuple with the PropertiesMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductConfiguratorCreateMatrixRequest) GetPropertiesModeOk() (*ProductconfiguratorPropertyMode, bool) {
	if o == nil || IsNil(o.PropertiesMode) {
		return nil, false
	}
	return o.PropertiesMode, true
}

// HasPropertiesMode returns a boolean if a field has been set.
func (o *ProductConfiguratorCreateMatrixRequest) HasPropertiesMode() bool {
	if o != nil && !IsNil(o.PropertiesMode) {
		return true
	}

	return false
}

// SetPropertiesMode gets a reference to the given ProductconfiguratorPropertyMode and assigns it to the PropertiesMode field.
func (o *ProductConfiguratorCreateMatrixRequest) SetPropertiesMode(v ProductconfiguratorPropertyMode) {
	o.PropertiesMode = &v
}

func (o ProductConfiguratorCreateMatrixRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductConfiguratorCreateMatrixRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConfiguratorId) {
		toSerialize["configuratorId"] = o.ConfiguratorId
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.GenericType) {
		toSerialize["genericType"] = o.GenericType
	}
	if !IsNil(o.PriceType) {
		toSerialize["priceType"] = o.PriceType
	}
	if !IsNil(o.WeightType) {
		toSerialize["weightType"] = o.WeightType
	}
	if !IsNil(o.Steps) {
		toSerialize["steps"] = o.Steps
	}
	if !IsNil(o.PropertiesMode) {
		toSerialize["propertiesMode"] = o.PropertiesMode
	}
	return toSerialize, nil
}

type NullableProductConfiguratorCreateMatrixRequest struct {
	value *ProductConfiguratorCreateMatrixRequest
	isSet bool
}

func (v NullableProductConfiguratorCreateMatrixRequest) Get() *ProductConfiguratorCreateMatrixRequest {
	return v.value
}

func (v *NullableProductConfiguratorCreateMatrixRequest) Set(val *ProductConfiguratorCreateMatrixRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductConfiguratorCreateMatrixRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductConfiguratorCreateMatrixRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductConfiguratorCreateMatrixRequest(val *ProductConfiguratorCreateMatrixRequest) *NullableProductConfiguratorCreateMatrixRequest {
	return &NullableProductConfiguratorCreateMatrixRequest{value: val, isSet: true}
}

func (v NullableProductConfiguratorCreateMatrixRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductConfiguratorCreateMatrixRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

