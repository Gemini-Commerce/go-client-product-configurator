/*
Product Configurator Service

 The Product Configurator Service is a versatile platform designed to manage dynamic product configurations.  It enables the creation, updating, and management of product configurations through steps, options, and dependencies,  ensuring granular control over customization.  ## Core Components 1. **Configurators**    - Create and manage configurators linked to products.    - Support for configurator states (`ACTIVE`, `DRAFT`, etc.) and versioning.  2. **Steps**    - Define logical sequences to guide users through the configuration process.    - Include localized labels, descriptions, and selection rules.  3. **Options**    - Add and manage options available for each step.    - Support for visual content (`Swatch`) and configurable quantities.  4. **Dependencies**    - Create rules between options and steps to control dynamic interactions.    - Manage complex conditions across configurations.  5. **Matrices**    - Use matrices to handle prices, weights, and other properties.    - Support for dynamic customization based on user selections.  6. **Properties**    - Add custom attributes and properties to configurators.  7. **Configuration Management**    - Retrieve available or user-specific configurations.    - Create optimized configurations to enhance the user experience.  ## Key Features - **Security**: Authenticate every request with JWT, ensuring safety and reliability. - **Flexibility**: Bulk operations (create, update, delete) for steps, options, and dependencies. - **Scalability**: Suitable for large volumes of configurations and complex personalization scenarios. - **Backward Compatibility**: Version management to minimize the impact of changes on existing clients.

API version: v1
Contact: info@gemini-commerce.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package productconfigurator

import (
	"encoding/json"
)

// checks if the ProductConfiguratorCreatePropertyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductConfiguratorCreatePropertyRequest{}

// ProductConfiguratorCreatePropertyRequest struct for ProductConfiguratorCreatePropertyRequest
type ProductConfiguratorCreatePropertyRequest struct {
	MatrixId             *string                                     `json:"matrixId,omitempty"`
	StepIdToOptionId     *map[string]string                          `json:"stepIdToOptionId,omitempty"`
	GenericProperty      *ProductconfiguratorpropertyGenericProperty `json:"genericProperty,omitempty"`
	PriceProperty        *ProductconfiguratorpropertyPriceProperty   `json:"priceProperty,omitempty"`
	WeightProperty       *ProductconfiguratorpropertyWeightProperty  `json:"weightProperty,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductConfiguratorCreatePropertyRequest ProductConfiguratorCreatePropertyRequest

// NewProductConfiguratorCreatePropertyRequest instantiates a new ProductConfiguratorCreatePropertyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductConfiguratorCreatePropertyRequest() *ProductConfiguratorCreatePropertyRequest {
	this := ProductConfiguratorCreatePropertyRequest{}
	return &this
}

// NewProductConfiguratorCreatePropertyRequestWithDefaults instantiates a new ProductConfiguratorCreatePropertyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductConfiguratorCreatePropertyRequestWithDefaults() *ProductConfiguratorCreatePropertyRequest {
	this := ProductConfiguratorCreatePropertyRequest{}
	return &this
}

// GetMatrixId returns the MatrixId field value if set, zero value otherwise.
func (o *ProductConfiguratorCreatePropertyRequest) GetMatrixId() string {
	if o == nil || IsNil(o.MatrixId) {
		var ret string
		return ret
	}
	return *o.MatrixId
}

// GetMatrixIdOk returns a tuple with the MatrixId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductConfiguratorCreatePropertyRequest) GetMatrixIdOk() (*string, bool) {
	if o == nil || IsNil(o.MatrixId) {
		return nil, false
	}
	return o.MatrixId, true
}

// HasMatrixId returns a boolean if a field has been set.
func (o *ProductConfiguratorCreatePropertyRequest) HasMatrixId() bool {
	if o != nil && !IsNil(o.MatrixId) {
		return true
	}

	return false
}

// SetMatrixId gets a reference to the given string and assigns it to the MatrixId field.
func (o *ProductConfiguratorCreatePropertyRequest) SetMatrixId(v string) {
	o.MatrixId = &v
}

// GetStepIdToOptionId returns the StepIdToOptionId field value if set, zero value otherwise.
func (o *ProductConfiguratorCreatePropertyRequest) GetStepIdToOptionId() map[string]string {
	if o == nil || IsNil(o.StepIdToOptionId) {
		var ret map[string]string
		return ret
	}
	return *o.StepIdToOptionId
}

// GetStepIdToOptionIdOk returns a tuple with the StepIdToOptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductConfiguratorCreatePropertyRequest) GetStepIdToOptionIdOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.StepIdToOptionId) {
		return nil, false
	}
	return o.StepIdToOptionId, true
}

// HasStepIdToOptionId returns a boolean if a field has been set.
func (o *ProductConfiguratorCreatePropertyRequest) HasStepIdToOptionId() bool {
	if o != nil && !IsNil(o.StepIdToOptionId) {
		return true
	}

	return false
}

// SetStepIdToOptionId gets a reference to the given map[string]string and assigns it to the StepIdToOptionId field.
func (o *ProductConfiguratorCreatePropertyRequest) SetStepIdToOptionId(v map[string]string) {
	o.StepIdToOptionId = &v
}

// GetGenericProperty returns the GenericProperty field value if set, zero value otherwise.
func (o *ProductConfiguratorCreatePropertyRequest) GetGenericProperty() ProductconfiguratorpropertyGenericProperty {
	if o == nil || IsNil(o.GenericProperty) {
		var ret ProductconfiguratorpropertyGenericProperty
		return ret
	}
	return *o.GenericProperty
}

// GetGenericPropertyOk returns a tuple with the GenericProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductConfiguratorCreatePropertyRequest) GetGenericPropertyOk() (*ProductconfiguratorpropertyGenericProperty, bool) {
	if o == nil || IsNil(o.GenericProperty) {
		return nil, false
	}
	return o.GenericProperty, true
}

// HasGenericProperty returns a boolean if a field has been set.
func (o *ProductConfiguratorCreatePropertyRequest) HasGenericProperty() bool {
	if o != nil && !IsNil(o.GenericProperty) {
		return true
	}

	return false
}

// SetGenericProperty gets a reference to the given ProductconfiguratorpropertyGenericProperty and assigns it to the GenericProperty field.
func (o *ProductConfiguratorCreatePropertyRequest) SetGenericProperty(v ProductconfiguratorpropertyGenericProperty) {
	o.GenericProperty = &v
}

// GetPriceProperty returns the PriceProperty field value if set, zero value otherwise.
func (o *ProductConfiguratorCreatePropertyRequest) GetPriceProperty() ProductconfiguratorpropertyPriceProperty {
	if o == nil || IsNil(o.PriceProperty) {
		var ret ProductconfiguratorpropertyPriceProperty
		return ret
	}
	return *o.PriceProperty
}

// GetPricePropertyOk returns a tuple with the PriceProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductConfiguratorCreatePropertyRequest) GetPricePropertyOk() (*ProductconfiguratorpropertyPriceProperty, bool) {
	if o == nil || IsNil(o.PriceProperty) {
		return nil, false
	}
	return o.PriceProperty, true
}

// HasPriceProperty returns a boolean if a field has been set.
func (o *ProductConfiguratorCreatePropertyRequest) HasPriceProperty() bool {
	if o != nil && !IsNil(o.PriceProperty) {
		return true
	}

	return false
}

// SetPriceProperty gets a reference to the given ProductconfiguratorpropertyPriceProperty and assigns it to the PriceProperty field.
func (o *ProductConfiguratorCreatePropertyRequest) SetPriceProperty(v ProductconfiguratorpropertyPriceProperty) {
	o.PriceProperty = &v
}

// GetWeightProperty returns the WeightProperty field value if set, zero value otherwise.
func (o *ProductConfiguratorCreatePropertyRequest) GetWeightProperty() ProductconfiguratorpropertyWeightProperty {
	if o == nil || IsNil(o.WeightProperty) {
		var ret ProductconfiguratorpropertyWeightProperty
		return ret
	}
	return *o.WeightProperty
}

// GetWeightPropertyOk returns a tuple with the WeightProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductConfiguratorCreatePropertyRequest) GetWeightPropertyOk() (*ProductconfiguratorpropertyWeightProperty, bool) {
	if o == nil || IsNil(o.WeightProperty) {
		return nil, false
	}
	return o.WeightProperty, true
}

// HasWeightProperty returns a boolean if a field has been set.
func (o *ProductConfiguratorCreatePropertyRequest) HasWeightProperty() bool {
	if o != nil && !IsNil(o.WeightProperty) {
		return true
	}

	return false
}

// SetWeightProperty gets a reference to the given ProductconfiguratorpropertyWeightProperty and assigns it to the WeightProperty field.
func (o *ProductConfiguratorCreatePropertyRequest) SetWeightProperty(v ProductconfiguratorpropertyWeightProperty) {
	o.WeightProperty = &v
}

func (o ProductConfiguratorCreatePropertyRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductConfiguratorCreatePropertyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MatrixId) {
		toSerialize["matrixId"] = o.MatrixId
	}
	if !IsNil(o.StepIdToOptionId) {
		toSerialize["stepIdToOptionId"] = o.StepIdToOptionId
	}
	if !IsNil(o.GenericProperty) {
		toSerialize["genericProperty"] = o.GenericProperty
	}
	if !IsNil(o.PriceProperty) {
		toSerialize["priceProperty"] = o.PriceProperty
	}
	if !IsNil(o.WeightProperty) {
		toSerialize["weightProperty"] = o.WeightProperty
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductConfiguratorCreatePropertyRequest) UnmarshalJSON(data []byte) (err error) {
	varProductConfiguratorCreatePropertyRequest := _ProductConfiguratorCreatePropertyRequest{}

	err = json.Unmarshal(data, &varProductConfiguratorCreatePropertyRequest)

	if err != nil {
		return err
	}

	*o = ProductConfiguratorCreatePropertyRequest(varProductConfiguratorCreatePropertyRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "matrixId")
		delete(additionalProperties, "stepIdToOptionId")
		delete(additionalProperties, "genericProperty")
		delete(additionalProperties, "priceProperty")
		delete(additionalProperties, "weightProperty")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ProductConfiguratorCreatePropertyRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *ProductConfiguratorCreatePropertyRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableProductConfiguratorCreatePropertyRequest struct {
	value *ProductConfiguratorCreatePropertyRequest
	isSet bool
}

func (v NullableProductConfiguratorCreatePropertyRequest) Get() *ProductConfiguratorCreatePropertyRequest {
	return v.value
}

func (v *NullableProductConfiguratorCreatePropertyRequest) Set(val *ProductConfiguratorCreatePropertyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductConfiguratorCreatePropertyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductConfiguratorCreatePropertyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductConfiguratorCreatePropertyRequest(val *ProductConfiguratorCreatePropertyRequest) *NullableProductConfiguratorCreatePropertyRequest {
	return &NullableProductConfiguratorCreatePropertyRequest{value: val, isSet: true}
}

func (v NullableProductConfiguratorCreatePropertyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductConfiguratorCreatePropertyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
