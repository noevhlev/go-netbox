/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.6.9 (3.6)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
)

// InterfaceRfRoleValue * `ap` - Access point * `station` - Station
type InterfaceRfRoleValue string

// List of Interface_rf_role_value
const (
	INTERFACERFROLEVALUE_AP      InterfaceRfRoleValue = "ap"
	INTERFACERFROLEVALUE_STATION InterfaceRfRoleValue = "station"
	INTERFACERFROLEVALUE_EMPTY   InterfaceRfRoleValue = ""
)

// All allowed values of InterfaceRfRoleValue enum
var AllowedInterfaceRfRoleValueEnumValues = []InterfaceRfRoleValue{
	"ap",
	"station",
	"",
}

func (v *InterfaceRfRoleValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InterfaceRfRoleValue(value)
	for _, existing := range AllowedInterfaceRfRoleValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InterfaceRfRoleValue", value)
}

// NewInterfaceRfRoleValueFromValue returns a pointer to a valid InterfaceRfRoleValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInterfaceRfRoleValueFromValue(v string) (*InterfaceRfRoleValue, error) {
	ev := InterfaceRfRoleValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InterfaceRfRoleValue: valid values are %v", v, AllowedInterfaceRfRoleValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InterfaceRfRoleValue) IsValid() bool {
	for _, existing := range AllowedInterfaceRfRoleValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Interface_rf_role_value value
func (v InterfaceRfRoleValue) Ptr() *InterfaceRfRoleValue {
	return &v
}

type NullableInterfaceRfRoleValue struct {
	value *InterfaceRfRoleValue
	isSet bool
}

func (v NullableInterfaceRfRoleValue) Get() *InterfaceRfRoleValue {
	return v.value
}

func (v *NullableInterfaceRfRoleValue) Set(val *InterfaceRfRoleValue) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceRfRoleValue) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceRfRoleValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceRfRoleValue(val *InterfaceRfRoleValue) *NullableInterfaceRfRoleValue {
	return &NullableInterfaceRfRoleValue{value: val, isSet: true}
}

func (v NullableInterfaceRfRoleValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceRfRoleValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
