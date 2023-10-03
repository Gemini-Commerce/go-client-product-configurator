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

// ProductconfiguratorpropertyBulkCreateRequestCreateEntity struct for ProductconfiguratorpropertyBulkCreateRequestCreateEntity
type ProductconfiguratorpropertyBulkCreateRequestCreateEntity struct {
	OptionIds []string `json:"optionIds,omitempty"`
	GenericProperty *ProductconfiguratorpropertyGenericProperty `json:"genericProperty,omitempty"`
	PriceProperty map[string]interface{} `json:"priceProperty,omitempty"`
	// coordinates of the entity in the matrix. The order matters. Example: [1, 2] means that the entity is located at the first row, second column. Example: [1, 2, 3] means that the entity is located at the first row, second column, third layer. Example: [1, 2, 3, 4] means that the entity is located at the first row, second column, third layer, fourth depth.
	Coordinates []int64 `json:"coordinates,omitempty"`
}

// NewProductconfiguratorpropertyBulkCreateRequestCreateEntity instantiates a new ProductconfiguratorpropertyBulkCreateRequestCreateEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductconfiguratorpropertyBulkCreateRequestCreateEntity() *ProductconfiguratorpropertyBulkCreateRequestCreateEntity {
	this := ProductconfiguratorpropertyBulkCreateRequestCreateEntity{}
	return &this
}

// NewProductconfiguratorpropertyBulkCreateRequestCreateEntityWithDefaults instantiates a new ProductconfiguratorpropertyBulkCreateRequestCreateEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductconfiguratorpropertyBulkCreateRequestCreateEntityWithDefaults() *ProductconfiguratorpropertyBulkCreateRequestCreateEntity {
	this := ProductconfiguratorpropertyBulkCreateRequestCreateEntity{}
	return &this
}

// GetOptionIds returns the OptionIds field value if set, zero value otherwise.
func (o *ProductconfiguratorpropertyBulkCreateRequestCreateEntity) GetOptionIds() []string {
	if o == nil || isNil(o.OptionIds) {
		var ret []string
		return ret
	}
	return o.OptionIds
}

// GetOptionIdsOk returns a tuple with the OptionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorpropertyBulkCreateRequestCreateEntity) GetOptionIdsOk() ([]string, bool) {
	if o == nil || isNil(o.OptionIds) {
    return nil, false
	}
	return o.OptionIds, true
}

// HasOptionIds returns a boolean if a field has been set.
func (o *ProductconfiguratorpropertyBulkCreateRequestCreateEntity) HasOptionIds() bool {
	if o != nil && !isNil(o.OptionIds) {
		return true
	}

	return false
}

// SetOptionIds gets a reference to the given []string and assigns it to the OptionIds field.
func (o *ProductconfiguratorpropertyBulkCreateRequestCreateEntity) SetOptionIds(v []string) {
	o.OptionIds = v
}

// GetGenericProperty returns the GenericProperty field value if set, zero value otherwise.
func (o *ProductconfiguratorpropertyBulkCreateRequestCreateEntity) GetGenericProperty() ProductconfiguratorpropertyGenericProperty {
	if o == nil || isNil(o.GenericProperty) {
		var ret ProductconfiguratorpropertyGenericProperty
		return ret
	}
	return *o.GenericProperty
}

// GetGenericPropertyOk returns a tuple with the GenericProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorpropertyBulkCreateRequestCreateEntity) GetGenericPropertyOk() (*ProductconfiguratorpropertyGenericProperty, bool) {
	if o == nil || isNil(o.GenericProperty) {
    return nil, false
	}
	return o.GenericProperty, true
}

// HasGenericProperty returns a boolean if a field has been set.
func (o *ProductconfiguratorpropertyBulkCreateRequestCreateEntity) HasGenericProperty() bool {
	if o != nil && !isNil(o.GenericProperty) {
		return true
	}

	return false
}

// SetGenericProperty gets a reference to the given ProductconfiguratorpropertyGenericProperty and assigns it to the GenericProperty field.
func (o *ProductconfiguratorpropertyBulkCreateRequestCreateEntity) SetGenericProperty(v ProductconfiguratorpropertyGenericProperty) {
	o.GenericProperty = &v
}

