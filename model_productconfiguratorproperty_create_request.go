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

// ProductconfiguratorpropertyCreateRequest struct for ProductconfiguratorpropertyCreateRequest
type ProductconfiguratorpropertyCreateRequest struct {
	TenantId *string `json:"tenantId,omitempty"`
	MatrixId *string `json:"matrixId,omitempty"`
	OptionIds []string `json:"optionIds,omitempty"`
	GenericProperty *ProductconfiguratorpropertyGenericProperty `json:"genericProperty,omitempty"`
	PriceProperty map[string]interface{} `json:"priceProperty,omitempty"`
	// coordinates of the entity in the matrix. The order matters. Example: [1, 2] means that the entity is located at the first row, second column. Example: [1, 2, 3] means that the entity is located at the first row, second column, third layer. Example: [1, 2, 3, 4] means that the entity is located at the first row, second column, third layer, fourth depth.
	Coordinates []int64 `json:"coordinates,omitempty"`
}

// NewProductconfiguratorpropertyCreateRequest instantiates a new ProductconfiguratorpropertyCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductconfiguratorpropertyCreateRequest() *ProductconfiguratorpropertyCreateRequest {
	this := ProductconfiguratorpropertyCreateRequest{}
	return &this
}

// NewProductconfiguratorpropertyCreateRequestWithDefaults instantiates a new ProductconfiguratorpropertyCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductconfiguratorpropertyCreateRequestWithDefaults() *ProductconfiguratorpropertyCreateRequest {
	this := ProductconfiguratorpropertyCreateRequest{}
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *ProductconfiguratorpropertyCreateRequest) GetTenantId() string {
	if o == nil || isNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorpropertyCreateRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || isNil(o.TenantId) {
    return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *ProductconfiguratorpropertyCreateRequest) HasTenantId() bool {
	if o != nil && !isNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *ProductconfiguratorpropertyCreateRequest) SetTenantId(v string) {
	o.TenantId = &v
}

// GetMatrixId returns the MatrixId field value if set, zero value otherwise.
func (o *ProductconfiguratorpropertyCreateRequest) GetMatrixId() string {
	if o == nil || isNil(o.MatrixId) {
		var ret string
		return ret
	}
	return *o.MatrixId
}

// GetMatrixIdOk returns a tuple with the MatrixId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorpropertyCreateRequest) GetMatrixIdOk() (*string, bool) {
	if o == nil || isNil(o.MatrixId) {
    return nil, false
	}
	return o.MatrixId, true
}

// HasMatrixId returns a boolean if a field has been set.
func (o *ProductconfiguratorpropertyCreateRequest) HasMatrixId() bool {
	if o != nil && !isNil(o.MatrixId) {
		return true
	}

	return false
}

// SetMatrixId gets a reference to the given string and assigns it to the MatrixId field.
func (o *ProductconfiguratorpropertyCreateRequest) SetMatrixId(v string) {
	o.MatrixId = &v
}

// GetOptionIds returns the OptionIds field value if set, zero value otherwise.
func (o *ProductconfiguratorpropertyCreateRequest) GetOptionIds() []string {
	if o == nil || isNil(o.OptionIds) {
		var ret []string
		return ret
	}
	return o.OptionIds
}

// GetOptionIdsOk returns a tuple with the OptionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorpropertyCreateRequest) GetOptionIdsOk() ([]string, bool) {
	if o == nil || isNil(o.OptionIds) {
    return nil, false
	}
	return o.OptionIds, true
}

// HasOptionIds returns a boolean if a field has been set.
func (o *ProductconfiguratorpropertyCreateRequest) HasOptionIds() bool {
	if o != nil && !isNil(o.OptionIds) {
		return true
	}

	return false
}

// SetOptionIds gets a reference to the given []string and assigns it to the OptionIds field.
func (o *ProductconfiguratorpropertyCreateRequest) SetOptionIds(v []string) {
	o.OptionIds = v
}

