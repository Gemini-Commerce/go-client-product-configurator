/*
Product Configurator Service

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
Contact: info@gemini-commerce.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package product_configurator

import (
	"encoding/json"
)

// ProductconfiguratoroptionCreateRequest struct for ProductconfiguratoroptionCreateRequest
type ProductconfiguratoroptionCreateRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	StepId *string `json:"stepId,omitempty"`
	Label *LocalisationLocalizedText `json:"label,omitempty"`
	Position *string `json:"position,omitempty"`
}

// NewProductconfiguratoroptionCreateRequest instantiates a new ProductconfiguratoroptionCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductconfiguratoroptionCreateRequest() *ProductconfiguratoroptionCreateRequest {
	this := ProductconfiguratoroptionCreateRequest{}
	return &this
}

// NewProductconfiguratoroptionCreateRequestWithDefaults instantiates a new ProductconfiguratoroptionCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductconfiguratoroptionCreateRequestWithDefaults() *ProductconfiguratoroptionCreateRequest {
	this := ProductconfiguratoroptionCreateRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *ProductconfiguratoroptionCreateRequest) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratoroptionCreateRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
    return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *ProductconfiguratoroptionCreateRequest) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *ProductconfiguratoroptionCreateRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetStepId returns the StepId field value if set, zero value otherwise.
func (o *ProductconfiguratoroptionCreateRequest) GetStepId() string {
	if o == nil || isNil(o.StepId) {
		var ret string
		return ret
	}
	return *o.StepId
}

// GetStepIdOk returns a tuple with the StepId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratoroptionCreateRequest) GetStepIdOk() (*string, bool) {
	if o == nil || isNil(o.StepId) {
    return nil, false
	}
	return o.StepId, true
}

// HasStepId returns a boolean if a field has been set.
func (o *ProductconfiguratoroptionCreateRequest) HasStepId() bool {
	if o != nil && !isNil(o.StepId) {
		return true
	}

	return false
}

// SetStepId gets a reference to the given string and assigns it to the StepId field.
func (o *ProductconfiguratoroptionCreateRequest) SetStepId(v string) {
	o.StepId = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ProductconfiguratoroptionCreateRequest) GetLabel() LocalisationLocalizedText {
	if o == nil || isNil(o.Label) {
		var ret LocalisationLocalizedText
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratoroptionCreateRequest) GetLabelOk() (*LocalisationLocalizedText, bool) {
	if o == nil || isNil(o.Label) {
    return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ProductconfiguratoroptionCreateRequest) HasLabel() bool {
	if o != nil && !isNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given LocalisationLocalizedText and assigns it to the Label field.
func (o *ProductconfiguratoroptionCreateRequest) SetLabel(v LocalisationLocalizedText) {
	o.Label = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *ProductconfiguratoroptionCreateRequest) GetPosition() string {
	if o == nil || isNil(o.Position) {
		var ret string
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratoroptionCreateRequest) GetPositionOk() (*string, bool) {
	if o == nil || isNil(o.Position) {
    return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *ProductconfiguratoroptionCreateRequest) HasPosition() bool {
	if o != nil && !isNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given string and assigns it to the Position field.
func (o *ProductconfiguratoroptionCreateRequest) SetPosition(v string) {
	o.Position = &v
}

func (o ProductconfiguratoroptionCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !isNil(o.StepId) {
		toSerialize["stepId"] = o.StepId
	}
	if !isNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !isNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	return json.Marshal(toSerialize)
}

type NullableProductconfiguratoroptionCreateRequest struct {
	value *ProductconfiguratoroptionCreateRequest
	isSet bool
}

func (v NullableProductconfiguratoroptionCreateRequest) Get() *ProductconfiguratoroptionCreateRequest {
	return v.value
}

func (v *NullableProductconfiguratoroptionCreateRequest) Set(val *ProductconfiguratoroptionCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductconfiguratoroptionCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductconfiguratoroptionCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductconfiguratoroptionCreateRequest(val *ProductconfiguratoroptionCreateRequest) *NullableProductconfiguratoroptionCreateRequest {
	return &NullableProductconfiguratoroptionCreateRequest{value: val, isSet: true}
}

func (v NullableProductconfiguratoroptionCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductconfiguratoroptionCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