// GetPriceProperty returns the PriceProperty field value if set, zero value otherwise.
func (o *ProductconfiguratorpropertyBulkCreateRequestCreateEntity) GetPriceProperty() map[string]interface{} {
	if o == nil || isNil(o.PriceProperty) {
		var ret map[string]interface{}
		return ret
	}
	return o.PriceProperty
}

// GetPricePropertyOk returns a tuple with the PriceProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorpropertyBulkCreateRequestCreateEntity) GetPricePropertyOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.PriceProperty) {
    return map[string]interface{}{}, false
	}
	return o.PriceProperty, true
}

// HasPriceProperty returns a boolean if a field has been set.
func (o *ProductconfiguratorpropertyBulkCreateRequestCreateEntity) HasPriceProperty() bool {
	if o != nil && !isNil(o.PriceProperty) {
		return true
	}

	return false
}

// SetPriceProperty gets a reference to the given map[string]interface{} and assigns it to the PriceProperty field.
func (o *ProductconfiguratorpropertyBulkCreateRequestCreateEntity) SetPriceProperty(v map[string]interface{}) {
	o.PriceProperty = v
}

// GetCoordinates returns the Coordinates field value if set, zero value otherwise.
func (o *ProductconfiguratorpropertyBulkCreateRequestCreateEntity) GetCoordinates() []int64 {
	if o == nil || isNil(o.Coordinates) {
		var ret []int64
		return ret
	}
	return o.Coordinates
}

// GetCoordinatesOk returns a tuple with the Coordinates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorpropertyBulkCreateRequestCreateEntity) GetCoordinatesOk() ([]int64, bool) {
	if o == nil || isNil(o.Coordinates) {
    return nil, false
	}
	return o.Coordinates, true
}

// HasCoordinates returns a boolean if a field has been set.
func (o *ProductconfiguratorpropertyBulkCreateRequestCreateEntity) HasCoordinates() bool {
	if o != nil && !isNil(o.Coordinates) {
		return true
	}

	return false
}

// SetCoordinates gets a reference to the given []int64 and assigns it to the Coordinates field.
func (o *ProductconfiguratorpropertyBulkCreateRequestCreateEntity) SetCoordinates(v []int64) {
	o.Coordinates = v
}

func (o ProductconfiguratorpropertyBulkCreateRequestCreateEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.OptionIds) {
		toSerialize["optionIds"] = o.OptionIds
	}
	if !isNil(o.GenericProperty) {
		toSerialize["genericProperty"] = o.GenericProperty
	}
	if !isNil(o.PriceProperty) {
		toSerialize["priceProperty"] = o.PriceProperty
	}
	if !isNil(o.Coordinates) {
		toSerialize["coordinates"] = o.Coordinates
	}
	return json.Marshal(toSerialize)
}

type NullableProductconfiguratorpropertyBulkCreateRequestCreateEntity struct {
	value *ProductconfiguratorpropertyBulkCreateRequestCreateEntity
	isSet bool
}

func (v NullableProductconfiguratorpropertyBulkCreateRequestCreateEntity) Get() *ProductconfiguratorpropertyBulkCreateRequestCreateEntity {
	return v.value
}

func (v *NullableProductconfiguratorpropertyBulkCreateRequestCreateEntity) Set(val *ProductconfiguratorpropertyBulkCreateRequestCreateEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableProductconfiguratorpropertyBulkCreateRequestCreateEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableProductconfiguratorpropertyBulkCreateRequestCreateEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductconfiguratorpropertyBulkCreateRequestCreateEntity(val *ProductconfiguratorpropertyBulkCreateRequestCreateEntity) *NullableProductconfiguratorpropertyBulkCreateRequestCreateEntity {
	return &NullableProductconfiguratorpropertyBulkCreateRequestCreateEntity{value: val, isSet: true}
}

func (v NullableProductconfiguratorpropertyBulkCreateRequestCreateEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductconfiguratorpropertyBulkCreateRequestCreateEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

