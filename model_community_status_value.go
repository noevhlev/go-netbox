/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.0.11 (4.0)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
)

// CommunityStatusValue * `active` - Active * `reserved` - Reserved * `deprecated` - Deprecated
type CommunityStatusValue string

// List of Community_status_value
const (
	COMMUNITYSTATUSVALUE_ACTIVE     CommunityStatusValue = "active"
	COMMUNITYSTATUSVALUE_RESERVED   CommunityStatusValue = "reserved"
	COMMUNITYSTATUSVALUE_DEPRECATED CommunityStatusValue = "deprecated"
)

// All allowed values of CommunityStatusValue enum
var AllowedCommunityStatusValueEnumValues = []CommunityStatusValue{
	"active",
	"reserved",
	"deprecated",
}

func (v *CommunityStatusValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CommunityStatusValue(value)
	for _, existing := range AllowedCommunityStatusValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CommunityStatusValue", value)
}

// NewCommunityStatusValueFromValue returns a pointer to a valid CommunityStatusValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCommunityStatusValueFromValue(v string) (*CommunityStatusValue, error) {
	ev := CommunityStatusValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CommunityStatusValue: valid values are %v", v, AllowedCommunityStatusValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CommunityStatusValue) IsValid() bool {
	for _, existing := range AllowedCommunityStatusValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Community_status_value value
func (v CommunityStatusValue) Ptr() *CommunityStatusValue {
	return &v
}

type NullableCommunityStatusValue struct {
	value *CommunityStatusValue
	isSet bool
}

func (v NullableCommunityStatusValue) Get() *CommunityStatusValue {
	return v.value
}

func (v *NullableCommunityStatusValue) Set(val *CommunityStatusValue) {
	v.value = val
	v.isSet = true
}

func (v NullableCommunityStatusValue) IsSet() bool {
	return v.isSet
}

func (v *NullableCommunityStatusValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommunityStatusValue(val *CommunityStatusValue) *NullableCommunityStatusValue {
	return &NullableCommunityStatusValue{value: val, isSet: true}
}

func (v NullableCommunityStatusValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommunityStatusValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}