/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.7.8 (3.7)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
)

// InterfaceRequestDuplex * `half` - Half * `full` - Full * `auto` - Auto
type InterfaceRequestDuplex string

// List of InterfaceRequest_duplex
const (
	INTERFACEREQUESTDUPLEX_HALF  InterfaceRequestDuplex = "half"
	INTERFACEREQUESTDUPLEX_FULL  InterfaceRequestDuplex = "full"
	INTERFACEREQUESTDUPLEX_AUTO  InterfaceRequestDuplex = "auto"
	INTERFACEREQUESTDUPLEX_EMPTY InterfaceRequestDuplex = ""
)

// All allowed values of InterfaceRequestDuplex enum
var AllowedInterfaceRequestDuplexEnumValues = []InterfaceRequestDuplex{
	"half",
	"full",
	"auto",
	"",
}

func (v *InterfaceRequestDuplex) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InterfaceRequestDuplex(value)
	for _, existing := range AllowedInterfaceRequestDuplexEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InterfaceRequestDuplex", value)
}

// NewInterfaceRequestDuplexFromValue returns a pointer to a valid InterfaceRequestDuplex
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInterfaceRequestDuplexFromValue(v string) (*InterfaceRequestDuplex, error) {
	ev := InterfaceRequestDuplex(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InterfaceRequestDuplex: valid values are %v", v, AllowedInterfaceRequestDuplexEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InterfaceRequestDuplex) IsValid() bool {
	for _, existing := range AllowedInterfaceRequestDuplexEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InterfaceRequest_duplex value
func (v InterfaceRequestDuplex) Ptr() *InterfaceRequestDuplex {
	return &v
}

type NullableInterfaceRequestDuplex struct {
	value *InterfaceRequestDuplex
	isSet bool
}

func (v NullableInterfaceRequestDuplex) Get() *InterfaceRequestDuplex {
	return v.value
}

func (v *NullableInterfaceRequestDuplex) Set(val *InterfaceRequestDuplex) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceRequestDuplex) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceRequestDuplex) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceRequestDuplex(val *InterfaceRequestDuplex) *NullableInterfaceRequestDuplex {
	return &NullableInterfaceRequestDuplex{value: val, isSet: true}
}

func (v NullableInterfaceRequestDuplex) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceRequestDuplex) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
