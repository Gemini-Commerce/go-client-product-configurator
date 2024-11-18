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

// checks if the ProductconfiguratoroptionBulkCreateRequestCreateEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductconfiguratoroptionBulkCreateRequestCreateEntity{}

// ProductconfiguratoroptionBulkCreateRequestCreateEntity struct for ProductconfiguratoroptionBulkCreateRequestCreateEntity
type ProductconfiguratoroptionBulkCreateRequestCreateEntity struct {
	Label                *LocalisationLocalizedText `json:"label,omitempty"`
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
func (o *ProductconfiguratoroptionBulkCreateRequestCreateEntity) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given LocalisationLocalizedText and assigns it to the Label field.
func (o *ProductconfiguratoroptionBulkCreateRequestCreateEntity) SetLabel(v LocalisationLocalizedText) {
	o.Label = &v
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
func (o *ProductconfiguratoroptionBulkCreateRequestCreateEntity) HasPosition() bool {
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
func (o *ProductconfiguratoroptionBulkCreateRequestCreateEntity) HasExternalReferenceId() bool {
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
func (o *ProductconfiguratoroptionBulkCreateRequestCreateEntity) HasSwatch() bool {
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
