/*
Product Configurator Service

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
Contact: info@gemini-commerce.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ProductconfiguratorstepBulkCreateRequestCreateEntity struct for ProductconfiguratorstepBulkCreateRequestCreateEntity
type ProductconfiguratorstepBulkCreateRequestCreateEntity struct {
	Label *LocalisationLocalizedText `json:"label,omitempty"`
	Description *LocalisationLocalizedText `json:"description,omitempty"`
	Position *string `json:"position,omitempty"`
	IsRequired *bool `json:"isRequired,omitempty"`
	SubjectToStepId *string `json:"subjectToStepId,omitempty"`
}

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
	if o == nil || isNil(o.Label) {
		var ret LocalisationLocalizedText
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorstepBulkCreateRequestCreateEntity) GetLabelOk() (*LocalisationLocalizedText, bool) {
	if o == nil || isNil(o.Label) {
    return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ProductconfiguratorstepBulkCreateRequestCreateEntity) HasLabel() bool {
	if o != nil && !isNil(o.Label) {
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
	if o == nil || isNil(o.Description) {
		var ret LocalisationLocalizedText
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorstepBulkCreateRequestCreateEntity) GetDescriptionOk() (*LocalisationLocalizedText, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ProductconfiguratorstepBulkCreateRequestCreateEntity) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given LocalisationLocalizedText and assigns it to the Description field.
func (o *ProductconfiguratorstepBulkCreateRequestCreateEntity) SetDescription(v LocalisationLocalizedText) {
	o.Description = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *ProductconfiguratorstepBulkCreateRequestCreateEntity) GetPosition() string {
	if o == nil || isNil(o.Position) {
		var ret string
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorstepBulkCreateRequestCreateEntity) GetPositionOk() (*string, bool) {
	if o == nil || isNil(o.Position) {
    return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *ProductconfiguratorstepBulkCreateRequestCreateEntity) HasPosition() bool {
	if o != nil && !isNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given string and assigns it to the Position field.
func (o *ProductconfiguratorstepBulkCreateRequestCreateEntity) SetPosition(v string) {
	o.Position = &v
}

// GetIsRequired returns the IsRequired field value if set, zero value otherwise.
func (o *ProductconfiguratorstepBulkCreateRequestCreateEntity) GetIsRequired() bool {
	if o == nil || isNil(o.IsRequired) {
		var ret bool
		return ret
	}
	return *o.IsRequired
}

// GetIsRequiredOk returns a tuple with the IsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorstepBulkCreateRequestCreateEntity) GetIsRequiredOk() (*bool, bool) {
	if o == nil || isNil(o.IsRequired) {
    return nil, false
	}
	return o.IsRequired, true
}

// HasIsRequired returns a boolean if a field has been set.
func (o *ProductconfiguratorstepBulkCreateRequestCreateEntity) HasIsRequired() bool {
	if o != nil && !isNil(o.IsRequired) {
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
	if o == nil || isNil(o.SubjectToStepId) {
		var ret string
		return ret
	}
	return *o.SubjectToStepId
}

// GetSubjectToStepIdOk returns a tuple with the SubjectToStepId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorstepBulkCreateRequestCreateEntity) GetSubjectToStepIdOk() (*string, bool) {
	if o == nil || isNil(o.SubjectToStepId) {
    return nil, false
	}
	return o.SubjectToStepId, true
}

// HasSubjectToStepId returns a boolean if a field has been set.
func (o *ProductconfiguratorstepBulkCreateRequestCreateEntity) HasSubjectToStepId() bool {
	if o != nil && !isNil(o.SubjectToStepId) {
		return true
	}

	return false
}

// SetSubjectToStepId gets a reference to the given string and assigns it to the SubjectToStepId field.
func (o *ProductconfiguratorstepBulkCreateRequestCreateEntity) SetSubjectToStepId(v string) {
	o.SubjectToStepId = &v
}

func (o ProductconfiguratorstepBulkCreateRequestCreateEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	if !isNil(o.IsRequired) {
		toSerialize["isRequired"] = o.IsRequired
	}
	if !isNil(o.SubjectToStepId) {
		toSerialize["subjectToStepId"] = o.SubjectToStepId
	}
	return json.Marshal(toSerialize)
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


