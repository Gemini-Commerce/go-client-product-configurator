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

// checks if the ProductConfiguratorCreateOptionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductConfiguratorCreateOptionRequest{}

// ProductConfiguratorCreateOptionRequest struct for ProductConfiguratorCreateOptionRequest
type ProductConfiguratorCreateOptionRequest struct {
	Label                *LocalisationLocalizedText `json:"label,omitempty"`
	Description          *LocalisationLocalizedText `json:"description,omitempty"`
	Position             *string                    `json:"position,omitempty"`
	ExternalReferenceId  *string                    `json:"externalReferenceId,omitempty"`
	Swatch               *OptionSwatch              `json:"swatch,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductConfiguratorCreateOptionRequest ProductConfiguratorCreateOptionRequest

// NewProductConfiguratorCreateOptionRequest instantiates a new ProductConfiguratorCreateOptionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductConfiguratorCreateOptionRequest() *ProductConfiguratorCreateOptionRequest {
	this := ProductConfiguratorCreateOptionRequest{}
	return &this
}

// NewProductConfiguratorCreateOptionRequestWithDefaults instantiates a new ProductConfiguratorCreateOptionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductConfiguratorCreateOptionRequestWithDefaults() *ProductConfiguratorCreateOptionRequest {
	this := ProductConfiguratorCreateOptionRequest{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ProductConfiguratorCreateOptionRequest) GetLabel() LocalisationLocalizedText {
	if o == nil || IsNil(o.Label) {
		var ret LocalisationLocalizedText
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductConfiguratorCreateOptionRequest) GetLabelOk() (*LocalisationLocalizedText, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ProductConfiguratorCreateOptionRequest) IsSetLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given LocalisationLocalizedText and assigns it to the Label field.
func (o *ProductConfiguratorCreateOptionRequest) SetLabel(v LocalisationLocalizedText) {
	o.Label = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ProductConfiguratorCreateOptionRequest) GetDescription() LocalisationLocalizedText {
	if o == nil || IsNil(o.Description) {
		var ret LocalisationLocalizedText
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductConfiguratorCreateOptionRequest) GetDescriptionOk() (*LocalisationLocalizedText, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ProductConfiguratorCreateOptionRequest) IsSetDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given LocalisationLocalizedText and assigns it to the Description field.
func (o *ProductConfiguratorCreateOptionRequest) SetDescription(v LocalisationLocalizedText) {
	o.Description = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *ProductConfiguratorCreateOptionRequest) GetPosition() string {
	if o == nil || IsNil(o.Position) {
		var ret string
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductConfiguratorCreateOptionRequest) GetPositionOk() (*string, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *ProductConfiguratorCreateOptionRequest) IsSetPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given string and assigns it to the Position field.
func (o *ProductConfiguratorCreateOptionRequest) SetPosition(v string) {
	o.Position = &v
}

// GetExternalReferenceId returns the ExternalReferenceId field value if set, zero value otherwise.
func (o *ProductConfiguratorCreateOptionRequest) GetExternalReferenceId() string {
	if o == nil || IsNil(o.ExternalReferenceId) {
		var ret string
		return ret
	}
	return *o.ExternalReferenceId
}

// GetExternalReferenceIdOk returns a tuple with the ExternalReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductConfiguratorCreateOptionRequest) GetExternalReferenceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalReferenceId) {
		return nil, false
	}
	return o.ExternalReferenceId, true
}

// HasExternalReferenceId returns a boolean if a field has been set.
func (o *ProductConfiguratorCreateOptionRequest) IsSetExternalReferenceId() bool {
	if o != nil && !IsNil(o.ExternalReferenceId) {
		return true
	}

	return false
}

// SetExternalReferenceId gets a reference to the given string and assigns it to the ExternalReferenceId field.
func (o *ProductConfiguratorCreateOptionRequest) SetExternalReferenceId(v string) {
	o.ExternalReferenceId = &v
}

// GetSwatch returns the Swatch field value if set, zero value otherwise.
func (o *ProductConfiguratorCreateOptionRequest) GetSwatch() OptionSwatch {
	if o == nil || IsNil(o.Swatch) {
		var ret OptionSwatch
		return ret
	}
	return *o.Swatch
}

// GetSwatchOk returns a tuple with the Swatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductConfiguratorCreateOptionRequest) GetSwatchOk() (*OptionSwatch, bool) {
	if o == nil || IsNil(o.Swatch) {
		return nil, false
	}
	return o.Swatch, true
}

// HasSwatch returns a boolean if a field has been set.
func (o *ProductConfiguratorCreateOptionRequest) IsSetSwatch() bool {
	if o != nil && !IsNil(o.Swatch) {
		return true
	}

	return false
}

// SetSwatch gets a reference to the given OptionSwatch and assigns it to the Swatch field.
func (o *ProductConfiguratorCreateOptionRequest) SetSwatch(v OptionSwatch) {
	o.Swatch = &v
}

func (o ProductConfiguratorCreateOptionRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductConfiguratorCreateOptionRequest) ToMap() (map[string]interface{}, error) {
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

func (o *ProductConfiguratorCreateOptionRequest) UnmarshalJSON(data []byte) (err error) {
	varProductConfiguratorCreateOptionRequest := _ProductConfiguratorCreateOptionRequest{}

	err = json.Unmarshal(data, &varProductConfiguratorCreateOptionRequest)

	if err != nil {
		return err
	}

	*o = ProductConfiguratorCreateOptionRequest(varProductConfiguratorCreateOptionRequest)

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
func (o *ProductConfiguratorCreateOptionRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *ProductConfiguratorCreateOptionRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableProductConfiguratorCreateOptionRequest struct {
	value *ProductConfiguratorCreateOptionRequest
	isSet bool
}

func (v NullableProductConfiguratorCreateOptionRequest) Get() *ProductConfiguratorCreateOptionRequest {
	return v.value
}

func (v *NullableProductConfiguratorCreateOptionRequest) Set(val *ProductConfiguratorCreateOptionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductConfiguratorCreateOptionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductConfiguratorCreateOptionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductConfiguratorCreateOptionRequest(val *ProductConfiguratorCreateOptionRequest) *NullableProductConfiguratorCreateOptionRequest {
	return &NullableProductConfiguratorCreateOptionRequest{value: val, isSet: true}
}

func (v NullableProductConfiguratorCreateOptionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductConfiguratorCreateOptionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
