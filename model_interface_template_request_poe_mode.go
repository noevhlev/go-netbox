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

// InterfaceTemplateRequestPoeMode * `pd` - PD * `pse` - PSE
type InterfaceTemplateRequestPoeMode string

// List of InterfaceTemplateRequest_poe_mode
const (
	INTERFACETEMPLATEREQUESTPOEMODE_PD    InterfaceTemplateRequestPoeMode = "pd"
	INTERFACETEMPLATEREQUESTPOEMODE_PSE   InterfaceTemplateRequestPoeMode = "pse"
	INTERFACETEMPLATEREQUESTPOEMODE_EMPTY InterfaceTemplateRequestPoeMode = ""
)

// All allowed values of InterfaceTemplateRequestPoeMode enum
var AllowedInterfaceTemplateRequestPoeModeEnumValues = []InterfaceTemplateRequestPoeMode{
	"pd",
	"pse",
	"",
}

func (v *InterfaceTemplateRequestPoeMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := InterfaceTemplateRequestPoeMode(value)
	for _, existing := range AllowedInterfaceTemplateRequestPoeModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid InterfaceTemplateRequestPoeMode", value)
}

// NewInterfaceTemplateRequestPoeModeFromValue returns a pointer to a valid InterfaceTemplateRequestPoeMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInterfaceTemplateRequestPoeModeFromValue(v string) (*InterfaceTemplateRequestPoeMode, error) {
	ev := InterfaceTemplateRequestPoeMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for InterfaceTemplateRequestPoeMode: valid values are %v", v, AllowedInterfaceTemplateRequestPoeModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v InterfaceTemplateRequestPoeMode) IsValid() bool {
	for _, existing := range AllowedInterfaceTemplateRequestPoeModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to InterfaceTemplateRequest_poe_mode value
func (v InterfaceTemplateRequestPoeMode) Ptr() *InterfaceTemplateRequestPoeMode {
	return &v
}

type NullableInterfaceTemplateRequestPoeMode struct {
	value *InterfaceTemplateRequestPoeMode
	isSet bool
}

func (v NullableInterfaceTemplateRequestPoeMode) Get() *InterfaceTemplateRequestPoeMode {
	return v.value
}

func (v *NullableInterfaceTemplateRequestPoeMode) Set(val *InterfaceTemplateRequestPoeMode) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceTemplateRequestPoeMode) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceTemplateRequestPoeMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceTemplateRequestPoeMode(val *InterfaceTemplateRequestPoeMode) *NullableInterfaceTemplateRequestPoeMode {
	return &NullableInterfaceTemplateRequestPoeMode{value: val, isSet: true}
}

func (v NullableInterfaceTemplateRequestPoeMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceTemplateRequestPoeMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
