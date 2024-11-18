/*
Product Configurator Service

## Introduction  This comprehensive guide will equip you with the knowledge to integrate and leverage our Product Configurator Service in your applications.  ## Quick Start  Get up and running in no time! Follow these steps to kickstart your integration:  1. **Authentication:** Obtain your integration JWT to authenticate your requests. 2. **Client Libraries:** Explore our GitHub repositories to grab client libraries in your preferred programming language. 3. **API Overview:** Familiarize yourself with our RESTful API using the OpenAPI specification.  ## Integration  ### API Overview  Our RESTful API is the gateway to unlocking the full potential of Product Configurator. Check out the detailed [API Reference](/docs/category/configurator) for a granular understanding of each endpoint and request/response format.  ### Client Libraries  To expedite your integration process, we provide client libraries for various programming languages. Find the one that suits your stack in our [GitHub repositories](https://github.com/GeminiCommerce).  ### Authentication  Security is paramount. Learn how to authenticate your requests using JWT. This ensures a secure and reliable connection between your application and Product Configurator.  ## Configuration Management  ### Configurator Lifecycle  Understand the lifecycle of configurators, from draft to active and deleted. This flexibility allows you to manage configurations at your own pace.  ### Steps and Options  Configure product steps with ease and define options effortlessly. Explore the power of dependencies to create dynamic and intuitive configurations.  ### Matrices  Delve into matrices—your secret weapon. Explore price and weight matrices, and learn how configured steps influence properties and pricing.  ### Price Management  Unleash dynamic pricing with our versatile price matrices. From fixed prices to incremental structures, adapt to diverse pricing models effortlessly.  ## Security  Your data is in safe hands. Discover how Product Configurator ensures security through JWT authentication, safeguarding your sensitive information.  ## Backward Compatibility  Stay ahead of the curve. Learn about our versioning strategy, providing backward compatibility while allowing our service to evolve seamlessly.  ## Developer Support  Have questions? Need assistance? Write to us at [info@geminicommerce.com](mailto:info@geminicommerce.com) and we will get back to you.

API version: v1
Contact: info@geminicommerce.com
*/

// Code generated by OpenAPI Generator (https://openapigenerator.tech); DO NOT EDIT.

package productconfigurator

import (
	"encoding/json"
	"fmt"
)

// ProductconfiguratorconfiguratorStatus the model 'ProductconfiguratorconfiguratorStatus'
type ProductconfiguratorconfiguratorStatus string

// List of productconfiguratorconfiguratorStatus
const (
	PRODUCTCONFIGURATORCONFIGURATORSTATUS_UNKNOWN  ProductconfiguratorconfiguratorStatus = "UNKNOWN"
	PRODUCTCONFIGURATORCONFIGURATORSTATUS_ACTIVE   ProductconfiguratorconfiguratorStatus = "ACTIVE"
	PRODUCTCONFIGURATORCONFIGURATORSTATUS_DRAFT    ProductconfiguratorconfiguratorStatus = "DRAFT"
	PRODUCTCONFIGURATORCONFIGURATORSTATUS_DISABLED ProductconfiguratorconfiguratorStatus = "DISABLED"
)

// All allowed values of ProductconfiguratorconfiguratorStatus enum
var AllowedProductconfiguratorconfiguratorStatusEnumValues = []ProductconfiguratorconfiguratorStatus{
	"UNKNOWN",
	"ACTIVE",
	"DRAFT",
	"DISABLED",
}

func (v *ProductconfiguratorconfiguratorStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProductconfiguratorconfiguratorStatus(value)
	for _, existing := range AllowedProductconfiguratorconfiguratorStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProductconfiguratorconfiguratorStatus", value)
}

// NewProductconfiguratorconfiguratorStatusFromValue returns a pointer to a valid ProductconfiguratorconfiguratorStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProductconfiguratorconfiguratorStatusFromValue(v string) (*ProductconfiguratorconfiguratorStatus, error) {
	ev := ProductconfiguratorconfiguratorStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProductconfiguratorconfiguratorStatus: valid values are %v", v, AllowedProductconfiguratorconfiguratorStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProductconfiguratorconfiguratorStatus) IsValid() bool {
	for _, existing := range AllowedProductconfiguratorconfiguratorStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to productconfiguratorconfiguratorStatus value
func (v ProductconfiguratorconfiguratorStatus) Ptr() *ProductconfiguratorconfiguratorStatus {
	return &v
}

type NullableProductconfiguratorconfiguratorStatus struct {
	value *ProductconfiguratorconfiguratorStatus
	isSet bool
}

func (v NullableProductconfiguratorconfiguratorStatus) Get() *ProductconfiguratorconfiguratorStatus {
	return v.value
}

func (v *NullableProductconfiguratorconfiguratorStatus) Set(val *ProductconfiguratorconfiguratorStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableProductconfiguratorconfiguratorStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableProductconfiguratorconfiguratorStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductconfiguratorconfiguratorStatus(val *ProductconfiguratorconfiguratorStatus) *NullableProductconfiguratorconfiguratorStatus {
	return &NullableProductconfiguratorconfiguratorStatus{value: val, isSet: true}
}

func (v NullableProductconfiguratorconfiguratorStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductconfiguratorconfiguratorStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
