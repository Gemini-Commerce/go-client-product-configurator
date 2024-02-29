/*
Product Configurator Service

## Introduction  This comprehensive guide will equip you with the knowledge to integrate and leverage our Product Configurator Service in your applications.  ## Quick Start  Get up and running in no time! Follow these steps to kickstart your integration:  1. **Authentication:** Obtain your integration JWT to authenticate your requests. 2. **Client Libraries:** Explore our GitHub repositories to grab client libraries in your preferred programming language. 3. **API Overview:** Familiarize yourself with our RESTful API using the OpenAPI specification.  ## Integration  ### API Overview  Our RESTful API is the gateway to unlocking the full potential of Product Configurator. Check out the detailed [API Reference](/docs/category/configurator) for a granular understanding of each endpoint and request/response format.  ### Client Libraries  To expedite your integration process, we provide client libraries for various programming languages. Find the one that suits your stack in our [GitHub repositories](https://github.com/Gemini-Commerce).  ### Authentication  Security is paramount. Learn how to authenticate your requests using JWT. This ensures a secure and reliable connection between your application and Product Configurator.  ## Configuration Management  ### Configurator Lifecycle  Understand the lifecycle of configurators, from draft to active and deleted. This flexibility allows you to manage configurations at your own pace.  ### Steps and Options  Configure product steps with ease and define options effortlessly. Explore the power of dependencies to create dynamic and intuitive configurations.  ### Matrices  Delve into matrices—your secret weapon. Explore price and weight matrices, and learn how configured steps influence properties and pricing.  ### Price Management  Unleash dynamic pricing with our versatile price matrices. From fixed prices to incremental structures, adapt to diverse pricing models effortlessly.  ## Security  Your data is in safe hands. Discover how Product Configurator ensures security through JWT authentication, safeguarding your sensitive information.  ## Backward Compatibility  Stay ahead of the curve. Learn about our versioning strategy, providing backward compatibility while allowing our service to evolve seamlessly.  ## Developer Support  Have questions? Need assistance? Write to us at [info@gemini-commerce.com](mailto:info@gemini-commerce.com) and we will get back to you.

API version: v1
Contact: info@gemini-commerce.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package product-configurator

import (
	"encoding/json"
)

// checks if the ConfigurationConfigurationStep type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigurationConfigurationStep{}

// ConfigurationConfigurationStep struct for ConfigurationConfigurationStep
type ConfigurationConfigurationStep struct {
	Id *string `json:"id,omitempty"`
	Grn *string `json:"grn,omitempty"`
	Label *LocalisationLocalizedText `json:"label,omitempty"`
	Options []ConfigurationStepOption `json:"options,omitempty"`
	HasMultipleSelection *bool `json:"hasMultipleSelection,omitempty"`
}

// NewConfigurationConfigurationStep instantiates a new ConfigurationConfigurationStep object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigurationConfigurationStep() *ConfigurationConfigurationStep {
	this := ConfigurationConfigurationStep{}
	return &this
}

// NewConfigurationConfigurationStepWithDefaults instantiates a new ConfigurationConfigurationStep object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurationConfigurationStepWithDefaults() *ConfigurationConfigurationStep {
	this := ConfigurationConfigurationStep{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ConfigurationConfigurationStep) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationConfigurationStep) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ConfigurationConfigurationStep) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ConfigurationConfigurationStep) SetId(v string) {
	o.Id = &v
}

// GetGrn returns the Grn field value if set, zero value otherwise.
func (o *ConfigurationConfigurationStep) GetGrn() string {
	if o == nil || IsNil(o.Grn) {
		var ret string
		return ret
	}
	return *o.Grn
}

// GetGrnOk returns a tuple with the Grn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationConfigurationStep) GetGrnOk() (*string, bool) {
	if o == nil || IsNil(o.Grn) {
		return nil, false
	}
	return o.Grn, true
}

// HasGrn returns a boolean if a field has been set.
func (o *ConfigurationConfigurationStep) HasGrn() bool {
	if o != nil && !IsNil(o.Grn) {
		return true
	}

	return false
}

// SetGrn gets a reference to the given string and assigns it to the Grn field.
func (o *ConfigurationConfigurationStep) SetGrn(v string) {
	o.Grn = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ConfigurationConfigurationStep) GetLabel() LocalisationLocalizedText {
	if o == nil || IsNil(o.Label) {
		var ret LocalisationLocalizedText
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationConfigurationStep) GetLabelOk() (*LocalisationLocalizedText, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ConfigurationConfigurationStep) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given LocalisationLocalizedText and assigns it to the Label field.
func (o *ConfigurationConfigurationStep) SetLabel(v LocalisationLocalizedText) {
	o.Label = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *ConfigurationConfigurationStep) GetOptions() []ConfigurationStepOption {
	if o == nil || IsNil(o.Options) {
		var ret []ConfigurationStepOption
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationConfigurationStep) GetOptionsOk() ([]ConfigurationStepOption, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ConfigurationConfigurationStep) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []ConfigurationStepOption and assigns it to the Options field.
func (o *ConfigurationConfigurationStep) SetOptions(v []ConfigurationStepOption) {
	o.Options = v
}

// GetHasMultipleSelection returns the HasMultipleSelection field value if set, zero value otherwise.
func (o *ConfigurationConfigurationStep) GetHasMultipleSelection() bool {
	if o == nil || IsNil(o.HasMultipleSelection) {
		var ret bool
		return ret
	}
	return *o.HasMultipleSelection
}

// GetHasMultipleSelectionOk returns a tuple with the HasMultipleSelection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationConfigurationStep) GetHasMultipleSelectionOk() (*bool, bool) {
	if o == nil || IsNil(o.HasMultipleSelection) {
		return nil, false
	}
	return o.HasMultipleSelection, true
}

// HasHasMultipleSelection returns a boolean if a field has been set.
func (o *ConfigurationConfigurationStep) HasHasMultipleSelection() bool {
	if o != nil && !IsNil(o.HasMultipleSelection) {
		return true
	}

	return false
}

// SetHasMultipleSelection gets a reference to the given bool and assigns it to the HasMultipleSelection field.
func (o *ConfigurationConfigurationStep) SetHasMultipleSelection(v bool) {
	o.HasMultipleSelection = &v
}

func (o ConfigurationConfigurationStep) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigurationConfigurationStep) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	if !IsNil(o.HasMultipleSelection) {
		toSerialize["hasMultipleSelection"] = o.HasMultipleSelection
	}
	return toSerialize, nil
}

type NullableConfigurationConfigurationStep struct {
	value *ConfigurationConfigurationStep
	isSet bool
}

func (v NullableConfigurationConfigurationStep) Get() *ConfigurationConfigurationStep {
	return v.value
}

func (v *NullableConfigurationConfigurationStep) Set(val *ConfigurationConfigurationStep) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurationConfigurationStep) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurationConfigurationStep) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurationConfigurationStep(val *ConfigurationConfigurationStep) *NullableConfigurationConfigurationStep {
	return &NullableConfigurationConfigurationStep{value: val, isSet: true}
}

func (v NullableConfigurationConfigurationStep) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurationConfigurationStep) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


