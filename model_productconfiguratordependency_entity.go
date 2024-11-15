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
	"time"
)

// checks if the ProductconfiguratordependencyEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductconfiguratordependencyEntity{}

// ProductconfiguratordependencyEntity struct for ProductconfiguratordependencyEntity
type ProductconfiguratordependencyEntity struct {
	Id *string `json:"id,omitempty"`
	Grn *string `json:"grn,omitempty"`
	OptionIds []string `json:"optionIds,omitempty"`
	Condition *DependencyCondition `json:"condition,omitempty"`
	StepId *string `json:"stepId,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductconfiguratordependencyEntity ProductconfiguratordependencyEntity

// NewProductconfiguratordependencyEntity instantiates a new ProductconfiguratordependencyEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductconfiguratordependencyEntity() *ProductconfiguratordependencyEntity {
	this := ProductconfiguratordependencyEntity{}
	return &this
}

// NewProductconfiguratordependencyEntityWithDefaults instantiates a new ProductconfiguratordependencyEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductconfiguratordependencyEntityWithDefaults() *ProductconfiguratordependencyEntity {
	this := ProductconfiguratordependencyEntity{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProductconfiguratordependencyEntity) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratordependencyEntity) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProductconfiguratordependencyEntity) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProductconfiguratordependencyEntity) SetId(v string) {
	o.Id = &v
}

// GetGrn returns the Grn field value if set, zero value otherwise.
func (o *ProductconfiguratordependencyEntity) GetGrn() string {
	if o == nil || IsNil(o.Grn) {
		var ret string
		return ret
	}
	return *o.Grn
}

// GetGrnOk returns a tuple with the Grn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratordependencyEntity) GetGrnOk() (*string, bool) {
	if o == nil || IsNil(o.Grn) {
		return nil, false
	}
	return o.Grn, true
}

// HasGrn returns a boolean if a field has been set.
func (o *ProductconfiguratordependencyEntity) HasGrn() bool {
	if o != nil && !IsNil(o.Grn) {
		return true
	}

	return false
}

// SetGrn gets a reference to the given string and assigns it to the Grn field.
func (o *ProductconfiguratordependencyEntity) SetGrn(v string) {
	o.Grn = &v
}

// GetOptionIds returns the OptionIds field value if set, zero value otherwise.
func (o *ProductconfiguratordependencyEntity) GetOptionIds() []string {
	if o == nil || IsNil(o.OptionIds) {
		var ret []string
		return ret
	}
	return o.OptionIds
}

// GetOptionIdsOk returns a tuple with the OptionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratordependencyEntity) GetOptionIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.OptionIds) {
		return nil, false
	}
	return o.OptionIds, true
}

// HasOptionIds returns a boolean if a field has been set.
func (o *ProductconfiguratordependencyEntity) HasOptionIds() bool {
	if o != nil && !IsNil(o.OptionIds) {
		return true
	}

	return false
}

// SetOptionIds gets a reference to the given []string and assigns it to the OptionIds field.
func (o *ProductconfiguratordependencyEntity) SetOptionIds(v []string) {
	o.OptionIds = v
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *ProductconfiguratordependencyEntity) GetCondition() DependencyCondition {
	if o == nil || IsNil(o.Condition) {
		var ret DependencyCondition
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratordependencyEntity) GetConditionOk() (*DependencyCondition, bool) {
	if o == nil || IsNil(o.Condition) {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *ProductconfiguratordependencyEntity) HasCondition() bool {
	if o != nil && !IsNil(o.Condition) {
		return true
	}

	return false
}

// SetCondition gets a reference to the given DependencyCondition and assigns it to the Condition field.
func (o *ProductconfiguratordependencyEntity) SetCondition(v DependencyCondition) {
	o.Condition = &v
}

// GetStepId returns the StepId field value if set, zero value otherwise.
func (o *ProductconfiguratordependencyEntity) GetStepId() string {
	if o == nil || IsNil(o.StepId) {
		var ret string
		return ret
	}
	return *o.StepId
}

// GetStepIdOk returns a tuple with the StepId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratordependencyEntity) GetStepIdOk() (*string, bool) {
	if o == nil || IsNil(o.StepId) {
		return nil, false
	}
	return o.StepId, true
}

// HasStepId returns a boolean if a field has been set.
func (o *ProductconfiguratordependencyEntity) HasStepId() bool {
	if o != nil && !IsNil(o.StepId) {
		return true
	}

	return false
}

// SetStepId gets a reference to the given string and assigns it to the StepId field.
func (o *ProductconfiguratordependencyEntity) SetStepId(v string) {
	o.StepId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ProductconfiguratordependencyEntity) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratordependencyEntity) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ProductconfiguratordependencyEntity) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ProductconfiguratordependencyEntity) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ProductconfiguratordependencyEntity) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratordependencyEntity) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ProductconfiguratordependencyEntity) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ProductconfiguratordependencyEntity) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o ProductconfiguratordependencyEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductconfiguratordependencyEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Grn) {
		toSerialize["grn"] = o.Grn
	}
	if !IsNil(o.OptionIds) {
		toSerialize["optionIds"] = o.OptionIds
	}
	if !IsNil(o.Condition) {
		toSerialize["condition"] = o.Condition
	}
	if !IsNil(o.StepId) {
		toSerialize["stepId"] = o.StepId
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

func (o *ProductconfiguratordependencyEntity) UnmarshalJSON(data []byte) (err error) {
	varProductconfiguratordependencyEntity := _ProductconfiguratordependencyEntity{}

	err = json.Unmarshal(data, &varProductconfiguratordependencyEntity)

	if err != nil {
		return err
	}

	*o = ProductconfiguratordependencyEntity(varProductconfiguratordependencyEntity)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "grn")
		delete(additionalProperties, "optionIds")
		delete(additionalProperties, "condition")
		delete(additionalProperties, "stepId")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "updatedAt")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ProductconfiguratordependencyEntity) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *ProductconfiguratordependencyEntity) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableProductconfiguratordependencyEntity struct {
	value *ProductconfiguratordependencyEntity
	isSet bool
}

func (v NullableProductconfiguratordependencyEntity) Get() *ProductconfiguratordependencyEntity {
	return v.value
}

func (v *NullableProductconfiguratordependencyEntity) Set(val *ProductconfiguratordependencyEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableProductconfiguratordependencyEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableProductconfiguratordependencyEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductconfiguratordependencyEntity(val *ProductconfiguratordependencyEntity) *NullableProductconfiguratordependencyEntity {
	return &NullableProductconfiguratordependencyEntity{value: val, isSet: true}
}

func (v NullableProductconfiguratordependencyEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductconfiguratordependencyEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