// GetGenericProperty returns the GenericProperty field value if set, zero value otherwise.
func (o *ProductconfiguratorpropertyCreateRequest) GetGenericProperty() ProductconfiguratorpropertyGenericProperty {
	if o == nil || isNil(o.GenericProperty) {
		var ret ProductconfiguratorpropertyGenericProperty
		return ret
	}
	return *o.GenericProperty
}

// GetGenericPropertyOk returns a tuple with the GenericProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorpropertyCreateRequest) GetGenericPropertyOk() (*ProductconfiguratorpropertyGenericProperty, bool) {
	if o == nil || isNil(o.GenericProperty) {
    return nil, false
	}
	return o.GenericProperty, true
}

// HasGenericProperty returns a boolean if a field has been set.
func (o *ProductconfiguratorpropertyCreateRequest) HasGenericProperty() bool {
	if o != nil && !isNil(o.GenericProperty) {
		return true
	}

	return false
}

// SetGenericProperty gets a reference to the given ProductconfiguratorpropertyGenericProperty and assigns it to the GenericProperty field.
func (o *ProductconfiguratorpropertyCreateRequest) SetGenericProperty(v ProductconfiguratorpropertyGenericProperty) {
	o.GenericProperty = &v
}

// GetPriceProperty returns the PriceProperty field value if set, zero value otherwise.
func (o *ProductconfiguratorpropertyCreateRequest) GetPriceProperty() map[string]interface{} {
	if o == nil || isNil(o.PriceProperty) {
		var ret map[string]interface{}
		return ret
	}
	return o.PriceProperty
}

// GetPricePropertyOk returns a tuple with the PriceProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorpropertyCreateRequest) GetPricePropertyOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.PriceProperty) {
    return map[string]interface{}{}, false
	}
	return o.PriceProperty, true
}

// HasPriceProperty returns a boolean if a field has been set.
func (o *ProductconfiguratorpropertyCreateRequest) HasPriceProperty() bool {
	if o != nil && !isNil(o.PriceProperty) {
		return true
	}

	return false
}

// SetPriceProperty gets a reference to the given map[string]interface{} and assigns it to the PriceProperty field.
func (o *ProductconfiguratorpropertyCreateRequest) SetPriceProperty(v map[string]interface{}) {
	o.PriceProperty = v
}

// GetCoordinates returns the Coordinates field value if set, zero value otherwise.
func (o *ProductconfiguratorpropertyCreateRequest) GetCoordinates() []int64 {
	if o == nil || isNil(o.Coordinates) {
		var ret []int64
		return ret
	}
	return o.Coordinates
}

// GetCoordinatesOk returns a tuple with the Coordinates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorpropertyCreateRequest) GetCoordinatesOk() ([]int64, bool) {
	if o == nil || isNil(o.Coordinates) {
    return nil, false
	}
	return o.Coordinates, true
}

// HasCoordinates returns a boolean if a field has been set.
func (o *ProductconfiguratorpropertyCreateRequest) HasCoordinates() bool {
	if o != nil && !isNil(o.Coordinates) {
		return true
	}

	return false
}

// SetCoordinates gets a reference to the given []int64 and assigns it to the Coordinates field.
func (o *ProductconfiguratorpropertyCreateRequest) SetCoordinates(v []int64) {
	o.Coordinates = v
}

func (o ProductconfiguratorpropertyCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if !isNil(o.MatrixId) {
		toSerialize["matrixId"] = o.MatrixId
	}
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

type NullableProductconfiguratorpropertyCreateRequest struct {
	value *ProductconfiguratorpropertyCreateRequest
	isSet bool
}

func (v NullableProductconfiguratorpropertyCreateRequest) Get() *ProductconfiguratorpropertyCreateRequest {
	return v.value
}

func (v *NullableProductconfiguratorpropertyCreateRequest) Set(val *ProductconfiguratorpropertyCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductconfiguratorpropertyCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductconfiguratorpropertyCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductconfiguratorpropertyCreateRequest(val *ProductconfiguratorpropertyCreateRequest) *NullableProductconfiguratorpropertyCreateRequest {
	return &NullableProductconfiguratorpropertyCreateRequest{value: val, isSet: true}
}

func (v NullableProductconfiguratorpropertyCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductconfiguratorpropertyCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


