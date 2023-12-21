/*
Product Configurator Service

## Introduction  This comprehensive guide will equip you with the knowledge to integrate and leverage our Product Configurator Service in your applications.  ## Quick Start  Get up and running in no time! Follow these steps to kickstart your integration:  1. **Authentication:** Obtain your integration JWT to authenticate your requests. 2. **Client Libraries:** Explore our GitHub repositories to grab client libraries in your preferred programming language. 3. **API Overview:** Familiarize yourself with our RESTful API using the OpenAPI specification.  ## Integration  ### API Overview  Our RESTful API is the gateway to unlocking the full potential of Product Configurator. Check out the detailed [API Reference](/docs/category/configurator) for a granular understanding of each endpoint and request/response format.  ### Client Libraries  To expedite your integration process, we provide client libraries for various programming languages. Find the one that suits your stack in our [GitHub repositories](https://github.com/Gemini-Commerce).  ### Authentication  Security is paramount. Learn how to authenticate your requests using JWT. This ensures a secure and reliable connection between your application and Product Configurator.  ## Configuration Management  ### Configurator Lifecycle  Understand the lifecycle of configurators, from draft to active and deleted. This flexibility allows you to manage configurations at your own pace.  ### Steps and Options  Configure product steps with ease and define options effortlessly. Explore the power of dependencies to create dynamic and intuitive configurations.  ### Matrices  Delve into matrices—your secret weapon. Explore price and weight matrices, and learn how configured steps influence properties and pricing.  ### Price Management  Unleash dynamic pricing with our versatile price matrices. From fixed prices to incremental structures, adapt to diverse pricing models effortlessly.  ## Security  Your data is in safe hands. Discover how Product Configurator ensures security through JWT authentication, safeguarding your sensitive information.  ## Backward Compatibility  Stay ahead of the curve. Learn about our versioning strategy, providing backward compatibility while allowing our service to evolve seamlessly.  ## Developer Support  Have questions? Need assistance? Write to us at [info@gemini-commerce.com](mailto:info@gemini-commerce.com) and we will get back to you.

API version: v1
Contact: info@gemini-commerce.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package product_configurator

import (
	"encoding/json"
	"fmt"
)

// ProductconfiguratorWeightUnit the model 'ProductconfiguratorWeightUnit'
type ProductconfiguratorWeightUnit string

// List of productconfiguratorWeightUnit
const (
	UNKNOWN ProductconfiguratorWeightUnit = "WEIGHT_UNIT_UNKNOWN"
	KILOGRAM ProductconfiguratorWeightUnit = "WEIGHT_UNIT_KILOGRAM"
	POUND ProductconfiguratorWeightUnit = "WEIGHT_UNIT_POUND"
)

// All allowed values of ProductconfiguratorWeightUnit enum
var AllowedProductconfiguratorWeightUnitEnumValues = []ProductconfiguratorWeightUnit{
	"WEIGHT_UNIT_UNKNOWN",
	"WEIGHT_UNIT_KILOGRAM",
	"WEIGHT_UNIT_POUND",
}

func (v *ProductconfiguratorWeightUnit) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProductconfiguratorWeightUnit(value)
	for _, existing := range AllowedProductconfiguratorWeightUnitEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProductconfiguratorWeightUnit", value)
}

// NewProductconfiguratorWeightUnitFromValue returns a pointer to a valid ProductconfiguratorWeightUnit
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProductconfiguratorWeightUnitFromValue(v string) (*ProductconfiguratorWeightUnit, error) {
	ev := ProductconfiguratorWeightUnit(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProductconfiguratorWeightUnit: valid values are %v", v, AllowedProductconfiguratorWeightUnitEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProductconfiguratorWeightUnit) IsValid() bool {
	for _, existing := range AllowedProductconfiguratorWeightUnitEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to productconfiguratorWeightUnit value
func (v ProductconfiguratorWeightUnit) Ptr() *ProductconfiguratorWeightUnit {
	return &v
}

type NullableProductconfiguratorWeightUnit struct {
	value *ProductconfiguratorWeightUnit
	isSet bool
}

func (v NullableProductconfiguratorWeightUnit) Get() *ProductconfiguratorWeightUnit {
	return v.value
}

func (v *NullableProductconfiguratorWeightUnit) Set(val *ProductconfiguratorWeightUnit) {
	v.value = val
	v.isSet = true
}

func (v NullableProductconfiguratorWeightUnit) IsSet() bool {
	return v.isSet
}

func (v *NullableProductconfiguratorWeightUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductconfiguratorWeightUnit(val *ProductconfiguratorWeightUnit) *NullableProductconfiguratorWeightUnit {
	return &NullableProductconfiguratorWeightUnit{value: val, isSet: true}
}

func (v NullableProductconfiguratorWeightUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductconfiguratorWeightUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

