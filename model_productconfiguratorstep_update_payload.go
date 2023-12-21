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

// checks if the ProductconfiguratorstepUpdatePayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductconfiguratorstepUpdatePayload{}

// ProductconfiguratorstepUpdatePayload struct for ProductconfiguratorstepUpdatePayload
type ProductconfiguratorstepUpdatePayload struct {
	Label *LocalisationLocalizedText `json:"label,omitempty"`
	Description *LocalisationLocalizedText `json:"description,omitempty"`
	Position *string `json:"position,omitempty"`
	IsRequired *bool `json:"isRequired,omitempty"`
	SubjectToStepId *string `json:"subjectToStepId,omitempty"`
	HasMultipleSelection *bool `json:"hasMultipleSelection,omitempty"`
	OptionsHaveQuantity *bool `json:"optionsHaveQuantity,omitempty"`
}

// NewProductconfiguratorstepUpdatePayload instantiates a new ProductconfiguratorstepUpdatePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductconfiguratorstepUpdatePayload() *ProductconfiguratorstepUpdatePayload {
	this := ProductconfiguratorstepUpdatePayload{}
	return &this
}

// NewProductconfiguratorstepUpdatePayloadWithDefaults instantiates a new ProductconfiguratorstepUpdatePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductconfiguratorstepUpdatePayloadWithDefaults() *ProductconfiguratorstepUpdatePayload {
	this := ProductconfiguratorstepUpdatePayload{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ProductconfiguratorstepUpdatePayload) GetLabel() LocalisationLocalizedText {
	if o == nil || IsNil(o.Label) {
		var ret LocalisationLocalizedText
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorstepUpdatePayload) GetLabelOk() (*LocalisationLocalizedText, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ProductconfiguratorstepUpdatePayload) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given LocalisationLocalizedText and assigns it to the Label field.
func (o *ProductconfiguratorstepUpdatePayload) SetLabel(v LocalisationLocalizedText) {
	o.Label = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ProductconfiguratorstepUpdatePayload) GetDescription() LocalisationLocalizedText {
	if o == nil || IsNil(o.Description) {
		var ret LocalisationLocalizedText
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorstepUpdatePayload) GetDescriptionOk() (*LocalisationLocalizedText, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ProductconfiguratorstepUpdatePayload) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given LocalisationLocalizedText and assigns it to the Description field.
func (o *ProductconfiguratorstepUpdatePayload) SetDescription(v LocalisationLocalizedText) {
	o.Description = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *ProductconfiguratorstepUpdatePayload) GetPosition() string {
	if o == nil || IsNil(o.Position) {
		var ret string
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorstepUpdatePayload) GetPositionOk() (*string, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *ProductconfiguratorstepUpdatePayload) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given string and assigns it to the Position field.
func (o *ProductconfiguratorstepUpdatePayload) SetPosition(v string) {
	o.Position = &v
}

// GetIsRequired returns the IsRequired field value if set, zero value otherwise.
func (o *ProductconfiguratorstepUpdatePayload) GetIsRequired() bool {
	if o == nil || IsNil(o.IsRequired) {
		var ret bool
		return ret
	}
	return *o.IsRequired
}

// GetIsRequiredOk returns a tuple with the IsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorstepUpdatePayload) GetIsRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRequired) {
		return nil, false
	}
	return o.IsRequired, true
}

// HasIsRequired returns a boolean if a field has been set.
func (o *ProductconfiguratorstepUpdatePayload) HasIsRequired() bool {
	if o != nil && !IsNil(o.IsRequired) {
		return true
	}

	return false
}

// SetIsRequired gets a reference to the given bool and assigns it to the IsRequired field.
func (o *ProductconfiguratorstepUpdatePayload) SetIsRequired(v bool) {
	o.IsRequired = &v
}

// GetSubjectToStepId returns the SubjectToStepId field value if set, zero value otherwise.
func (o *ProductconfiguratorstepUpdatePayload) GetSubjectToStepId() string {
	if o == nil || IsNil(o.SubjectToStepId) {
		var ret string
		return ret
	}
	return *o.SubjectToStepId
}

// GetSubjectToStepIdOk returns a tuple with the SubjectToStepId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorstepUpdatePayload) GetSubjectToStepIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubjectToStepId) {
		return nil, false
	}
	return o.SubjectToStepId, true
}

