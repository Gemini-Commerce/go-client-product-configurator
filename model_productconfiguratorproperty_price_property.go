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

// checks if the ProductconfiguratorpropertyPriceProperty type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductconfiguratorpropertyPriceProperty{}

// ProductconfiguratorpropertyPriceProperty struct for ProductconfiguratorpropertyPriceProperty
type ProductconfiguratorpropertyPriceProperty struct {
	Price                *ProductconfiguratorMoney `json:"price,omitempty"`
	PricelistGrn         *string                   `json:"pricelistGrn,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductconfiguratorpropertyPriceProperty ProductconfiguratorpropertyPriceProperty

// NewProductconfiguratorpropertyPriceProperty instantiates a new ProductconfiguratorpropertyPriceProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductconfiguratorpropertyPriceProperty() *ProductconfiguratorpropertyPriceProperty {
	this := ProductconfiguratorpropertyPriceProperty{}
	return &this
}

// NewProductconfiguratorpropertyPricePropertyWithDefaults instantiates a new ProductconfiguratorpropertyPriceProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductconfiguratorpropertyPricePropertyWithDefaults() *ProductconfiguratorpropertyPriceProperty {
	this := ProductconfiguratorpropertyPriceProperty{}
	return &this
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *ProductconfiguratorpropertyPriceProperty) GetPrice() ProductconfiguratorMoney {
	if o == nil || IsNil(o.Price) {
		var ret ProductconfiguratorMoney
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorpropertyPriceProperty) GetPriceOk() (*ProductconfiguratorMoney, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *ProductconfiguratorpropertyPriceProperty) IsSetPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given ProductconfiguratorMoney and assigns it to the Price field.
func (o *ProductconfiguratorpropertyPriceProperty) SetPrice(v ProductconfiguratorMoney) {
	o.Price = &v
}

// GetPricelistGrn returns the PricelistGrn field value if set, zero value otherwise.
func (o *ProductconfiguratorpropertyPriceProperty) GetPricelistGrn() string {
	if o == nil || IsNil(o.PricelistGrn) {
		var ret string
		return ret
	}
	return *o.PricelistGrn
}

// GetPricelistGrnOk returns a tuple with the PricelistGrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorpropertyPriceProperty) GetPricelistGrnOk() (*string, bool) {
	if o == nil || IsNil(o.PricelistGrn) {
		return nil, false
	}
	return o.PricelistGrn, true
}

// HasPricelistGrn returns a boolean if a field has been set.
func (o *ProductconfiguratorpropertyPriceProperty) IsSetPricelistGrn() bool {
	if o != nil && !IsNil(o.PricelistGrn) {
		return true
	}

	return false
}

// SetPricelistGrn gets a reference to the given string and assigns it to the PricelistGrn field.
func (o *ProductconfiguratorpropertyPriceProperty) SetPricelistGrn(v string) {
	o.PricelistGrn = &v
}

func (o ProductconfiguratorpropertyPriceProperty) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductconfiguratorpropertyPriceProperty) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.PricelistGrn) {
		toSerialize["pricelistGrn"] = o.PricelistGrn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductconfiguratorpropertyPriceProperty) UnmarshalJSON(data []byte) (err error) {
	varProductconfiguratorpropertyPriceProperty := _ProductconfiguratorpropertyPriceProperty{}

	err = json.Unmarshal(data, &varProductconfiguratorpropertyPriceProperty)

	if err != nil {
		return err
	}

	*o = ProductconfiguratorpropertyPriceProperty(varProductconfiguratorpropertyPriceProperty)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "price")
		delete(additionalProperties, "pricelistGrn")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ProductconfiguratorpropertyPriceProperty) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *ProductconfiguratorpropertyPriceProperty) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableProductconfiguratorpropertyPriceProperty struct {
	value *ProductconfiguratorpropertyPriceProperty
	isSet bool
}

func (v NullableProductconfiguratorpropertyPriceProperty) Get() *ProductconfiguratorpropertyPriceProperty {
	return v.value
}

func (v *NullableProductconfiguratorpropertyPriceProperty) Set(val *ProductconfiguratorpropertyPriceProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableProductconfiguratorpropertyPriceProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableProductconfiguratorpropertyPriceProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductconfiguratorpropertyPriceProperty(val *ProductconfiguratorpropertyPriceProperty) *NullableProductconfiguratorpropertyPriceProperty {
	return &NullableProductconfiguratorpropertyPriceProperty{value: val, isSet: true}
}

func (v NullableProductconfiguratorpropertyPriceProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductconfiguratorpropertyPriceProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
