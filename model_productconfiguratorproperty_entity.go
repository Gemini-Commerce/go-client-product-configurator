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
	"time"
)

// checks if the ProductconfiguratorpropertyEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductconfiguratorpropertyEntity{}

// ProductconfiguratorpropertyEntity struct for ProductconfiguratorpropertyEntity
type ProductconfiguratorpropertyEntity struct {
	Id                   *string                          `json:"id,omitempty"`
	Grn                  *string                          `json:"grn,omitempty"`
	StepIdToOptionId     *map[string]string               `json:"stepIdToOptionId,omitempty"`
	PropertyKey          *string                          `json:"propertyKey,omitempty"`
	PropertyValue        *string                          `json:"propertyValue,omitempty"`
	PropertyType         *ProductconfiguratorPropertyType `json:"propertyType,omitempty"`
	CreatedAt            *time.Time                       `json:"createdAt,omitempty"`
	UpdatedAt            *time.Time                       `json:"updatedAt,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductconfiguratorpropertyEntity ProductconfiguratorpropertyEntity

// NewProductconfiguratorpropertyEntity instantiates a new ProductconfiguratorpropertyEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductconfiguratorpropertyEntity() *ProductconfiguratorpropertyEntity {
	this := ProductconfiguratorpropertyEntity{}
	var propertyType ProductconfiguratorPropertyType = PRODUCTCONFIGURATORPROPERTYTYPE_UNKNOWN
	this.PropertyType = &propertyType
	return &this
}

// NewProductconfiguratorpropertyEntityWithDefaults instantiates a new ProductconfiguratorpropertyEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductconfiguratorpropertyEntityWithDefaults() *ProductconfiguratorpropertyEntity {
	this := ProductconfiguratorpropertyEntity{}
	var propertyType ProductconfiguratorPropertyType = PRODUCTCONFIGURATORPROPERTYTYPE_UNKNOWN
	this.PropertyType = &propertyType
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProductconfiguratorpropertyEntity) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorpropertyEntity) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProductconfiguratorpropertyEntity) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProductconfiguratorpropertyEntity) SetId(v string) {
	o.Id = &v
}

// GetGrn returns the Grn field value if set, zero value otherwise.
func (o *ProductconfiguratorpropertyEntity) GetGrn() string {
	if o == nil || IsNil(o.Grn) {
		var ret string
		return ret
	}
	return *o.Grn
}

// GetGrnOk returns a tuple with the Grn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorpropertyEntity) GetGrnOk() (*string, bool) {
	if o == nil || IsNil(o.Grn) {
		return nil, false
	}
	return o.Grn, true
}

// HasGrn returns a boolean if a field has been set.
func (o *ProductconfiguratorpropertyEntity) IsSetGrn() bool {
	if o != nil && !IsNil(o.Grn) {
		return true
	}

	return false
}

// SetGrn gets a reference to the given string and assigns it to the Grn field.
func (o *ProductconfiguratorpropertyEntity) SetGrn(v string) {
	o.Grn = &v
}

// GetStepIdToOptionId returns the StepIdToOptionId field value if set, zero value otherwise.
func (o *ProductconfiguratorpropertyEntity) GetStepIdToOptionId() map[string]string {
	if o == nil || IsNil(o.StepIdToOptionId) {
		var ret map[string]string
		return ret
	}
	return *o.StepIdToOptionId
}

// GetStepIdToOptionIdOk returns a tuple with the StepIdToOptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorpropertyEntity) GetStepIdToOptionIdOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.StepIdToOptionId) {
		return nil, false
	}
	return o.StepIdToOptionId, true
}

// HasStepIdToOptionId returns a boolean if a field has been set.
func (o *ProductconfiguratorpropertyEntity) IsSetStepIdToOptionId() bool {
	if o != nil && !IsNil(o.StepIdToOptionId) {
		return true
	}

	return false
}

// SetStepIdToOptionId gets a reference to the given map[string]string and assigns it to the StepIdToOptionId field.
func (o *ProductconfiguratorpropertyEntity) SetStepIdToOptionId(v map[string]string) {
	o.StepIdToOptionId = &v
}

// GetPropertyKey returns the PropertyKey field value if set, zero value otherwise.
func (o *ProductconfiguratorpropertyEntity) GetPropertyKey() string {
	if o == nil || IsNil(o.PropertyKey) {
		var ret string
		return ret
	}
	return *o.PropertyKey
}

// GetPropertyKeyOk returns a tuple with the PropertyKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorpropertyEntity) GetPropertyKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PropertyKey) {
		return nil, false
	}
	return o.PropertyKey, true
}

// HasPropertyKey returns a boolean if a field has been set.
func (o *ProductconfiguratorpropertyEntity) IsSetPropertyKey() bool {
	if o != nil && !IsNil(o.PropertyKey) {
		return true
	}

	return false
}

// SetPropertyKey gets a reference to the given string and assigns it to the PropertyKey field.
func (o *ProductconfiguratorpropertyEntity) SetPropertyKey(v string) {
	o.PropertyKey = &v
}

// GetPropertyValue returns the PropertyValue field value if set, zero value otherwise.
func (o *ProductconfiguratorpropertyEntity) GetPropertyValue() string {
	if o == nil || IsNil(o.PropertyValue) {
		var ret string
		return ret
	}
	return *o.PropertyValue
}

// GetPropertyValueOk returns a tuple with the PropertyValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorpropertyEntity) GetPropertyValueOk() (*string, bool) {
	if o == nil || IsNil(o.PropertyValue) {
		return nil, false
	}
	return o.PropertyValue, true
}

// HasPropertyValue returns a boolean if a field has been set.
func (o *ProductconfiguratorpropertyEntity) IsSetPropertyValue() bool {
	if o != nil && !IsNil(o.PropertyValue) {
		return true
	}

	return false
}

// SetPropertyValue gets a reference to the given string and assigns it to the PropertyValue field.
func (o *ProductconfiguratorpropertyEntity) SetPropertyValue(v string) {
	o.PropertyValue = &v
}

// GetPropertyType returns the PropertyType field value if set, zero value otherwise.
func (o *ProductconfiguratorpropertyEntity) GetPropertyType() ProductconfiguratorPropertyType {
	if o == nil || IsNil(o.PropertyType) {
		var ret ProductconfiguratorPropertyType
		return ret
	}
	return *o.PropertyType
}

// GetPropertyTypeOk returns a tuple with the PropertyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorpropertyEntity) GetPropertyTypeOk() (*ProductconfiguratorPropertyType, bool) {
	if o == nil || IsNil(o.PropertyType) {
		return nil, false
	}
	return o.PropertyType, true
}

// HasPropertyType returns a boolean if a field has been set.
func (o *ProductconfiguratorpropertyEntity) IsSetPropertyType() bool {
	if o != nil && !IsNil(o.PropertyType) {
		return true
	}

	return false
}

// SetPropertyType gets a reference to the given ProductconfiguratorPropertyType and assigns it to the PropertyType field.
func (o *ProductconfiguratorpropertyEntity) SetPropertyType(v ProductconfiguratorPropertyType) {
	o.PropertyType = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ProductconfiguratorpropertyEntity) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorpropertyEntity) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ProductconfiguratorpropertyEntity) IsSetCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ProductconfiguratorpropertyEntity) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ProductconfiguratorpropertyEntity) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductconfiguratorpropertyEntity) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ProductconfiguratorpropertyEntity) IsSetUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ProductconfiguratorpropertyEntity) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o ProductconfiguratorpropertyEntity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductconfiguratorpropertyEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Grn) {
		toSerialize["grn"] = o.Grn
	}
	if !IsNil(o.StepIdToOptionId) {
		toSerialize["stepIdToOptionId"] = o.StepIdToOptionId
	}
	if !IsNil(o.PropertyKey) {
		toSerialize["propertyKey"] = o.PropertyKey
	}
	if !IsNil(o.PropertyValue) {
		toSerialize["propertyValue"] = o.PropertyValue
	}
	if !IsNil(o.PropertyType) {
		toSerialize["propertyType"] = o.PropertyType
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductconfiguratorpropertyEntity) UnmarshalJSON(data []byte) (err error) {
	varProductconfiguratorpropertyEntity := _ProductconfiguratorpropertyEntity{}

	err = json.Unmarshal(data, &varProductconfiguratorpropertyEntity)

	if err != nil {
		return err
	}

	*o = ProductconfiguratorpropertyEntity(varProductconfiguratorpropertyEntity)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "grn")
		delete(additionalProperties, "stepIdToOptionId")
		delete(additionalProperties, "propertyKey")
		delete(additionalProperties, "propertyValue")
		delete(additionalProperties, "propertyType")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "updatedAt")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ProductconfiguratorpropertyEntity) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *ProductconfiguratorpropertyEntity) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableProductconfiguratorpropertyEntity struct {
	value *ProductconfiguratorpropertyEntity
	isSet bool
}

func (v NullableProductconfiguratorpropertyEntity) Get() *ProductconfiguratorpropertyEntity {
	return v.value
}

func (v *NullableProductconfiguratorpropertyEntity) Set(val *ProductconfiguratorpropertyEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableProductconfiguratorpropertyEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableProductconfiguratorpropertyEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductconfiguratorpropertyEntity(val *ProductconfiguratorpropertyEntity) *NullableProductconfiguratorpropertyEntity {
	return &NullableProductconfiguratorpropertyEntity{value: val, isSet: true}
}

func (v NullableProductconfiguratorpropertyEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductconfiguratorpropertyEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
