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

// PrefixStatusValue * `container` - Container * `active` - Active * `reserved` - Reserved * `deprecated` - Deprecated
type PrefixStatusValue string

// List of Prefix_status_value
const (
	PREFIXSTATUSVALUE_CONTAINER  PrefixStatusValue = "container"
	PREFIXSTATUSVALUE_ACTIVE     PrefixStatusValue = "active"
	PREFIXSTATUSVALUE_RESERVED   PrefixStatusValue = "reserved"
	PREFIXSTATUSVALUE_DEPRECATED PrefixStatusValue = "deprecated"
)

// All allowed values of PrefixStatusValue enum
var AllowedPrefixStatusValueEnumValues = []PrefixStatusValue{
	"container",
	"active",
	"reserved",
	"deprecated",
}

func (v *PrefixStatusValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PrefixStatusValue(value)
	for _, existing := range AllowedPrefixStatusValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PrefixStatusValue", value)
}

// NewPrefixStatusValueFromValue returns a pointer to a valid PrefixStatusValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPrefixStatusValueFromValue(v string) (*PrefixStatusValue, error) {
	ev := PrefixStatusValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PrefixStatusValue: valid values are %v", v, AllowedPrefixStatusValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PrefixStatusValue) IsValid() bool {
	for _, existing := range AllowedPrefixStatusValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Prefix_status_value value
func (v PrefixStatusValue) Ptr() *PrefixStatusValue {
	return &v
}

type NullablePrefixStatusValue struct {
	value *PrefixStatusValue
	isSet bool
}

func (v NullablePrefixStatusValue) Get() *PrefixStatusValue {
	return v.value
}

func (v *NullablePrefixStatusValue) Set(val *PrefixStatusValue) {
	v.value = val
	v.isSet = true
}

func (v NullablePrefixStatusValue) IsSet() bool {
	return v.isSet
}

func (v *NullablePrefixStatusValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrefixStatusValue(val *PrefixStatusValue) *NullablePrefixStatusValue {
	return &NullablePrefixStatusValue{value: val, isSet: true}
}

func (v NullablePrefixStatusValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrefixStatusValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
