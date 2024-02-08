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

// InterfaceDuplexLabel the model 'InterfaceDuplexLabel'
type InterfaceDuplexLabel string

// List of Interface_duplex_label
const (
	INTERFACEDUPLEXLABEL_HALF InterfaceDuplexLabel = "Half"
	INTERFACEDUPLEXLABEL_FULL InterfaceDuplexLabel = "Full"
	INTERFACEDUPLEXLABEL_AUTO InterfaceDuplexLabel = "Auto"
)

// All allowed values of InterfaceDuplexLabel enum
var AllowedInterfaceDuplexLabelEnumValues = []InterfaceDuplexLabel{
	"Half",
	"Full",
	"Auto",
}

func (v *InterfaceDuplexLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InterfaceDuplexLabel(value)
	for _, existing := range AllowedInterfaceDuplexLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InterfaceDuplexLabel", value)
}

// NewInterfaceDuplexLabelFromValue returns a pointer to a valid InterfaceDuplexLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInterfaceDuplexLabelFromValue(v string) (*InterfaceDuplexLabel, error) {
	ev := InterfaceDuplexLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InterfaceDuplexLabel: valid values are %v", v, AllowedInterfaceDuplexLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InterfaceDuplexLabel) IsValid() bool {
	for _, existing := range AllowedInterfaceDuplexLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Interface_duplex_label value
func (v InterfaceDuplexLabel) Ptr() *InterfaceDuplexLabel {
	return &v
}

type NullableInterfaceDuplexLabel struct {
	value *InterfaceDuplexLabel
	isSet bool
}

func (v NullableInterfaceDuplexLabel) Get() *InterfaceDuplexLabel {
	return v.value
}

func (v *NullableInterfaceDuplexLabel) Set(val *InterfaceDuplexLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceDuplexLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceDuplexLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceDuplexLabel(val *InterfaceDuplexLabel) *NullableInterfaceDuplexLabel {
	return &NullableInterfaceDuplexLabel{value: val, isSet: true}
}

func (v NullableInterfaceDuplexLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceDuplexLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