// HasSubjectToStepId returns a boolean if a field has been set.
func (o *ProductconfiguratorstepUpdatePayload) HasSubjectToStepId() bool {
	if o != nil && !IsNil(o.SubjectToStepId) {
		return true
	}

	return false
}

// SetSubjectToStepId gets a reference to the given string and assigns it to the SubjectToStepId field.
func (o *ProductconfiguratorstepUpdatePayload) SetSubjectToStepId(v string) {
	o.SubjectToStepId = &v
}

// GetHasMultipleSelection returns the HasMultipleSelection field value if set, zero value otherwise.
func (o *ProductconfiguratorstepUpdatePayload) GetHasMultipleSelection() bool {
	if o == nil || IsNil(o.HasMultipleSelection) {
		var ret bool
		return ret
	}
	return *o.HasMultipleSelection
}

// GetHasMultipleSelectionOk returns a tuple with the HasMultipleSelection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorstepUpdatePayload) GetHasMultipleSelectionOk() (*bool, bool) {
	if o == nil || IsNil(o.HasMultipleSelection) {
		return nil, false
	}
	return o.HasMultipleSelection, true
}

// HasHasMultipleSelection returns a boolean if a field has been set.
func (o *ProductconfiguratorstepUpdatePayload) HasHasMultipleSelection() bool {
	if o != nil && !IsNil(o.HasMultipleSelection) {
		return true
	}

	return false
}

// SetHasMultipleSelection gets a reference to the given bool and assigns it to the HasMultipleSelection field.
func (o *ProductconfiguratorstepUpdatePayload) SetHasMultipleSelection(v bool) {
	o.HasMultipleSelection = &v
}

// GetOptionsHaveQuantity returns the OptionsHaveQuantity field value if set, zero value otherwise.
func (o *ProductconfiguratorstepUpdatePayload) GetOptionsHaveQuantity() bool {
	if o == nil || IsNil(o.OptionsHaveQuantity) {
		var ret bool
		return ret
	}
	return *o.OptionsHaveQuantity
}

// GetOptionsHaveQuantityOk returns a tuple with the OptionsHaveQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorstepUpdatePayload) GetOptionsHaveQuantityOk() (*bool, bool) {
	if o == nil || IsNil(o.OptionsHaveQuantity) {
		return nil, false
	}
	return o.OptionsHaveQuantity, true
}

// HasOptionsHaveQuantity returns a boolean if a field has been set.
func (o *ProductconfiguratorstepUpdatePayload) HasOptionsHaveQuantity() bool {
	if o != nil && !IsNil(o.OptionsHaveQuantity) {
		return true
	}

	return false
}

// SetOptionsHaveQuantity gets a reference to the given bool and assigns it to the OptionsHaveQuantity field.
func (o *ProductconfiguratorstepUpdatePayload) SetOptionsHaveQuantity(v bool) {
	o.OptionsHaveQuantity = &v
}

func (o ProductconfiguratorstepUpdatePayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductconfiguratorstepUpdatePayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
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
	return toSerialize, nil
}

type NullableProductconfiguratorstepUpdatePayload struct {
	value *ProductconfiguratorstepUpdatePayload
	isSet bool
}

func (v NullableProductconfiguratorstepUpdatePayload) Get() *ProductconfiguratorstepUpdatePayload {
	return v.value
}

func (v *NullableProductconfiguratorstepUpdatePayload) Set(val *ProductconfiguratorstepUpdatePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableProductconfiguratorstepUpdatePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableProductconfiguratorstepUpdatePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductconfiguratorstepUpdatePayload(val *ProductconfiguratorstepUpdatePayload) *NullableProductconfiguratorstepUpdatePayload {
	return &NullableProductconfiguratorstepUpdatePayload{value: val, isSet: true}
}

func (v NullableProductconfiguratorstepUpdatePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductconfiguratorstepUpdatePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


