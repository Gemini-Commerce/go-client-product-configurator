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

// checks if the ProductconfiguratoroptionBulkCreateRequestCreateEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductconfiguratoroptionBulkCreateRequestCreateEntity{}

// ProductconfiguratoroptionBulkCreateRequestCreateEntity struct for ProductconfiguratoroptionBulkCreateRequestCreateEntity
type ProductconfiguratoroptionBulkCreateRequestCreateEntity struct {
	Label                *LocalisationLocalizedText `json:"label,omitempty"`
	Description          *LocalisationLocalizedText `json:"description,omitempty"`
	Position             *string                    `json:"position,omitempty"`
	ExternalReferenceId  *string                    `json:"externalReferenceId,omitempty"`
	Swatch               *OptionSwatch              `json:"swatch,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductconfiguratoroptionBulkCreateRequestCreateEntity ProductconfiguratoroptionBulkCreateRequestCreateEntity

// NewProductconfiguratoroptionBulkCreateRequestCreateEntity instantiates a new ProductconfiguratoroptionBulkCreateRequestCreateEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductconfiguratoroptionBulkCreateRequestCreateEntity() *ProductconfiguratoroptionBulkCreateRequestCreateEntity {
	this := ProductconfiguratoroptionBulkCreateRequestCreateEntity{}
	return &this
}

// NewProductconfiguratoroptionBulkCreateRequestCreateEntityWithDefaults instantiates a new ProductconfiguratoroptionBulkCreateRequestCreateEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductconfiguratoroptionBulkCreateRequestCreateEntityWithDefaults() *ProductconfiguratoroptionBulkCreateRequestCreateEntity {
	this := ProductconfiguratoroptionBulkCreateRequestCreateEntity{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ProductconfiguratoroptionBulkCreateRequestCreateEntity) GetLabel() LocalisationLocalizedText {
	if o == nil || IsNil(o.Label) {
		var ret LocalisationLocalizedText
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratoroptionBulkCreateRequestCreateEntity) GetLabelOk() (*LocalisationLocalizedText, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ProductconfiguratoroptionBulkCreateRequestCreateEntity) IsSetLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given LocalisationLocalizedText and assigns it to the Label field.
func (o *ProductconfiguratoroptionBulkCreateRequestCreateEntity) SetLabel(v LocalisationLocalizedText) {
	o.Label = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ProductconfiguratoroptionBulkCreateRequestCreateEntity) GetDescription() LocalisationLocalizedText {
	if o == nil || IsNil(o.Description) {
		var ret LocalisationLocalizedText
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratoroptionBulkCreateRequestCreateEntity) GetDescriptionOk() (*LocalisationLocalizedText, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ProductconfiguratoroptionBulkCreateRequestCreateEntity) IsSetDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given LocalisationLocalizedText and assigns it to the Description field.
func (o *ProductconfiguratoroptionBulkCreateRequestCreateEntity) SetDescription(v LocalisationLocalizedText) {
	o.Description = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *ProductconfiguratoroptionBulkCreateRequestCreateEntity) GetPosition() string {
	if o == nil || IsNil(o.Position) {
		var ret string
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratoroptionBulkCreateRequestCreateEntity) GetPositionOk() (*string, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *ProductconfiguratoroptionBulkCreateRequestCreateEntity) IsSetPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given string and assigns it to the Position field.
func (o *ProductconfiguratoroptionBulkCreateRequestCreateEntity) SetPosition(v string) {
	o.Position = &v
}

// GetExternalReferenceId returns the ExternalReferenceId field value if set, zero value otherwise.
func (o *ProductconfiguratoroptionBulkCreateRequestCreateEntity) GetExternalReferenceId() string {
	if o == nil || IsNil(o.ExternalReferenceId) {
		var ret string
		return ret
	}
	return *o.ExternalReferenceId
}

// GetExternalReferenceIdOk returns a tuple with the ExternalReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratoroptionBulkCreateRequestCreateEntity) GetExternalReferenceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalReferenceId) {
		return nil, false
	}
	return o.ExternalReferenceId, true
}

// HasExternalReferenceId returns a boolean if a field has been set.
func (o *ProductconfiguratoroptionBulkCreateRequestCreateEntity) IsSetExternalReferenceId() bool {
	if o != nil && !IsNil(o.ExternalReferenceId) {
		return true
	}

	return false
}

// SetExternalReferenceId gets a reference to the given string and assigns it to the ExternalReferenceId field.
func (o *ProductconfiguratoroptionBulkCreateRequestCreateEntity) SetExternalReferenceId(v string) {
	o.ExternalReferenceId = &v
}

// GetSwatch returns the Swatch field value if set, zero value otherwise.
func (o *ProductconfiguratoroptionBulkCreateRequestCreateEntity) GetSwatch() OptionSwatch {
	if o == nil || IsNil(o.Swatch) {
		var ret OptionSwatch
		return ret
	}
	return *o.Swatch
}

// GetSwatchOk returns a tuple with the Swatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratoroptionBulkCreateRequestCreateEntity) GetSwatchOk() (*OptionSwatch, bool) {
	if o == nil || IsNil(o.Swatch) {
		return nil, false
	}
	return o.Swatch, true
}

// HasSwatch returns a boolean if a field has been set.
func (o *ProductconfiguratoroptionBulkCreateRequestCreateEntity) IsSetSwatch() bool {
	if o != nil && !IsNil(o.Swatch) {
		return true
	}

	return false
}

// SetSwatch gets a reference to the given OptionSwatch and assigns it to the Swatch field.
func (o *ProductconfiguratoroptionBulkCreateRequestCreateEntity) SetSwatch(v OptionSwatch) {
	o.Swatch = &v
}

func (o ProductconfiguratoroptionBulkCreateRequestCreateEntity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductconfiguratoroptionBulkCreateRequestCreateEntity) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ExternalReferenceId) {
		toSerialize["externalReferenceId"] = o.ExternalReferenceId
	}
	if !IsNil(o.Swatch) {
		toSerialize["swatch"] = o.Swatch
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductconfiguratoroptionBulkCreateRequestCreateEntity) UnmarshalJSON(data []byte) (err error) {
	varProductconfiguratoroptionBulkCreateRequestCreateEntity := _ProductconfiguratoroptionBulkCreateRequestCreateEntity{}

	err = json.Unmarshal(data, &varProductconfiguratoroptionBulkCreateRequestCreateEntity)

	if err != nil {
		return err
	}

	*o = ProductconfiguratoroptionBulkCreateRequestCreateEntity(varProductconfiguratoroptionBulkCreateRequestCreateEntity)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "label")
		delete(additionalProperties, "description")
		delete(additionalProperties, "position")
		delete(additionalProperties, "externalReferenceId")
		delete(additionalProperties, "swatch")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ProductconfiguratoroptionBulkCreateRequestCreateEntity) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *ProductconfiguratoroptionBulkCreateRequestCreateEntity) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableProductconfiguratoroptionBulkCreateRequestCreateEntity struct {
	value *ProductconfiguratoroptionBulkCreateRequestCreateEntity
	isSet bool
}

func (v NullableProductconfiguratoroptionBulkCreateRequestCreateEntity) Get() *ProductconfiguratoroptionBulkCreateRequestCreateEntity {
	return v.value
}

func (v *NullableProductconfiguratoroptionBulkCreateRequestCreateEntity) Set(val *ProductconfiguratoroptionBulkCreateRequestCreateEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableProductconfiguratoroptionBulkCreateRequestCreateEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableProductconfiguratoroptionBulkCreateRequestCreateEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductconfiguratoroptionBulkCreateRequestCreateEntity(val *ProductconfiguratoroptionBulkCreateRequestCreateEntity) *NullableProductconfiguratoroptionBulkCreateRequestCreateEntity {
	return &NullableProductconfiguratoroptionBulkCreateRequestCreateEntity{value: val, isSet: true}
}

func (v NullableProductconfiguratoroptionBulkCreateRequestCreateEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductconfiguratoroptionBulkCreateRequestCreateEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
