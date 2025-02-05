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

// checks if the ConfigurationStepOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigurationStepOption{}

// ConfigurationStepOption struct for ConfigurationStepOption
type ConfigurationStepOption struct {
	Id                   *string                       `json:"id,omitempty"`
	Grn                  *string                       `json:"grn,omitempty"`
	Label                *LocalisationLocalizedText    `json:"label,omitempty"`
	Description          *LocalisationLocalizedText    `json:"description,omitempty"`
	Swatch               *OptionSwatch                 `json:"swatch,omitempty"`
	OptionProperties     []ConfigurationOptionProperty `json:"optionProperties,omitempty"`
	HasQuantity          *bool                         `json:"hasQuantity,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConfigurationStepOption ConfigurationStepOption

// NewConfigurationStepOption instantiates a new ConfigurationStepOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigurationStepOption() *ConfigurationStepOption {
	this := ConfigurationStepOption{}
	return &this
}

// NewConfigurationStepOptionWithDefaults instantiates a new ConfigurationStepOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurationStepOptionWithDefaults() *ConfigurationStepOption {
	this := ConfigurationStepOption{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ConfigurationStepOption) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationStepOption) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ConfigurationStepOption) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ConfigurationStepOption) SetId(v string) {
	o.Id = &v
}

// GetGrn returns the Grn field value if set, zero value otherwise.
func (o *ConfigurationStepOption) GetGrn() string {
	if o == nil || IsNil(o.Grn) {
		var ret string
		return ret
	}
	return *o.Grn
}

// GetGrnOk returns a tuple with the Grn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationStepOption) GetGrnOk() (*string, bool) {
	if o == nil || IsNil(o.Grn) {
		return nil, false
	}
	return o.Grn, true
}

// HasGrn returns a boolean if a field has been set.
func (o *ConfigurationStepOption) IsSetGrn() bool {
	if o != nil && !IsNil(o.Grn) {
		return true
	}

	return false
}

// SetGrn gets a reference to the given string and assigns it to the Grn field.
func (o *ConfigurationStepOption) SetGrn(v string) {
	o.Grn = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ConfigurationStepOption) GetLabel() LocalisationLocalizedText {
	if o == nil || IsNil(o.Label) {
		var ret LocalisationLocalizedText
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationStepOption) GetLabelOk() (*LocalisationLocalizedText, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ConfigurationStepOption) IsSetLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given LocalisationLocalizedText and assigns it to the Label field.
func (o *ConfigurationStepOption) SetLabel(v LocalisationLocalizedText) {
	o.Label = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ConfigurationStepOption) GetDescription() LocalisationLocalizedText {
	if o == nil || IsNil(o.Description) {
		var ret LocalisationLocalizedText
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationStepOption) GetDescriptionOk() (*LocalisationLocalizedText, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ConfigurationStepOption) IsSetDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given LocalisationLocalizedText and assigns it to the Description field.
func (o *ConfigurationStepOption) SetDescription(v LocalisationLocalizedText) {
	o.Description = &v
}

// GetSwatch returns the Swatch field value if set, zero value otherwise.
func (o *ConfigurationStepOption) GetSwatch() OptionSwatch {
	if o == nil || IsNil(o.Swatch) {
		var ret OptionSwatch
		return ret
	}
	return *o.Swatch
}

// GetSwatchOk returns a tuple with the Swatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationStepOption) GetSwatchOk() (*OptionSwatch, bool) {
	if o == nil || IsNil(o.Swatch) {
		return nil, false
	}
	return o.Swatch, true
}

// HasSwatch returns a boolean if a field has been set.
func (o *ConfigurationStepOption) IsSetSwatch() bool {
	if o != nil && !IsNil(o.Swatch) {
		return true
	}

	return false
}

// SetSwatch gets a reference to the given OptionSwatch and assigns it to the Swatch field.
func (o *ConfigurationStepOption) SetSwatch(v OptionSwatch) {
	o.Swatch = &v
}

// GetOptionProperties returns the OptionProperties field value if set, zero value otherwise.
func (o *ConfigurationStepOption) GetOptionProperties() []ConfigurationOptionProperty {
	if o == nil || IsNil(o.OptionProperties) {
		var ret []ConfigurationOptionProperty
		return ret
	}
	return o.OptionProperties
}

// GetOptionPropertiesOk returns a tuple with the OptionProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationStepOption) GetOptionPropertiesOk() ([]ConfigurationOptionProperty, bool) {
	if o == nil || IsNil(o.OptionProperties) {
		return nil, false
	}
	return o.OptionProperties, true
}

// HasOptionProperties returns a boolean if a field has been set.
func (o *ConfigurationStepOption) IsSetOptionProperties() bool {
	if o != nil && !IsNil(o.OptionProperties) {
		return true
	}

	return false
}

// SetOptionProperties gets a reference to the given []ConfigurationOptionProperty and assigns it to the OptionProperties field.
func (o *ConfigurationStepOption) SetOptionProperties(v []ConfigurationOptionProperty) {
	o.OptionProperties = v
}

// GetHasQuantity returns the HasQuantity field value if set, zero value otherwise.
func (o *ConfigurationStepOption) GetHasQuantity() bool {
	if o == nil || IsNil(o.HasQuantity) {
		var ret bool
		return ret
	}
	return *o.HasQuantity
}

// GetHasQuantityOk returns a tuple with the HasQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationStepOption) GetHasQuantityOk() (*bool, bool) {
	if o == nil || IsNil(o.HasQuantity) {
		return nil, false
	}
	return o.HasQuantity, true
}

// HasHasQuantity returns a boolean if a field has been set.
func (o *ConfigurationStepOption) IsSetHasQuantity() bool {
	if o != nil && !IsNil(o.HasQuantity) {
		return true
	}

	return false
}

// SetHasQuantity gets a reference to the given bool and assigns it to the HasQuantity field.
func (o *ConfigurationStepOption) SetHasQuantity(v bool) {
	o.HasQuantity = &v
}

func (o ConfigurationStepOption) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigurationStepOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Grn) {
		toSerialize["grn"] = o.Grn
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Swatch) {
		toSerialize["swatch"] = o.Swatch
	}
	if !IsNil(o.OptionProperties) {
		toSerialize["optionProperties"] = o.OptionProperties
	}
	if !IsNil(o.HasQuantity) {
		toSerialize["hasQuantity"] = o.HasQuantity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConfigurationStepOption) UnmarshalJSON(data []byte) (err error) {
	varConfigurationStepOption := _ConfigurationStepOption{}

	err = json.Unmarshal(data, &varConfigurationStepOption)

	if err != nil {
		return err
	}

	*o = ConfigurationStepOption(varConfigurationStepOption)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "grn")
		delete(additionalProperties, "label")
		delete(additionalProperties, "description")
		delete(additionalProperties, "swatch")
		delete(additionalProperties, "optionProperties")
		delete(additionalProperties, "hasQuantity")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ConfigurationStepOption) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *ConfigurationStepOption) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableConfigurationStepOption struct {
	value *ConfigurationStepOption
	isSet bool
}

func (v NullableConfigurationStepOption) Get() *ConfigurationStepOption {
	return v.value
}

func (v *NullableConfigurationStepOption) Set(val *ConfigurationStepOption) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurationStepOption) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurationStepOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurationStepOption(val *ConfigurationStepOption) *NullableConfigurationStepOption {
	return &NullableConfigurationStepOption{value: val, isSet: true}
}

func (v NullableConfigurationStepOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurationStepOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
