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

// checks if the MatrixListMatricesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MatrixListMatricesResponse{}

// MatrixListMatricesResponse struct for MatrixListMatricesResponse
type MatrixListMatricesResponse struct {
	Matrices             []ProductconfiguratormatrixEntity `json:"matrices,omitempty"`
	NextPageToken        *string                           `json:"nextPageToken,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MatrixListMatricesResponse MatrixListMatricesResponse

// NewMatrixListMatricesResponse instantiates a new MatrixListMatricesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMatrixListMatricesResponse() *MatrixListMatricesResponse {
	this := MatrixListMatricesResponse{}
	return &this
}

// NewMatrixListMatricesResponseWithDefaults instantiates a new MatrixListMatricesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMatrixListMatricesResponseWithDefaults() *MatrixListMatricesResponse {
	this := MatrixListMatricesResponse{}
	return &this
}

// GetMatrices returns the Matrices field value if set, zero value otherwise.
func (o *MatrixListMatricesResponse) GetMatrices() []ProductconfiguratormatrixEntity {
	if o == nil || IsNil(o.Matrices) {
		var ret []ProductconfiguratormatrixEntity
		return ret
	}
	return o.Matrices
}

// GetMatricesOk returns a tuple with the Matrices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatrixListMatricesResponse) GetMatricesOk() ([]ProductconfiguratormatrixEntity, bool) {
	if o == nil || IsNil(o.Matrices) {
		return nil, false
	}
	return o.Matrices, true
}

// HasMatrices returns a boolean if a field has been set.
func (o *MatrixListMatricesResponse) IsSetMatrices() bool {
	if o != nil && !IsNil(o.Matrices) {
		return true
	}

	return false
}

// SetMatrices gets a reference to the given []ProductconfiguratormatrixEntity and assigns it to the Matrices field.
func (o *MatrixListMatricesResponse) SetMatrices(v []ProductconfiguratormatrixEntity) {
	o.Matrices = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *MatrixListMatricesResponse) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatrixListMatricesResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *MatrixListMatricesResponse) IsSetNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *MatrixListMatricesResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o MatrixListMatricesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MatrixListMatricesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Matrices) {
		toSerialize["matrices"] = o.Matrices
	}
	if !IsNil(o.NextPageToken) {
		toSerialize["nextPageToken"] = o.NextPageToken
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MatrixListMatricesResponse) UnmarshalJSON(data []byte) (err error) {
	varMatrixListMatricesResponse := _MatrixListMatricesResponse{}

	err = json.Unmarshal(data, &varMatrixListMatricesResponse)

	if err != nil {
		return err
	}

	*o = MatrixListMatricesResponse(varMatrixListMatricesResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "matrices")
		delete(additionalProperties, "nextPageToken")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *MatrixListMatricesResponse) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *MatrixListMatricesResponse) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableMatrixListMatricesResponse struct {
	value *MatrixListMatricesResponse
	isSet bool
}

func (v NullableMatrixListMatricesResponse) Get() *MatrixListMatricesResponse {
	return v.value
}

func (v *NullableMatrixListMatricesResponse) Set(val *MatrixListMatricesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMatrixListMatricesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMatrixListMatricesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMatrixListMatricesResponse(val *MatrixListMatricesResponse) *NullableMatrixListMatricesResponse {
	return &NullableMatrixListMatricesResponse{value: val, isSet: true}
}

func (v NullableMatrixListMatricesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMatrixListMatricesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
