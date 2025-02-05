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

// checks if the ProductconfiguratorstepUpdatePayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductconfiguratorstepUpdatePayload{}

// ProductconfiguratorstepUpdatePayload struct for ProductconfiguratorstepUpdatePayload
type ProductconfiguratorstepUpdatePayload struct {
	Label                *LocalisationLocalizedText `json:"label,omitempty"`
	Description          *LocalisationLocalizedText `json:"description,omitempty"`
	Position             *string                    `json:"position,omitempty"`
	IsRequired           *bool                      `json:"isRequired,omitempty"`
	SubjectToStepId      *string                    `json:"subjectToStepId,omitempty"`
	HasMultipleSelection *bool                      `json:"hasMultipleSelection,omitempty"`
	OptionsHaveQuantity  *bool                      `json:"optionsHaveQuantity,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductconfiguratorstepUpdatePayload ProductconfiguratorstepUpdatePayload

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
func (o *ProductconfiguratorstepUpdatePayload) IsSetLabel() bool {
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
func (o *ProductconfiguratorstepUpdatePayload) IsSetDescription() bool {
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
func (o *ProductconfiguratorstepUpdatePayload) IsSetPosition() bool {
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
func (o *ProductconfiguratorstepUpdatePayload) IsSetIsRequired() bool {
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
func (o *ProductconfiguratorstepUpdatePayload) IsSetSubjectToStepId() bool {
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
func (o *ProductconfiguratorstepUpdatePayload) IsSetHasMultipleSelection() bool {
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
func (o *ProductconfiguratorstepUpdatePayload) IsSetOptionsHaveQuantity() bool {
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
	toSerialize, err := o.ToMap()
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductconfiguratorstepUpdatePayload) UnmarshalJSON(data []byte) (err error) {
	varProductconfiguratorstepUpdatePayload := _ProductconfiguratorstepUpdatePayload{}

	err = json.Unmarshal(data, &varProductconfiguratorstepUpdatePayload)

	if err != nil {
		return err
	}

	*o = ProductconfiguratorstepUpdatePayload(varProductconfiguratorstepUpdatePayload)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "label")
		delete(additionalProperties, "description")
		delete(additionalProperties, "position")
		delete(additionalProperties, "isRequired")
		delete(additionalProperties, "subjectToStepId")
		delete(additionalProperties, "hasMultipleSelection")
		delete(additionalProperties, "optionsHaveQuantity")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ProductconfiguratorstepUpdatePayload) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *ProductconfiguratorstepUpdatePayload) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
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
