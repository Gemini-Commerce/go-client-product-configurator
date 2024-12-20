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

// checks if the ProductconfiguratorstepBulkCreateRequestCreateEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductconfiguratorstepBulkCreateRequestCreateEntity{}

// ProductconfiguratorstepBulkCreateRequestCreateEntity struct for ProductconfiguratorstepBulkCreateRequestCreateEntity
type ProductconfiguratorstepBulkCreateRequestCreateEntity struct {
	Label                *LocalisationLocalizedText `json:"label,omitempty"`
	Description          *LocalisationLocalizedText `json:"description,omitempty"`
	IsRequired           *bool                      `json:"isRequired,omitempty"`
	SubjectToStepId      *string                    `json:"subjectToStepId,omitempty"`
	HasMultipleSelection *bool                      `json:"hasMultipleSelection,omitempty"`
	OptionsHaveQuantity  *bool                      `json:"optionsHaveQuantity,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductconfiguratorstepBulkCreateRequestCreateEntity ProductconfiguratorstepBulkCreateRequestCreateEntity

// NewProductconfiguratorstepBulkCreateRequestCreateEntity instantiates a new ProductconfiguratorstepBulkCreateRequestCreateEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductconfiguratorstepBulkCreateRequestCreateEntity() *ProductconfiguratorstepBulkCreateRequestCreateEntity {
	this := ProductconfiguratorstepBulkCreateRequestCreateEntity{}
	return &this
}

// NewProductconfiguratorstepBulkCreateRequestCreateEntityWithDefaults instantiates a new ProductconfiguratorstepBulkCreateRequestCreateEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductconfiguratorstepBulkCreateRequestCreateEntityWithDefaults() *ProductconfiguratorstepBulkCreateRequestCreateEntity {
	this := ProductconfiguratorstepBulkCreateRequestCreateEntity{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ProductconfiguratorstepBulkCreateRequestCreateEntity) GetLabel() LocalisationLocalizedText {
	if o == nil || IsNil(o.Label) {
		var ret LocalisationLocalizedText
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorstepBulkCreateRequestCreateEntity) GetLabelOk() (*LocalisationLocalizedText, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ProductconfiguratorstepBulkCreateRequestCreateEntity) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given LocalisationLocalizedText and assigns it to the Label field.
func (o *ProductconfiguratorstepBulkCreateRequestCreateEntity) SetLabel(v LocalisationLocalizedText) {
	o.Label = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ProductconfiguratorstepBulkCreateRequestCreateEntity) GetDescription() LocalisationLocalizedText {
	if o == nil || IsNil(o.Description) {
		var ret LocalisationLocalizedText
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorstepBulkCreateRequestCreateEntity) GetDescriptionOk() (*LocalisationLocalizedText, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ProductconfiguratorstepBulkCreateRequestCreateEntity) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given LocalisationLocalizedText and assigns it to the Description field.
func (o *ProductconfiguratorstepBulkCreateRequestCreateEntity) SetDescription(v LocalisationLocalizedText) {
	o.Description = &v
}

// GetIsRequired returns the IsRequired field value if set, zero value otherwise.
func (o *ProductconfiguratorstepBulkCreateRequestCreateEntity) GetIsRequired() bool {
	if o == nil || IsNil(o.IsRequired) {
		var ret bool
		return ret
	}
	return *o.IsRequired
}

// GetIsRequiredOk returns a tuple with the IsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorstepBulkCreateRequestCreateEntity) GetIsRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRequired) {
		return nil, false
	}
	return o.IsRequired, true
}

// HasIsRequired returns a boolean if a field has been set.
func (o *ProductconfiguratorstepBulkCreateRequestCreateEntity) HasIsRequired() bool {
	if o != nil && !IsNil(o.IsRequired) {
		return true
	}

	return false
}

// SetIsRequired gets a reference to the given bool and assigns it to the IsRequired field.
func (o *ProductconfiguratorstepBulkCreateRequestCreateEntity) SetIsRequired(v bool) {
	o.IsRequired = &v
}

// GetSubjectToStepId returns the SubjectToStepId field value if set, zero value otherwise.
func (o *ProductconfiguratorstepBulkCreateRequestCreateEntity) GetSubjectToStepId() string {
	if o == nil || IsNil(o.SubjectToStepId) {
		var ret string
		return ret
	}
	return *o.SubjectToStepId
}

// GetSubjectToStepIdOk returns a tuple with the SubjectToStepId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorstepBulkCreateRequestCreateEntity) GetSubjectToStepIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubjectToStepId) {
		return nil, false
	}
	return o.SubjectToStepId, true
}

// HasSubjectToStepId returns a boolean if a field has been set.
func (o *ProductconfiguratorstepBulkCreateRequestCreateEntity) HasSubjectToStepId() bool {
	if o != nil && !IsNil(o.SubjectToStepId) {
		return true
	}

	return false
}

// SetSubjectToStepId gets a reference to the given string and assigns it to the SubjectToStepId field.
func (o *ProductconfiguratorstepBulkCreateRequestCreateEntity) SetSubjectToStepId(v string) {
	o.SubjectToStepId = &v
}

// GetHasMultipleSelection returns the HasMultipleSelection field value if set, zero value otherwise.
func (o *ProductconfiguratorstepBulkCreateRequestCreateEntity) GetHasMultipleSelection() bool {
	if o == nil || IsNil(o.HasMultipleSelection) {
		var ret bool
		return ret
	}
	return *o.HasMultipleSelection
}

// GetHasMultipleSelectionOk returns a tuple with the HasMultipleSelection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorstepBulkCreateRequestCreateEntity) GetHasMultipleSelectionOk() (*bool, bool) {
	if o == nil || IsNil(o.HasMultipleSelection) {
		return nil, false
	}
	return o.HasMultipleSelection, true
}

// HasHasMultipleSelection returns a boolean if a field has been set.
func (o *ProductconfiguratorstepBulkCreateRequestCreateEntity) HasHasMultipleSelection() bool {
	if o != nil && !IsNil(o.HasMultipleSelection) {
		return true
	}

	return false
}

// SetHasMultipleSelection gets a reference to the given bool and assigns it to the HasMultipleSelection field.
func (o *ProductconfiguratorstepBulkCreateRequestCreateEntity) SetHasMultipleSelection(v bool) {
	o.HasMultipleSelection = &v
}

// GetOptionsHaveQuantity returns the OptionsHaveQuantity field value if set, zero value otherwise.
func (o *ProductconfiguratorstepBulkCreateRequestCreateEntity) GetOptionsHaveQuantity() bool {
	if o == nil || IsNil(o.OptionsHaveQuantity) {
		var ret bool
		return ret
	}
	return *o.OptionsHaveQuantity
}

// GetOptionsHaveQuantityOk returns a tuple with the OptionsHaveQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorstepBulkCreateRequestCreateEntity) GetOptionsHaveQuantityOk() (*bool, bool) {
	if o == nil || IsNil(o.OptionsHaveQuantity) {
		return nil, false
	}
	return o.OptionsHaveQuantity, true
}

// HasOptionsHaveQuantity returns a boolean if a field has been set.
func (o *ProductconfiguratorstepBulkCreateRequestCreateEntity) HasOptionsHaveQuantity() bool {
	if o != nil && !IsNil(o.OptionsHaveQuantity) {
		return true
	}

	return false
}

// SetOptionsHaveQuantity gets a reference to the given bool and assigns it to the OptionsHaveQuantity field.
func (o *ProductconfiguratorstepBulkCreateRequestCreateEntity) SetOptionsHaveQuantity(v bool) {
	o.OptionsHaveQuantity = &v
}

func (o ProductconfiguratorstepBulkCreateRequestCreateEntity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductconfiguratorstepBulkCreateRequestCreateEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.IsRequired) {
		toSerialize["isRequired"] = o.IsRequired
	}
	if !IsNil(o.SubjectToStepId) {
		toSerialize["subjectToStepId"] = o.SubjectToStepId
	}
	if !IsNil(o.HasMultipleSelection) {
		toSerialize["hasMultipleSelection"] = o.HasMultipleSelection
	}
	if !IsNil(o.OptionsHaveQuantity) {
		toSerialize["optionsHaveQuantity"] = o.OptionsHaveQuantity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductconfiguratorstepBulkCreateRequestCreateEntity) UnmarshalJSON(data []byte) (err error) {
	varProductconfiguratorstepBulkCreateRequestCreateEntity := _ProductconfiguratorstepBulkCreateRequestCreateEntity{}

	err = json.Unmarshal(data, &varProductconfiguratorstepBulkCreateRequestCreateEntity)

	if err != nil {
		return err
	}

	*o = ProductconfiguratorstepBulkCreateRequestCreateEntity(varProductconfiguratorstepBulkCreateRequestCreateEntity)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "label")
		delete(additionalProperties, "description")
		delete(additionalProperties, "isRequired")
		delete(additionalProperties, "subjectToStepId")
		delete(additionalProperties, "hasMultipleSelection")
		delete(additionalProperties, "optionsHaveQuantity")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ProductconfiguratorstepBulkCreateRequestCreateEntity) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *ProductconfiguratorstepBulkCreateRequestCreateEntity) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableProductconfiguratorstepBulkCreateRequestCreateEntity struct {
	value *ProductconfiguratorstepBulkCreateRequestCreateEntity
	isSet bool
}

func (v NullableProductconfiguratorstepBulkCreateRequestCreateEntity) Get() *ProductconfiguratorstepBulkCreateRequestCreateEntity {
	return v.value
}

func (v *NullableProductconfiguratorstepBulkCreateRequestCreateEntity) Set(val *ProductconfiguratorstepBulkCreateRequestCreateEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableProductconfiguratorstepBulkCreateRequestCreateEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableProductconfiguratorstepBulkCreateRequestCreateEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductconfiguratorstepBulkCreateRequestCreateEntity(val *ProductconfiguratorstepBulkCreateRequestCreateEntity) *NullableProductconfiguratorstepBulkCreateRequestCreateEntity {
	return &NullableProductconfiguratorstepBulkCreateRequestCreateEntity{value: val, isSet: true}
}

func (v NullableProductconfiguratorstepBulkCreateRequestCreateEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductconfiguratorstepBulkCreateRequestCreateEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
