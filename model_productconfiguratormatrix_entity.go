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
	"time"
)

// checks if the ProductconfiguratormatrixEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductconfiguratormatrixEntity{}

// ProductconfiguratormatrixEntity struct for ProductconfiguratormatrixEntity
type ProductconfiguratormatrixEntity struct {
	Id                   *string                          `json:"id,omitempty"`
	Grn                  *string                          `json:"grn,omitempty"`
	ConfiguratorId       *string                          `json:"configuratorId,omitempty"`
	Label                *string                          `json:"label,omitempty"`
	DefaultPropertyId    *string                          `json:"defaultPropertyId,omitempty"`
	PropertiesMode       *ProductconfiguratorPropertyMode `json:"propertiesMode,omitempty"`
	GenericType          *MatrixGenericType               `json:"genericType,omitempty"`
	PriceType            *MatrixPriceType                 `json:"priceType,omitempty"`
	WeightType           *MatrixWeightType                `json:"weightType,omitempty"`
	Steps                []ProductconfiguratormatrixStep  `json:"steps,omitempty"`
	CreatedAt            *time.Time                       `json:"createdAt,omitempty"`
	UpdatedAt            *time.Time                       `json:"updatedAt,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductconfiguratormatrixEntity ProductconfiguratormatrixEntity

// NewProductconfiguratormatrixEntity instantiates a new ProductconfiguratormatrixEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductconfiguratormatrixEntity() *ProductconfiguratormatrixEntity {
	this := ProductconfiguratormatrixEntity{}
	var propertiesMode ProductconfiguratorPropertyMode = PRODUCTCONFIGURATORPROPERTYMODE_UNKNOWN
	this.PropertiesMode = &propertiesMode
	return &this
}

// NewProductconfiguratormatrixEntityWithDefaults instantiates a new ProductconfiguratormatrixEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductconfiguratormatrixEntityWithDefaults() *ProductconfiguratormatrixEntity {
	this := ProductconfiguratormatrixEntity{}
	var propertiesMode ProductconfiguratorPropertyMode = PRODUCTCONFIGURATORPROPERTYMODE_UNKNOWN
	this.PropertiesMode = &propertiesMode
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProductconfiguratormatrixEntity) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratormatrixEntity) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProductconfiguratormatrixEntity) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProductconfiguratormatrixEntity) SetId(v string) {
	o.Id = &v
}

// GetGrn returns the Grn field value if set, zero value otherwise.
func (o *ProductconfiguratormatrixEntity) GetGrn() string {
	if o == nil || IsNil(o.Grn) {
		var ret string
		return ret
	}
	return *o.Grn
}

// GetGrnOk returns a tuple with the Grn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratormatrixEntity) GetGrnOk() (*string, bool) {
	if o == nil || IsNil(o.Grn) {
		return nil, false
	}
	return o.Grn, true
}

// HasGrn returns a boolean if a field has been set.
func (o *ProductconfiguratormatrixEntity) HasGrn() bool {
	if o != nil && !IsNil(o.Grn) {
		return true
	}

	return false
}

// SetGrn gets a reference to the given string and assigns it to the Grn field.
func (o *ProductconfiguratormatrixEntity) SetGrn(v string) {
	o.Grn = &v
}

// GetConfiguratorId returns the ConfiguratorId field value if set, zero value otherwise.
func (o *ProductconfiguratormatrixEntity) GetConfiguratorId() string {
	if o == nil || IsNil(o.ConfiguratorId) {
		var ret string
		return ret
	}
	return *o.ConfiguratorId
}

// GetConfiguratorIdOk returns a tuple with the ConfiguratorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratormatrixEntity) GetConfiguratorIdOk() (*string, bool) {
	if o == nil || IsNil(o.ConfiguratorId) {
		return nil, false
	}
	return o.ConfiguratorId, true
}

// HasConfiguratorId returns a boolean if a field has been set.
func (o *ProductconfiguratormatrixEntity) HasConfiguratorId() bool {
	if o != nil && !IsNil(o.ConfiguratorId) {
		return true
	}

	return false
}

// SetConfiguratorId gets a reference to the given string and assigns it to the ConfiguratorId field.
func (o *ProductconfiguratormatrixEntity) SetConfiguratorId(v string) {
	o.ConfiguratorId = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ProductconfiguratormatrixEntity) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratormatrixEntity) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ProductconfiguratormatrixEntity) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ProductconfiguratormatrixEntity) SetLabel(v string) {
	o.Label = &v
}

// GetDefaultPropertyId returns the DefaultPropertyId field value if set, zero value otherwise.
func (o *ProductconfiguratormatrixEntity) GetDefaultPropertyId() string {
	if o == nil || IsNil(o.DefaultPropertyId) {
		var ret string
		return ret
	}
	return *o.DefaultPropertyId
}

// GetDefaultPropertyIdOk returns a tuple with the DefaultPropertyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratormatrixEntity) GetDefaultPropertyIdOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultPropertyId) {
		return nil, false
	}
	return o.DefaultPropertyId, true
}

// HasDefaultPropertyId returns a boolean if a field has been set.
func (o *ProductconfiguratormatrixEntity) HasDefaultPropertyId() bool {
	if o != nil && !IsNil(o.DefaultPropertyId) {
		return true
	}

	return false
}

// SetDefaultPropertyId gets a reference to the given string and assigns it to the DefaultPropertyId field.
func (o *ProductconfiguratormatrixEntity) SetDefaultPropertyId(v string) {
	o.DefaultPropertyId = &v
}

// GetPropertiesMode returns the PropertiesMode field value if set, zero value otherwise.
func (o *ProductconfiguratormatrixEntity) GetPropertiesMode() ProductconfiguratorPropertyMode {
	if o == nil || IsNil(o.PropertiesMode) {
		var ret ProductconfiguratorPropertyMode
		return ret
	}
	return *o.PropertiesMode
}

// GetPropertiesModeOk returns a tuple with the PropertiesMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratormatrixEntity) GetPropertiesModeOk() (*ProductconfiguratorPropertyMode, bool) {
	if o == nil || IsNil(o.PropertiesMode) {
		return nil, false
	}
	return o.PropertiesMode, true
}

// HasPropertiesMode returns a boolean if a field has been set.
func (o *ProductconfiguratormatrixEntity) HasPropertiesMode() bool {
	if o != nil && !IsNil(o.PropertiesMode) {
		return true
	}

	return false
}

// SetPropertiesMode gets a reference to the given ProductconfiguratorPropertyMode and assigns it to the PropertiesMode field.
func (o *ProductconfiguratormatrixEntity) SetPropertiesMode(v ProductconfiguratorPropertyMode) {
	o.PropertiesMode = &v
}

// GetGenericType returns the GenericType field value if set, zero value otherwise.
func (o *ProductconfiguratormatrixEntity) GetGenericType() MatrixGenericType {
	if o == nil || IsNil(o.GenericType) {
		var ret MatrixGenericType
		return ret
	}
	return *o.GenericType
}

// GetGenericTypeOk returns a tuple with the GenericType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratormatrixEntity) GetGenericTypeOk() (*MatrixGenericType, bool) {
	if o == nil || IsNil(o.GenericType) {
		return nil, false
	}
	return o.GenericType, true
}

// HasGenericType returns a boolean if a field has been set.
func (o *ProductconfiguratormatrixEntity) HasGenericType() bool {
	if o != nil && !IsNil(o.GenericType) {
		return true
	}

	return false
}

// SetGenericType gets a reference to the given MatrixGenericType and assigns it to the GenericType field.
func (o *ProductconfiguratormatrixEntity) SetGenericType(v MatrixGenericType) {
	o.GenericType = &v
}

// GetPriceType returns the PriceType field value if set, zero value otherwise.
func (o *ProductconfiguratormatrixEntity) GetPriceType() MatrixPriceType {
	if o == nil || IsNil(o.PriceType) {
		var ret MatrixPriceType
		return ret
	}
	return *o.PriceType
}

// GetPriceTypeOk returns a tuple with the PriceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratormatrixEntity) GetPriceTypeOk() (*MatrixPriceType, bool) {
	if o == nil || IsNil(o.PriceType) {
		return nil, false
	}
	return o.PriceType, true
}

// HasPriceType returns a boolean if a field has been set.
func (o *ProductconfiguratormatrixEntity) HasPriceType() bool {
	if o != nil && !IsNil(o.PriceType) {
		return true
	}

	return false
}

// SetPriceType gets a reference to the given MatrixPriceType and assigns it to the PriceType field.
func (o *ProductconfiguratormatrixEntity) SetPriceType(v MatrixPriceType) {
	o.PriceType = &v
}

// GetWeightType returns the WeightType field value if set, zero value otherwise.
func (o *ProductconfiguratormatrixEntity) GetWeightType() MatrixWeightType {
	if o == nil || IsNil(o.WeightType) {
		var ret MatrixWeightType
		return ret
	}
	return *o.WeightType
}

// GetWeightTypeOk returns a tuple with the WeightType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratormatrixEntity) GetWeightTypeOk() (*MatrixWeightType, bool) {
	if o == nil || IsNil(o.WeightType) {
		return nil, false
	}
	return o.WeightType, true
}

// HasWeightType returns a boolean if a field has been set.
func (o *ProductconfiguratormatrixEntity) HasWeightType() bool {
	if o != nil && !IsNil(o.WeightType) {
		return true
	}

	return false
}

// SetWeightType gets a reference to the given MatrixWeightType and assigns it to the WeightType field.
func (o *ProductconfiguratormatrixEntity) SetWeightType(v MatrixWeightType) {
	o.WeightType = &v
}

// GetSteps returns the Steps field value if set, zero value otherwise.
func (o *ProductconfiguratormatrixEntity) GetSteps() []ProductconfiguratormatrixStep {
	if o == nil || IsNil(o.Steps) {
		var ret []ProductconfiguratormatrixStep
		return ret
	}
	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratormatrixEntity) GetStepsOk() ([]ProductconfiguratormatrixStep, bool) {
	if o == nil || IsNil(o.Steps) {
		return nil, false
	}
	return o.Steps, true
}

// HasSteps returns a boolean if a field has been set.
func (o *ProductconfiguratormatrixEntity) HasSteps() bool {
	if o != nil && !IsNil(o.Steps) {
		return true
	}

	return false
}

// SetSteps gets a reference to the given []ProductconfiguratormatrixStep and assigns it to the Steps field.
func (o *ProductconfiguratormatrixEntity) SetSteps(v []ProductconfiguratormatrixStep) {
	o.Steps = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ProductconfiguratormatrixEntity) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratormatrixEntity) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ProductconfiguratormatrixEntity) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ProductconfiguratormatrixEntity) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ProductconfiguratormatrixEntity) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratormatrixEntity) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ProductconfiguratormatrixEntity) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ProductconfiguratormatrixEntity) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o ProductconfiguratormatrixEntity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductconfiguratormatrixEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Grn) {
		toSerialize["grn"] = o.Grn
	}
	if !IsNil(o.ConfiguratorId) {
		toSerialize["configuratorId"] = o.ConfiguratorId
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.DefaultPropertyId) {
		toSerialize["defaultPropertyId"] = o.DefaultPropertyId
	}
	if !IsNil(o.PropertiesMode) {
		toSerialize["propertiesMode"] = o.PropertiesMode
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
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductconfiguratormatrixEntity) UnmarshalJSON(data []byte) (err error) {
	varProductconfiguratormatrixEntity := _ProductconfiguratormatrixEntity{}

	err = json.Unmarshal(data, &varProductconfiguratormatrixEntity)

	if err != nil {
		return err
	}

	*o = ProductconfiguratormatrixEntity(varProductconfiguratormatrixEntity)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "grn")
		delete(additionalProperties, "configuratorId")
		delete(additionalProperties, "label")
		delete(additionalProperties, "defaultPropertyId")
		delete(additionalProperties, "propertiesMode")
		delete(additionalProperties, "genericType")
		delete(additionalProperties, "priceType")
		delete(additionalProperties, "weightType")
		delete(additionalProperties, "steps")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "updatedAt")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of wellknown types
func (o *ProductconfiguratormatrixEntity) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of wellknown types
func (o *ProductconfiguratormatrixEntity) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableProductconfiguratormatrixEntity struct {
	value *ProductconfiguratormatrixEntity
	isSet bool
}

func (v NullableProductconfiguratormatrixEntity) Get() *ProductconfiguratormatrixEntity {
	return v.value
}

func (v *NullableProductconfiguratormatrixEntity) Set(val *ProductconfiguratormatrixEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableProductconfiguratormatrixEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableProductconfiguratormatrixEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductconfiguratormatrixEntity(val *ProductconfiguratormatrixEntity) *NullableProductconfiguratormatrixEntity {
	return &NullableProductconfiguratormatrixEntity{value: val, isSet: true}
}

func (v NullableProductconfiguratormatrixEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductconfiguratormatrixEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
